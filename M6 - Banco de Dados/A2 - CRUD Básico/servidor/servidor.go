package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converte o usuário para struct"))
		return
	}

	fmt.Println(usuario)

	db, erro := banco.ConectarNoBancoDeDados()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco!"))
	}
	defer db.Close()

	// PREPARE STATEMENT (EVITAR SQL INJECTION)

	statement, erro := db.Prepare(
		"INSERT INTO USUARIOS (nome, email) VALUES (?, ?)") // A ? significa que não vamos colocar os valores agora
	// ISSO EVITARÁ QUE SEJA POSSÍVEL REALIZAR UM SQL INJECTION

	if erro != nil {
		w.Write([]byte("Falha ao Criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id:%d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.ConectarNoBancoDeDados()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	// SELECT * from usuarios

	linhas, erro := db.Query("SELECT * FROM usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {

		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)

	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Error encode in JSON"))
		return
	}

	/*

		O JSON é retornado porque o Encode escreve diretamente no w, que é o http.ResponseWriter fornecido ao manipulador HTTP.
		http.ResponseWriter é responsável por enviar a resposta HTTP de volta ao cliente. Assim, qualquer coisa escrita no http.
		ResponseWriter será incluída no corpo da resposta HTTP.

	*/
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32) // 2nd BASE, 3rd BITS

	if erro != nil {
		w.Write([]byte("Error ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.ConectarNoBancoDeDados()
	if erro != nil {
		w.Write([]byte("Error ao conectar no banco de dados!"))
		return
	}

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Error ao buscar usuário!"))
		return
	}

	var usuario usuario

	if linha.Next() {
		if erro != linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email) {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	if usuario.ID == 0 {
		w.Write([]byte("Usuário não existe!"))
		return
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Error encode in JSON"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Error ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converte o usuário para struct"))
		return
	}

	db, erro := banco.ConectarNoBancoDeDados()

	if erro != nil {
		w.Write([]byte("Error ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Falha ao Criar o statement"))
		return
	}

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Error ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Error ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.ConectarNoBancoDeDados()

	if erro != nil {
		w.Write([]byte("Error ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Falha ao Criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Falha ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
