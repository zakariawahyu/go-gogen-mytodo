package getalltodo

import (
	"zakariawahyu.com/go-gogen-mytodo/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Page int
	Size int
}

type InportResponse struct {
	Count int64
	Items []any
}
