package main

import (
	"fmt"
)

// type Season int

// const (
// 	_ 		Season = iota
// 	Spring
// 	Summer
// 	Autumn
// 	Winter
// )

// func (s Season) String() string {
// 	switch s {
// 		case Spring: return "весна"
// 		case Summer: return "лето"
// 		case Autumn: return "осень"
// 		case Winter: return "зима"
// 		default: return  "неизвестно"
// 	}
// }

// fmt.Println(int(Winter)) // 1
// fmt.Println(Winter)      // "весна"

// Задача 3
// type FileMode int

// const (
//     ModeRead    FileMode = 1 << iota
//     ModeWrite
//     ModeExecute
// )

// func Describe(m FileMode) string {
// 	var s []string

// 	if m & ModeRead != 0 {
// 		s = append(s, "read")
// 	}
// 	if m & ModeWrite != 0 {
// 		s = append(s, "write")
// 	}
// 	if m & ModeExecute != 0 {
// 		s = append(s, "execute")
// 	}

// 	if len(s) == 0 {
// 		return "нет прав"
// 	}

// 	return strings.Join(s, " ")
// }

// switch m {
// 	case 1: return "read"
// 	case 2: return "write"
// 	case 3: return "read write"
// 	case 4: return "execute"
// 	case 5: return "read execute"
// 	case 6: return "write execute"
// 	case 7: return "write execute read"
// 	default: return "неизвестно"
// }

// perms := ModeRead | ModeWrite
// fmt.Println(Describe(perms)) // "read write"

// 4 задача

const (
	KB = 1 << (10 * iota) 
	MB 
	GB 
	TB
)

func FormatSize(bytes int) string {
	
}


func main() {

	fmt.Println(FormatSize(1536))       // "1.50 KB"
	fmt.Println(FormatSize(2621440))    // "2.50 MB"
	fmt.Println(FormatSize(5368709120)) // "5.00 GB"

}