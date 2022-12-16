package runtodocheck

import "zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/repository"

type Outport interface {
	repository.FindOneTodoByIdRepo
	repository.SaveTodoRepo
}
