package runusercreate

import (
	"zakariawahyu.com/go-gogen-mytodo/domain_usercore/model/repository"
)

type Outport interface {
	repository.SaveUserRepo
}
