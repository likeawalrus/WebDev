package main
import (
		"fmt"
		"net/http"
		"strconv"
)
func counter(res http.ResponseWriter,req *http.Request){
	myCookie, err := req.Cookie("counter")
	if err == http.ErrNoCookie{
		myCookie = &http.Cookie{
			Name:"counter",
			Value:"0",
		}
	}
	numVisits, _ := strconv.Atoi(myCookie.Value)
	numVisits++
	myCookie.Value = strconv.Itoa(numVisits)
	http.SetCookie(res, myCookie)
	fmt.Fprint(res,myCcookie.Value)
}
func main(){
	http.HandleFunc("/favicon.ico",func(res http.ResponseWriter,req *http.Request){})
	http.HandleFunc("/",counter)
	http.ListenAndServe(":8080",nil)
}
