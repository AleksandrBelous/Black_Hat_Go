package main

import (
	"fmt"
	"log"
	"os"
)

// FooReader определяет io.Reader для чтения из stdin
type FooReader struct{}

// Read считывает данные из stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// FooWriter определяет io.Writer для записи в stdout
type FooWriter struct{}

// Write записывает данные в stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out < ")
	return os.Stdout.Write(b)
}

func main() {
	// Создаём экземпляры Reader, Writer
	var (
		reader FooReader
		writer FooWriter
	)

	// Буфер для хранения ввода/вывода
	input := make([]byte, 4096)

	// Используем reader для чтения ввода
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Используем writer для записи вывода
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
