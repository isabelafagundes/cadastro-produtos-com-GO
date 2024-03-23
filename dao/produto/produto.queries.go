package produto

const OBTER_PRODUTOS = "SELECT * FROM itens"
const OBTER_PRODUTO = "SELECT * FROM itens WHERE itens.id = ?"

const SALVAR_PRODUTO = "INSERT INTO itens (nome, descricao, preco, quantidade) VALUES (?,?,?,?)"

const DELETAR_PRODUTO = "DELETE FROM itens WHERE itens.id = ?"

const ATUALIZAR_PRODUTO = "UPDATE itens SET nome=?, descricao=?, preco=?, quantidade=? WHERE itens.id = ?"
