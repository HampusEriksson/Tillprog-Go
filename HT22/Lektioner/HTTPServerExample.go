// https://yourbasic.org/golang/http-server-example/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// The call to http.HandleFunc tells the net.http package to handle all requests to the web root with the HelloServer function.
	http.HandleFunc("/", HelloServer)

	// The call to http.ListenAndServe tells the server to listen on the TCP network address :8080. This function blocks until the program is terminated.
	http.ListenAndServe(":8080", nil)
}

// Writing to an http.ResponseWriter sends data to the HTTP client.
// Writing to an http.ResponseWriter sends data to the HTTP client.An http.Request is a data structure that represents a client HTTP request.
func HelloServer(w http.ResponseWriter, r *http.Request) {

	//r.URL.Path is the path component of the requested URL. In this case, "/world" is the path component of "http://localhost:8080/world".

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
