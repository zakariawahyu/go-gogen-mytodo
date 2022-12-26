package application

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/controller/userapi"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/gateway/withgorm"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/getalluser"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/getoneuser"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/runuserdelete"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/runuserlogin"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/runuserregister"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/usecase/runuserupdate"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/config"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/server"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/token"
)

type userapp struct{}

func NewUserapp() gogen.Runner {
	return &userapp{}
}

func (userapp) Run() error {

	const appName = "userapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := userapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		getalluser.NewUsecase(datasource),
		getoneuser.NewUsecase(datasource),
		runuserupdate.NewUsecase(datasource),
		runuserdelete.NewUsecase(datasource),
		runuserregister.NewUsecase(datasource),
		runuserlogin.NewUsecase(datasource, jwtToken),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
