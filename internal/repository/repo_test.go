package repository

import (
	"context"
	"testing"
	"weatherapi/v2/cmd/svr/config"
	"weatherapi/v2/external/models"
)

func TestRepository_GetGridInfo(t *testing.T) {
	type fields struct {
		Config *config.ServiceConfig
	}
	type args struct {
		ctx     context.Context
		request models.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Happy path",
			fields: fields{
				Config: &config.ServiceConfig{
					Name:    "WeatherApi",
					URL:     "https://api.weather.gov/gridpoints/TOP/32,81/forecast",
					Timeout: 60,
				},
			},
			args: args{
				ctx: context.Background(),
				request: models.Request{
					Longitude: "15",
					Latitude:  "-95",
				},
			},
			want:    "", // Replace with the expected result
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Config: tt.fields.Config,
			}
			got, err := r.GetGridInfo(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGridInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetGridInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
