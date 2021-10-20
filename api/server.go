package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KaiqueSantosDev/gocrud/api/core/price"
	"github.com/KaiqueSantosDev/gocrud/api/database"
	"github.com/KaiqueSantosDev/gocrud/api/handler"
	"github.com/gorilla/mux"
)

func Run() {
	database.InitDatabase()
	service := price.NewService(database.DB)
	r := mux.NewRouter()

	handler.MakeProductHandlers(r, service)

	srv := &http.Server{

		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8000",
		Handler:      handler.LoadCors(r),
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
