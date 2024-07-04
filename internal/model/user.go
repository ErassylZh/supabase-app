package model

import (
	"fmt"
)

type User struct {
}

func (u User) SchemaName() string {
	return fmt.Sprintf("auth")
}
