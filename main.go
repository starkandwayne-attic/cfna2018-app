package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Todo")
	panic("aaaa")
	a := App{}
	a.Initialize(os.Getenv("APP_VERSION"), os.Getenv("DB_PATH"))
	a.Run(":8080")
}
