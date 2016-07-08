package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ChangeOperation int

const (
	insert ChangeOperation = iota //insert in front of given character
	delete ChangeOperation = iota //delete on inclusive range
)

type Change struct {
	Operation ChangeOperation
	Start     int
	End       int
	Text      string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var operations []Change
	//read in list of changes here
	diffFile, err := os.Open("diff.txt")
	check(err)

	defer diffFile.Close()

	//parse diff file into Go Change structs
	//NOTE: scanner doesn't handle lines over ~65000 chars each
	//depending on how we store diffs (assuming hexdecimal characters) this means we can store a max
	//of about 30000 bytes 30k in a single diff line
	scanner = bufio.NewScanner(diffFile)

	for scanner.Scan() {
		row := scanner.Text()
		splitRow = strings.Split(row, "|")
		rowChange := Change{splitRow[0], splitRow[1], splitRow[2], splitRow[3]}
		operations = append(operations, rowChange)
	}

	//operations[0] = Change{insert, 0, -1, }
	//operations[0] = Change{delete, 0, 4, ""}
	//operations[2] = Change{insert, 17, -1, "insert2"}
	operations = append(operations, Change{insert, 0, -1, "abc"})
	operations = append(operations, Change{delete, 3, 6, "abc"})

	fmt.Println([]byte("abc"))

	//read in reference text here
	refFile, err := os.Open("tmpfile2.txt")
	check(err)

	defer refFile.Close()

	referenceBytes, err := ioutil.ReadFile("tmpfile2.txt")
	fmt.Println("Before: ", (referenceBytes))
	check(err)

	var finishedBytes []byte

	for index, current := range referenceBytes {

		fmt.Println("Index: ", index)
		fmt.Println("-----")
		fmt.Println(finishedBytes)
		fmt.Println("-----")

		if len(operations) == 0 {
			finishedBytes = append(finishedBytes, referenceBytes[index:]...)
			break
		}

		if operations[0].Start <= index {

			fmt.Println("Start <= index")

			if operations[0].Operation == insert {

				fmt.Println("Inserting: ", []byte(operations[0].Text))

				finishedBytes = append(finishedBytes, []byte(operations[0].Text)...)
				operations = operations[1:]

			} else if operations[0].Operation == delete {
				fmt.Println("in delete")

				if operations[0].End > index {
					fmt.Println("continued")
					continue
				} else {
					fmt.Println("ended delete")
					operations = operations[1:]
				}

			}

		}

		finishedBytes = append(finishedBytes, current)

	}

	/*
		for _, current := range operations {
			fmt.Println(current)
			if current.Operation == insert {
				referenceBytes = append(referenceBytes[:current.Start], append([]byte(current.Text), referenceBytes[current.Start:]...)...)
			} else if current.Operation == delete {
				referenceBytes = append(referenceBytes[:current.Start], referenceBytes[current.End:]...)
			}
		}
	*/
	fmt.Println("After:  ", finishedBytes)

	//dump bytes into file
	newFile, err := os.Create("new.bin")
	check(err)

	defer newFile.Close()

	newFile.Write(finishedBytes)

}
