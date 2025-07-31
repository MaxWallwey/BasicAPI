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

	ctx := r.Context()

	session := cassandra.SetupCassandra()
	defer session.Close()

	err := session.Query("INSERT INTO store.users (id, name, email_address, last_updated_timestamp) VALUES (?,?,?,?)", user.ID, user.Name, user.EmailAddress, user.LastUpdatedTimestamp).WithContext(ctx).Exec()
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h UsersHandler) ListUsers(w http.ResponseWriter, r *http.Request) {}
func (h UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ctx := r.Context()

	session := cassandra.SetupCassandra()
	defer session.Close()

	var user users.User
	queryString := "SELECT id, name, email_address, last_updated_timestamp FROM store.users WHERE id = ? LIMIT 1"
	err := session.Query(queryString, id).Consistency(gocql.One).WithContext(ctx).Scan(&user.ID, &user.Name, &user.EmailAddress, &user.LastUpdatedTimestamp)
	if err != nil {
		NotFoundHandler(w, r)
		return
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
	_, err := w.Write([]byte("500 Internal Server Error"))
	if err != nil {
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("404 Not Found"))
	if err != nil {
		return
	}
}
