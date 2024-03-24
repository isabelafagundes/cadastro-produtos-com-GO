package produto

import (
	"trabalhando-golang/model"
	"trabalhando-golang/service/database"
)

func ObterProdutos() ([]model.Produto, error) {
	db := database.ConcetarComBancoDeDados()
	defer db.Close() //executa toda a função e apenas depois executa o defer
	produtosQuery, err := db.Query(OBTER_PRODUTOS)
	if err != nil {
		panic(err)
	}

	produto := model.Produto{}
	produtos := []model.Produto{}

	for produtosQuery.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtosQuery.Scan(
			&id, &nome, &descricao, &preco, &quantidade,
		)
		if err != nil {
			panic(err)
		}

		produto = obterProduto(nome, id, descricao, preco, quantidade)

		produtos = append(produtos, produto) //Adicionando item no slice

	}
	return produtos, nil
}

func CadastrarProduto(nome string, descricao string, preco float64, quantidade int) {
	db := database.ConcetarComBancoDeDados()
	inserirProduto, erro := db.Prepare(SALVAR_PRODUTO)

	if erro != nil {
		panic(erro)
	}
	inserirProduto.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(id int) {
	db := database.ConcetarComBancoDeDados()
	deletarProduto, erro := db.Prepare(DELETAR_PRODUTO)
	if erro != nil {
		panic(erro)
	}
	deletarProduto.Exec(id)
	defer db.Close()
}

func AtualizarProduto(id string, nome string, descricao string, preco float64, quantidade int) {
	db := database.ConcetarComBancoDeDados()
	atualizarProduto, erro := db.Prepare(ATUALIZAR_PRODUTO)
	if erro != nil {
		panic(erro)
	}
	atualizarProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

func ObterProduto(id string) model.Produto {
	db := database.ConcetarComBancoDeDados()
	produtoQuery, erro := db.Query(OBTER_PRODUTO, id)

	if erro != nil {
		panic(erro)
	}

	produto := model.Produto{}

	for produtoQuery.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro = produtoQuery.Scan(
			&id, &nome, &descricao, &preco, &quantidade,
		)

		if erro != nil {
			panic(erro)
		}

		produto = obterProduto(nome, id, descricao, preco, quantidade)
	}
	defer db.Close()
	return produto
}

func obterProduto(nome string, id int, descricao string, preco float64, quantidade int) model.Produto {
	produto := model.Produto{}
	produto.Nome = nome
	produto.Id = id
	produto.Descricao = descricao
	produto.Preco = preco
	produto.Quantidade = quantidade
	return produto
}
