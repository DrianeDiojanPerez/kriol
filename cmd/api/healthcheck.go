// Filename: cmd/api/Healthcheck.go
package main

import (
	"fmt"
	"net/http"
)

func (app *application) healcheckHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "status: avaialble")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n",version)

}