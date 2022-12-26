package runuserlogin

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
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

	userObj, err := r.outport.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	isValidPass := comparePassword(userObj.Password, []byte(req.Password))
	if !isValidPass {
		return nil, fmt.Errorf("wrong email and password")
	}

	token, err := r.token.CreateToken([]byte(userObj.Email), time.Hour)
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
