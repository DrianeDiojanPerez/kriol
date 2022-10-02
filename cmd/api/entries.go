// Filename: cmd/api/entries.go
package main

import (
	
	"fmt"
	"net/http"
)

//create entires hander for the POST /v1/entries endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request){
	//our target decode destination
	var input struct{
		Name string `json:"name"`
		String string `json:"string"`
		Translate string `json:"translate"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)

	//fmt.Fprintln(w, "Create a New Entry")
}
//create showentires hander for the GET /v1/entries/:id endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request){
	
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Display the entires
	fmt.Fprintf(w, "Show The Details For Entry %d\n", id)
}