package custom

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func IsRecordFoundError(err error) bool {
	fmt.Println(err)
	fmt.Println(gorm.ErrRecordNotFound)
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value")
}
