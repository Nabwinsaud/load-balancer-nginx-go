package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintln(w, "Hello from server :%", hostname)
}
func main() {
	http.HandleFunc("/", handler)
	log.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)

}
