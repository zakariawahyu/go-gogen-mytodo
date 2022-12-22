package gatewayuser

import (
	"context"

	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/repository"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/vo"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/config"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) FindAllUser(ctx context.Context, req repository.FindAllUserFilterRequest) ([]*entity.User, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneUserByID(ctx context.Context, userID vo.UserID) (*entity.User, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveUser(ctx context.Context, obj *entity.User) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteUser(ctx context.Context, userID vo.UserID) error {
	r.log.Info(ctx, "called")

	return nil
}
