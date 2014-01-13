package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	ctx "github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func API(w http.ResponseWriter, r *http.Request) {
	if u, err := json.Marshal(ctx.Get(r, "user")); err == nil {
		writeJSON(w, u)
	} else {
		http.Error(w, "Server Error", 500)
	}
}

//API_Campaigns returns a list of campaigns if requested via GET.
//If requested via POST, API_Campaigns creates a new campaign and returns a reference to it.
func API_Campaigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method == "GET":

	case r.Method == "POST":
		fmt.Fprintf(w, "Hello POST!")
	}
	//fmt.Fprintf(w, "Hello api")
}

//API_Campaigns_Id returns details about the requested campaign. If the campaign is not
//valid, API_Campaigns_Id returns null.
func API_Campaigns_Id(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fmt.Fprintf(w, "{\"method\" : \""+r.Method+"\", \"id\" : "+vars["id"]+"}")
}

//API_Doc renders a template describing the API documentation.
func API_Doc(w http.ResponseWriter, r *http.Request) {
	getTemplate(w, "api_doc").ExecuteTemplate(w, "base", nil)
}

func writeJSON(w http.ResponseWriter, c []byte) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", c)
}
