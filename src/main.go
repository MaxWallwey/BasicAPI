package main

import (
	"basic-api/src/cassandra"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	Session := cassandra.SetupCassandra()
	defer Session.Close()

	handleRequest()
}

func handleRequest() {
	router := mux.NewRouter()

	router.HandleFunc("/users/findMany", findManyUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/fineOne/{id}", findOneUser).Methods(http.MethodGet)
	router.HandleFunc("/users/add", addOneUser).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// findManyUsers responds with the list of all users as JSON.
func findManyUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("get users")

	//err := json.NewEncoder(w).Encode(users)
	//if err != nil {
	//	return
	//}
}

// findOneUser locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func findOneUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	//err := json.NewEncoder(w).Encode(user)
	//if err != nil {
	//	return
	//}

	log.Println(fmt.Sprintf("get user with id: %s", id))
}

// addOneUser adds a user from JSON received in the request body.
func addOneUser(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	log.Println(fmt.Sprintf("post user"))
}
