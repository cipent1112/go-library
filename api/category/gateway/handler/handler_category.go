package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validate "github.com/go-playground/validator/v10"
	"go-library/api/category/entity"
	"go-library/lib/binder"
	"go-library/lib/presenter"
	"net/http"
)

type CategoryContract interface {
	Get(ctx *gin.Context)
	GetList(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func (h *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.Usecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presenter.Default{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_GET, h.Name),
		Data:    res,
	})
}

func (h *Handler) GetList(ctx *gin.Context) {
	param, err := binder.ValidateRequestList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}
	res, meta, err := h.Usecase.GetList(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presenter.List{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_GET_LIST, h.Name),
		Meta:    meta,
		Data:    res,
	})
}

func (h *Handler) Store(ctx *gin.Context) {
	param := entity.Category{}
	if err := ctx.Bind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	v := validate.New()
	if err := v.Struct(param); err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	if err := h.Usecase.Store(param); err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, presenter.Default{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_CREATE, h.Name),
	})
}

func (h *Handler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	param := entity.Category{}
	if err := ctx.Bind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	v := validate.New()
	if err := v.Struct(param); err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	if err := h.Usecase.Update(id, param); err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presenter.Default{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_UPDATE, h.Name),
	})
}

func (h *Handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.Usecase.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presenter.Default{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_DELETE, h.Name),
	})
}
