package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

//To access application/json request fields they have to be exported (capitalize)
type body struct {
	Search string `json:"search" example:"madgyvers"`
}

// Search godoc
// @Summary Create user
// @Description Creates an user
// @Accept x-www-form-urlencoded
// @Produce json
// @Param main body search.body true "The string you want to search by"
// @Success 200 {string} string "Everything went OK"
// @Router /search [post]
func Search(res http.ResponseWriter, req *http.Request) {
	//if we accepted an application/json
	//var search body
	//decoder := json.NewDecoder(req.Body)
	//err := decoder.Decode(&search)
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	//cant define an example value for response in swaggo
	//need to deploy it publicly
	//https://blog.golang.org/error-handling-and-go Simplifying repetitive error handling

	req.ParseForm()
	search := req.FormValue("text")
	var mention string

	splt := regexp.MustCompile(" <@").Split(search, -1)

	if len(splt) > 1 {
		mention = "<@" + regexp.MustCompile("\\|").Split(splt[len(splt) - 1], -1)[0] + ">"
		search = strings.Join(splt[:len(splt) - 1], "")
	}

	if link, ok := retrieveLink(search); ok {

		jsonResp, _ := json.Marshal(struct {
			Type string `json:"response_type"`
			Text string `json:"text"`
		}{
			Type: "in_channel",
			Text: fmt.Sprintf("%s", mention + " let me wiki that for you... \n" + link),
		})

		res.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(res, string(jsonResp))
	} else {
		http.Error(res, "Empty search!!!", http.StatusInternalServerError)
		return
	}
}

func retrieveLink(search string) (string, bool) {
	if search == "" {
		return "", false
	}
	s := strings.Join(strings.Split(search, " "), "+")

	return "https://wiki.web.zooplus.de/dosearchsite.action?cql=siteSearch+~+" +
		"\"" + url.PathEscape(s) + "\"" +
		"&queryString=" + url.PathEscape(s), true

}