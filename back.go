package backrepo

import (
	"io/ioutil"
	"log"
)

type MyOwnWriter struct {}

func (w * MyOwnWriter) Write(p []byte) (n int, err error) {
	ioutil.WriteFile("/tmp/log.log", p, 0644)

	return len(p), nil
}


func RunMe() {
	log.SetOutput(&MyOwnWriter{})
	log.Println("Am I nice log line?")
}
