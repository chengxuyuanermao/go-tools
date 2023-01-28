package goth

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"html/template"
	"log"
	"net/http"
)

var (
	ptTemplate *template.Template
)

func init() {
	githubProvider := github.New("63095142e73555d256b0", "c64d6490fe6f0abcdbd244855eb611f3b9c2c8d9", "http://localhost:8080/auth/github/callback")
	goth.UseProviders(githubProvider)
}

func Use() {
	ptTemplate = template.Must(template.New("").ParseGlob("./goth/tpls/*.tpl"))

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login/github", LoginHandler)
	r.HandleFunc("/logout/github", LogoutHandler)
	r.HandleFunc("/auth/github", AuthHandler)
	r.HandleFunc("/auth/github/callback", CallbackHandler)

	log.Println("listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Redirect(w, r, "/login/github", http.StatusTemporaryRedirect)
		return
	}
	ptTemplate.ExecuteTemplate(w, "home.tpl", user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := ptTemplate.ExecuteTemplate(w, "login.tpl", nil)
	if err != nil {
		fmt.Println("ddd")
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
}

// 点击登录，由AuthHandler处理请求：
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

// GitHub 验证完成后，浏览器会重定向到/auth/github/callback处理：
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	ptTemplate.ExecuteTemplate(w, "home.tpl", user)
}
