package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
		return false
	}
	s1Count := [26]int{}
	s2Count := [26]int{}
	for i := 0; i < len(s1); i++ {
		s1Count[ s1[i] -'a'] +=1
		s2Count[ s2[i] -'a'] +=1
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	l :=0
	i :=0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		i = int(s2[r] - 'a');
		s2Count[i] ++;
		if s1Count[i] == s2Count[i] {
			matches ++;
		}else if  s1Count[i] + 1 == s2Count[i] {
			matches --;
		}
		i = int(s2[l] - 'a');
		s2Count[i] --;
		if s1Count[i] == s2Count[i] {
			matches ++;
		}else if  s1Count[i] - 1 == s2Count[i] {
			matches --;
		}
		l++;

	}

	if matches == 26 {
		return true
	}
	return false
}

func main(){
	fmt.Println(checkInclusion("ab", "eidbaooo")) // Output: true
	fmt.Println(checkInclusion("ab", "eidboaoo")) // Output: false
}
0235341084
0235342122