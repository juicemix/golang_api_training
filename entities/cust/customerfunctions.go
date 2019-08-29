package cust

import (
	"golang_api/common"
	"golang_api/entities/auth"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Add(c echo.Context) error {
	aut := auth.Authenticate(c.Get("user").(*jwt.Token))
	if aut == "expired" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "Access Unathorized", ""}})
	}

	var cu customerInput
	err := c.Bind(&cu)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_99", err.Error(), ""}})
	}

	a := common.ValidateStructValues(cu)
	if a != "" {
		a = strings.TrimRight(a, ",")
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_02", "The field(s) is/are required", a}})
	}

	if !common.ValidateString(common.EMAIL_REGEX, cu.Email) {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_03", "Invalid email", "email"}})
	}

	if _, e := UserExistsById(cu.Id); e {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_12", "This userid is already used", "id"}})
	}

	if _, e := UserExistsByEmail(cu.Email); e {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_04", "This email is already used", "email"}})
	}

	cin := InsertPreparation(cu)

	err = insert(cin)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_99", err.Error(), ""}})
	}

	return c.JSON(http.StatusCreated, common.ErrorReturn{common.Error{"001", "USR_00", "Success", ""}})
}

func ViewAll(c echo.Context) error {
	aut := auth.Authenticate(c.Get("user").(*jwt.Token))
	if aut == "expired" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "Access Unathorized", ""}})
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil && c.QueryParam("limit") != "" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_10", "Wrong limit input", "limit"}})
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil && c.QueryParam("page") != "" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_11", "Wrong page input", "limit"}})
	}

	ret, _ := getAll([]string{"timestamp"}, limit, page)

	return c.JSON(http.StatusOK, common.ErrorWithData{common.Error{"001", "USR_00", "Success", ""}, ret})
}

func Find(c echo.Context) error {
	aut := auth.Authenticate(c.Get("user").(*jwt.Token))
	if aut == "expired" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "Access Unathorized", ""}})
	}

	id := c.Param("id")

	if u, f := UserExistsById(id); f {
		return c.JSON(http.StatusOK, common.ErrorWithData{common.Error{"001", "USR_00", "Success", ""}, u})
	} else {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "User not found", ""}})
	}
}

func GetLoggedInData(c echo.Context) error {
	t := c.Get("user").(*jwt.Token)
	aut := auth.Authenticate(t)
	if aut == "expired" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "Access Unathorized", ""}})
	}

	x := t.Claims.(jwt.MapClaims)
	id := x["id"].(string)

	if u, f := UserExistsById(id); f {
		return c.JSON(http.StatusOK, common.ErrorWithData{common.Error{"001", "USR_00", "Success", ""}, u})
	} else {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "User not found", ""}})
	}
}

func Edit(c echo.Context) error {
	aut := auth.Authenticate(c.Get("user").(*jwt.Token))
	if aut == "expired" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "Access Unathorized", ""}})
	}

	var cu customerInput
	err := c.Bind(&cu)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_99", err.Error(), ""}})
	}

	cin := InsertPreparation(cu)
	err = update(cin)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_99", err.Error(), ""}})
	}

	return c.JSON(http.StatusCreated, common.ErrorReturn{common.Error{"001", "USR_00", "Success", ""}})
}

func UserExistsById(id string) (customer, bool) {
	c, _ := getOne("id", id)

	if c.Firstname == "" {
		return c, false
	} else {
		return c, true
	}
}

func UserExistsByEmail(email string) (customer, bool) {
	c, err := getOne("email", email)

	if err != nil && err.Error() == "not found" {
		return c, false
	} else {
		return c, true
	}
}
