package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar Página de Usuários"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página não existe bro!"))
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE - SERVIDOR => CLIENTE FAZ UMA REQUISIÇÃO, SERVIDOR PROCESSA E DEVOLVE UMA RESPOS	TA AO CLIENTE.
	// REQUEST - RESPONSE

	// ROTAS => Maneiras de identificar quais tipos de requisição e processamento devem ser utilizados.
	// URI - Identificador do Recurso
	// MÉTODOS - GET (Buscar Dados), POST (Enviar Dados), PUT (Atualizar Dados), DELETE (Deletar Dados)

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":5000", nil)) // Servidor rodando na porta 5000
}
