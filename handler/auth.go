package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/goetian/gogen/pkg/sb"
	"github.com/goetian/gogen/view/auth"
	"github.com/nedpals/supabase-go"
)

func HandleLogInUser(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)

}

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if !sb.ValidateEmail(credentials.Email) {
		return render(w, r, auth.LoginForm(credentials, auth.LogInErrors{InvalidInput: "Invalid E-Mail"}))
	}

	passWordInfo, isValidPW := sb.IsValidPassword(credentials.Password)
	if !isValidPW {
		return render(w, r, auth.LoginForm(credentials, auth.LogInErrors{
			InvalidInput: passWordInfo,
		}))
	}

	responds, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		slog.Error("Login Error SB-Client", "err", "err")
		fmt.Println(err.Error())
	}
	fmt.Println(credentials)
	fmt.Println(responds)

	cookie := &http.Cookie{
		Value:    responds.AccessToken,
		Name:     "at",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	// return render(w, r, auth.LoginForm(credentials, auth.LogInErrors{
	// 	InvalidInput: "Invalid Credentials",
	// }))
	return nil

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
