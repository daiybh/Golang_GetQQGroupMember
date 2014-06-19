// CustomHttpPost project main.go
package main

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("\r\nFatal error \r\n", err.Error())
		os.Exit(0)
		return
	}
}
func main() {
	var cookiesStr string
	main_tcpPost(cookiesStr)
	httpPost_readFile_main()
}
