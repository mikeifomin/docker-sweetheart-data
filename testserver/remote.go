package testserver

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type RemoteAPI struct {
	Port int
	Host string
}

func NewRemoteAPI(port int, host string) (result *RemoteAPI, err error) {
	result = &RemoteAPI{port, host}
	err = result.Check()
	return
}
func (r *RemoteAPI) Call(path string, params interface{}, out interface{}) error {
	url := "http://" + r.Host + ":" + strconv.Itoa(r.Port) + path
	var body io.ReadWriter
	errE := json.NewEncoder(body).Encode(params)
	if errE != nil {
		return errE
	}
	resp, errR := http.Post(url, "application/json", body)
	if errR != nil {
		return errR
	}
	defer resp.Body.Close()
	errP := json.NewDecoder(resp.Body).Decode(out)
	if errP != nil {
		return errP
	}
	return nil
}
func (r *RemoteAPI) Check() error {

	var out RespHealth
	err := r.Call("/health", ParamsHealth{}, &out)
	if err != nil {
		return err
	} else if out.Status != "ok" {
		return errors.New("wrong status")
	}
	return nil
}
