package models

type DbMchtInfo struct {
	MchtCd       string `gorm:"type:varchar(15);primary_key"`
	MchtNm       string `gorm:"type:varchar(40)"`
	MccCd18      string `gorm:"type:varchar(4);column:mcc_cd_18"`
	BankBelongCd string `gorm:"type:varchar(11)"`
	AipBranCd    string `gorm:"type:varchar(11)"`
	UcBcCd32     string `gorm:"type:varchar(11);column:uc_bc_cd_32"`
	DbBase
}

func (d DbMchtInfo) TableName() string {
	return "mcht_infos"
}
