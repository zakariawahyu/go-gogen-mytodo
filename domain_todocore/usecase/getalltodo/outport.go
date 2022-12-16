package getalltodo

import "zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/repository"

type Outport interface {
	repository.FindAllTodoRepo
}
