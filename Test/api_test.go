package Test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func CallAPI(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func TestAPI(t *testing.T) {
	res, err := CallAPI("localhost:8888", "/v1/prods", "GET")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
