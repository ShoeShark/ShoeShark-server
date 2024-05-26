package schema

type Contracts struct {
	ContractName    string `gorm:"type:varchar(60);not null;primary_key"`
	ContractAddress string `gorm:"type:varchar(42);not null"`
	IsDelete        bool   `gorm:"default:false"`
}

func (Contracts) TableName() string {
	return "sst.sst_contracts"
}
