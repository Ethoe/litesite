package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"cmd/server/main.go/internal/core"
	"cmd/server/main.go/internal/db"
	"cmd/server/main.go/pkg/middleware"

	"github.com/gorilla/mux"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!\n")
	io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	db.SetupDB()
	defer db.Master.Close()

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web/app/build"))

	r.PathPrefix("/static/").Handler(fs)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/", MainHandler)
	apiRouter.HandleFunc("/user/register", core.AddUserHandler).Methods("POST")
	apiRouter.HandleFunc("/user/login", core.LoginHandler).Methods("POST")
	apiRouter.HandleFunc("/user/logout", core.LogoutHandler).Methods("GET")
	apiRouter.HandleFunc("/user/self", core.SelfHandlers).Methods("GET")
	apiRouter.HandleFunc("/user/list/all", core.GetAllUsersHandler).Methods("GET")

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/app/build/index.html")
	})

	r.Use(middleware.AuthMiddleware)

	fmt.Println("Listening on port 5050...")

	http.Handle("/", r)
	http.ListenAndServe(":5050", r)
}
