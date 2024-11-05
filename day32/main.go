package main

import (
    "sort"
    "fmt"
)



// split splits a string by a given delimiter and returns a slice of substrings
func split(s, delimiter string) []string {
    var result []string
    start := 0
    for i := 0; i < len(s); i++ {
        if string(s[i]) == delimiter {
            result = append(result, s[start:i])
            start = i + 1
        }
    }
    result = append(result, s[start:])
    return result
}

// func removeSubfolders(folder []string) []string {
//     var result []string
//     m := make(map[string]bool)

//     for _, v := range folder {
//         m[v] = true
//     }

//     for i := 0; i < len(folder); i++ {
//         path := folder[i]
//         found := false
//         for j := 1; j < len(path); j++ {
//             if path[j] == '/' {
//                 if _, ok := m[path[:j]]; ok {
//                     found = true
//                 }
//             }
//         }
//         if !found {
//             result = append(result, path)
//         }
//     }
//     return result
// }


func removeSubfolders(folder []string) []string {
    // Sort the folders lexicographically
    sort.Strings(folder)
    
    var result []string
    result = append(result, folder[0])

    for i := 1; i < len(folder); i++ {
        lastAdded := result[len(result)-1]
        if !(len(folder[i]) > len(lastAdded) && folder[i][:len(lastAdded)] == lastAdded && folder[i][len(lastAdded)] == '/') {
            result = append(result, folder[i])
        }
    }
    return result
}


func main() {
    folders := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
    result := removeSubfolders(folders)
    fmt.Println(result) // Output should be ["/a", "/c/d", "/c/f"]
}
