package main

import (
    "strings"
    "os"
    "fmt"
    "io/ioutil"
)

type ChangeOperation int

const (
  insert ChangeOperation = iota //insert in front of given character
  delete ChangeOperation = iota //delete on inclusive range
)

type Change struct {
    Operation ChangeOperation
    Start int
    End int
    Text string
}

func main(){

  //read in list of changes here
  var operations [3]Change
  //operations[0] = Change{insert, 0, -1, "insert"}
  operations[0] = Change{delete, 0,4, ""}
  //operations[2] = Change{insert, 17, -1, "insert2"}

  fmt.Println(operations)

  //read in reference text here
  var referenceBytes, err = ioutil.ReadFile("reference.txt")
  if err != nil {
    fmt.Println(err)
    //log.Fatal(err)
  }
  var reference = string(referenceBytes)

  fmt.Printf("%s", reference)

  for _, current := range operations {
    if current.Operation == insert {
      reference = strings.Join([]string{reference[:current.Start], current.Text, reference[current.Start:]},"")
    } else if current.Operation == delete {
      reference = strings.Join([]string{reference[:current.Start], reference[current.End+1:]},"")
    }
  }

  fmt.Printf("%s", reference)

  //save new string
  var newFile, err1  = os.Create("new.txt")
  if err1 != nil {
    fmt.Println(err1)
    //log.Fatal(err1)
  }
  newFile.WriteString(reference)


}
