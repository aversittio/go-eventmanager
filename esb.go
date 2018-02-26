package main

import (
	"net/http"
	"encoding/json"
	"strconv"
	"os"
	"bytes"
	"esb/listener"
);


type myData struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type serverConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
}

var config serverConfig


func PrepareRequest(writer http.ResponseWriter, request *http.Request) {

	/*var contentType = request.Header.Get("Content-type");
	var smg myData;

	if contentType == "application/x-www-form-urlencoded" {
		var data = request.FormValue("data");
		json.Unmarshal([]byte(data), &smg)

	} else if contentType == "application/json" {
		json.NewDecoder(request.Body).Decode(&smg)
	}

	writer.Write([]byte(strconv.Itoa(smg.Id)));
	writer.Write([]byte("\n"));
	writer.Write([]byte(smg.Name));*/


	var listener listener.Listener
	listener.Init(1, "test", "http://127.0.0.1", 8480)
	message := myData{Id:123, Name:"Max"}

	msg, _ := json.Marshal(message)
	code := listener.Notify(msg)
	writer.Write([]byte(strconv.Itoa(code)));

}

func prepareHttp()  {
	http.HandleFunc("/", PrepareRequest);
}

func parseConfig()  {
	file, error := os.Open("esb-config.json");

	if error != nil {
		panic(error)
	}

	json.NewDecoder(file).Decode(&config)

}

func formAddress() (string, error)  {
	var addr bytes.Buffer
	var err error

	if config.Host == "" {
		panic("Config host not found!\n")
	}

	if strconv.Itoa(config.Port) == "" {
		panic("Config port not found!\n")
	}

	addr.WriteString(config.Host)
	addr.WriteString(":");
	addr.WriteString(strconv.Itoa(config.Port))

	var address = addr.String()
	return address, err
}

/**
Start server
 */
func main() {

	parseConfig()

	var address, err = formAddress();
	if err != nil {
		panic(err)
	}

	prepareHttp()

	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err);
	}

}
