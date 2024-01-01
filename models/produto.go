package models

import (
	"api_go/data"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := data.ConectaComBancoDados()
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY ID ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := data.ConectaComBancoDados()
	queryInsert := "INSERT INTO produtos(nome, descricao, preco, quantidade) values ($1,$2,$3,$4)"
	insereDadosNoBanco, err := db.Prepare(queryInsert)

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := data.ConectaComBancoDados()
	queryDelete := "DELETE FROM produtos WHERE id=$1"
	deletarProduto, err := db.Prepare(queryDelete)
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := data.ConectaComBancoDados()
	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade

		defer db.Close()

	}
	return produtoParaAtualizar
}

func AtualizaProduto(nome, descricao string, preco float64, quantidade, id int) {
	db := data.ConectaComBancoDados()
	atualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
}
