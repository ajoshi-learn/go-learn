package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	// n, err = io.Copy(io.Writer, io.Reader)

	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)


	// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

	// formatted:
	// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	//%t: boolean (true или false)
	//
	//%b: 01110100
	//
	//%c: numeric code
	//
	//%d: digital
	//
	//%o: 0x36532
	//
	//%q: single quotes
	//
	//%x: 0xa43b
	//
	//%X: 0xA43B
	//
	//%U: Unicode U+1234
	//
	//%e: -1.234456e+78
	//
	//%E: -1.234456E+78
	//
	//%f: float 123.456
	//
	//%F: %f
	//
	//%s: string
	//
	//%p: pointer in hex

	// Float numbers:
	// %f: default
	//
	// %9f: 9, default
	//
	// %.2f: default,2
	//
	// %9.2f: 9,2
	//
	// %9.f: 9,0

	// to the console:
	// Println(), Print() и Printf()

	// read data:
	// func Fscan(r io.Reader, a ...interface{}) (n int, err error)
	// func Fscanln(r io.Reader, a ...interface{}) (n int, err error)

	// formatted:
	// func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

	// read from console:
	// func Scan(a ...interface{}) (n int, err error)
	// func Scanf(format string, a ...interface{}) (n int, err error)
	// func Scanln(a ...interface{}) (n int, err error)

}
