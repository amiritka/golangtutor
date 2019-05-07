package main

import (
    "fmt"
    "net/http"
)


func main(){
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", http.FileServer(http.Dir("c:/users/admin/desktop/goworkspace/static")))
}