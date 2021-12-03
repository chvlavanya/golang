/* RZFeeser | Alta3 Research
   HTTP File server  
   create a new program that utilizes both of these HTML documents(public/index.html,public/about.html)
   go run ~/fileserver01.go
   In the new pane, try sending an HTTP request to our Go server. This should return the index.html
   curl localhost:8044/
   In the new pane, try sending an HTTP request to our Go server to get the about.html
   The response should be the home page
   curl localhost:8044/about.html
   r prints the header
   1 application should run at time on 1 port
   Using log.Fatal - if we run same job second time it throws below error(print error on screen)
   Server started at port 8089
    2021/12/03 10:43:06 listen tcp :8089: bind: address already in use
    exit status 1
   */
 
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {

    // A file server is registered with Handle() and serves files from the public/ directory
    fileServer := http.FileServer(http.Dir("./public"))
    http.Handle("/", fileServer)

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello there!\n")
    })

    fmt.Printf("Starting server at port 8044\n")
    log.Fatal(http.ListenAndServe(":8044", nil))
}     

