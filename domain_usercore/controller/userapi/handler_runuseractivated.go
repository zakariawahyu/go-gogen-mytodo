package userapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/runuseractivated"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
	"zakariawahyu.com/go-gogen-mytodo/shared/model/payload"
	"zakariawahyu.com/go-gogen-mytodo/shared/util"
)

func (r *ginController) runuseractivatedHandler() gin.HandlerFunc {

	type InportRequest = runuseractivated.InportRequest
	type InportResponse = runuseractivated.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		jsonReq := InportRequest{
			Email:           c.Param("email"),
			ActivationToken: c.Param("token"),
		}

		if err := c.BindUri(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Email = jsonReq.Email
		req.ActivationToken = jsonReq.ActivationToken

		r.log.Info(ctx, util.MustJSON(req))

		_, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Message = "User successfully activated"

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
