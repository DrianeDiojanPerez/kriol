// Filename: cmd/api/entries.go
package main
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)
//create entires hander for the POST /v1/entries endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Create a New Entry")
}
//create showentires hander for the GET /v1/entries/:id endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request){
	//use the "paramsfromcontext()" function to get the request context a slice
	params := httprouter.ParamsFromContext(r.Context())
	//GET the value of the "id" parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Display the entires
	fmt.Fprintf(w, "Show The Details For Emtry %d\n", id)
}