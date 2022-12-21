package application

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/controller/restapi"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/gateway/withgorm"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/usecase/getalltodo"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/usecase/runtodocheck"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/usecase/runtodocreate"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/usecase/runtododelete"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/config"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/server"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := restapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
		runtododelete.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
