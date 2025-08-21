package main

import (
	"fmt"
	"strings"
)

func main() {
text:="go is simple and go is powerfull"
	word := strings.Fields(text)
	freq := make(map[string]int)

	for _,v:=range word{
		freq[v]++
	}
	for i,k:=range freq{
		fmt.Printf("%s -> %d \n",i,k)
	}

}