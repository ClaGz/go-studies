package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(true)
	generica("aloo")
	generica(678)
	generica(map[string]string{"asduaw": "ahusdhuas"})
}
