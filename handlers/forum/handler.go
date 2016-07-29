package forum

import (
)

func Create(w http.ResponseWriter, r * http.Request){
    w.Header().Set("Server","A Go WebServer")
    w.Header().Set("Content-Type", "text/html")

    common.Clear()

    hostname := r.URL.Query()["hostname"]
    //w.Write([]byte(hostname[0]))
    fmt.Printf("%s", hostname)
}
