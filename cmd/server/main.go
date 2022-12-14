package main

import (
	"github.com/gin-gonic/gin"
	"prova/cmd/server/handler"
	"prova/pkg/store/pacienteStore"
	"prova/internal/paciente"
	"prova/pkg/store/dentistaStore"
	"prova/internal/dentista"
	"prova/pkg/store/consultaStore"
	"prova/internal/consulta"
)

// função principal da aplicação
func main() {

	//camadas para os pacientes
	pacienteRepo := paciente.NewRepository(pacienteStore.NewSQLStore())
	pacienteService := paciente.NewService(pacienteRepo)
	pacienteHandler := handler.NewPacientHandler(pacienteService)

	//camadas para os dentistas
	dentistaRepo := dentista.NewRepository(dentistaStore.NewSQLStore())
	dentistaService := dentista.NewService(dentistaRepo)
	dentistaHandler := handler.NewDentistaHandler(dentistaService)

	//camadas para os consultas
	consultaRepo := consulta.NewRepository(consultaStore.NewSQLStore())
	consultaService := consulta.NewService(consultaRepo)
	consultaHandler := handler.NewConsultaHandler(consultaService)

	// criando a aplicação web com o gin
	r := gin.Default()


	// criando rota /paciente
	pacientes := r.Group("/paciente")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PUT(":id", pacienteHandler.Put())
		pacientes.PATCH(":id", pacienteHandler.Patch())
	}

	// criando rota /dentista
	dentistas := r.Group("/dentista")
	{
		dentistas.GET(":id", dentistaHandler.GetByID())
		dentistas.POST("", dentistaHandler.Post())
		dentistas.DELETE(":id", dentistaHandler.Delete())
		dentistas.PUT(":id", dentistaHandler.Put())
		dentistas.PATCH(":id", dentistaHandler.Patch())
	}

	// criando rota /consulta
	consultas := r.Group("/consulta")
	{
		consultas.GET(":id", consultaHandler.GetByID())
		consultas.POST("", consultaHandler.Post())
		consultas.DELETE(":id", consultaHandler.Delete())
		consultas.PUT(":id", consultaHandler.Put())
		consultas.PATCH(":id", consultaHandler.Patch())
	}

	// porta para rodar a aplicação
	r.Run(":8080")
}