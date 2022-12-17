package withgorm

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(entity.Todo{})
	if err != nil {
		return nil
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		db:      db,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	var todoObjs []*entity.Todo
	var count int64

	if err := r.db.
		Model(entity.Todo{}).
		Count(&count).
		Limit(size).
		Offset((page - 1) * size).
		Find(&todoObjs).Error; err != nil {
		return nil, 0, err
	}

	return todoObjs, count, nil
}

func (r *gateway) FindOneTodoById(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	var todoObj entity.Todo

	if err := r.db.First(&todoObj, "id = ? ", todoID).Error; err != nil {
		return nil, err
	}

	return &todoObj, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	if err := r.db.Save(obj).Error; err != nil {
		return err
	}

	return nil
}
