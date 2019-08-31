package prod

/*
func Add(c echo.Context) error {
	aut := auth.Authenticate(c.Get("user").(*jwt.Token))
	if aut == "expired" {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "AUT_09", "Access Unathorized", ""}})
	}

	var cu productInput
	err := c.Bind(&cu)
	if err != nil {
		return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{"100", "PRO_99", err.Error(), ""}})
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

	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func AddCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func RemoveCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func ViewAll(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func Find(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func FindByCategories(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func UpdateName(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func UpdateStock(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func UpdatePrice(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func UpdateCategories(c echo.Context) error {
	return c.JSON(http.StatusOK, common.ErrorReturn{common.Error{Status: "100", Code: "PRO_99", Message: "Unimplemented", Field: ""}})
}

func ProdExistsById(id string) (product, bool) {
	c, _ := getOne("id", id)

	if c.Product_name == "" {
		return c, false
	} else {
		return c, true
	}
}
*/
