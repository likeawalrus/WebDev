package main
import (
    "net/http"
    "fmt"
    "log"
)

func page(res http.ResponseWriter, req *http.Request){
    fmt.Fprint(res,"%v",req.URL.Path)
}

func U_path(res http.ResponseWriter, req *http.Request){
    fmt.Fprintf(res, "My name is me")
}

func main(){
    http.HandleFunc("/",page)
    http.HandleFunc("/name", U_path)
    http.ListenAndServe(":8080",nil)
    if err != nil{
        log.Fatal("Error: ",err)
    }
}
        