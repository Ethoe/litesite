package users_api

import "cmd/server/main.go/internal/core/users_api/handlers"

var AddUserHandler = handlers.AddUser
var GetAllUsersHandler = handlers.GetAllUsers
var LoginHandler = handlers.LoginUser
var LogoutHandler = handlers.LogoutUser
var SelfHandlers = handlers.SelfUser
