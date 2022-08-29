package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"wallester_test/controllers"
	"wallester_test/database"
)

func main() {
	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	router := mux.NewRouter().StrictSlash(true)
	RegisterCustomerRoutes(router)
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterCustomerRoutes(router *mux.Router) {
	router.HandleFunc("/", index)
	router.HandleFunc("/customer/getAll", controllers.GetAll).Methods("GET")
	//router.HandleFunc("/customer/getCustomer/{FirstName}+{LastName}", controllers.GetCustomerByFullName).Methods("GET")
	router.HandleFunc("/customer/getFirst/{FirstName}", controllers.GetCustomerByFirstName).Methods("GET")
	router.HandleFunc("/customer/getLast/{LastName}", controllers.GetCustomerByLastName).Methods("GET")
	router.HandleFunc("/customer/create", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/delete/{id}", controllers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customer/update/{id}", controllers.UpdateCustomer).Methods("PUT")

}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("templates/index.gohtml")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, "index")
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
}
