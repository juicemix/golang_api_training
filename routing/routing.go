package routing

import (
	"net/http"

	"golang_api/entities/auth"
	"golang_api/entities/cat"
	"golang_api/entities/cust"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start the API server
func RoutingStart() {
	e := echo.New()
	e.HideBanner = true

	// Documentations
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h3>Available routes : </h3><br/><table><th><td>Routes</td></th><tr><td>/login</td></tr><tr><td>/customers</td></tr><table>")
	})

	// Authentication
	h := &auth.Handler{}
	e.GET("/login", h.Login)

	isLoggedIn := middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("P@ssw0rd")})

	// Customers functions
	e.GET("/customers", cust.ViewAll, isLoggedIn)
	e.GET("/customers", cust.GetLoggedInData, isLoggedIn)
	e.GET("/customers/:id", cust.Find, isLoggedIn)
	e.POST("/customers", cust.Add, isLoggedIn)
	e.PUT("/customers", cust.Edit, isLoggedIn)

	// Categories functions
	e.GET("/categories", cat.ViewAll, isLoggedIn)
	e.GET("/categories/:id", cat.Find, isLoggedIn)

	// Products functions

	// Start
	e.Logger.Fatal(e.Start(":6969"))
}
