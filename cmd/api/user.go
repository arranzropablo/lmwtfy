package user

import (
	"io"
	"net/http"
)

// CreateUser godoc
// @Summary Create user
// @Description Creates an user
// @Accept  json
// @Produce  json
// @Param user body string true "username"
// @Param password body string true "password"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "no ok"
// @Router /user/new [get]
func CreateUser(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Create user")
}

func EditUser(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Edit user")
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Getting user: " + req.URL.Path)
}