package main

import (
    //"io"
    "io/ioutil"
    "fmt"
    //"math"
)

var SIZE int = 10

type Operation struct {
    Type int
    StartPos int
    EndPos int
    Chars string
}

func powerfulPow(base uint64, pow int) uint64 {
    if pow == 0 {
        return 1
    } else  {
        return base * powerfulPow(base, pow - 1)
    }
}

func computeHash(h uint64, input []byte, oldChar byte, newChar byte) uint64 {
    PRIME_BASE := 257
    //uOldChar := uint64(oldChar)
    //uNewChar := uint64(newChar)

    size := SIZE

    if h != 0   { // rolling case
        h = (h - uint64(oldChar) * powerfulPow(uint64(PRIME_BASE), size - 1) )// % PRIME_MOD

        h = (h * uint64(PRIME_BASE))// % PRIME_MOD

        h = (h + uint64(newChar))// % PRIME_MOD
    } else    { // initial case
        for j := 0; j < size; j++ {
            h += uint64(uint64(input[j]) * powerfulPow(uint64(PRIME_BASE), (size - 1) - j))
            //h = h % PRIME_MOD
        }
    }

    return h
}

func computeRollingHash(input []byte, hmap map[uint64][]int) {
    PRIME_BASE := 257
    //PRIME_MOD := 1000000007

    length := len(input)

    size := SIZE

    var h uint64 = 0

    var oldChar, newChar uint64

    for i := 0; i < length; i++ {
        if i != 0   { // rolling case
            h = (h - uint64(oldChar) * powerfulPow(uint64(PRIME_BASE), size - 1) )// % PRIME_MOD

            h = (h * uint64(PRIME_BASE))// % PRIME_MOD

            h = (h + uint64(newChar))// % PRIME_MOD
        } else    { // initial case
            for j := 0; j < size; j++ {
                h += uint64(uint64(input[j]) * powerfulPow(uint64(PRIME_BASE), (size - 1) - j))
                //h = h % PRIME_MOD
                fmt.Printf("%v", string(input[j]))
            }
            fmt.Printf("\n")

        }
        fmt.Printf("%v %v\n", string(oldChar), string(newChar))
        oldChar = uint64(input[i])
        if i + size < length    {
            newChar = uint64(input[i + (size)])
            fmt.Printf("%d %d", i, i + size)
        } else  {
            newChar = 0
        }

        if hmap[h] == nil   {
            hmap[h] = make([]int, 0)
        }
        hmap[h] = append(hmap[h], i)
        fmt.Printf("%v\n", hmap[h])
    }
}

func main() {
    // Read in old file
    oldBytes, err := ioutil.ReadFile("tmpfile1.txt")

    if err != nil {
        return
    }

    newBytes, err2 := ioutil.ReadFile("tmpfile2.txt")

    if err2 != nil {
        return
    }
    oldStart, oldEnd := -1, -1
    newStart, newEnd := -1, -1

    var operationList []Operation

    // preprocess old file, hash 5-10 byte chunks for stopping the diff on replacements:

    hmap := make(map[uint64][]int)

    computeRollingHash(oldBytes, hmap)

    fmt.Printf("%v\n", hmap)

    oldPtr, newPtr := 0, 0

    // slice of changes

    for oldPtr < len(oldBytes) && newPtr < len(newBytes)   {
        if oldBytes[oldPtr] == newBytes[newPtr]   {
            oldPtr++
            newPtr++
        } else    {
            // DIFF TIME
            oldStart = oldPtr
            newStart = newPtr
            var h uint64 = 0
            var newSlice []byte

            var startRune, endRune byte = 0, 0
            var found bool = false
            var temp []int
            for (found != true && newPtr < len(newBytes))    {
                fmt.Printf("%d %d\n", oldPtr, newPtr)
                fmt.Printf("%d\n", h)
                if newPtr + SIZE < len(newBytes)   {
                    newSlice = newBytes[newPtr:newPtr+SIZE]
                } else {
                    newSlice = newBytes[newPtr:]
                }
                h = computeHash(h, newSlice, startRune, endRune)
                temp =  hmap[h]
                fmt.Printf("%v %s\n", temp, newSlice)
                for i := range temp  {
                    if i > oldStart {
                        oldEnd = i
                        newEnd = newPtr
                        insertString := oldBytes[oldStart:oldEnd]

                        fmt.Println("Found it!")
                        deleteOp := Operation{0, newStart, newEnd, ""}
                        insertOp := Operation{1, oldStart, oldEnd, string(insertString)}
                        operationList = append(operationList, deleteOp, insertOp)
                        found = true
                        oldPtr = oldEnd + SIZE
                        newPtr = newEnd + SIZE
                        break
                    }
                }
                if !found   {
                    newPtr++
                    startRune = newSlice[0]
                    endRune = newSlice[len(newSlice)-1]
                }
            }
       }

    }

    fmt.Println("shit didn't break maybe?")

    fmt.Printf("%v\n", operationList)
    ///////////////////////////////////////////////////////////////////
    // When we reach the end of one file, make a deletion/insertion of the remaining
    // bytes of the other file (operation depending on which has remaining bytes)
    ///////////////////////////////////////////////////////////////////
    fmt.Printf("%d %d %d\n", oldPtr, newPtr,len(newBytes))
    // while not EoF

    // compare each byte

    // set old_s and new_s once a difference occurs
}

