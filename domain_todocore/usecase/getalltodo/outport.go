package getalltodo

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_todocore/model/repository"
	repository2 "zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/repository"
)

type Outport interface {
	repository.FindAllTodoByUserIDRepo
}

type Outport2 interface {
	repository2.FindUserByEmailRepo
}
