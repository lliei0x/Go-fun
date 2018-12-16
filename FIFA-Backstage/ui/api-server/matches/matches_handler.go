package matches

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

var (
	ErrorMatchNumber = errors.New("number not allow")
)

// MatchHandler will list one match
// @Summary List Match
// @Accept json
// @Tags Match
// @Security Bearer
// @Produce  json
// @Param matchID query string false "matchID"
// @Resource Match
// @Router /v1/api/matches/:matchID [get]
// @Success 200 {object} model.MatchSerializer
func HandlerGetMatchesByID(ctx *gin.Context) {
	id := ctx.Param("matchID")

	number, _ := strconv.Atoi(id)

	if number > 64 || number < 0 {
		ctx.JSON(400, ctx.AbortWithError(400, ErrorMatchNumber))
		return
	}

	var match model.Match
	if dbError := initiator.POSTGRES.Where("id = ?", id).First(&match).Error; dbError != nil {
		ctx.JSON(400, ctx.AbortWithError(400, dbError))
		return
	}

	ctx.JSON(200, match.Serializer())

}
