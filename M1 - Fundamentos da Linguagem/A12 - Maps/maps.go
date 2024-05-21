package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":  "pedro",
		"teste": "davi",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	// TIPAGEM DO MAP POUCO DINÂMICA. TODAS AS KEYS
	// E VALORES SÃO DO MESMO TIPO, EXEMPLO
	// TODAS AS CHAVES STRING E VALORES INT
	// SERÁ ASSIM ATÉ O FINAL DO MAP

	usuario2 := map[string]map[string]string{
		"teste": {
			"primeiro": "homem",
			"ultimo":   "mulher",
		},
		"teste2": {
			"segundo":       "homem",
			"antepenultimo": "mulher",
		},
	}

	fmt.Println(usuario2)

	// DELETAR CHAVES

	delete(usuario2, "teste")

	fmt.Println(usuario2)

	// Adicionar chaves

	usuario2["Signo"] = map[string]string{
		"Cancer": "Algum mes ai",
		"Touro":  "Algum mes ai v2",
	}

	fmt.Println(usuario2)
}
