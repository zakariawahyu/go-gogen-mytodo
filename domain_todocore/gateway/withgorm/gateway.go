package withgorm

import (
	"context"

	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/vo"
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

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneTodoById(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	return nil
}
