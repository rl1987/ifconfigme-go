package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ip", ipHandler)
	http.HandleFunc("/ua", uaHandler)
	http.HandleFunc("/port", portHandler)
	http.HandleFunc("/json", jsonHandler)
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
		for h := range hosts {
			fmt.Fprintf(w, "Host: %s<br>", h)
		}
	}

	fmt.Fprintf(w, "<br>Headers:</br>")

	for h, c := range r.Header {
		for _, cc := range c {
			fmt.Fprintf(w, "%s : %s<br>", h, cc)
		}
	}

	fmt.Fprintln(w, "</html>")
}

type UserInfo struct {
	IP        string
	Hostname  string
	Port      uint16
	UserAgent string
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	spew.Dump(r)

	var userInfo UserInfo

	if strings.Contains(r.RemoteAddr, ":") {
		substrings := strings.Split(r.RemoteAddr, ":")

		userInfo.IP = substrings[0]
		if port, err := strconv.Atoi(substrings[1]); err == nil {
			userInfo.Port = uint16(port)
		}
	}

	if hosts, err := net.LookupAddr(userInfo.IP); err != nil {
		userInfo.Hostname = hosts[0]
	}

	if userAgentHeader, ok := r.Header["User-Agent"]; ok {
		userInfo.UserAgent = userAgentHeader[0]
	}

	json, err := json.Marshal(&userInfo)
	if err != nil {
		spew.Dump(err)
		return
	}

	if _, err := io.WriteString(w, string(json)); err != nil {
		spew.Dump(err)
	}
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
