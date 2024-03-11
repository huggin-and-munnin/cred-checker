package get_credentials

import (
	"context"
	"errors"
	"fmt"

	"github.com/hugin-and-munin/cred-checker/internal/details/cred_checkers"
	"github.com/hugin-and-munin/cred-checker/internal/model"
)

type CredentialChecker interface {
	SearchCompany(ctx context.Context, inn string) (*model.Company, error)
}

type useCase struct {
	credentialChecker CredentialChecker
}

func NewUseCase(credentialChecker CredentialChecker) *useCase {
	return &useCase{
		credentialChecker: credentialChecker,
	}
}

func (u *useCase) CheckCredentials(ctx context.Context, inn string) (bool, error) {
	_, err := u.credentialChecker.SearchCompany(ctx, inn)

	if errors.Is(err, cred_checkers.ErrNotFound) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("search for compeny failed: %w", err)
	}

	return true, err
}
