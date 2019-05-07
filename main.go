package main

import (
    "fmt"
    "net/http"
    "html/template"
    //"github.com/gorilla/mux"
)


type ViewData struct{
    Title string
    Message string
}

func main(){

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        data := ViewData{
            Title : "Champ Boris Britva",
            Message : "Congratulations!!!",
        }
        tmpl, _ := template.ParseFiles("c:/users/admin/desktop/goworkspace/templates/userdie.html")
        tmpl.Execute(w, data)
        
    })



    fmt.Println("Server...")
    http.ListenAndServe(":8181", nil)
}