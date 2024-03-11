package app

import (
	"context"
	"fmt"

	cred_checker "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/cred-checker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetDigitalMinistryCreditsState implements cred_checker.CredCheckerServer.
func (i *Implementation) GetDigitalMinistryCreditsState(ctx context.Context, req *cred_checker.GetDigitalMinistryCreditsStateRequest) (*cred_checker.GetDigitalMinistryCreditsStateResponse, error) {

	result, err := i.checkCredentialsUseCase.CheckCredentials(ctx, req.GetInn())

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error getting credential state: %v", err))
	}

	state := cred_checker.CreditState_NOT_CREDITED
	if result {
		state = cred_checker.CreditState_CREDITED
	}

	return &cred_checker.GetDigitalMinistryCreditsStateResponse{
		Inn:   req.GetInn(),
		State: state,
	}, nil
}
