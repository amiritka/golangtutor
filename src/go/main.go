package main

import (
    "fmt"
    "net/http"
)

//$GOPATH = 


func main(){
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", http.FileServer(http.Dir("static")))
}