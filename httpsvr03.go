/* RZFeeser | Alta3 Research
   HTTP Server with r.URL.Query() 
   go run ~/httpsvr03.go
   In the new pane, try sending an HTTP request to our Go server.
   curl localhost:8055/?name=Batman
   Response should be, Hello, Batman
   */

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", handler)

    fmt.Printf("Starting server at port 8055\n")
    log.Fatal(http.ListenAndServe(":8055", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

    // accept the name parameter
    keys, ok := r.URL.Query()["name"]

    name := "guest"

    if ok {

        name = keys[0]
    }

    fmt.Fprintf(w, "Hello %s!\n", name)
}

