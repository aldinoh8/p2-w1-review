package main

import (
	"demo/config"
	"demo/controller"
	"demo/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.InitDatabase()
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "internal server error",
			"detail":  i,
		})
	}

	branchRepository := repository.Branch{DB: db}
	branchController := controller.Branch{
		Repository: branchRepository,
	}
	router.GET("/branches", branchController.Index)
	router.GET("/branches/:id", branchController.Detail)
	router.POST("/branches/", branchController.Create)
	router.PUT("/branches/:id", branchController.Update)
	router.DELETE("/branches/:id", branchController.Delete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
