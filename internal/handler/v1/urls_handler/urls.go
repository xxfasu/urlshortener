package urls_handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxfasu/urlshortener/internal/service/urls_service"
	"github.com/xxfasu/urlshortener/internal/validation"
	"github.com/xxfasu/urlshortener/pkg/logs"
	"github.com/xxfasu/urlshortener/pkg/response"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	service urls_service.Service
}

func New(service urls_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateURL(ctx *gin.Context) {
	req := new(validation.CreateURL)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	userID, _ := ctx.Get("userID")
	req.UserID = userID.(int)

	resp, err := h.service.CreateURL(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, resp)
}

func (h *Handler) RedirectURL(ctx *gin.Context) {
	shortCode := ctx.Param("code")
	fmt.Println(shortCode)

	originalURL, err := h.service.GetURL(ctx, shortCode)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	go func() {
		if err := h.service.IncrViews(context.Background(), shortCode); err != nil {
			logs.Log.Error("failed to incr view ", zap.String("shortCode", shortCode))
		}
	}()
	ctx.Redirect(http.StatusPermanentRedirect, originalURL)
}

func (h *Handler) GetURLs(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	req := new(validation.GetURLs)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.Size == 0 {
		req.Size = 10
	}

	req.UserID = userID.(int)

	resp, err := h.service.GetURLs(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, resp)
}

func (h *Handler) DeleteURL(ctx *gin.Context) {
	shortCode := ctx.Param("code")

	if err := h.service.DeleteURL(ctx, shortCode); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}

func (h *Handler) UpdateURLDuration(ctx *gin.Context) {
	req := new(validation.UpdateURLDuration)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	req.Code = ctx.Param("code")

	if err := h.service.UpdateURLDuration(ctx, req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}
