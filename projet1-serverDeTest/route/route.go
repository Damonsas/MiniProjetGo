package route

import "net/http"

func init() {
	http.HandleFunc("/asset/scss/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
}
