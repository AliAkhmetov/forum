package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"forum/repository"
	"forum/service"
)

var (
	tpl        *template.Template
	artistsUrl string = "https://groupietrackers.herokuapp.com/api/artists"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

var client *http.Client

type Handler struct {
	repos *repository.Repository
}

func NewHandler(repos *repository.Repository) *Handler {
	return &Handler{repos: repos}
}

func Server(h *Handler) {
	http.HandleFunc("/", h.homePage)
	http.HandleFunc("/registration", h.registration)
	http.HandleFunc("/login", h.login)
	http.HandleFunc("/posts", h.posts)
	http.HandleFunc("/postPage", h.posts)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	log.Println("Starting a web server on http://localhost:8081/ ")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		Errors(w, http.StatusNotFound, "NOT FOUND")
		return
	}
	if r.Method == "GET" {
		if err := tpl.ExecuteTemplate(w, "registration.html", nil); err != nil {
			Errors(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}
	if r.Method == "POST" {
		userName := r.FormValue("registerUsername")
		email := r.FormValue("registerEmail")
		password := r.FormValue("registerPassword")

		user, code, err := service.Registration(h.repos, userName, email, password)
		if err != nil || code != http.StatusCreated {
			if code == http.StatusInternalServerError {
				Errors(w, http.StatusInternalServerError, err.Error())
				return
			}
			w.WriteHeader(code)
			if err := tpl.ExecuteTemplate(w, "registration.html", nil); err != nil {
				Errors(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		if err := tpl.ExecuteTemplate(w, "registered.html", user); err != nil {
			Errors(w, http.StatusInternalServerError, "Internal Error")
			return
		}
	}
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		Errors(w, http.StatusNotFound, "NOT FOUND")
		return
	}
	if r.Method == "GET" {
		if err := tpl.ExecuteTemplate(w, "login.html", nil); err != nil {
			Errors(w, http.StatusInternalServerError, "Internal Error")
			return
		}
	}
	if r.Method == "POST" {
		email := r.FormValue("loginEmail")
		password := r.FormValue("loginPassword")
		passHash, err := service.GeneratePassHash(password)
		if err != nil {
			fmt.Println(err.Error())
			Errors(w, http.StatusInternalServerError, "Password Hash Error")
		}
		user, err := h.repos.Authorization.GetUser(email, passHash)
		if err != nil {
			fmt.Print(err.Error())
		}

		fmt.Println("rabotaet \n ", user)
		fmt.Println(email, password)
		if err := tpl.ExecuteTemplate(w, "login.html", nil); err != nil {
			Errors(w, http.StatusInternalServerError, "Internal Error")
			return
		}
	}
}

func (h *Handler) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, "NOT FOUND")
		return
	}
	if r.Method != "GET" {
		Errors(w, http.StatusMethodNotAllowed, "NOT GET METHOD")
		return
	}

	if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		Errors(w, http.StatusInternalServerError, "Internal Error")
		return
	}
}

func (h *Handler) posts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts" {
		Errors(w, http.StatusNotFound, "NOT FOUND")
		return
	}
	if r.Method == "GET" {
		if err := tpl.ExecuteTemplate(w, "create_post.html", nil); err != nil {
			Errors(w, http.StatusInternalServerError, "Internal Error")
			return
		}
	}
	if r.Method == "POST" {
		postName := r.FormValue("postName")
		postText := r.FormValue("postText")

		fmt.Println(postName, postText)
		if err := tpl.ExecuteTemplate(w, "create_post.html", nil); err != nil {
			Errors(w, http.StatusInternalServerError, "Internal Error")
			return
		}
	}
}

func Errors(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	if code == 500 {
		return
	}
	d := struct {
		ErrorCode int
		ErrorText string
	}{
		ErrorCode: code,
		ErrorText: message,
	}

	if err := tpl.ExecuteTemplate(w, "error.html", d); err != nil {
		Errors(w, http.StatusInternalServerError, "Internal Error")
		return
	}
}
