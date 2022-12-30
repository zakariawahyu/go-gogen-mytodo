package getprofile

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/entity"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Email string `json:"email"`
}

type InportResponse struct {
	*entity.User
}
