package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

var restaurants []Restaurant



func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(restaurants)
}

func GetSingleRestaurant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range restaurants {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Restaurant{})
}

func CreateRestaurants(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Restaurant
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	restaurants = append(restaurants, person)
	json.NewEncoder(w).Encode(restaurants)
}
func EditRestaurants(w http.ResponseWriter, r *http.Request){
          params := mux.Vars(r)
	 for index, item := range restaurants {

		 if item.ID == params["id"]{
			 item.Name=params["name"];
		 }
	 }

	
}
func DeleteRestaurants(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range restaurants {
		if item.ID == params["id"] {
			restaurants = append(restaurants[:index], restaurants[index+1:]...)
			break;
		}
		json.NewEncoder(w).Encode(restaurants)
	}
}


type Restaurant struct {
	ID        string   `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// our main function
func main() {
	router := mux.NewRouter()
	//fmt.Println(reflect.TypeOf(router))

	restaurants = append(restaurants, Restaurant{ID: "1", Name: "John", Address: &Address{City: "City X", State: "State X"}})
	restaurants = append(restaurants, Restaurant{ID: "2", Name: "Koko",  Address: &Address{City: "City Z", State: "State Y"}})
	restaurants = append(restaurants, Restaurant{ID: "3", Name: "Francis"})


	router.HandleFunc("/restaurants", GetRestaurants).Methods("GET")
	router.HandleFunc("/restaurants/{id}", GetSingleRestaurant).Methods("GET")
	router.HandleFunc("/restaurants/{id}", CreateRestaurants).Methods("POST")
	router.HandleFunc("/restaurants/{id}",EditRestaurants).Methods("PUT")
	router.HandleFunc("/restaurants/{id}", DeleteRestaurants).Methods("DELETE")
	//http.ListenAndServe(":4646", router)

	log.Fatal(http.ListenAndServe(":4646", router)) // 4646 on PhoneDial stands for GoGo
}



