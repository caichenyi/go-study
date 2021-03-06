package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
				<input type="text" name="user"/>
				<input type="text" name="password"/>
				<input type="submit" value="submit"/>
				</form></body></html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>Hello, world!</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.Form["user"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("password"))
	}
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("http listen failed.")
	}
}