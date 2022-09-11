// Filename: cmd/api/entries.go
package main
import (
	"fmt"
	"net/http"
)
//create entires hander for the POST /v1/entries endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Create a New Entry")
}
//create showentires hander for the GET /v1/entries/:id endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request){
	
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w,r)
		return
	}
	// Display the entires
	fmt.Fprintf(w, "Show The Details For Entry %d\n", id)
}