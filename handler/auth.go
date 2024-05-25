package handler

import (
	"net/http"

	"github.com/goetian/gogen/view/auth"
)

func HandleLogInUser(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)

}

// TO-DO i want to toggle passwort with HTMX
// func TogglePassword(w http.ResponseWriter, r *http.Request) error {
// 	// Get the current type of the password input field
// 	currentType := r.URL.Query().Get("type")

// 	// Toggle the type between 'password' and 'text'
// 	newType := "text"
// 	if currentType == "text" {
// 		newType = "password"
// 	}
// 	fmt.Println("Hier in Auth Handler")
// 	// Respond with the new input field
// 	response := `<input type="` + newType + `" id="password" class="grow" value="" placeholder="Password"/>`
// 	w.Write([]byte(response))
// 	return nil
// }
