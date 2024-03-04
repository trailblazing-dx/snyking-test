package main 

import (
	"os"
)


func main() {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	ret, err := os.ReadFile(password)
	if err != nil {
		return 
	}
}

