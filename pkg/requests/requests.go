package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Post(jsonData []byte, addr string) (error, []byte) {
	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	return err, body
}

func Get(addr string) (error, []byte) {
	resp, err := http.Get(addr)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	return err, body
}
