package userapi

import (
	"github.com/gin-gonic/gin"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/config"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/token"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log      logger.Logger
	cfg      *config.Config
	jwtToken token.JWTToken
}

func NewGinController(log logger.Logger, cfg *config.Config, tk token.JWTToken) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
		jwtToken:       tk,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {
	resource := router.Group("/api/v1")
	resource.GET("/user", r.authentication(), r.getAllUserHandler())
	resource.GET("/user/:user_id", r.authentication(), r.getOneUserHandler())
	resource.PUT("/user/:user_id", r.authentication(), r.runUserUpdateHandler())
	resource.DELETE("/user/:user_id", r.authentication(), r.runUserDeleteHandler())
	resource.POST("/register", r.runuserregisterHandler())
	resource.POST("/login", r.runuserloginHandler())
	resource.GET("/activated/:email/:token", r.runuseractivatedHandler())
	resource.GET("/profile", r.authentication(), r.getprofileHandler())
}
