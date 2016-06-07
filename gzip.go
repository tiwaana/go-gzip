/*
** parse the gzip header and dump some information
*/

package main

import (
//    "bufio"
	"fmt"
//   "io"
	"io/ioutil"
	"os"
)

type gzip_header_fixed struct {
	gzip_id1 byte
	gzip_id2 byte
	gzip_cm byte
	gzip_flg byte
	gzip_mtime[4] byte
	gzip_xfl byte
	gzip_os byte

}

/* from gobyexample.com */
func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	if len(os.Args) == 1 {
		fmt.Printf("usage: %s\n", os.Args[0]);
	}

	data, err := ioutil.ReadFile("tests/fail.txt.gz")
	check(err)
	fmt.Print(string(data))

	f, err := os.Open("tests/fail.txt.gz")
	check(err)

	f.Close()
}
