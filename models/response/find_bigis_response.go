package response

import "github.com/tensuqiuwulu/be-service-data-bigis/models/entity"

type FindBigisResponse struct {
	Nik       string `json:"nik"`
	Name      string `json:"name"`
	Alamat    string `json:"alamat"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	IdKelu    int    `json:"idkelu"`
}

func ToFindBigisResponse(responses *entity.Responses, kelurahan string, kecamatan string) (bigisResponse FindBigisResponse) {
	bigisResponse.Nik = responses.Nik
	bigisResponse.Name = responses.RespondentName
	bigisResponse.Alamat = responses.Alamat
	bigisResponse.Kecamatan = kecamatan
	bigisResponse.Kelurahan = kelurahan
	bigisResponse.IdKelu = responses.IdKelu
	return bigisResponse
}
