// Filename: internal/data/entries.go
package data

import "kriol.DrianePerez.net/internal/validator"

type EntriesData struct {
	Name string `json:"name"`
	String string `json:"string"`
	Translate string `json:"translate"`
	// Status string `json:"status,omitempty"`
	// Enviornment string `json:"enviornment,omitempty"`
	// Version string `json:"version,omitempty"`
}
func ValidateEntires(v *validator.Validator, entries *EntriesData)  {
	//use the check method to execute our validation checks
	v.Check(entries.Name != "", "name", "must be provided")
	v.Check(len(entries.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(entries.String != "", "string", "must be provided")
	v.Check(len(entries.String) <= 200, "string", "must not be more than 200 bytes long")

	v.Check(entries.Translate != "", "translate", "must be provided")
	v.Check(len(entries.Translate) <= 200, "translate", "must not be more than 200 bytes long")
	
}