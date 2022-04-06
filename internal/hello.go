package internal

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "URL:", r.URL)
	fmt.Fprintln(w, "URL Path:", r.URL.Path)
	fmt.Fprintln(w, "GET Parameters:", r.URL.Query().Get("token"))
	fmt.Fprintln(w, "Header:", r.Header)
	fmt.Fprintln(w, "Body:", r.Body)
}
