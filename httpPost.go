package main

import (
	"net/http"
)

func main_httpPost(cookies string) {
	clinet := &http.Client{}
	request, err := http.NewRequest("POST", "http://qun.qq.com/cgi-bin/qun_mgr/search_group_members", nil)

	checkErr(err)

	request.Header.Add("Host", "qun.qq.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0")
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Referer", "http://qun.qq.com/member.html")
	request.Header.Add("Cookie", cookies)
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Pragma", "no-cache")
	request.Header.Add("Cache-Control", "no-cache")

	//request.Body = ioutil.NopCloser()
	//request.Body.Read()
	response, err := clinet.Do(request)
	defer response.Body.Close()

}
