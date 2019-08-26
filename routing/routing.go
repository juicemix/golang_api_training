package routing

import (
	"net/http"

	"golang_api/cust"

	"github.com/labstack/echo"
)

// Start the API server
func RoutingStart() {
	e := echo.New()

	e.GET("/HELLO", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "{\"message\":"+id+"}")
	})

	e.GET("/customers", cust.ViewAll)
	e.POST("customers", cust.Add)

	e.Logger.Fatal(e.Start(":6969"))
}
