package users

import "time"

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	RegDate  time.Time `json:"regDate"`
}

type Session struct {
	ID           int       `json:"id"`
	UserID       int       `json:"userId"`
	SessionToken string    `json:"sessionToken"`
	Creation     time.Time `json:"creation"`
}
