package front

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) error {
	home := struct {
		Title string
	}{
		Title: "HEllo",
	}
	return render(w, r, "home.html", home, http.StatusOK)
}
