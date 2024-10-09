package main

import "fmt"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    s1 := split(sentence1)
	s2 := split(sentence2)
	fmt.Println(s1)
	fmt.Println(s2)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	l := 0
	for l < len(s1) && s1[l] == s2[l]{
		l++
	}
	r1, r2:= len(s1) -1 , len(s2) -1

	for r1 >= l && r2 >= 0 && s1[r1] == s2[r2]{
		r2--
		r1--
	}
	return l > r1

}
func split(words string) []string {
        w := []string{}
        startIndex := 0
        inWord := false

        for i := range words {
                if words[i] != ' ' {
                        if !inWord {
                                startIndex = i
                                inWord = true
                        }
                } else {
                        if inWord {
                                w = append(w, words[startIndex:i])
                                inWord = false
                        }
                }
        }
        if inWord {
                w = append(w, words[startIndex:])
        }
        return w
}


func main(){
	fmt.Println(areSentencesSimilar("I love programming", "I love coding"))
	fmt.Println(areSentencesSimilar("Hello world", "Hello"))
	fmt.Println(areSentencesSimilar("Go is great", "Go is awesome"))
	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
}