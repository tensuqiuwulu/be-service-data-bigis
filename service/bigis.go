package service

import (
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-data-bigis/exceptions"
	"github.com/tensuqiuwulu/be-service-data-bigis/models/request"
	"github.com/tensuqiuwulu/be-service-data-bigis/models/response"
	"github.com/tensuqiuwulu/be-service-data-bigis/repository"
	"gorm.io/gorm"
)

type BigisServiceInterface interface {
	FindResponseByNik(requestId string, requestBigisResponse *request.FindBigisResponsesRequest) (responses response.FindBigisResponse)
}

type BigisServiceImplementation struct {
	DB                       *gorm.DB
	Validate                 *validator.Validate
	Logger                   *logrus.Logger
	BigisRepositoryInterface repository.BigisRepositoryInterface
}

func NewBigisService(
	db *gorm.DB,
	validate *validator.Validate,
	logger *logrus.Logger,
	bigisServiceInterface repository.BigisRepositoryInterface,
) BigisServiceInterface {
	return &BigisServiceImplementation{
		DB:                       db,
		Validate:                 validate,
		Logger:                   logger,
		BigisRepositoryInterface: bigisServiceInterface,
	}
}

func (service *BigisServiceImplementation) FindResponseByNik(requestId string, requestBigisResponse *request.FindBigisResponsesRequest) (responses response.FindBigisResponse) {
	var err error

	request.ValidateRequest(service.Validate, requestBigisResponse, requestId, service.Logger)

	// Check product if exist in cart
	responseResult, err := service.BigisRepositoryInterface.FindResponseByNik(service.DB, requestBigisResponse.Nik)
	exceptions.PanicIfError(err, requestId, service.Logger)

	kecamatanResult, err := service.BigisRepositoryInterface.FindKecamatanById(service.DB, responseResult.IdKeca)
	exceptions.PanicIfError(err, requestId, service.Logger)

	kelurahanResult, err := service.BigisRepositoryInterface.FindKelurahanById(service.DB, responseResult.IdKelu)
	exceptions.PanicIfError(err, requestId, service.Logger)
	// log.Println(dataValue.Value)

	responses = response.ToFindBigisResponse(responseResult, kecamatanResult.NamaKeca, kelurahanResult.NamaKelu)
	return responses
}
