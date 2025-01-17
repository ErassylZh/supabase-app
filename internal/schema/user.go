package schema

import (
	"errors"
	"regexp"
)

type UserUpdate struct {
	UserID   string `json:"-"`
	Nickname string `json:"nickname"`
}

func (u *UserUpdate) Validate() error {
	var nicknameRegex = regexp.MustCompile(`^(?i)(?=.{1,10}$)(?!.*\b(лох|чорт|черт|нигер|негр)\b).+$`)

	if !nicknameRegex.MatchString(u.Nickname) {
		return errors.New("nickname is invalid: must be up to 10 characters and not contain prohibited words")
	}

	return nil
}
