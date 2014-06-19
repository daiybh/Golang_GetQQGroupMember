package main

import (
	"bufio"
	"compress/gzip"

	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UnZipData(Body io.ReadCloser) {
	/*
		body, err := ioutil.ReadAll(Body)
		checkErr(err)

		fmt.Printf("%v", body)
		//	r, err := gzip.NewReader(&body)
		/**/
	r, err := gzip.NewReader(Body)
	//
	defer r.Close()
	checkErr(err)
	unDatas, _ := ioutil.ReadAll(r)
	ParseJson(unDatas)
}
func json_unMarshal(r *bufio.Reader) {
	//data, err := ioutil.ReadAll(r)

	//checkErr(err)
	//err = json.NewDecoder(r).Decode(value)
}
func CovertToResponse(r *bufio.Reader) {

	resp, err := http.ReadResponse(r, nil)
	checkErr(err)
	fmt.Printf("\r\nConver:-->%v\r\n", resp)
	UnZipData(resp.Body)

}
func httpPost_readFile_main() {
	fread, err := os.Open("r.txt")
	defer fread.Close()
	checkErr(err)
	/*
		buf := make([]byte, 1024)
		for {
			n, _ := fread.Read(buf)
			if 0 == n {
				break
			}
			//os.Stdout.Write(buf[:n])

			break
		}
		/**/
	br := bufio.NewReader(fread)
	/*nCount := 0
	for {
		line, err := br.ReadString(`\n`)
		if err == io.EOF {
			break
		}
		fmt.Printf("%d-->%v", nCount, line)
		nCount += 1
	}/**/
	//tp := textproto.NewReader(br)
	CovertToResponse(br)
}
