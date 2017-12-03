package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ip", ipHandler)
	http.HandleFunc("/ua", uaHandler)
	http.HandleFunc("/port", portHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	spew.Dump(r)

	fmt.Fprintln(w, "<html>Hello there!<br><br>")

	var ip string

	if strings.Contains(r.RemoteAddr, ":") {
		ip = strings.Split(r.RemoteAddr, ":")[0]

		fmt.Fprintf(w, "IP: %s<br>", ip)
		fmt.Fprintf(w, "Port: %s<br>", strings.Split(r.RemoteAddr, ":")[1])
	} else {
		fmt.Fprintf(w, "RemoteAddr: %s<br>", r.RemoteAddr)
	}

	if hosts, err := net.LookupAddr(ip); err != nil {
		for h := range(hosts) {
			fmt.Fprintf(w, "Host: %s<br>", h)
		}
	}

	fmt.Fprintf(w, "<br>Headers:</br>")

	for h, c := range(r.Header) {
		for _, cc := range(c) {
			fmt.Fprintf(w, "%s : %s<br>", h, cc)
		}
	}

	fmt.Fprintln(w, "</html>")
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	spew.Dump(r)

	fmt.Fprintln(w, strings.Split(r.RemoteAddr, ":")[0])
}

func uaHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.Header["User-Agent"]) == 1 {
		fmt.Fprintln(w, r.Header["User-Agent"][0])
	}
}


func portHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.RemoteAddr, ":") {
		fmt.Fprintln(w, strings.Split(r.RemoteAddr, ":")[1])
	}
}

