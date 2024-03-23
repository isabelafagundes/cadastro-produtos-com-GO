package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"trabalhando-golang/dao/produto"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func ExibirProdutos(w http.ResponseWriter, r *http.Request) {
	produtos, erro := produto.ObterProdutos()
	if erro != nil {
		return
	}
	erro = templates.ExecuteTemplate(w, "index", produtos)
}

func ExibirNovoProduto(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new", nil)
}

func CadastrarProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoConvertido, erro := strconv.ParseFloat(preco, 64)
		if erro != nil {
			log.Println("Erro ao obter o preço do produto: ", erro)
		}
		quantidadeConvertida, erro := strconv.Atoi(quantidade)
		if erro != nil {
			log.Println("Erro ao obter a quantidade do produto: ", erro)
		}

		produto.CadastrarProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	idConvertido, erro := strconv.Atoi(idProduto)
	if erro != nil {
		log.Println("Erro ao obter o id do produto: ", erro)
	}
	produto.DeletarProduto(idConvertido)
	http.Redirect(w, r, "/", 301)
}

func ExibirEdicaoProduto(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	produto := produto.ObterProduto(id)
	templates.ExecuteTemplate(w, "edit", produto)
}

func EditarProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")
		idProduto := r.FormValue("id")
		precoConvertido, erro := strconv.ParseFloat(preco, 64)
		if erro != nil {
			log.Println("Erro ao obter o preço do produto: ", erro)
		}
		quantidadeConvertida, erro := strconv.Atoi(quantidade)
		if erro != nil {
			log.Println("Erro ao obter a quantidade do produto: ", erro)
		}

		produto.AtualizarProduto(idProduto, nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}
