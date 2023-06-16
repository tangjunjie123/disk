package main

import (
	"disk/sql"
	"fmt"
)

func main() {
	key := "1213"
	get := sql.RedSadd(key, "123123", 60)
	fmt.Println(get)
}
