package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit....Homepage")
	fmt.Fprintf(w, "Hi Aravind fghjk")

}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("End Point hit ... return All products ")
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	//key := r.URL.Path[len("/product/"):]
	for _, i := range Products {
		if string(i.Id) == key {
			json.NewEncoder(w).Encode(i)
		}

	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/product/{id}", getProduct)
	myRouter.HandleFunc("/products", returnAllProducts)

	myRouter.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", myRouter)

}
func main() {
	fmt.Println("j")
	Products = []Product{
		Product{Id: "1", Name: "chair", Quantity: 3, Price: 100.5},
		Product{Id: "2", Name: "table", Quantity: 6, Price: 100.5},
		Product{Id: "3", Name: "scale", Quantity: 3, Price: 100.5},
		Product{Id: "4", Name: "pens", Quantity: 8, Price: 100.5},
	}
	handleRequests()
}
