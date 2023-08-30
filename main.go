package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	f, err := os.Open("sample.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// read bytes from file
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	m, b := b[:2], b[2:]
	fmt.Printf("%x\n", m)
	// print 16bytes per line
	for len(b) > 0 {
		var m, l, sm []byte

		fmt.Printf("========================================\n")
		m, b = b[:2], b[2:]
		fmt.Printf("marker=`%x`\n", m)

		l, b = b[:2], b[2:]

		length := int(l[0])<<8 + int(l[1])
		length = length - 2
		fmt.Printf("len=`%d`\n", length-2)

		sm, b = b[:length], b[length:]
		fmt.Printf("data=`%x`\n", sm)

		if reflect.DeepEqual(m, []byte{0xff, 0xda}) {
			fmt.Printf("========================================\n")
			fmt.Printf("start of scan\n")
			fmt.Printf("image=`%x`\n", b[:len(b)-2])
			fmt.Printf("========================================\n")
			fmt.Printf("end of scan: %x\n", b[len(b)-2:])
			break
		}
	}
}
