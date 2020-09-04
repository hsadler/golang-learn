package main

import (
	"fmt"
	"net/http"
	"web_app/api"
)


func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}


func main() {

	http.HandleFunc("/headers", headers)

	http.HandleFunc("/hello", api.Hello)
	http.HandleFunc("/get_user", api.GetUser)

	http.ListenAndServe(":8090", nil)

}