package getprofile

import (
	"context"
	"fmt"
)

type getprofileInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getprofileInteractor{
		outport: outputPort,
	}
}

func (r *getprofileInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	existingUser, err := r.outport.FindAllTodoByUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	fmt.Println(existingUser)

	res.User = existingUser

	return res, nil
}
