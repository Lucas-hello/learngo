package main

//import (
//	"fmt"
//	"strconv"
//)
//
//type Element interface {}//定义一个空的接口
//type List []Element//定义一个空接口类型的类
//type Person struct {
//	name string
//	age int
//}
//
//func (p Person)String()string {
//	return "(name:"+p.name+"-age:"+strconv.Itoa(p.age)+"years)"}
//
//func main() {
//	list:=make(List,3)
//	list[0] = 1
//	list[1] = "Hello"
//	list[2]=Person{"Dennis",70}
//	for index,element:=range list{
//		if value,ok:=element.(int);ok{
//			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
//		} else if value,ok:=element.(string);ok{
//			fmt.Printf("list[%d] is an string and its value is %s\n",index,value)
//		} else if value,ok:=element.(Person);ok {
//				fmt.Printf("list[%d] is an Persion and its calue is %s\n ",index,value)
//		}else {fmt.Printf("list[%d] is of a different type",index)
//
//		}
//	}
//
//
//
//
//}
