package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func datehandler(w http.ResponseWriter, r *http.Request) {

	var b bytes.Buffer
	flusher := w.(http.Flusher)

    // declare SSE MIME type
	w.Header().Set("Content-type", "text/event-stream")

	for {
		b.Write([]byte("data: "))   // SSE message header

        // run and take all ouput from the "date" command
		cmd := exec.Command("date")
		cmd.Stdout = &b
		cmd.Stderr = &b
		cmd.Run()

		b.Write([]byte("\n\n"))     // SSE message coda

		b.WriteTo(w)
		flusher.Flush()

		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))     // for index.html
	http.HandleFunc("/date", datehandler)   // for SSE stream

	fmt.Println("Starting server on localhost:8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
