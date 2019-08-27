package auth

import (
	"net/http"
	"strings"
	"time"

	"golang_api/common"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type login struct {
	Userid   string `bson:"id" json:"id"`
	Password string `bson:"password" json:"password"`
}

type Handler struct{}

func (h *Handler) Login(c echo.Context) error {
	var l login
	err := c.Bind(&l)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_99", err.Error(), ""}})
	}

	a := common.ValidateStructValues(l)
	if a != "" {
		a = strings.TrimRight(a, ",")
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_03", "The field(s) is/are required", a}})
	}

	cu, e := getOne("id", l.Userid)
	if e != nil || cu.Password != l.Password {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_04", "Invalid userid / password", ""}})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = cu.Userid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("P@ssw0rd"))
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_99", err.Error(), ""}})
	}

	return c.JSON(http.StatusOK, common.ErrorWithData{common.Error{"100", "AUT_00", "Success", ""}, map[string]string{"token": t}})
}

func getOne(fkey string, fvalue string) (i login, err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB("golang_training2").C("customers")
	q := bson.M{fkey: fvalue}
	err = c.Find(q).One(&i)

	return
}

func Authenticate(claims *jwt.Token) string {
	c := claims.Claims.(jwt.MapClaims)
	s := c["exp"].(float64)

	if float64(time.Now().Unix()) > s {
		return "expired"
	}

	return "ok"
}
