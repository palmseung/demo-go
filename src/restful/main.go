package main

import (
	"net/http"
	"restful/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
