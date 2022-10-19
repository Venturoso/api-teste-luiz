package controller

import (
	"net/http"
	"strconv"

	"myapp/storage"

	"myapp/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Função para retornar todos os produtos para a api
func GetProdutos(c echo.Context) error {
	//Chama a função que retorna os produtos
	produtos, erro := GetRepoProdutos(c)
	//Verifica se houve algum erro
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ocorreu um erro ao buscar o produto"})
	}

	//retorna o json dos produtos
	return c.JSON(http.StatusOK, produtos)
}

func GetRepoProdutos(c echo.Context) ([]models.Produtos, error) {
	// Inicia a instância do bd
	db := storage.GetDBInstance()
	//Cria uma instância do model de produtos
	produtos := []models.Produtos{}

	//Verifica se foi passado uma pagina
	limit, error := strconv.Atoi(c.QueryParam("limit"))
	if error != nil || limit == 0 {
		limit = 10
	}
	start, error := strconv.Atoi(c.QueryParam("start"))
	if error != nil {
		start = 0
	}

	log.Error(limit)
	//seleciona todos os produtos
	if err := db.Limit(limit).Find(&produtos, "id > ?", start).Error; err != nil {
		return nil, err
	}
	//Retorna os produtos
	return produtos, nil
}

// Função para retornar um produto específico para a api
func GetProduto(c echo.Context) error {
	//Chama a função que retorna o produto
	produtos, erro := GetRepoProduto(c)
	//Verifica se houve algum erro
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ocorreu um erro ao buscar o produto"})
	}
	//retorna o json do produto
	return c.JSON(http.StatusOK, produtos)
}

// Função para retornar um produto específico
func GetRepoProduto(c echo.Context) ([]models.Produtos, error) {
	// Inicia a instância do bd
	db := storage.GetDBInstance()
	//Cria uma instância do model de produtos
	produtos := []models.Produtos{}
	//Pega a ID do produto
	id_produto, _ := strconv.Atoi(c.Param("id"))
	//seleciona o o produto pela id
	if err := db.First(&produtos, "ID = ?", id_produto).Error; err != nil {
		return nil, err
	}
	//Retorna o produto
	return produtos, nil
}

// Função para retornar um produto específico
func DeleteProduto(c echo.Context) error {
	// Inicia a instância do bd
	db := storage.GetDBInstance()
	//Cria uma instância do model de produtos
	produtos := []models.Produtos{}
	//Pega a ID do produto
	id_produto, _ := strconv.Atoi(c.Param("id"))
	//seleciona o o produto pela id
	if err := db.Delete(&produtos, id_produto).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ocorreu um erro ao deletar o produto"})
	}
	//retorna o json do produto
	return c.JSON(http.StatusOK, map[string]string{"message": "Produto deletado com sucesso!"})
}

// Atualiza o cadastro do Produto
func SetUpdateProduto(c echo.Context) error {
	// Inicia a instância do bd
	db := storage.GetDBInstance()
	//Cria uma instância do model de produtos
	produtos := []models.Produtos{}
	//Pega a ID do produto
	id_produto, _ := strconv.Atoi(c.Param("id"))
	//Pega a situação do produto
	//situacao, _ := strconv.Atoi(c.Param("ativo"))
	//Cria uma instância do model de produtos
	p := new(models.Produtos)

	//Faz o Bind dos parâmetros
	if err := c.Bind(p); err != nil {
		///registra no log o erro.
		log.Error(err)
		//retorna o erro para o requesitante
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erro ao fazer o Bind"})
	}

	//Atualiza as informações do banco de dados
	if err := db.Model(&produtos).Where("ID = ?", id_produto).Updates(models.Produtos{Nome: p.Nome, Descricao: p.Descricao, Valor: p.Valor, Ativo: p.Ativo}).Error; err != nil {
		return nil
	}

	///Retorna o produto atualizado
	produtos, _ = GetRepoProduto(c)

	///Retorna o JSON para o requesitante
	return c.JSON(http.StatusOK, produtos)
}

// Cadastra  do Produto
func SetCadProdutos(c echo.Context) error {
	// Inicia a instância do bd
	db := storage.GetDBInstance()
	//Cria uma instância do model de produtos
	produtos := new(models.Produtos)

	//Faz o Bindo dos parâmetros
	if err := c.Bind(produtos); err != nil {
		///registra no log o erro.
		log.Error(err)
		//retorna o erro para o requesitante
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erro ao fazer o Bind"})
	}

	//Verifica se foi passado o nome
	if produtos.Nome == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Nome do produto não pode ser vazio"})
	}
	//Verifica se foi passado a Descricao
	if produtos.Descricao == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Descricao do produto não pode ser vazio"})
	}
	//Verifica se foi passado o Valor
	if produtos.Valor == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Valor do produto não pode ser 0"})
	}

	///Cria o registro no BD
	db.Create(&produtos)
	///Retorna o JSON para o requesitante
	return c.JSON(http.StatusCreated, produtos)
}
