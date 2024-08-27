package config

import (
	"fmt"
	"net/http"
	"time"
)

func (s *ServiceConfig) InitHTTPClient() {
	s.Client = &http.Client{
		Timeout: time.Duration(s.Timeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:    100,
			IdleConnTimeout: 90,
		},
	}
}

func (s *ServiceConfig) ApplyMergedConfig() error {
	if s.mergedComponentConfigs == nil {
		return fmt.Errorf("no merged component configs available")
	}

	return nil
}
