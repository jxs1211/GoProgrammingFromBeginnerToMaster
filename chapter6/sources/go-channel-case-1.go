package main

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}

func showGoChanCase1() {
	go f()
	c <- 5
	println(a)
}

func main() {
	showGoChanCase1()

}
