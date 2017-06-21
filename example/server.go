package main

import (
	"log"
	"net/http"

	reqid "../"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ReqId=%s\n", reqid.ReqIdFromRequest(r))
	w.Write([]byte(r.URL.Path))
}

func main() {
	handle := http.HandlerFunc(handler)

	http.Handle("/scrub/", reqid.ScrubRequestIdHeader(reqid.RandomId(handle)))
	http.Handle("/", reqid.RandomId(handle))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
