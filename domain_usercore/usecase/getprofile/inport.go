package getprofile

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/vo"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Email string `json:"email"`
}

type InportResponse struct {
	ID     vo.UserID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Status bool      `json:"status"`
}
