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

type DeleteTodoByIdRepo interface {
	DeleteTodoById(ctx context.Context, todoID vo.TodoID) error
}

type FindAllTodoByUserIDRepo interface {
	FindAllTodoByUserID(ctx context.Context, page, size int, userID string) ([]*entity.Todo, int64, error)
}
