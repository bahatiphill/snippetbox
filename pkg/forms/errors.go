package forms


//error type which will hold the validation errors
type errors map[string][]string

// Add method to ADD error message for a given field map
func (e errors) Add(field, message string){
	e[field] = append(e[field], message)
}

//Get the first error message for a given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}