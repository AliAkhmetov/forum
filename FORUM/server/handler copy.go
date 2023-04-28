package server

import (
	"fmt"
	"net/http"
)

func (h *Handler) postssss(w http.ResponseWriter, r *http.Request) {
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
