package models

type DbBizCtl struct {
	Prod_cd    string `gorm:"type:varchar(10);primary_key"`
	Tran_cd    string `gorm:"type:varchar(10);primary_key"`
	Biz_cd     string `gorm:"type:varchar(10);primary_key"`
	Prod_nm    string `gorm:"type:varchar(50)"`
	Cancel_flg string `gorm:"type:varchar(10)"`
	Revel_flg  string `gorm:"type:varchar(10)"`
	DbBase
}

func (t DbBizCtl) TableName() string {
	return "biz_ctl"
}
