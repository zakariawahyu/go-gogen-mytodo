package runuserregister

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/repository"
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/service"
)

type Outport interface {
	repository.SaveUserRepo
	service.HashAndSaltPasswordServices
}
