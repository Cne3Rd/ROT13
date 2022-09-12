package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

var final []rune

const (
    tl = 13
    addtl = 14
    space = ' '

)

func main() {

    alpha := map[int]rune{
      1: 'A',
      2: 'B',
      3: 'C',
      4: 'D',
      5: 'E',
      6: 'F',
      7: 'G',
      8: 'H',
      9: 'I',
      10: 'J',
      11: 'K',
      12: 'L',
      13: 'M',
      14: 'N',
      15: 'O',
      16: 'P',
      17: 'Q',
      18: 'R',
      19: 'S',
      20: 'T',
      21: 'U',
      22: 'V',
      23: 'W',
      24: 'X',
      25: 'Y',
      26: 'Z',
    }

    fmt.Print("Message: ")
    bval := bufio.NewReader(os.Stdin)
    val, err := bval.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }

    val = strings.ToUpper(val)

    for _, v := range val {
        for k, nv := range alpha {
            if v == nv {
                tp := k
                fp := 0
                if tp < addtl {
                    fp = tp + tl
                }else {
                    fp = tp - tl
                }

                fval := alpha[fp]
                final = append(final, fval)
                break
            }else if v == space {
                final = append(final, space)
                break 
            }
        }
    }

    fmt.Println(string(final))

}
