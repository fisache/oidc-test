package main

import (
          "os"
          "fmt"
          "net/http"
          "golang.org/x/oauth2"
          "github.com/dgrijalva/jwt-go"
          "encoding/json"
       )


const htmlIndex = `<html><body>
<a href="/KeycloakLogin">Log in with Keycloak</a>
</body></html>
`
var keycloakUrl = os.Getenv("keycloakUrl")
var realm = os.Getenv("realm")

var (
    keycloakOauthConfig = &oauth2.Config{
        RedirectURL:    "http://localhost:3000/KeycloakCallback",
        ClientID:     os.Getenv("clientId"),
        ClientSecret: os.Getenv("clientSecret"),
        Scopes:       []string{keycloakUrl+"/auth/realms/"+realm+"/protocol/openid-connect/userinfo",},
        Endpoint:     oauth2.Endpoint{
                AuthURL:  keycloakUrl+"/auth/realms/"+realm+"/protocol/openid-connect/auth",
                TokenURL: keycloakUrl+"/auth/realms/"+realm+"/protocol/openid-connect/token",
        },
    }
// Some random string, random for each request
    oauthStateString = "random"
)

func main() {
    http.HandleFunc("/", handleMain)
    http.HandleFunc("/KeycloakLogin", handleKeycloakLogin)
    http.HandleFunc("/KeycloakCallback", handleKeycloakCallback)
    fmt.Println(http.ListenAndServe(":3000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, htmlIndex)
}

func handleKeycloakLogin(w http.ResponseWriter, r *http.Request) {
    url := keycloakOauthConfig.AuthCodeURL(oauthStateString)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleKeycloakCallback(w http.ResponseWriter, r *http.Request) {
    state := r.FormValue("state")
    if state != oauthStateString {
        fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    code := r.FormValue("code")
    token, err := keycloakOauthConfig.Exchange(oauth2.NoContext, code)
    if err != nil {
        fmt.Println("Code exchange failed with '%s'\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    decode, err := jwt.Parse(token.AccessToken, func(token *jwt.Token) (interface{}, error) {
            return []byte("AllYourBase"), nil
    })

    contents, err := json.Marshal(decode)

    fmt.Fprintf(w, "Content: %s\n", contents)
}
