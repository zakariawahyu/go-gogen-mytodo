package getalluser

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/repository"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	repository.FindAllUserFilterRequest
}

type InportResponse struct {
	Count int64
	Items []any
}
