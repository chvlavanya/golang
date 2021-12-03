/* Alta3 Research | RZFeeser
   Prometheus - exposing metrics in a go application requires about
     HTTP endpoint. 
     cd ~/prom/
     go mod init #prom
     go get -u github.com/prometheus/client_golang/prometheus
     go get -u github.com/prometheus/client_golang/prometheus/promauto
     go get -u github.com/prometheus/client_golang/prometheus/promhttp
     go run ~/prom/prometheus01.go
     new screen: curl localhost:2112/metrics
     curl localhost:2112/metrics
     one more time curl localhost:2112/metrics
     exit(stop ur server with ctl+c)
     The server will stop and display the following on the screen.
     ^Csignal: interrupt
     
     */
   
package main

import (
        "net/http"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2112", nil)
}

