package example

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func handlerV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q, URL : %v\n", r.URL.Path, r.URL)
}

func StudWebServerV1() {
	http.HandleFunc("/", handlerV1)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//V2
var m sync.Mutex
var count int

func handlerV2(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	count++
	m.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}

func counter(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	m.Unlock()
}
func StudWebServerV2() {
	http.HandleFunc("/", handlerV2)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//V3
func handlerV3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}
func StudWebServerV3() {
	http.HandleFunc("/", handlerV3)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//V4

func StudWebServerV4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Call(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
