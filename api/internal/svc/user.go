package svc

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

func (s *ServiceContext) CheckIfUsernameExist(ctx context.Context, userName string) (bool, error) {
	user, err := s.UserModel.FindOneByUsername(ctx, userName)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	if user.Id > 0 {
		return true, nil
	}
	return false, nil
}
