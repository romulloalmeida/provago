package handler

import (
	"prova/pkg/web"
	"errors"
	//"math"
	"net/http"
	"strconv"
	//"strings"
	"prova/internal/domain"
	"prova/internal/dentista"
	"github.com/gin-gonic/gin"
)

// comunicação com o service
type dentistaHandler struct {
	s dentista.Service
}

//instanciando o controlador
func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

//funcao getByID
func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(c, 400, "error", "An error occurred while trying to get a dentista")
			return
		}
		pReturn, err := h.s.GetByID(id)
		if err != nil {
			web.Failed(c, http.StatusNotFound, "error", "dentista not found!")
			return
		}
		web.Success(c, http.StatusOK, pReturn)
	}
} 


//funcao para validar que os campos não estão vazios
func validateEmptysDentista(dentista *domain.Dentista) (bool, error) {
	if (dentista.Nome == "" || dentista.Sobrenome == "" || dentista.Matricula == "") {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

//funcao post
func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentista domain.Dentista
		err := ctx.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "dentista invalido")
			return
		}
		valid, err := validateEmptysDentista(&dentista)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		d, err := h.s.Create(dentista)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}

		web.Success(ctx, http.StatusOK, d)
	}
}

//funcao delete
func (h *dentistaHandler) Delete() gin.HandlerFunc {
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
		ctx.JSON(200, gin.H{"success": "dentista deletedo"})
	}
}

//funcao put
func (h *dentistaHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "id invalido")
			return
		}
		var dentista domain.Dentista
		err = ctx.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "dentista invalido")
			return
		}
		valid, err := validateEmptysDentista(&dentista)
		if !valid {
			web.Failed(ctx, http.StatusBadRequest, "error", err.Error())
			return
		}
		d, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failed(ctx, http.StatusConflict, "error", err.Error())
			return
		}
		web.Success(ctx, http.StatusOK, d)
	}
}

//funcao patch
func (h *dentistaHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "id invalido")
			return
		}
		var dentista domain.Dentista
		err = ctx.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, "error", "dentista invalido")
			return
		}
		d, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failed(ctx, http.StatusConflict, "error", err.Error())
			return
		}
		web.Success(ctx, http.StatusOK, d)
	}
}