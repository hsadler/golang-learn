package api

import(
	"fmt"
	"net/http"
	"encoding/json"
)


type User struct {
	Age int
}


func GetUser(w http.ResponseWriter, req *http.Request) {
	user := &User{Age:1}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}
