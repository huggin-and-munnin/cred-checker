package app

import (
	"context"
	"net/http"

	"github.com/hugin-and-munin/cred-checker/internal/details/cred_checkers"
	"github.com/hugin-and-munin/cred-checker/internal/model"
	"github.com/hugin-and-munin/cred-checker/internal/use_cases/get_credentials"
	cred_checker "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/cred-checker"
)

type CheckCredentialsUseCase interface {
	CheckCredentials(ctx context.Context, inn string) (*model.Company, error)
}

type Implementation struct {
	checkCredentialsUseCase CheckCredentialsUseCase

	cred_checker.UnimplementedCredCheckerServer
}

func NewCredChecker(httpClient *http.Client) cred_checker.CredCheckerServer {
	return &Implementation{
		checkCredentialsUseCase: get_credentials.NewUseCase(
			cred_checkers.NewGosuslugiCredsChecker(httpClient),
		),
	}
}
