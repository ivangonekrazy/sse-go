Quick PoC for SSE from Go
=========================
Quick example of how to use Go and its built-in
libraries to generate a Server-sent Event stream.

Also demonstrates the HTML5 EventSource Javascript API
reading from such a stream.

1. Start up `sse.go` by running `go run sse.go`
2. Point your browser to http://localhost:8080
3. Watch the date update from the server's `date`
   command every second.

References
----------
* [Server-sent Events](http://dev.w3.org/html5/eventsource/)
* [Stream Updates with Server-Sent Events](http://www.html5rocks.com/en/tutorials/eventsource/basics/)
