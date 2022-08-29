package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/mail"
	"strconv"
	"wallester_test/database"
	"wallester_test/entities"
)

/*func GetCustomerByFullName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer []entities.Customer
	database.Instance.Where("first_name = ? AND last_name = ? ", params["FirstName"], params["LastName"]).Find(&customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}*/

func GetCustomerByFirstName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["FirstName"]
	var customers []entities.Customer
	if firstNameValidator(params) == false {
		json.NewEncoder(w).Encode(fmt.Sprintf("Customer with name: %s not found", params))
		return
	}
	database.Instance.Where("first_name = ?", params).Find(&customers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerByLastName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["LastName"]
	var customer []entities.Customer
	if lastNameValidator(params) == false {
		json.NewEncoder(w).Encode(fmt.Sprintf("Customer with name: %s not found", params))
		return
	}
	database.Instance.Where("last_name = ?", params).Find(&customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var customer []entities.Customer
	database.Instance.Find(&customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customerId := mux.Vars(r)["id"]
	if idValidator(customerId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(fmt.Sprintf("Customer with name: %s not found", customerId))
		return
	}
	var product entities.Customer
	database.Instance.Delete(&product, customerId)
	json.NewEncoder(w).Encode("Customer Deleted Successfully!")
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer entities.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	if _, err := strconv.Atoi(customer.FirstName); err == nil {
		json.NewEncoder(w).Encode("First name must not contain numbers!")
		return
	} else if _, err := strconv.Atoi(customer.LastName); err == nil {
		json.NewEncoder(w).Encode("Last name must not contain numbers!")
		return
	} else if len(customer.FirstName) == 0 {
		json.NewEncoder(w).Encode("FirstName is empty!")
		return
	} else if len(customer.LastName) == 0 {
		json.NewEncoder(w).Encode("LastName is empty!")
		return
	} else if emailValidator(customer.Email) != true {
		json.NewEncoder(w).Encode("Email is incorrect!")
		return
	}
	database.Instance.Create(&customer)
	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)["id"]
	if idValidator(customerId) == false {
		json.NewEncoder(w).Encode(fmt.Sprintf("Customer with name: %s not found", customerId))
		return
	}
	var customer entities.Customer
	tx := database.Instance.Begin()
	tx.First(&customer, customerId)
	json.NewDecoder(r.Body).Decode(&customer)
	if _, err := strconv.Atoi(customer.FirstName); err == nil {
		json.NewEncoder(w).Encode("First name must not contain numbers!")
		return
	} else if _, err := strconv.Atoi(customer.LastName); err == nil {
		json.NewEncoder(w).Encode("Last name must not contain numbers!")
		return
	} else if len(customer.FirstName) == 0 {
		json.NewEncoder(w).Encode("FirstName is empty!")
		tx.Rollback()
		return
	} else if len(customer.LastName) == 0 {
		json.NewEncoder(w).Encode("LastName is empty!")
		tx.Rollback()
		return
	} else if emailValidator(customer.Email) != true {
		json.NewEncoder(w).Encode("Email is incorrect!")
		tx.Rollback()
		return
	} else {
		tx.Save(&customer)
		tx.Commit()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

func idValidator(id string) bool {
	var customer entities.Customer
	database.Instance.Find(&customer, id)
	if customer.ID == 0 {
		return false
	}
	return true
}

func firstNameValidator(name string) bool {
	var exists bool
	err := database.Instance.Model(entities.Customer{FirstName: name}).
		Select("count(*) > 0").
		Where("first_name = ?", name).
		Find(&exists).
		Error
	if err != nil {
		return false
	}
	if exists == false {
		return false
	}
	return true
}

func lastNameValidator(name string) bool {
	var exists bool
	err := database.Instance.Model(entities.Customer{LastName: name}).
		Select("count(*) > 0").
		Where("last_name = ?", name).
		Find(&exists).
		Error
	if err != nil {
		return false
	}
	if exists == false {
		return false
	}
	return true
}

func emailValidator(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
