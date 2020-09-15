package api

import (
	"io/ioutil"
	"net/http"
	"web_app/service/jsonapi"
)

const fileName = "file_write_file.txt"

// WriteValueToFile : endpoint to write a value to a file
func WriteValueToFile(w http.ResponseWriter, req *http.Request) {
	writeValue := req.FormValue("write_value")
	// format data for write
	dataToWrite := []byte(writeValue)
	// attempt to write to file
	err := ioutil.WriteFile(fileName, dataToWrite, 0777)
	if err != nil {
		jsonapi.SendStringAPIResponse(w, false, "problem writing to file")
		return
	}
	jsonapi.SendStringAPIResponse(w, true, "value has been written to file")
}

// ReadValueFromFile : endpoint to read a value from a file
func ReadValueFromFile(w http.ResponseWriter, req *http.Request) {
	// attempt to read file contents
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		jsonapi.SendStringAPIResponse(w, false, "no file at location")
		return
	}
	jsonapi.SendStringAPIResponse(w, true, string(data))
}
