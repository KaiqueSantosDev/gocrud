package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/KaiqueSantosDev/gocrud/api/core/price"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, service price.UseCase) {
	r.Handle("/", getAllProducts(service)).Methods("GET", "OPTIONS")
	r.Handle("/{id}", getProduct(service)).Methods("GET", "OPTIONS")
	r.Handle("/", createProduct(service)).Methods("POST", "OPTIONS")
	r.Handle("/{id}", updateProduct(service)).Methods("PUT", "OPTIONS")
	r.Handle("/{id}", deleteProduct(service)).Methods("DELETE", "OPTIONS")
	r.Handle("/compare/", compareProduct(service)).Methods("POST", "OPTIONS")

}

func getAllProducts(service price.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		all, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(all)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
}

func getProduct(service price.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(&p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	})
}

func createProduct(service price.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		p := price.Product{}
		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		service.Create(&p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func updateProduct(service price.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		product := price.Product{}
		product.ID = id
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Fatal("unable to decode")
		}
		service.Update(&product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
}

func deleteProduct(service price.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		service.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func compareProduct(service price.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p []*price.Product
		var temp, largerNumber float64
		json.NewDecoder(r.Body).Decode(&p)

		for _, p := range p {
			if float64(p.Price) > temp {
				temp = float64(p.Price)
				largerNumber = temp
			}
		}
		json.NewEncoder(w).Encode(&largerNumber)

	})
}
