package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeimedee/loverboy/internal/models"
	"github.com/zeimedee/loverboy/internal/services"
)

type TapHandler struct {
	TapService *services.LoversService
}

func NewTapHandler(loverService *services.LoversService) *TapHandler {
	return &TapHandler{
		TapService: loverService,
	}
}

func Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func (l *TapHandler) Register(ctx *gin.Context) {
	taps := new(models.Register)
	if err := ctx.ShouldBindBodyWithJSON(taps); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	l.TapService.StoreWord(taps.Id, taps.Token)
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (l *TapHandler) Tap(ctx *gin.Context) {
	id := ctx.Param("id")

	token, err := l.TapService.GetToken(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	err = services.SendPushNotifs(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "sent"})
}

func (l *TapHandler) GetAll(ctx *gin.Context) {
	ids := l.TapService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{"data": ids})

}
