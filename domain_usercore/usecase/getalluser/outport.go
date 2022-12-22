package getalluser

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/repository"
)

type Outport interface {
	repository.FindAllUserRepo
}
