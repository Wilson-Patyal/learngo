package main

import (
	"fmt"
	"os"
	"testing"
	"io"
	"strings"
)
func Test_main(t *testing.T){
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	fmt.Println(r)
	fmt.Println(w)

	os.Stdout = w

	main()

	_ =  w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "288860"){
	    t.Error("wrong bank balance returned")
	}

}