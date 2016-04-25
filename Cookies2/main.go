package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/nu7hatch/gouuid"
)

func name(res http.ResponseWriter, req *http.Request) {
	myCookie, err := req.Cookie("session")
	if err != nil {
		fmt.Fprint(res,"This does whatever")
		id,_ := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res,myCookie)
	}
	fmt.Fprint(res,myCookie)
}

func main() {
	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
}
