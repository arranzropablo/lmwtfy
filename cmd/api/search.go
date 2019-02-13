package search

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

//To access application/json request fields they have to be exported (capitalize)

type body struct {
	Search string `json:"search" example:"madgyvers"`
}

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

	//cant define an example value for response in swaggo
	//https://blog.golang.org/error-handling-and-go Simplifying repetitive error handling

	if link, ok := retrieveLink(search.Search); ok {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(link))
	} else {
		http.Error(res, "Empty search!!!", http.StatusInternalServerError)
		return
	}
}


//scape chars before writing em into the body
func retrieveLink(search string) (string, bool) {
	if search == "" {
		return "", false
	}
	s := strings.Join(strings.Split(search, " "), "+")

	return "https://wiki.web.zooplus.de/dosearchsite.action?cql=siteSearch+~+" +
		"\"" + url.PathEscape(s) + "\"" +
		"&queryString=" + url.PathEscape(s), true

}