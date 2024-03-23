package routes

import (
	"net/http"
	"trabalhando-golang/controller"
)

func CarregarRotas() {
	http.HandleFunc("/", controller.ExibirProdutos)
	http.HandleFunc("/new", controller.ExibirNovoProduto)
	http.HandleFunc("/insert", controller.CadastrarProduto)
	http.HandleFunc("/delete", controller.DeletarProduto)
	http.HandleFunc("/edit", controller.ExibirEdicaoProduto)
	http.HandleFunc("/update", controller.EditarProduto)
}
