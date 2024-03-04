package main 

import (
	"os"
)


func main() {
	password := "12334567"
	ret, err := os.ReadFile(password)
	if err != nil {
		return 
	}
}

