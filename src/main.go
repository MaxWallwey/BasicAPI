package main

import (
	"basic-api/cassandra"
	"basic-api/users"
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	usersHandler := NewUsersHandler()
	// Create the router
	router := mux.NewRouter()

	// Register the routes
	router.HandleFunc("/users", usersHandler.ListUsers).Methods("GET")
	router.HandleFunc("/users", usersHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", usersHandler.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", usersHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", usersHandler.DeleteUser).Methods("DELETE")

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		return
	}
}

type userStore interface {
	Add(name string, user users.User) error
	Get(name string) (users.User, error)
	List() (map[string]users.User, error)
	Update(name string, user users.User) error
	Remove(name string) error
}

type UsersHandler struct {
	store userStore
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (h UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
}

func (h UsersHandler) ListUsers(w http.ResponseWriter, r *http.Request) {}
func (h UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// users.User{ID: gocql.UUIDFromTime(time.Now()), Name: "Max", EmailAddress: "test@test.com", Birthday: time.Now()}
	session := cassandra.SetupCassandra()
	defer session.Close()

	var user users.User

	err := session.Query("SELECT * FROM store.users WHERE id = ? LIMIT 1", id).Consistency(gocql.One).Scan(&user)
	if err != nil {
		InternalServerErrorHandler(w, r)
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		return
	}
}
func (h UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
func (h UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
