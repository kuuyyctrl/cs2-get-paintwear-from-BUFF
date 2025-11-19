package creatclient

import (
	"net/http"
)

func Cclient(ul string) *http.Response {
	req, _ := http.NewRequest("GET", ul, nil)
	//req.Header.Set("set-cookie", "")
	//req.Header.Add("set-cookie", "")
	//req.Header.Set("cookie", "")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0")
	req.Header.Set("connection", "keep-alive")
	rsp, _ := http.DefaultClient.Do(req)
	return rsp
}
