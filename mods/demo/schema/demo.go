package schema

type Demo struct {
	ID   string `json:"id" gorm:"size:20;primarykey;autoIncrement"`
	Code string `json:"code" gorm:"size:32;index;"`
	Name string `json:"name" gorm:"size:128;index"`
}

func (a *Demo) TableName() string {
	return "sst.demo"
}

type DemoForm struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
