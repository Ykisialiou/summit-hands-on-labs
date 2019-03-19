package main

import (
	"fmt"
	"net/http"
	"os"
)

func CFSummit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Cloud Foundry Summit @ Bazel!");
}

func main() {
	http.HandleFunc("/", CFSummit);

	var port string;
	port = os.Getenv("PORT");
	if len(port) == 0 {
		port = "8080"
	}

	http.ListenAndServe(":" + port, nil);
}