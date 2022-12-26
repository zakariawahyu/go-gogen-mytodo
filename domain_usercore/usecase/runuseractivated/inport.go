package runuseractivated

import "zakariawahyu.com/go-gogen-mytodo/shared/gogen"

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Email           string `json:"email"`
	ActivationToken string `json:"token"`
}

type InportResponse struct {
	Message string `json:"message"`
}
