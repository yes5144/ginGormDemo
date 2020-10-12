package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yes5144/ginGormDemo/models"
	"github.com/yes5144/ginGormDemo/routers"
	"github.com/yes5144/ginGormDemo/utils"
)

func main() {
	// initConfig
	utils.InitConfig()

	// initDb
	db := models.InitDb()
	defer db.Close()

	// initRouter
	r := routers.InitRouter()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 88),
		Handler: r,
		// ReadHeaderTimeout: 60,
		// WriteTimeout:      60,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("http_port: ", 88)
	s.ListenAndServe()
}
