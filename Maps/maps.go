package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Ze",
		"sobrenome": "Gilza",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "zezim",
			"segundo":  "manoel",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "do jordao",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}

	fmt.Println(usuario2)
}
