package userapi

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/runupdateuser"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
	"zakariawahyu.com/go-gogen-mytodo/shared/model/payload"
	"zakariawahyu.com/go-gogen-mytodo/shared/util"
)

func (r *ginController) runupdateuserHandler() gin.HandlerFunc {

	type InportRequest = runupdateuser.InportRequest
	type InportResponse = runupdateuser.InportResponse

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

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Name = jsonReq.Name
		req.CurrentEmail = c.MustGet("currentUser").(string)
		req.Email = jsonReq.Email
		req.Password = jsonReq.Password
		req.Now = time.Now()

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
