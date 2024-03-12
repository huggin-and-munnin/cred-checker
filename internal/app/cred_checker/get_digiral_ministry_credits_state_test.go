package cred_checker

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hugin-and-munin/cred-checker/internal/app/cred_checker/mocks"
	cred_checker "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/cred-checker"
	"github.com/stretchr/testify/assert"
)

func Test_GetDigitalMinistryCreditsState(t *testing.T) {
	type dependencies struct {
		checkCredentialsUseCase CheckCredentialsUseCase
	}

	type dependecyProvider func(ctrl *gomock.Controller) dependencies

	type args struct {
		ctx context.Context
		req *cred_checker.GetDigitalMinistryCreditsStateRequest
	}

	type want struct {
		errChecker func(t *testing.T, actualErr error, args ...interface{})
		result     *cred_checker.GetDigitalMinistryCreditsStateResponse
	}

	type testCase struct {
		name              string
		dependecyProvider dependecyProvider
		args              args
		want              want
	}

	tests := []testCase{
		{
			name: "credited",
			dependecyProvider: func(ctrl *gomock.Controller) dependencies {
				mock := mocks.NewMockCheckCredentialsUseCase(ctrl)

				mock.EXPECT().CheckCredentials(gomock.Any(), "1111").Return(
					true,
					nil,
				)

				return dependencies{
					checkCredentialsUseCase: mock,
				}
			},
			args: args{
				ctx: context.Background(),
				req: &cred_checker.GetDigitalMinistryCreditsStateRequest{
					Inn: "1111",
				},
			},
			want: want{
				errChecker: func(t *testing.T, actualErr error, args ...interface{}) {
					assert.NoError(t, actualErr, args...)
				},
				result: &cred_checker.GetDigitalMinistryCreditsStateResponse{
					Inn:   "1111",
					State: cred_checker.CreditState_CREDITED,
				},
			},
		},

		{
			name: "not credited",
			dependecyProvider: func(ctrl *gomock.Controller) dependencies {
				mock := mocks.NewMockCheckCredentialsUseCase(ctrl)

				mock.EXPECT().CheckCredentials(gomock.Any(), "1111").Return(
					false,
					nil,
				)

				return dependencies{
					checkCredentialsUseCase: mock,
				}
			},
			args: args{
				ctx: context.Background(),
				req: &cred_checker.GetDigitalMinistryCreditsStateRequest{
					Inn: "1111",
				},
			},
			want: want{
				errChecker: func(t *testing.T, actualErr error, args ...interface{}) {
					assert.NoError(t, actualErr, args...)
				},
				result: &cred_checker.GetDigitalMinistryCreditsStateResponse{
					Inn:   "1111",
					State: cred_checker.CreditState_NOT_CREDITED,
				},
			},
		},

		{
			name: "err",
			dependecyProvider: func(ctrl *gomock.Controller) dependencies {
				mock := mocks.NewMockCheckCredentialsUseCase(ctrl)

				mock.EXPECT().CheckCredentials(gomock.Any(), "1111").Return(
					false,
					errors.New("some err"),
				)

				return dependencies{
					checkCredentialsUseCase: mock,
				}
			},
			args: args{
				ctx: context.Background(),
				req: &cred_checker.GetDigitalMinistryCreditsStateRequest{
					Inn: "1111",
				},
			},
			want: want{
				errChecker: func(t *testing.T, actualErr error, args ...interface{}) {
					assert.ErrorContains(t, actualErr, "error getting credential state: some err", args)
				},
				result: nil,
			},
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			deps := tt.dependecyProvider(ctrl)

			service := &Implementation{
				checkCredentialsUseCase: deps.checkCredentialsUseCase,
			}

			result, err := service.GetDigitalMinistryCreditsState(tt.args.ctx, tt.args.req)
			tt.want.errChecker(t, err)

			if !cmp.Equal(result, tt.want.result, cmpopts.IgnoreUnexported(cred_checker.GetDigitalMinistryCreditsStateResponse{})) {
				t.Errorf("response is different from expected: %v", cmp.Diff(result, tt.want.result, cmpopts.IgnoreUnexported(cred_checker.GetDigitalMinistryCreditsStateResponse{})))
			}
		})
	}
}
