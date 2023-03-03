package entity

type Responses struct {
	Id             string `gorm:"primaryKey;column:id;"`
	Nik            string `gorm:"column:nik;"`
	ProjectId      string `gorm:"column:project_id;"`
	AuditorId      string `gorm:"column:auditor_id;"`
	RespondentName string `gorm:"column:respondent_name;"`
	Alamat         string `gorm:"column:alamat;"`
	IdKelu         int    `gorm:"column:idkelu;"`
	IdKeca         int    `gorm:"column:idkeca;"`
}

func (Responses) TableName() string {
	return "responses"
}

type Kecamatan struct {
	IdKeca   int    `gorm:"primaryKey;column:idkeca;"`
	IdKabu   int    `gorm:"column:idkabu;"`
	IdProp   int    `gorm:"column:idprop;"`
	KdKeca   string `gorm:"column:kdkeca;"`
	NamaKeca string `gorm:"column:nama_keca;"`
}

func (Kecamatan) TableName() string {
	return "master_kecamatan"
}

type Kelurahan struct {
	IdKelu   int    `gorm:"primaryKey;column:idkelu;"`
	IdKeca   int    `gorm:"column:idkeca;"`
	IdKabu   int    `gorm:"column:idkabu;"`
	IdProp   int    `gorm:"column:idprop;"`
	KdKelu   string `gorm:"column:kdkelu;"`
	NamaKelu string `gorm:"column:nama_kelu;"`
}

func (Kelurahan) TableName() string {
	return "master_kelurahan"
}
