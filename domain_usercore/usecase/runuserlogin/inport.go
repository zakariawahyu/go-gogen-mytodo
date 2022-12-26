package runuserlogin

import (
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InportResponse struct {
	Token string `json:"token"`
}
