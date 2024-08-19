package routes

import (
	"github.com/DiogoHumberto/api-go-gin-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/alunos", controllers.ExibirTodosAlunos)
	r.POST("/alunos", controllers.CadastrarAluno)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)
	r.PUT("/alunos/:id", controllers.AtualizarAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCpf)
	r.Run(":5000")
}
