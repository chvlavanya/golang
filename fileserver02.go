/* Alta3 Research | RZFeeser
   Simple File Server - simple file server with Go
   go run ~/fileserver02.go
   job runs form.html which has post request method.
   localhost:2224
   u see the html page askiing name and occupation with sumbit button. once u enter the 2 fields and sumbit - u can see the <name> is a <occupation> message in browser
   */

package main

import (
    "fmt"
    "log"
    "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    // GET request - send a web page with a form
    // POST request - process the data from the form attached on the POST
    switch r.Method {
    case "GET":

        http.ServeFile(w, r, "form.html")
    case "POST":

        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

        name := r.FormValue("name")
        occupation := r.FormValue("occupation")

        fmt.Fprintf(w, "%s is a %s\n", name, occupation)

    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func main() {

    http.HandleFunc("/", process)

    fmt.Printf("Starting server at port 2224\n")
    log.Fatal(http.ListenAndServe(":2224", nil))
}

