//编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格(参考 unicode.IsSpace)替换成一个空格返回
package main

func main() {

}
func space(s []byte) {
	for i := 1; i < len(s)-1; i++ {
		if s[i] == s[i+1] && s[i] == ' ' {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		}
	}

}
