package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
    word := ""
    maxL := 0
    for _, v := range s {
        if v == ' ' && word != "" {
            maxL = len(word)
            word = ""
        } else if v != ' ' {
            word += string(v) 
        }
    }

    
    if word != "" {
        maxL = len(word)
    }

    return maxL
}

func main (){
	s := "   fly me   to   the moon  "
	l := lengthOfLastWord(s)
	fmt.Println(l)
}