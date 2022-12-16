package runtodocreate

import "zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/repository"

type Outport interface {
	repository.SaveTodoRepo
}
