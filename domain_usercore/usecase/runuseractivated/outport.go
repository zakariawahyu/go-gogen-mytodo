package runuseractivated

import "zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/repository"

type Outport interface {
	repository.FindUserByEmailRepo
	repository.SaveUserRepo
}
