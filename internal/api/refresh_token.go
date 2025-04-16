package api

import (
	"net/http"

	"github.com/adityasuryadi/ewallet/constants"
	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	claim, ok := c.Get("token")
	refreshToken := c.Request.Header.Get("Authorization")
	if !ok {
		log.Error("failed to get claim in context")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse claim to claimToken")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Error("failed on RefreshToken service: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}
