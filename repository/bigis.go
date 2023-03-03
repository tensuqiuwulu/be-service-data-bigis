package repository

import (
	"github.com/tensuqiuwulu/be-service-data-bigis/config"
	"github.com/tensuqiuwulu/be-service-data-bigis/models/entity"
	"gorm.io/gorm"
)

type BigisRepositoryInterface interface {
	FindResponseByNik(db *gorm.DB, nik string) (*entity.Responses, error)
	FindKelurahanById(db *gorm.DB, idKelu int) (*entity.Kelurahan, error)
	FindKecamatanById(db *gorm.DB, idKeca int) (*entity.Kecamatan, error)
}

type BigisRepositoryImplementation struct {
	DB *config.Database
}

func NewBigisRepository(db *config.Database) BigisRepositoryInterface {
	return &BigisRepositoryImplementation{
		DB: db,
	}
}

func (service *BigisRepositoryImplementation) FindResponseByNik(db *gorm.DB, nik string) (*entity.Responses, error) {
	response := &entity.Responses{}
	result := db.Where("nik = ?", nik).Find(&response)
	return response, result.Error
}

func (service *BigisRepositoryImplementation) FindKelurahanById(db *gorm.DB, idKelu int) (*entity.Kelurahan, error) {
	kelurahan := &entity.Kelurahan{}
	result := db.Where("idkelu = ?", idKelu).Find(&kelurahan)
	return kelurahan, result.Error
}

func (service *BigisRepositoryImplementation) FindKecamatanById(db *gorm.DB, idKeca int) (*entity.Kecamatan, error) {
	kecamatan := &entity.Kecamatan{}
	result := db.Where("idkeca = ?", idKeca).Find(&kecamatan)
	return kecamatan, result.Error
}
