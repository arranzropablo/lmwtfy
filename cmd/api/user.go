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
// @Param user body string false "name search by q"
// @Param password query string false "name search by q"
// @Success 200 {array} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [get]
func CreateUser(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Create user")
}

func EditUser(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Edit user")
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Getting user: " + req.URL.Path)
}