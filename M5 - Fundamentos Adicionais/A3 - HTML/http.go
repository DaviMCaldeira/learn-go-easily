package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {

	// 1 PARAMETRO: Variável de escrita; 2 PARAMETRO: Página que deseja renderizar; 3 PARAMETRO: O que vai mandar de dados para a página utilizar

	user := usuario{
		"Geralt Of Rivia",
		"geraltrivia@witcher.us",
	}

	if erro := templates.ExecuteTemplate(w, "home.html", user); erro != nil {
		log.Fatal(erro)
	}
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE - SERVIDOR => CLIENTE FAZ UMA REQUISIÇÃO, SERVIDOR PROCESSA E DEVOLVE UMA RESPOS	TA AO CLIENTE.
	// REQUEST - RESPONSE

	// ROTAS => Maneiras de identificar quais tipos de requisição e processamento devem ser utilizados.
	// URI - Identificador do Recurso
	// MÉTODOS - GET (Buscar Dados), POST (Enviar Dados), PUT (Atualizar Dados), DELETE (Deletar Dados)

	templates = template.Must(template.ParseGlob("*.html")) // Todos os templates que forem criados vão para essa variável

	http.HandleFunc("/home", home)

	fmt.Println("Escutando na porta 5000")

	log.Fatal(http.ListenAndServe(":5000", nil)) // Servidor rodando na porta 5000
}
