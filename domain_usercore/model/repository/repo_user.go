package repository

import (
	"context"

	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/vo"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}

type FindAllUserFilterRequest struct {
	Page int64
	Size int64
	// add other field to filter here ...
}

type FindAllUserRepo interface {
	FindAllUser(ctx context.Context, req FindAllUserFilterRequest) ([]*entity.User, int64, error)
}

type DeleteUserRepo interface {
	DeleteUser(ctx context.Context, userID vo.UserID) error
}

type FindOneUserByIDRepo interface {
	FindOneUserByID(ctx context.Context, userID vo.UserID) (*entity.User, error)
}

type FindUserByEmailRepo interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
