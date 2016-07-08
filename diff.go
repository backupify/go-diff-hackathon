package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	//read in list of changes here

	var operations [1]Change
	//operations[0] = Change{insert, 0, -1, }
	//operations[0] = Change{delete, 0, 4, ""}
	//operations[2] = Change{insert, 17, -1, "insert2"}
	operations[0] = Change{insert, 6, -1, "abc"}

	fmt.Println([]byte("abc"))

	//read in reference text here
	var referenceBytes, err = ioutil.ReadFile("tmpfile2.txt")
	fmt.Println(string(referenceBytes))
	check(err)

	fmt.Println(referenceBytes)

	for _, current := range operations {
		fmt.Println(current)
		if current.Operation == insert {
			referenceBytes = append(referenceBytes[:current.Start], append([]byte(current.Text), referenceBytes[current.Start:]...)...)
		} else if current.Operation == delete {
			referenceBytes = append(referenceBytes[:current.Start], referenceBytes[current.End+1:]...)
		}
	}

	fmt.Println(referenceBytes)

	//dump bytes into file
	var newFile, err1 = os.Create("new.bin")
	check(err1)

	defer newFile.Close()

	newFile.Write(referenceBytes)

}
