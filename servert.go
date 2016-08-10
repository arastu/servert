package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU and Hostname: %s", runtime.GOOS, runtime.GOARCH, hostName)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
