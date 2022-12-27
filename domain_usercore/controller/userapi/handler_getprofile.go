package userapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/getprofile"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
	"zakariawahyu.com/go-gogen-mytodo/shared/model/payload"
	"zakariawahyu.com/go-gogen-mytodo/shared/util"
)

func (r *ginController) getprofileHandler() gin.HandlerFunc {

	type InportRequest = getprofile.InportRequest
	type InportResponse = getprofile.InportResponse

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

		var req InportRequest

		req.Email = c.MustGet("currentUser").(string)

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		r.log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(res, traceID))
	}
}
