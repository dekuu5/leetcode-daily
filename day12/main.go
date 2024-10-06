package main

import "fmt"

type Prefix struct {
	count int
	index []int
}

func sumPrefixScores(words []string) []int {
	m := make(map[string]Prefix)
	
	for ind, word := range words {
		for i := 1; i <= len(word); i++ {
			prefix := word[:i]

			if _, ok := m[prefix]; !ok {
				m[prefix] = Prefix{
					count: 1,
					index: []int{ind},
				}
			} else {
				p := m[prefix]
				p.count += 1
				p.index = append(p.index, ind)
				m[prefix] = p
			}
			fmt.Println(prefix, m[prefix])
		}
	}

	result := make([]int, len(words))

	for _, v := range m {
		index := v.index
		count := v.count
		for i := 0; i < len(index); i++ {
			result[index[i]] += count
		}
	}

	return result
}

func main() {
	words := []string{"abc", "ab", "bc", "b"}
	fmt.Println(sumPrefixScores(words)) // Expected result will vary based on input
}
