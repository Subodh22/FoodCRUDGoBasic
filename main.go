package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Food struct {
	ID       string `json:"ID"`
	Title    string `json:"title"`
	Calories string `json:"calories"`
	Protein  string `json:"protein"`
	Fat      string `json:"fat"`
	Carbs    string `json:"carbs"`
	Diet     *Diet
}

type Diet struct {
	Dietname string `json:"dietname"`
	Type     string `json:"type"`
}

var foods []Food

func getFoods(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(foods)
	return

}
func deleteFood(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range foods {
		if item.ID == params["ID"] {
			foods = append(foods[:index], foods[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(foods)
	return
}

func createFood(w http.ResponseWriter, r *http.Request) {
	var food Food

	json.NewDecoder(r.Body).Decode(&food)
	food.ID = strconv.Itoa(rand.Intn(100000))
	foods = append(foods, food)
	json.NewEncoder(w).Encode(food)
	return

}
func getFoodById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range foods {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func updateFood(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var food Food

	for index, item := range foods {
		if item.ID == params["ID"] {
			foods = append(foods[:index], foods[index+1:]...)
			food.ID = item.ID
			json.NewDecoder(r.Body).Decode(&food)
			foods = append(foods, food)

		}
	}
	json.NewEncoder(w).Encode(foods)
	return

}

func main() {
	r := mux.NewRouter()
	foods = append(foods, Food{ID: "1", Title: "Chix", Calories: "100", Protein: "30", Fat: "12", Carbs: "0", Diet: &Diet{Dietname: "Nonvg", Type: "carni"}})
	foods = append(foods, Food{ID: "2", Title: "Rice", Calories: "100", Protein: "1", Fat: "2", Carbs: "20", Diet: &Diet{Dietname: "veg", Type: "omni"}})

	fmt.Printf("Starting a server is 8000")
	r.HandleFunc("/foods", getFoods).Methods("GET")
	r.HandleFunc("/food/{ID}", getFoodById).Methods("GET")
	r.HandleFunc("/food/{ID}", deleteFood).Methods("DELETE")
	r.HandleFunc("/food", createFood).Methods("POST")
	r.HandleFunc("/food/{ID}", updateFood).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
