package search

import (
	"io"
	"net/http"
)

// Search godoc
// @Summary Create user
// @Description Creates an user
// @Accept json
// @Produce json
// @Param main body search.body true "The string you want to search by"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "no ok"
// @Router /search [post]
func Search(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Search: " + req.FormValue("search"))
}


type body struct {
	search string `json:"search" example:"madgyvers"`
}