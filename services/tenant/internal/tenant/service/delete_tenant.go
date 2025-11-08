package service

import (
	"context"
)

func (s *TenantService) Delete(ctx context.Context, id string) error {
	_, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
