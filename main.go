package main

import (
	"net/http"

	"myapp/controller"
	"myapp/storage"

	"github.com/labstack/echo/v4"
)

func handLeIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Está funcionando"})
}

func main() {

	// instancia o echo framework
	e := echo.New()
	storage.NewDB()

	// Rotas
	e.GET("/", inicio)
	//Retorna todos os produtos {Pode passar os parâmetros start: Inicio da paginação e Limit: Quantidade de dados a serem retornados}
	e.GET("/produtos", controller.GetProdutos)
	//Retorna um produto específico
	e.GET("/produtos/:id", controller.GetProduto)
	//Atualiza o cadastro do produto
	e.PUT("/produtos/:id", controller.SetUpdateProduto)
	//Cadastra um produto
	e.POST("/produtos", controller.SetCadProdutos)
	//Deleta um produto
	e.DELETE("/produtos/:id", controller.DeleteProduto)

	// Inicia o servidor
	e.Logger.Fatal(e.Start(":8085"))
}

// Retorno inicial
func inicio(c echo.Context) error {
	return c.String(http.StatusOK, "Tudo funcionando")

}
