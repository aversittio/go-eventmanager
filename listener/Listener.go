package listener

import (
	"net/http"
	"io"
	"bytes"
	"strconv"
)

type Listener struct {
	Id int
	Name string
	Host string
	Port int
}

func (l *Listener) Init(id int, name string, host string, port int) {
	l.Id = id
	l.Name = name
	l.Host = host
	if (port != 0) {
		l.Port = port
	} else {
		l.Port= 8480
	}
}

func (l *Listener) Notify(message []byte) (int) {

	var vals io.Reader = bytes.NewReader(message)
	addr := l.Host + ":" + strconv.Itoa(l.Port)
	response, err := http.Post(addr, "application/json", vals)
	if err != nil {
		panic(err)
	}

	return response.StatusCode

}