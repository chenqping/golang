package main

import (
	"fmt"
	"golang/chapter5/links"
	"log"
	"os"
)

func breadthFirst(f func(item string) []string, worklist []string){
	seen := make(map[string]bool)
	for len(worklist) > 0{
		items := worklist
		worklist = nil
		for _, item := range items{
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func craw(url string) []string{
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil{
		log.Print(err)
	}
	return list
}

func main(){
	breadthFirst(craw, os.Args[1:])
}