package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/subeenchung/cath-palug/config"
	"github.com/subeenchung/cath-palug/models"
	"github.com/subeenchung/cath-palug/routes"
)

var cfg config.MainConfig

type Env struct {
	db models.DataStore
}

func main() {
	//Load configuration
	cfg = config.LoadConfig("./config.toml")
	//Load database connection
	db, err := models.NewDB(fmt.Sprintf("%s://%s:%s@%s:%d/%s", cfg.DB.Type, cfg.DB.User, cfg.DB.Password, cfg.DB.IP, cfg.DB.Port, cfg.DB.Dbname))
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	//env := &Env{db}
	e := echo.New()
	e.GET("/test", routes.Handler)
	e.GET("/request", func(c echo.Context) error {
		req := c.Request()
		format := `
			<code>
				Protocol: %s<br>
				Host: %s<br>
				Remote Address: %s<br>
				Method: %s<br>
				Path: %s<br>
			</code>
		`
		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	})
	e.Logger.Fatal(e.StartTLS(":3000", "./ssl/selfsigned.crt", "./ssl/selfsigned.key"))
}
