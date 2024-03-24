/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo örnekte text bir dosya karakter karakter okunup stdout'a bastırılmıştır

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"SampleGoLand/csd/err"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		err.ExitFailure("wrong number of arguments!...")
	}

	f, e := os.Open(os.Args[1])

	if e != nil {
		err.ExitFailureError("Open:", e)
	}

	defer func() { _ = f.Close() }()

	reader := bufio.NewReader(f)

	for {
		line, e := reader.ReadString('\n')

		if e == io.EOF {
			break
		}

		if e != nil {
			err.ExitFailureError("ReadString", e)
		}

		fmt.Print(line)
	}
}
