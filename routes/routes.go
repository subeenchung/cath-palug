package routes

import 
(
//"encoding/json"
"net/http"
"github.com/labstack/echo"
)

func Handler(c echo.Context) error {
	return c.JSON(http.StatusOK, "kek")
}
