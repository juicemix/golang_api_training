package cat

import (
	"golang_api/common"
	"golang_api/entities/auth"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

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

	if u, f := CatExistsById(id); f {
		return c.JSON(http.StatusOK, common.ErrorWithData{common.Error{"001", "USR_00", "Success", ""}, u})
	} else {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "User not found", ""}})
	}
}

func CatExistsById(id string) (category, bool) {
	c, _ := getOne("id", id)

	if c.Category_name == "" {
		return c, false
	} else {
		return c, true
	}
}
