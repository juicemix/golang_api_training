package cust

import (
	"golang_api/common"
	"golang_api/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func Add(c echo.Context) error {
	cu := new(customerInput)
	err := c.Bind(cu)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_99", err.Error(), ""}})
	}

	a := common.ValidateStructValues(cu)
	if a != "" {
		a = strings.TrimRight(a, ",")
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_02", "The field(s) is/are required", a}})
	}

	if UserExistsById(cu.Id) {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_12", "This userid is already used", "id"}})
	}

	if UserExistsByEmail(cu.Email) {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_04", "This email is already used", "email"}})
	}

	err = db.Insert("golang_training2", "customers", cu)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_99", err.Error(), ""}})
	}

	return c.JSON(http.StatusCreated, common.ErrorReturn{common.Error{"001", "USR_00", "Success", ""}})
}

func ViewAll(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil && c.QueryParam("limit") != "" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_10", "Wrong limit input", "limit"}})
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil && c.QueryParam("page") != "" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "USR_11", "Wrong page input", "limit"}})
	}

	return c.String(http.StatusOK, strconv.Itoa(limit+page))
}

func UserExistsById(id string) bool {
	c, err := findByKey("id", id)

	if err != nil {
		return true
	}

	if c.Id == "" {
		return false
	} else {
		return true
	}
}

func UserExistsByEmail(email string) bool {
	c, err := findByKey("email", email)

	if err != nil {
		return true
	}

	if c.Id == "" {
		return false
	} else {
		return true
	}
}
