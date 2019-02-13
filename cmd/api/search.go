package search

import (
	"encoding/json"
	"net/http"
)

// Search godoc
// @Summary Create user
// @Description Creates an user
// @Accept json
// @Produce json
// @Param main body search.body true "The string you want to search by"
// @Success 200 {string} string "Everything went OK"
// @Router /search [post]
func Search(res http.ResponseWriter, req *http.Request) {
	var search body
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&search)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	//https://blog.golang.org/error-handling-and-go Simplifying repetitive error handling
	_, err = res.Write([]byte(search.Search))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}

type body struct {
	//To access application/json request fields they have to be exported
	Search string `json:"search" example:"madgyvers"`
}