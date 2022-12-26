package userapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/runuserregister"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
	"zakariawahyu.com/go-gogen-mytodo/shared/model/payload"
	"zakariawahyu.com/go-gogen-mytodo/shared/util"
)

func (r *ginController) runuserregisterHandler() gin.HandlerFunc {

	type InportRequest = runuserregister.InportRequest
	type InportResponse = runuserregister.InportResponse

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
		req.Email = jsonReq.Email
		req.Password = jsonReq.Password
		req.Now = time.Now()
		req.RandomString = util.GenerateID(5)

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.User = res.User

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
