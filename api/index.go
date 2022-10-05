package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.ayo.so/moby", 301)
}
