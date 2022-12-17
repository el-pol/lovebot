package main

import (
	"fmt"
	"log"

	"github.com/el-pol/lovebot/fetch"
)

func main() {
	result, err := fetch.Fetch()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)

}
