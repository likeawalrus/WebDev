package main
import (
    "net/http"
    "fmt"
    "log"
)

func page(res http.ResponseWriter, req *http.Request){
    fmt.Fprint(res,"%v",req.URL.Path)
}

func main(){
    http.HandleFunc("/",page)
    http.ListenAndServe(":8080",nil)
    if err != nil{
        log.Fatal("Error: ",err)
    }
}
        