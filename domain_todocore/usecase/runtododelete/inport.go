package runtododelete

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/vo"
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID `uri:"todo_id"`
}

type InportResponse struct {
	Message string `json:"message"`
}
