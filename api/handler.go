package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

type AddResp struct {
	A   int `json:"a"`
	B   int `json:"b"`
	Sum int `json:"sum"`
}

func add(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	a, _ := strconv.Atoi(q.Get("a"))
	b, _ := strconv.Atoi(q.Get("b"))
	resp := AddResp{A: a, B: b, Sum: a + b}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	// We'll use forwarded path (see host.json setting enableForwardingHttpRequest)
	mux.HandleFunc("/api/add", add)

	// When hosted behind Azure Functions, the host sets this env var
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		// fallback for local manual runs
		port = "8080"
	}
	log.Printf("Listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
