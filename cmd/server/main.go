package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"cmd/server/main.go/internal/core"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var Master *sql.DB

func MainHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!\n")
	io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	core.AddUser(Master, username, "", "", "")
	io.WriteString(w, fmt.Sprintf("Hello %v", vars["username"]))
}

func ListUserHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := core.GetAllUsers(Master)
	io.WriteString(w, fmt.Sprintf("%v", users))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DB")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDB)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	Master = db
	defer Master.Close()

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web/app/build/static"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/", MainHandler)
	apiRouter.HandleFunc("/user/{username}", UserHandler)
	apiRouter.HandleFunc("/user", ListUserHandler)

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/app/build/index.html")
	})

	fmt.Println("Listening on port 5050...")

	http.Handle("/", r)
	http.ListenAndServe(":5050", r)
}
