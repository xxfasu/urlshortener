package user_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xxfasu/urlshortener/internal/service/user_service"
	"github.com/xxfasu/urlshortener/internal/validation"
	"github.com/xxfasu/urlshortener/pkg/response"
)

type Handler struct {
	service user_service.Service
}

func New(service user_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(ctx *gin.Context) {
	req := new(validation.Login)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	resp, err := h.service.Login(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, resp)
}

func (h *Handler) Register(ctx *gin.Context) {
	req := new(validation.Register)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if err := h.service.IsEmailAvailable(ctx, req.Email); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	resp, err := h.service.Register(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, resp)
}

func (h *Handler) ForgetPassword(ctx *gin.Context) {
	req := new(validation.ForgetPassword)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err := h.service.IsEmailAvailable(ctx, req.Email)

	resp, err := h.service.ResetPassword(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, resp)
}

func (h *Handler) SendEmailCode(ctx *gin.Context) {
	email := ctx.Param("email")

	if err := h.service.SendEmailCode(ctx, email); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}
