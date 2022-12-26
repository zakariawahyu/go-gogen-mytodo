package runuserregister

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.UserRegisterRequest
}

type InportResponse struct {
	User *entity.User `json:"user"`
}
