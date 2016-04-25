package main

import (
	"net/http"
	"strconv"
	"github.com/nu7hatch/gouuid"
)

func thepage(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("uuidCookie")
	myid, _ := uuid.NewV4()
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "cookie_count",
			Value: myid.String(),
		}
	}
	http.SetCookie(res, cookie)
}

func main() {
	http.HandleFunc("/", thepage)
	http.ListenAndServeTLS(":8090", "cert.pem", "key.pem", nil)
}
