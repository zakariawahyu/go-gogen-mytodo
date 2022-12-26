package runuserlogin

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/shared/infrastructure/token"
)

type runuserloginInteractor struct {
	outport Outport
	token   token.JWTToken
}

func NewUsecase(outputPort Outport, jwtToken token.JWTToken) Inport {
	return &runuserloginInteractor{
		outport: outputPort,
		token:   jwtToken,
	}
}

func (r *runuserloginInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	obj, err := entity.NewLoginUser(req.UserLoginRequest)
	if err != nil {
		return nil, err
	}

	user, err := r.outport.FindUserByEmail(ctx, obj.Email)
	if err != nil {
		return nil, err
	}

	isValidPass := comparePassword(user.Password, []byte(obj.Password))
	if !isValidPass {
		return nil, fmt.Errorf("wrong email and password")
	}

	token, err := r.token.CreateToken([]byte(user.Email), time.Hour)
	if err != nil {
		return nil, err
	}

	res.Token = token

	return res, nil
}

func comparePassword(hashPass string, plainPass []byte) bool {
	byteHash := []byte(hashPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
