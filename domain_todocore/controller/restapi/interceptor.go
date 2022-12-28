package restapi

import (
	"net/http"
	"strings"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/errorenum"
	"zakariawahyu.com/go-gogen-mytodo/shared/model/payload"
	"zakariawahyu.com/go-gogen-mytodo/shared/util"

	"github.com/gin-gonic/gin"
)

func (r *ginController) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := util.GenerateID(16)
		token, err := ExtractToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, payload.NewErrorResponse(err, traceID))
			return
		}

		tokenInBytes, err := r.jwtToken.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, payload.NewErrorResponse(err, traceID))
			return
		}

		c.Set("currentUser", string(tokenInBytes))
		return
	}
}

func (r *ginController) authorization() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorized := true

		if !authorized {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

func ExtractToken(c *gin.Context) (string, error) {
	bearToken := c.GetHeader("Authorization")
	if bearToken == "" {
		return "", errorenum.NoTokenProvided
	}

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1], nil
	}

	return "", errorenum.NoTokenProvided
}
