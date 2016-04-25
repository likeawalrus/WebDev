package main
import (
		"net/http"
		"fmt"
		"crypto/hmac"
		"crypto/sha256"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("raka02asd2j01djmk"))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func theFunc(res http.ResponseWriter,req *http.Request){
	myCookie, err := req.Cookie("Session")
	if err == http.ErrNoCookie {
		myCookie = &http.Cookie{
			Name:  "Session",
			Value: "",
			// Secure: true,
			HttpOnly: true,
		}
	}
	if req.FormValue("name") != "" {
		thing := req.FormValue("name")
		myCookie.Value = thing + "|" + getCode(thing) 
	}
	fmt.Fprint(res,`
		<!DOCTYPE html>
		<html>
			<head></head>
			<body>
				<form method="POST">
					`+myCookie.Value+`
					<input type="text" name="name">
					<input type="submit">
				</form>
			</body>
		</html>
	`)
}

func main(){
	http.HandleFunc("/",theFunc)
	http.ListenAndServe(":8080",nil)
}
