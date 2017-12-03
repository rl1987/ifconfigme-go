package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	spew.Dump(r)

	fmt.Fprintln(w, "<html>Hello there!<br>")
	fmt.Fprintf(w, "IP: %s<br>", strings.Split(r.RemoteAddr, ":")[0])
	fmt.Fprintln(w, "</html>")
}

