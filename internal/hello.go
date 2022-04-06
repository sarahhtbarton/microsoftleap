package internal

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL)
	fmt.Println("URL Path:", r.URL.Path)
	fmt.Println("Header:", r.Header)
	fmt.Println("Body:", r.Body)
}
