package vo

import (
	"fmt"
	"time"
)

type UserID string

func NewUserID(randomStringID string, now time.Time) (UserID, error) {
	var obj = UserID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))
	return obj, nil
}

func (r UserID) String() string {
	return string(r)
}
