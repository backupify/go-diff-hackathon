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
	var operations [3]Change
	//operations[0] = Change{insert, 0, -1, "insert"}
	operations[0] = Change{delete, 0, 4, ""}
	//operations[2] = Change{insert, 17, -1, "insert2"}

	fmt.Println(operations)

	//read in reference text here
	var referenceBytes, err = ioutil.ReadFile("reference.txt")
	check(err)

	for _, current := range operations {
		if current.Operation == insert {
			referenceBytes = append(append(referenceBytes[:current.Start], []byte(current.Text)...), referenceBytes[current.Start:]...)
		} else if current.Operation == delete {
			referenceBytes = append(referenceBytes[:current.Start], referenceBytes[current.End+1:]...)
		}
	}

	//dump bytes into file
	var newFile, err1 = os.Create("new.txt")
	check(err1)

	defer newFile.Close()

	newFile.Write(referenceBytes)

}
