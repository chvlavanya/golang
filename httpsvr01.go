/* RZFeeser | Alata3 Research
   HTTP Server   
   go run ~/httpsvr01.go
   in other window curl localhost:8089 or in browser localhost:8089
   response should be, Hello, there */

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    // The HandleFunc registers the handler function for the given URL pattern
    http.HandleFunc("/", HelloHandler)
    fmt.Println("Server started at port 8089")
    log.Fatal(http.ListenAndServe(":8089", nil))      // listens on TCP network address for incoming HTTP requests
}

// describes how to respond to the HTTP request
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, there\n")
}

