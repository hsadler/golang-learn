package api

import(
	"strconv"
	"net/http"
	"web_app/service/json_api"
)


func Add(w http.ResponseWriter, req *http.Request) {
	n1 := req.URL.Query().Get("n1")
	n1int, _ := strconv.Atoi(n1)
	n2 := req.URL.Query().Get("n2")
	n2int, _ := strconv.Atoi(n2)
	addResult := n1int + n2int
	success := true
	payload := strconv.Itoa(addResult)
	json_api.SendRes(w, success, payload)
}
