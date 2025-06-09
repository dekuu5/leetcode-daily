package main

import (
	"fmt"
)
// 205. Isomorphic Strings
func isIsomorphic(s string, t string) bool {

	smap := make(map[byte]byte)
	tmap := make(map[byte]byte)
	n := len(s)
	for i := 0; i < n; i++ {
		sc, tc := s[i], t[i]

		if mapped, ok := smap[sc]; ok {
			if mapped != tc {
				return false
			}
		} else {
			smap[sc] = tc
		}

		if mapped, ok := tmap[tc]; ok {
			if mapped != sc {
				return false
			}
		} else {
			tmap[tc] = sc
		}
	}
	return true

}

func main(){
	s, t := "bbbaaaba", "aaabbbba"

	fmt.Println(isIsomorphic(s,t))
}


