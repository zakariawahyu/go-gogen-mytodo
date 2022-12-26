package service

import "context"

type HashAndSaltPasswordServices interface {
	HashAndSaltPassword(ctx context.Context, password []byte) (string, error)
}

type ComparePasswordServices interface {
	ComparePassword(ctx context.Context, hashPass string, plainPass []byte) bool
}
