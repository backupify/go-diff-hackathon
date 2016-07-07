package main

import (
    //"io"
    "io/ioutil"
    "fmt"
    //"math"
)

type Operation struct {
    Type, StartType, EndType int
    Chars string
}

func powerfulPow(base uint64, pow int) uint64 {
    if pow == 0 {
        return 1
    } else  {
        return base * powerfulPow(base, pow - 1)
    }
}

func computeRollingHash(input []byte, hmap map[uint64][]int) {
    PRIME_BASE := 257
    //PRIME_MOD := 1000000007

    length := len(input)

    size := 10;

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
            }

        }
        fmt.Printf("%d\n", h)
        oldChar = uint64(input[i])
        if i + size < length    {
            newChar = uint64(input[i + (size - 1)])
        } else  {
            newChar = 0
        }

        if hmap[h] == nil   {
            hmap[h] = make([]int, 5)
        }
        hmap[h] = append(hmap[h], i)
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
    //oldStart, oldEnd := -1, -1
    //newStart, newEnd := -1, -1

    // preprocess old file, hash 5-10 byte chunks for stopping the diff on replacements:

    hmap := make(map[uint64][]int)

    computeRollingHash(oldBytes, hmap)

    for i := 0; i < len(oldBytes); i++  {
        // grab size 10 chunks at a time
        //get hash for each
        //store in hashMap
    }

    oldPtr, newPtr := 0, 0

    for i := 0; i < len(oldBytes); i++   {
        //if oldBytes[i] == newBytes[i]   {
            //continue
        //} else    {
           // 
        //}

        oldPtr++
        newPtr++
    }

    fmt.Printf("%d %d %d\n", oldPtr, newPtr,len(newBytes))
    // while not EoF



    // compare each byte

    // set old_s and new_s once a difference occurs
}

