package repository

import (
	"context"
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/vo"

	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/entity"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FindAllTodoRepo interface {
	FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error)
}

type FindOneTodoByIdRepo interface {
	FindOneTodoById(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error)
}
