package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	intrange := flag.Int("n", 1000, "int range")
	flag.Parse()
	fmt.Println("n=", *intrange)
	sl := NewSkipList()
	for _, v := range rand.Perm(*intrange) {
		//		fmt.Println("insert ", v)
		sl.Insert(v)
		//fmt.Printf("%v\n", sl.Next)
		//		sl.Print()
	}

	for _, v := range rand.Perm(*intrange) {
		//		fmt.Println("delete", v)
		sl.Delete(v)
		//fmt.Printf("%v\n", sl.Next)
		//		sl.Print()
	}
}
