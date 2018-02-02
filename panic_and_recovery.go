package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	f()
	fmt.Println("Returned normally from f....")
}

func f(){
	//defer func(){
	//	if r:= recover(); r != nil {
	//		fmt.Println("Recovered", r)
	//	}
	//}()
	text, err := ioutil.ReadFile("a.txt")
	if err != nil{
		fmt.Println("file doesn't exists")
		//panic("File doesn't exists")
	}

	fmt.Println(text)

}
