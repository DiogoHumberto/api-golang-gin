package controllers

import (
	"net/http"

	database "github.com/DiogoHumberto/api-go-gin-rest/dataBase"
	"github.com/DiogoHumberto/api-go-gin-rest/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExibirTodosAlunos(c *gin.Context) {

	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)

}

func CadastrarAluno(c *gin.Context) {
	var aluno models.Aluno

	//faz o bind JSON para o struct
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func AtualizarAluno(c *gin.Context) {

	var alunoDto models.Aluno

	var alunoModel models.Aluno

	id := c.Param("id")

	result := database.DB.First(&alunoModel, id)

	if !isValidDatabaseResult(result, c) {
		return
	}

	if err := c.ShouldBindJSON(&alunoDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	// Atualiza os campos do aluno no banco de dados
	updates := models.Aluno{
		Nome:  alunoDto.Nome,
		Curso: alunoDto.Curso,
		// Adicione aqui outros campos que você deseja atualizar
	}

	// Realiza a atualização
	result = database.DB.Model(&alunoModel).Updates(updates)

	// Verifica se houve erro na atualização
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao atualizar o aluno",
		})
		return
	}

	// Retorna o aluno atualizado
	c.JSON(http.StatusOK, alunoModel)

}

func BuscarAlunoPorCpf(c *gin.Context) {

	var aluno models.Aluno
	cpf := c.Param("cpf")

	result := database.DB.Where(&models.Aluno{Cpf: cpf}).First(&aluno)

	// Interrompe a execução se a verificação falhar
	if !isValidDatabaseResult(result, c) {
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Param("id")
	database.DB.Delete(&aluno, id)

	c.Status(http.StatusNoContent)
}

func isValidDatabaseResult(result *gorm.DB, c *gin.Context) bool {
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Aluno não encontrado"})
		return false

	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": result.Error.Error()})
		return false
	}

	return true
}
