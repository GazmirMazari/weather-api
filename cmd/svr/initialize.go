package main

import "context"

func NewService(ctx context.Context, cfg *config.Config) (services routes.Handler, err error) {
	//init api client here

	return routes.Handler{
		Config: cfg,
	}, nil
}
