package main

type Reader interface {
	read()
}

type Writer interface {
	write()
}

type ReaderWriter interface {
	Reader
	Writer
}

func main() {

}
