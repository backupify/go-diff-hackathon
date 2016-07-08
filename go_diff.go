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

    size := len(input)

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
                //fmt.Printf("%v", string(input[j]))
            }
            //fmt.Printf("\n")

        }
        //fmt.Printf("%v %v\n", string(oldChar), string(newChar))
        oldChar = uint64(input[i])
        if i + size < length    {
            newChar = uint64(input[i + (size)])
            //fmt.Printf("%d %d", i, i + size)
        } else  {
            newChar = 0
        }

        if hmap[h] == nil   {
            hmap[h] = make([]int, 0)
        }
        hmap[h] = append(hmap[h], i)
        //fmt.Printf("%v\n", hmap[h])
    }
}

func main() {
    // Read in old file
    oldBytes, err := ioutil.ReadFile("tmpfile2.txt")

    if err != nil {
        return
    }

    newBytes, err2 := ioutil.ReadFile("tmpfile1.txt")

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
            for (found != true && newPtr < len(newBytes) && oldPtr < len(oldBytes))    {
                //fmt.Printf("%d %d ||| %d %d\n", oldPtr, newPtr, len(oldBytes), len(newBytes))
                //fmt.Printf("%d\n", h)
                if newPtr + SIZE < len(newBytes)   {
                    newSlice = newBytes[newPtr:newPtr+SIZE]
                } else {
                    newSlice = newBytes[newPtr:]
                }
                //fmt.Printf("%d %v\n", h, string(newSlice))
                h = computeHash(h, newSlice, startRune, endRune)
                temp =  hmap[h]
                //fmt.Printf("%v %s %s %s\n", temp, newSlice, string(startRune), string(endRune))
                for i := 0; i < len(temp); i++  {
                    //fmt.Printf("%d %d\n", temp[i], oldStart)
                    if temp[i] >= oldStart {
                        oldEnd = temp[i]
                        newEnd = newPtr
                        insertString := oldBytes[oldStart:oldEnd]

                        //fmt.Println("Found it!")
                        deleteOp := Operation{0, newStart, newEnd, ""}
                        //fmt.Printf("%d %d %s\n", oldStart, oldEnd, insertString)
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
                    oldPtr++
                    startRune = newBytes[newPtr-1]
                    if newPtr + SIZE - 1< len(newBytes)   {
                        endRune = newBytes[newPtr+SIZE-1]
                    } else  {
                        endRune = 0
                    }

                }
            }
       }

    }

    //fmt.Println("shit didn't break maybe?")

    ///////////////////////////////////////////////////////////////////
    // When we reach the end of one file, make a deletion/insertion of the remaining
    // bytes of the other file (operation depending on which has remaining bytes)
    ///////////////////////////////////////////////////////////////////
    //fmt.Printf("%d %d %d\n", oldPtr, newPtr,len(newBytes))

    if oldPtr < len(oldBytes)   {
        // do an insertion of remaining oldBytes
        insertOp := Operation{1, oldPtr, len(oldBytes), string(oldBytes[oldPtr:])}

        fmt.Printf("%x\n", string(newBytes[len(newBytes)-1]))
        operationList = append(operationList, insertOp)
    }
    if newPtr < len(newBytes){
        // do a deletion of remaining newBytes
        deleteOp := Operation{0, newPtr, len(newBytes)-1, ""}
        operationList = append(operationList, deleteOp)
    }

    fmt.Printf("%v\n", operationList)

}

