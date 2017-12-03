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

	fmt.Fprintln(w, "<html>Hello there!<br><br>")

	if strings.Contains(r.RemoteAddr, ":") {
		fmt.Fprintf(w, "IP: %s<br>", strings.Split(r.RemoteAddr, ":")[0])
		fmt.Fprintf(w, "Port: %s<br>", strings.Split(r.RemoteAddr, ":")[1])
	} else {
		fmt.Fprintf(w, "RemoteAddr: %s<br>", r.RemoteAddr)
	}

	fmt.Fprintf(w, "<br>Headers:</br>")

	for h, c := range(r.Header) {
		for _, cc := range(c) {
			fmt.Fprintf(w, "%s : %s<br>", h, cc)
		}
	}

	fmt.Fprintln(w, "</html>")
}

