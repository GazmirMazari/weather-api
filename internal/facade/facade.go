package facade

import "context"

import "context"
      "weatherapi/v2/external/models"


type ServiceI interface {
	GetWeatherData(ctx context.Context, request models.Request) (response models.WeatherApiRes, err error)
}

func GetWeatherData(ctx context.Context, request Request) (response models.Response, err error) {
	//TODO implement me
	//vadlidate request
	err = validateRequest(request); err != nil{
		log.Errorf("validate request error: %v", err)
		return response, err
	}

	response, err = r.Repository.SearchWeatherApi(ctx, request)
	if err != nil {
		log.Errorf("GetWeatherApiData error: %v", err)
		return response, err
	}

	//mapping here


	//TODO implement me
	panic("implement me")
	return response, err

}


func validateRequest(request models.Request) (err error) {
	//TODO implement me
	panic("implement me")
}
