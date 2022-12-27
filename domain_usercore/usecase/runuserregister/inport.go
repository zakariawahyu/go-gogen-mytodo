package runuserregister

import (
	"time"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/vo"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.UserRequest
}

type InportResponse struct {
	ID              vo.UserID `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Status          bool      `json:"status"`
	ActivationToken string    `json:"activation_token"`
	Created         time.Time `json:"created"`
}
