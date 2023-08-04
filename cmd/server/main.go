package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"cmd/server/main.go/internal/core/files_api"
	"cmd/server/main.go/internal/core/users_api"
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
	apiRouter.Use(middleware.AuthMiddleware)
	apiRouter.HandleFunc("/", MainHandler)
	apiRouter.HandleFunc("/user/register", users_api.AddUserHandler).Methods("POST")
	apiRouter.HandleFunc("/user/login", users_api.LoginHandler).Methods("POST")
	apiRouter.HandleFunc("/user/logout", users_api.LogoutHandler).Methods("GET")
	apiRouter.HandleFunc("/user/self", users_api.SelfHandlers).Methods("GET")
	apiRouter.HandleFunc("/user/list/all", users_api.GetAllUsersHandler).Methods("GET")

	apiRouter.HandleFunc("/file", files_api.UploadFileHandler).Methods("POST")
	apiRouter.HandleFunc("/file/{id}", files_api.DeleteFileHandler).Methods("DELETE")
	apiRouter.HandleFunc("/file/list/all", files_api.GetAllFilesHandler).Methods("GET")

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/app/build/index.html")
	})

	fmt.Println("Listening on port 5050...")

	http.Handle("/", r)
	http.ListenAndServe(":5050", r)
}
