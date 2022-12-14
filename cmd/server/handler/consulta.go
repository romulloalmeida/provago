package handler

import (
	"prova/pkg/web"
	"errors"
	//"math"
	"net/http"
	"strconv"
	//"strings"
	"prova/internal/domain"
	"prova/internal/consulta"
	"github.com/gin-gonic/gin"
)

// comunicação com o service
type consultaHandler struct {
	s consulta.Service
}

//instanciando o controlador
func NewConsultaHandler(s consulta.Service) *consultaHandler {
	return &consultaHandler{
		s: s,
	}
}

//funcao getByID
func (h *consultaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(c, 400, "error", "An error occurred while trying to get a consulta")
			return
		}
		pReturn, err := h.s.GetByID(id)
		if err != nil {
			web.Failed(c, http.StatusNotFound, "error", "consulta not found!")
			return
		}
		web.Success(c, http.StatusOK, pReturn)
	}
} 


//funcao para validar que os campos não estão vazios
func validateEmptysConsulta(consulta *domain.Consulta) (bool, error) {
	if (consulta.Paciente == "" || consulta.Dentista == "" || consulta.DataHora == "") {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

//funcao post
func (h *consultaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var consulta domain.Consulta
		err := ctx.ShouldBindJSON(&consulta)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "consulta invalido")
			return
		}
		valid, err := validateEmptysConsulta(&consulta)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		d, err := h.s.Create(consulta)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}

		web.Success(ctx, http.StatusOK, d)
	}
}

//funcao delete
func (h *consultaHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "invalid id")
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failed(ctx, http.StatusNotFound, "error", err.Error())
			return
		}
		ctx.JSON(200, gin.H{"success": "consulta deletedo"})
	}
}

//funcao put
func (h *consultaHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "id invalido")
			return
		}
		var consulta domain.Consulta
		err = ctx.ShouldBindJSON(&consulta)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "consulta invalido")
			return
		}
		valid, err := validateEmptysConsulta(&consulta)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		d, err := h.s.Update(id, consulta)
		if err != nil {
			web.Failed(ctx, http.StatusConflict, "error", err.Error())
			return
		}
		web.Success(ctx, http.StatusOK, d)
	}
}

//funcao patch
func (h *consultaHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "id invalido")
			return
		}
		var consulta domain.Consulta
		err = ctx.ShouldBindJSON(&consulta)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "consulta invalido")
			return
		}
		d, err := h.s.Update(id, consulta)
		if err != nil {
			web.Failed(ctx, http.StatusConflict, "error", err.Error())
			return
		}
		web.Success(ctx, http.StatusOK, d)
	}
}