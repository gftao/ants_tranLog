package models

import (
	"time"
	"golib/modules/logr"
	"golib/modules/gormdb"
	"golib/defs"
	"tranLog/models/sutil"
	"tranLog/godefs"
)

type DbTranLog struct {
	Trans_dt          string `gorm:"type:varchar(8);index:idx_trans_dt"`
	Tran_cd           string `gorm:"type:varchar(4)"`
	Tran_nm           string `gorm:"type:varchar(20)"`
	Prod_cd           string `gorm:"type:varchar(4)"`
	Biz_cd            string `gorm:"type:varchar(7)"`
	Mcht_cd           string `gorm:"type:varchar(15);unique_index:uidx_term_seq;primary_key"`
	Mcht_nm           string `gorm:"type:text(40)"`
	Term_id           string `gorm:"type:varchar(8);unique_index:uidx_term_seq;primary_key"`
	Term_seq          string `gorm:"type:varchar(6);unique_index:uidx_term_seq"`
	Term_batch        string `gorm:"type:varchar(6);unique_index:uidx_term_seq"`
	Order_id          string `gorm:"type:varchar(40);primary_key"`
	Order_desc        string `gorm:"type:text(32)"`
	Tran_dt_tm        string `gorm:"type:varchar(14)"`
	Resp_cd           string `gorm:"type:varchar(10)"`
	Resp_msg          string `gorm:"type:varchar(200)"`
	Trans_st          string `gorm:"type:varchar(10)"`
	Cld_order_id      string `gorm:"type:varchar(14);unique_index:uidx_cld_order_id"`
	Sys_order_id      string `gorm:"type:varchar(14);index:idx_sys_order_id"`
	Orig_sys_order_id string `gorm:"type:varchar(14)"`
	Orig_trans_dt     string `gorm:"type:varchar(8)"`
	Order_timeout     string `gorm:"type:varchar(10)"`
	Pre_auth_id       string `gorm:"type:varchar(6)"`
	Curr_cd           string `gorm:"type:varchar(3)"`
	Pri_acct_no       string `gorm:"type:varchar(30)"`
	Tran_amt          string `gorm:"type:varchar(12)"`
	Sett_dt           string `gorm:"type:varchar(8)"`
	Check_flg         string `gorm:"type:varchar(8)"`
	Acq_ins_id_cd     string `gorm:"type:varchar(11)"`
	Fwd_ins_id_cd     string `gorm:"type:varchar(11)"`
	Iss_ins_id_cd     string `gorm:"type:varchar(20)"`
	Sign_img          string `gorm:"type:varchar(3000)"`
	// order_cust_info   行业订单信息
	// customer_info    银行卡验证信息
	Cert_tp     string `gorm:"type:varchar(2)"`
	Cert_id     string `gorm:"type:text(20)"`
	Customer_nm string `gorm:"type:text(32)"`
	Phone_no    string `gorm:"type:varchar(20)"`
	Sms_code    string `gorm:"type:varchar(6)"`
	// card_trans_data   有卡交易域
	Pos_entry_cd   string `gorm:"type:varchar(3)"`
	Term_entry_cap string `gorm:"type:text(2)"`
	//instal_trans_info  分期付款信息
	Instal_num       string `gorm:"type:varchar(2)"`
	Instal_rate      string `gorm:"type:varchar(6)"`
	Mcht_fee_subsidy string `gorm:"type:varchar(6)"`
	// qr_code_info   扫码信息域
	Auth_code    string `gorm:"type:text(130)"`
	Qr_type      string `gorm:"type:text(20)"`
	Time_out     string `gorm:"type:varchar(3)"`
	Buyer_user   string `gorm:"type:text(30)"`
	Ins_order_id string `gorm:"type:text(70)"`
	Goods_id     string `gorm:"type:text(40)"`
	Goods_nm     string `gorm:"type:text(300)"`
	Goods_num    string `gorm:"type:text(10)"`
	Goods_price  string `gorm:"type:text(12)"`
	//  转入账号
	Trans_in_acct_no string `gorm:"type:varchar(30)"`
	//终端相关信息
	Ip_addr  string `gorm:"type:text(130)"`
	Gps_addr string `gorm:"type:text(64)"`
	//交易信息
	Cancel_flg string `gorm:"type:varchar(10)"`
	//Cust_order_id string `gorm:"type:varchar(100)"`
	Order_stat       string `gorm:"type:varchar(10)"`
	Order_info       string `gorm:"type:varchar(200)"`
	Be_order_id      string `gorm:"type:varchar(100)"`
	Order_notiy_info string `gorm:"type:varchar(100)"`
	//目的渠道信息
	DesAcqInsId string `gorm:"type:varchar(100)"`
	DesMchntCd  string `gorm:"type:varchar(100)"`
	CustomExtInfo
	Be_biz_cd   string `gorm:"type:varchar(7)"`
	DbBase
}

func (t DbTranLog) TableName() string {
	return "tran_rang_logs"
}

func (t *DbTranLog) InitByTransMsg(tmsg TransMessage) {
	t.Trans_dt = time.Now().Format("20060102")
	t.Tran_cd = tmsg.MsgBody.Tran_cd
	t.Prod_cd = tmsg.MsgBody.Prod_cd
	t.Biz_cd = tmsg.MsgBody.Biz_cd
	t.Mcht_cd = tmsg.MsgBody.Mcht_cd
	t.Mcht_nm = tmsg.MsgBody.Mcht_nm
	t.Term_id = tmsg.MsgBody.Term_id
	t.Term_seq = tmsg.MsgBody.Term_seq
	t.Term_batch = tmsg.MsgBody.Term_batch

	t.Order_id = tmsg.MsgBody.Order_id
	t.Order_desc = tmsg.MsgBody.Order_desc
	t.Tran_dt_tm = tmsg.MsgBody.Tran_dt_tm
	t.Cld_order_id = tmsg.MsgBody.Cld_order_id
	if tmsg.MsgBody.Tran_order_id != "" {
		t.Orig_sys_order_id = tmsg.MsgBody.Orig_tran_info.Sys_order_id
	} else {
		t.Orig_sys_order_id = tmsg.MsgBody.Orig_sys_order_id
	}
	t.Order_timeout = tmsg.MsgBody.Order_timeout
	t.Pre_auth_id = tmsg.MsgBody.Pre_auth_id
	t.Curr_cd = tmsg.MsgBody.Curr_cd
	//t.Pri_acct_no = tmsg.MsgBody.Pri_acct_no
	t.Tran_amt = tmsg.MsgBody.Tran_amt
	if len(t.Cld_order_id) >= 8 {
		t.Sett_dt = t.Cld_order_id[:8]
	} else if len(t.Order_id) >= 8 {
		t.Sett_dt = t.Order_id[:8]
	}

	t.Trans_in_acct_no = tmsg.MsgBody.Trans_in_acct_no
	if tmsg.MsgBody.Mcht_info != nil {
		t.Acq_ins_id_cd = tmsg.MsgBody.Mcht_info.UcBcCd32
		t.Fwd_ins_id_cd = tmsg.MsgBody.Mcht_info.BankBelongCd
	}
	if tmsg.MsgBody.Customer_info != nil {
		t.Cert_tp = tmsg.MsgBody.Customer_info.Cert_tp
		t.Cert_id = tmsg.MsgBody.Customer_info.Cert_id
		t.Customer_nm = tmsg.MsgBody.Customer_info.Customer_nm
		t.Phone_no = tmsg.MsgBody.Customer_info.Phone_no
		t.Sms_code = tmsg.MsgBody.Customer_info.Sms_code
	}
	if tmsg.MsgBody.Card_Trans_data != nil {
		t.Pos_entry_cd = tmsg.MsgBody.Card_Trans_data.Pos_entry_cd
		t.Term_entry_cap = tmsg.MsgBody.Card_Trans_data.Term_entry_cap
	}
	if tmsg.MsgBody.Instal_trans_info != nil {
		t.Instal_num = tmsg.MsgBody.Instal_trans_info.Instal_num
		t.Instal_rate = tmsg.MsgBody.Instal_trans_info.Instal_rate
		t.Mcht_fee_subsidy = tmsg.MsgBody.Instal_trans_info.Mcht_fee_subsidy
	}
	if tmsg.MsgBody.Qr_code_info != nil {
		t.Auth_code = tmsg.MsgBody.Qr_code_info.Auth_code
		t.Qr_type = tmsg.MsgBody.Qr_code_info.Qr_type
		t.Time_out = tmsg.MsgBody.Qr_code_info.Time_out
		t.Buyer_user = tmsg.MsgBody.Qr_code_info.Buyer_user
		t.Ins_order_id = tmsg.MsgBody.Qr_code_info.Ins_order_id
		t.Goods_id = tmsg.MsgBody.Qr_code_info.Goods_id
		t.Goods_nm = tmsg.MsgBody.Qr_code_info.Goods_nm
		t.Goods_num = tmsg.MsgBody.Qr_code_info.Goods_num
		t.Goods_price = tmsg.MsgBody.Qr_code_info.Goods_price
	}
	if tmsg.MsgBody.Term_inf != nil {
		t.Ip_addr = tmsg.MsgBody.Term_inf.Ip_addr
		t.Gps_addr = tmsg.MsgBody.Term_inf.Gps_addr
	}

	if tmsg.MsgBody.Custom_ext_info != nil {
		t.CustomExtInfo = *tmsg.MsgBody.Custom_ext_info
	}
	t.Orig_trans_dt = tmsg.MsgBody.Orig_trans_dt
	t.Be_biz_cd = tmsg.MsgBody.Be_biz_cd
	if t.Be_biz_cd == "" {
		t.Be_biz_cd = tmsg.MsgBody.Biz_cd
	}
	t.Trans_st = "0000000000"
	//取交易名称, 撤销标志
	dbc := gormdb.GetInstance()
	bizCtl := DbBizCtl{Prod_cd: t.Prod_cd, Tran_cd: t.Tran_cd, Biz_cd: t.Biz_cd}
	err := dbc.Find(&bizCtl).Error
	if err == nil {
		t.Tran_nm = bizCtl.Prod_nm
		t.Cancel_flg = bizCtl.Cancel_flg
	} else {
		logr.Error("取交易信息失败", bizCtl, err)
		t.Cancel_flg = "0"
	}
	if tmsg.MsgBody.Tran_cd == "1151" {
		tmsg.MsgBody.Resp_cd = "00"
		tmsg.MsgBody.Resp_msg = "交易成功"
		t.Resp_cd = "00"
		t.Resp_msg = "交易成功"
		t.Trans_st = "1000000000"
		t.Sys_order_id = tmsg.MsgBody.Sys_order_id
		if tmsg.MsgBody.Custom_ext_info.Cust_order_id != "" {
			t.Cust_order_id = tmsg.MsgBody.Custom_ext_info.Cust_order_id
		} else {
			t.Cust_order_id = tmsg.MsgBody.Sys_order_id
		}
	} else if tmsg.MsgBody.Tran_cd == "3151" {
		tmsg.MsgBody.Tran_cd = "3152"
	}
}

func (t *DbTranLog) UpdateTransSt() {
	if t.Resp_cd == defs.TRN_SUCCESS {
		t.Trans_st = sutil.ReplaceAtIndex(t.Trans_st, godefs.TRN_STAT_SUCC, godefs.LOC_IDX)
	} else {
		t.Trans_st = sutil.ReplaceAtIndex(t.Trans_st, godefs.TRN_STAT_FAIL, godefs.LOC_IDX)
	}
}
