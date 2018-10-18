package models

import (
	"encoding/json"
	"log"
 )

type TransParams struct {
	commonParams
	Term_key_inf *KeyInfo      `json:"term_key_info,omitempty"`
	Term_inf     *TermInfo     `json:"term_inf,omitempty"`

	CustomerInfo      string           `json:"customer_info,omitempty"` //银行卡验证信息
	Customer_info     *CustomerInfo    `json:"-"`
	Risk_info         *RiskInfo        `json:"risk_info,omitempty"`
	Instal_trans_info *InstalTransInfo `json:"instal_trans_info,omitempty"`
	CardTransData     string           `json:"card_trans_data,omitempty"` //有卡交易信息
	Card_Trans_data   *CardTransData   `json:"-"`
	//scanInfo
	Qr_code_info *QrCodeInfo `json:"qr_code_info,omitempty"`
	//app资源
	App_source_info []AppSourceInfo   `json:"app_source_info,omitempty"`
	App_params_info map[string]string `json:"app_params_info,omitempty"`
	//PBOC IC 参数
	IcParam_info []IcParam `json:"ic_param_info,omitempty"`
	//PBOC IC 公钥
	CaParam_info []CaParam `json:"ca_param_info,omitempty"`
	//商户信息
	Mcht_info *DbMchtInfo `json:"-"`
	//原交易信息
	Orig_sys_order_id string     `json:"orig_sys_order_id,omitempty"`
	Orig_term_seq     string     `json:"orig_term_seq,omitempty"`
	Orig_resp_cd      string     `json:"orig_resp_cd,omitempty"`
	Orig_resp_msg     string     `json:"orig_resp_msg,omitempty"`
	Orig_trans_dt     string     `json:"orig_trans_dt,omitempty"`
	Orig_tran_cd      string     `json:"orig_tran_cd,omitempty"`
	Orig_tran_nm      string     `json:"orig_tran_nm,omitempty"`
	Orig_tran_info    *DbTranLog `json:"-"`
	Cancel_flg        string     `json:"cancel_flg,omitempty"`
	Orig_prod_cd      string     `json:"orig_prod_cd,omitempty"`
	Orig_biz_cd       string     `json:"orig_biz_cd,omitempty"`
	//电子签名
	Sign_img string `json:"sign_img,omitempty"`

	//打印信息
	Ext_print_info string `json:"ext_print_info,omitempty"`
	//结算信息
	Clear_txn      string        `json:"clear_txn,omitempty"`
	//Tran_sett_list *TranSettList `json:"tran_sett_list,omitempty"`
	//明细查询信息
	//Tran_list []TranList `json:"tran_list,omitempty"`
	TranList      string       `json:"tran_list,omitempty"`
	//Tran_sel_info *TranSelInfo `json:"tran_sel_info,omitempty"`
	//app_sdk
	App_sdk_id  string `json:"app_sdk_id,omitempty"`
	App_sdk_key string `json:"app_sdk_key,omitempty"`
	//日志上送
	Log_info string `json:"log_info,omitempty"`
	//终端激活
	Bank_belong_cd string `json:"bank_belong_cd,omitempty"`
	Uc_bc_cd_32    string `json:"uc_bc_cd_32,omitempty"`
	Aip_bran_cd    string `json:"aip_bran_cd,omitempty"`
	//终端注销
	Reset_type       string            `json:"reset_type,omitempty"`
	Order_Info       string            `json:"orderInfo,omitempty"`
	//Order_info       *Order_info_total `json:"-"`
	Be_order_id      string            `json:"be_order_id,omitempty"`
	Be_biz_cd        string            `json:"be_biz_cd,omitempty"`
	Custom_ext_info  *CustomExtInfo    `json:"custom_ext_info,omitempty"`
	Tran_nm          string            `json:"tran_nm,omitempty"`
	Act_inf          string            `json:"act_inf,omitempty"`
	User_name        string            `json:"user_name,omitempty"`
	User_passwd      string            `json:"user_passwd,omitempty"`
	Orig_user_passwd string            `json:"orig_user_passwd,omitempty"`
	Nike_name        string            `json:"nike_name,omitempty"`
	//Opt_manage_list  []OptManagerList  `json:"opt_manage_list,omitempty"`
	DesMchntCd       string            `json:"des_mchnt_cd,omitempty"`
	DesTermCd        string            `json:"des_term_cd,omitempty"`
	Gift_info        string            `json:"gift_info,omitempty"`
	Card_tp          string            `json:"card_tp,omitempty"` //卡类型
	Ref_no           string            `json:"ref_no,omitempty"`  //和检索参考号
	Pos_sign         string            `json:"pos_sign,omitempty"`
	ForeignInfo      string            `json:"foreign_info,omitempty"` //外卡交易信息
	Foreign_info     *ForeignInfo      `json:"-"`
}

type TransMessage struct {
	MessageHead
	Msg_body     string                     `json:"msg_body"`
	MsgBody      *TransParams               `json:"-"`

}

func (t *TransMessage) ToString() string {
	t.SetMsgBody()
	res, err := json.Marshal(t)
	if err != nil {
		log.Println("TransMessage marshal fail:", err)
		return "{}"
	}
 	return string(res)
}

func (t *TransMessage) SetMsgBody() {
	btMsgBody, err := json.Marshal(t.MsgBody)
	if err != nil {
		log.Println("TransMessage marshal fail:", err)
		t.Msg_body = "{}"
		return
	}
	t.Msg_body = string(btMsgBody)
}
