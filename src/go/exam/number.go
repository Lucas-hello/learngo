package main

import (
	"fmt"
)

func main() {
	fmt.Println(Deduplication("aadddddffffff"))
	//b := "adddd"
	//c := b[:2] + b[3:]
	//fmt.Println(c)
	killRepeat("SSFDFSFSF")

}
func Deduplication(s string) map[string]int {
	mapS := make(map[string]int)

	for _, a := range s {
		severalS := 0
		fmt.Println("a:", a)

		for i := 0; i < len(s); i++ {

			sn := int32(s[i])
			fmt.Println("si", s[i])

			if sn == a {
				severalS = severalS + 1
				fmt.Println("severalS:", severalS)

			}
			s1 := fmt.Sprintf("%q", a)
			mapS[s1] = severalS

		}
	}
	return mapS
}
func killRepeat(s string) string {

	for _, vie := range s {
		num:=++
		for i :=0 ;i<len(s);i++{
			p:=int32(s[i])
			if vie==
		}

	}
	return s
}

//func Deduplication(s string) string {
//	fmt.Println("s:", s)
//	sl := make([]string, 0, len(s))
//	var ss string
//	for i := 0; i < len(s); i++ {
//		sl = append(sl, s[i:i+1])
//		fmt.Println("sl:", sl)
//		for i1 := 0; i1 <= len(sl)-1; i1++ {
//			for i2 := i + 1; i2 <= len(sl)-1; i2++ {
//				if sl[i1] == sl[i2] {
//					copy(sl[:i1], sl[i1+1:])
//					sl = sl[:len(sl)-1]
//					fmt.Println("sl2:", sl)
//					for i3 := 0; i3 < len(sl); i3++ {
//						ss += sl[i]
//						fmt.Println("ss:", ss)
//					}
//				}
//			}
//		}
//	}
//	return ss
//
//}

//func Deduplication(s string) string {
//	var b *string = &s
//	for i := 0; i <= len(s)-1; i++ {
//		for i1 := i + 1; i1 <= len(s)-1; i1++ {
//
//			//fmt.Println(s)
//			if s[i] == s[i1] {
//				*b = s[:i1] + s[i+1:]
//				//fmt.Println(b)
//			}
//
//		}
//	}
//	return *b
//
//}

//func SameLetter(s string)(string,int) {
//	for i:=0;i<len(s);i++{
//		for _,i1 :=range s{
//			if s[i]==i1
//		}
//	}
//
//}
