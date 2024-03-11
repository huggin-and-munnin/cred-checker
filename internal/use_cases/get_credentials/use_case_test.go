package get_credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hugin-and-munin/cred-checker/internal/details/cred_checkers"
	"github.com/hugin-and-munin/cred-checker/internal/model"
	"github.com/hugin-and-munin/cred-checker/internal/use_cases/get_credentials/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_CheckCredentials(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type dependencies struct {
		credentialChecker CredentialChecker
	}

	type dependecyProvider func(ctrl *gomock.Controller) dependencies

	type args struct {
		ctx context.Context
		inn string
	}

	type want struct {
		errChecker func(t *testing.T, actualErr error, args ...interface{})
		result     bool
	}

	type testCase struct {
		name              string
		dependecyProvider dependecyProvider
		args              args
		want              want
	}

	tests := []testCase{
		{
			name: "has creds",
			dependecyProvider: func(ctrl *gomock.Controller) dependencies {
				mock := mocks.NewMockCredentialChecker(ctrl)

				mock.EXPECT().SearchCompany(gomock.Any(), "1111").Return(
					&model.Company{},
					nil,
				)

				return dependencies{
					credentialChecker: mock,
				}
			},
			args: args{
				ctx: ctx,
				inn: "1111",
			},
			want: want{
				errChecker: func(t *testing.T, actualErr error, args ...interface{}) {
					assert.NoError(t, actualErr, args...)
				},
				result: true,
			},
		},

		{
			name: "has NO creds",
			dependecyProvider: func(ctrl *gomock.Controller) dependencies {
				mock := mocks.NewMockCredentialChecker(ctrl)

				mock.EXPECT().SearchCompany(gomock.Any(), "1111").Return(
					nil,
					cred_checkers.ErrNotFound,
				)

				return dependencies{
					credentialChecker: mock,
				}
			},
			args: args{
				ctx: ctx,
				inn: "1111",
			},
			want: want{
				errChecker: func(t *testing.T, actualErr error, args ...interface{}) {
					assert.NoError(t, actualErr, args...)
				},
				result: false,
			},
		},

		{
			name: "cred checker returns error",
			dependecyProvider: func(ctrl *gomock.Controller) dependencies {
				mock := mocks.NewMockCredentialChecker(ctrl)

				mock.EXPECT().SearchCompany(gomock.Any(), "1111").Return(
					nil,
					errors.New("gosuslugi error"),
				)

				return dependencies{
					credentialChecker: mock,
				}
			},
			args: args{
				ctx: ctx,
				inn: "1111",
			},
			want: want{
				errChecker: func(t *testing.T, actualErr error, args ...interface{}) {
					assert.ErrorContains(t, actualErr, "gosuslugi error")
					assert.ErrorContains(t, actualErr, "search for compeny failed:")
				},
				result: false,
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

			usecase := NewUseCase(deps.credentialChecker)

			result, err := usecase.CheckCredentials(tt.args.ctx, tt.args.inn)
			tt.want.errChecker(t, err)

			assert.Equalf(
				t,
				tt.want.result, result,
				"result is different from expected",
			)
		})
	}
}
