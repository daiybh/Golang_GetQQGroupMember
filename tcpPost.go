package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main_tcpPost(cookies string) {
	//因为post方法属于HTTP协议，HTTP协议以TCP为基础，所以先建立一个
	//TCP连接，通过这个TCP连接来发送我们的post请求
	conn, err := net.Dial("tcp", "qun.qq.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	//构造post请求
	var post string
	/*post += "POST /postpage HTTP/1.1\r\n"
		post += "Content-Type: application/x-www-form-urlencoded\r\n"
		post += "Content-Length: 37\r\n"
		post += "Connection: keep-alive\r\n"
		post += "Accept-Language: zh-CN,zh;q=0.8,en;q=0.6\r\n"
		post += "\r\n"
		post += "key=this is key&value=this is value\r\n"
	/**/
	post += "POST http://qun.qq.com/cgi-bin/qun_mgr/search_group_members HTTP/1.1\r\n"
	post += "Host: qun.qq.com\r\n"
	post += "User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0\r\n"
	post += "Accept: application/json, text/javascript, */*; q=0.01\r\n"
	post += "Accept-Language: en-US,en;q=0.5\r\n"
	post += "Accept-Encoding: gzip, deflate\r\n"
	post += "Content-Type: application/x-www-form-urlencoded; charset=UTF-8\r\n"
	post += "X-Requested-With: XMLHttpRequest\r\n"
	post += "Referer: http://qun.qq.com/member.html\r\n"
	//post += "Content-Length: 46\r\n"
	post += "Content-Length: 44\r\n"
	post += "Cookie:" + cookies + "\r\n"
	post += "Connection: keep-alive\r\n"
	post += "Pragma: no-cache\r\n"
	post += "Cache-Control: no-cache\r\n"
	post += "\r\n"
	post += "gc=9978616&st=0&end=19&sort=0&bkn=1121614792\r\n"
	if _, err := conn.Write([]byte(post)); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("post send success.")

	var buff = make([]byte, 102400)
	iCount := 0
	for {

		fmt.Printf("Read start[%d].", iCount)
		n, err := conn.Read(buff)

		fmt.Printf("Read end[%d].", iCount)
		iCount += 1
		if err == io.EOF {
			fmt.Printf("remote connect shudown\n")
			return
		}
		if err != nil {
			fmt.Printf("read error:%s\n", err.Error())
			return
		}

		if n > 0 {
			fmt.Printf("Msg:\n%s", string(buff[:n]))
			userFile := "r.txt"
			fout, err := os.Create(userFile)
			defer fout.Close()
			if err != nil {
				return
			}
			//fout.Write(buff)
			fout.WriteString(string(buff[:n]))
		}
	}

	fmt.Println("Exit.")
}
