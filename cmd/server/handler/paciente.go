package handler

import (
	"prova/pkg/web"
	"errors"
	//"math"
	"net/http"
	"strconv"
	"strings"
	"prova/internal/domain"
	"prova/internal/paciente"
	"github.com/gin-gonic/gin"
)

// comunicação com o service
type pacienteHandler struct {
	s paciente.Service
}

//instanciando o controlador
func NewPacientHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

//funcao getByID
func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(c, 400, "error", "An error occurred while trying to get a paciente")
			return
		}
		pReturn, err := h.s.GetByID(id)
		if err != nil {
			web.Failed(c, http.StatusNotFound, "error", "Paciente not found!")
			return
		}
		web.Success(c, http.StatusOK, pReturn)
	}
} 

//funcao para validar data
func validateExpiration(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

//funcao para validar que os campos não estão vazios
func validateEmptys(paciente *domain.Paciente) (bool, error) {
	if (paciente.Nome == "" || paciente.Sobrenome == "" || paciente.RG == "" || paciente.DataCadastro == "") {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

//funcao post
func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paciente domain.Paciente
		err := ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "paciente invalido")
			return
		}
		valid, err := validateEmptys(&paciente)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		valid, err = validateExpiration(paciente.DataCadastro)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		p, err := h.s.Create(paciente)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}

		web.Success(ctx, http.StatusOK, p)
	}
}

//funcao delete
func (h *pacienteHandler) Delete() gin.HandlerFunc {
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
		ctx.JSON(200, gin.H{"success": "paciente deletedo"})
	}
}

//funcao put
func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "id invalido")
			return
		}
		var paciente domain.Paciente
		err = ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "paciente invalido")
			return
		}
		valid, err := validateEmptys(&paciente)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		valid, err = validateExpiration(paciente.DataCadastro)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		p, err := h.s.Update(id, paciente)
		if err != nil {
			web.Failed(ctx, http.StatusConflict, "error", err.Error())
			return
		}
		web.Success(ctx, http.StatusOK, p)
	}
}

//funcao patch
func (h *pacienteHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "id invalido")
			return
		}
		var paciente domain.Paciente
		err = ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "paciente invalido")
			return
		}
		valid, err := validateExpiration(paciente.DataCadastro)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		p, err := h.s.Update(id, paciente)
		if err != nil {
			web.Failed(ctx, http.StatusConflict, "error", err.Error())
			return
		}
		web.Success(ctx, http.StatusOK, p)
	}
}