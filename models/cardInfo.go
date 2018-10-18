package models

type CustomerInfo struct {
	Cert_tp     string `json:"cert_tp,omitempty"`
	Cert_id     string `json:"cert_id,omitempty"`
	Customer_nm string `json:"customer_nm,omitempty"`
	Phone_no    string `json:"phone_no,omitempty"`
	Sms_code    string `json:"sms_code,omitempty"`
	Pin         string `json:"pin,omitempty"`
	Cvn2        string `json:"cvn2,omitempty"`
	Expired     string `json:"expired,omitempty"`
}

type RiskInfo struct {
}

type InstalTransInfo struct {
	Instal_num       string `json:"instal_num,omitempty"`
	Instal_rate      string `json:"instal_rate,omitempty"`
	Mcht_fee_subsidy string `json:"mcht_fee_subsidy,omitempty"`
}

type CardTransData struct {
	Ic_card_data   string `json:"ic_card_data,omitempty"`
	Ic_card_seq    string `json:"ic_card_seq,omitempty"`
	Track2         string `json:"track2,omitempty"`
	Track3         string `json:"track3,omitempty"`
	Pos_entry_cd   string `json:"pos_entry_cd,omitempty"`
	Term_entry_cap string `json:"term_entry_cap,omitempty"`
	Chip_cond_code string `json:"chip_cond_code,omitempty"`
	TransSendMode  string `json:"transSendMode,omitempty"`
	Pos_pin_cd     string `json:"pos_pin_cd,omitempty"`
	Sec_pin_inf    string `json:"sec_pin_inf,omitempty"`
}

type CardAmtInfo struct {
	Acc_tp       string `json:"acc_tp,omitempty"`
	Balace_tp    string `json:"balace_tp,omitempty"`
	Curr_cd      string `json:"curr_cd,omitempty"`
	Balance_sign string `json:"balance_sign,omitempty"`
	Balance      string `json:"balance,omitempty"`
}

type CaParam struct {
	PBOCIdX      string `json:"PBOCIDX,omitempty"`
	RID_TAG      string `json:"9F06,omitempty"`
	RIDEX_TAG    string `json:"9F22,omitempty"`
	DEADLINE_TAG string `json:"DF05,omitempty"`
	HSALID_TAG   string `json:"DF06,omitempty"`
	CAALID_TAG   string `json:"DF07,omitempty"`
	MODULUS_TAG  string `json:"DF02,omitempty"`
	EXPONENT_TAG string `json:"DF04,omitempty"`
	CHKVAL_TAG   string `json:"DF03,omitempty"`
}

type IcParam struct {
	PBOCIdX string `json:"PBOCIDX,omitempty"`
	A_ID    string `json:"9F06,omitempty"`
	A_IDX   string `json:"DF01,omitempty"`
	VERSION string `json:"9F08,omitempty"`
	IC_DF11 string `json:"DF11,omitempty"`
	IC_DF12 string `json:"DF12,omitempty"`
	IC_DF13 string `json:"DF13,omitempty"`
	IC_9F1B string `json:"9F1B,omitempty"`
	IC_DF15 string `json:"DF15,omitempty"`
	IC_DF16 string `json:"DF16,omitempty"`
	IC_DF17 string `json:"DF17,omitempty"`
	IC_DF14 string `json:"DF14,omitempty"`
	IC_DF18 string `json:"DF18,omitempty"`
	IC_9F7B string `json:"9F7B,omitempty"`
	IC_DF19 string `json:"DF19,omitempty"`
	IC_DF20 string `json:"DF20,omitempty"`
	IC_DF21 string `json:"DF21,omitempty"`
}

type ForeignInfo struct {
	Foreign_flag          string `json:"foreign_flag,omitempty"`
	Down_ic_param_flag    string `json:"down_ic_param_flag,omitempty"`
	Foreign_amt           string `json:"foreign_amt,omitempty"`
	Foreign_rate          string `json:"foreign_rate,omitempty"`
	Foreign_curr_cd       string `json:"foreign_curr_cd,omitempty"`
	Foreign_sys_order_id  string `json:"foreign_sys_order_id,omitempty"`
	Foreign_term_seq      string `json:"foreign_term_seq,omitempty"`
	Foreign_tran_dt_tm    string `json:"foreign_tran_dt_tm,omitempty"`
	Foreign_acq_ins_id_cd string `json:"foreign_acq_ins_id_cd,omitempty"`
	Foreign_fwd_ins_id_cd string `json:"foreign_fwd_ins_id_cd,omitempty"`
}
