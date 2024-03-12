package health

import (
	"context"

	"github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/health"
)

func (i *Implementation) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}
