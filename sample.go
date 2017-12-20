package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := "localhost"
	portPtr := flag.Int("port", 3000, "")
	namePtr := flag.String("name", "default", "")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v %v\n", *namePtr, r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", host, *portPtr), nil))
}
