/* RZFeeser | Alta3 Research
   HTTP Server - Working with URL Path param 
   go run ~/httpsvr02.go
   In the new pane, try sending an HTTP request to our Go server. When we make the request, we're sending the data Alta3 to /

student@bchd:~$ curl localhost:8079/Alta3

The response should be, Hello, Alta3
   
   */

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", HelloServer)
    fmt.Println("Server started at port 8079")
    log.Fatal(http.ListenAndServe(":8079", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    // get the URL path value with r.URL.Path[1:] and build a message with the data
    fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}    

