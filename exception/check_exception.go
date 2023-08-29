package exception

import (
	"errors"

	"github.com/tae2089/devops-platform-backend/ent"
)

func IsEntityNotFound(err error) error {
	if errors.Is(err, &ent.NotFoundError{}) {
		return NotFoundEntity
	}
	return err
}

func IsAdminRole(role string) error {
	if role == "admin" {
		return nil
	}
	return Forbidden
}
