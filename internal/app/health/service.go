package health

import "github.com/hugin-and-munin/cred-checker/pb/github.com/hugin-and-munin/health"

type Implementation struct {
	health.UnimplementedHealthServer
}

func NewHealthProbe() health.HealthServer {
	return &Implementation{}
}
