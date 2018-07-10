package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", test)

	fmt.Println("SERVING...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Tokopedia Tower"))
}
