package main

func sum(a []int ,c chan int)  {
	sum:=0
	for _,v:=range (a){
		sum+=v
	}
	c<-sum

	
}

func main() {
	a:=[]int{7,2,8,9,-4,0}
	c:=make(chan int)

}
