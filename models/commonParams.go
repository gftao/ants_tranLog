package models

import (
	"golib/modules/idManager"
	"golib/modules/run"
	"sync"
)

type MessageHead struct {
	Encoding    string `json:"encoding"`
	Sign_method string `json:"sign_method"`
	Signature   string `json:"signature"`
	Version     string `json:"version"`
}

type commonParams struct {
	Tran_cd          string `json:"tran_cd,omitempty"`
	Prod_cd          string `json:"prod_cd,omitempty"`
	Biz_cd           string `json:"biz_cd,omitempty"`
	Term_seq         string `json:"term_seq,omitempty"`
	Term_batch       string `json:"term_batch,omitempty"`
	Mcht_cd          string `json:"mcht_cd,omitempty"`
	Mcht_nm          string `json:"mcht_nm,omitempty"`
	Term_id          string `json:"term_id,omitempty"`
	Term_flag        string `json:"term_flag,omitempty"`
	Tran_dt_tm       string `json:"tran_dt_tm,omitempty"`
	Order_id         string `json:"order_id,omitempty"`
	Order_timeout    string `json:"order_timeout,omitempty"`
	Sys_order_id     string `json:"sys_order_id,omitempty"` //产品平台订单号
	Cld_order_id     string `json:"-"`                      //云前置自己的订单号
	Tran_order_id    string `json:"tran_order_id,omitempty"` //交易订单号，汇宜系统内唯一
	Acct_order_id    string `json:"-"`
	Order_desc       string `json:"order_desc,omitempty"`
	Req_reserved     string `json:"req_reserved,omitempty"`
	Resp_cd          string `json:"resp_cd,omitempty"`
	Resp_msg         string `json:"resp_msg,omitempty"`
	ActiveCode       string `json:"active_code,omitempty"`
	Pri_acct_no      string `json:"pri_acct_no,omitempty"`
	Tran_amt         string `json:"tran_amt,omitempty"`
	Curr_cd          string `json:"curr_cd,omitempty"`
	Pre_auth_id      string `json:"pre_auth_id,omitempty"`
	Sett_dt          string `json:"sett_dt"`
	Ins_id_cd        string `json:"ins_id_cd,omitempty"`
	Iss_ins_id_cd    string `json:"iss_ins_id_cd"`
	Trans_in_acct_no string `json:"trans_in_acct_no,omitempty"`
	Chn_ins_id_cd    string `json:"chn_ins_id_cd,omitempty"`
}

type SystemParams struct {
	SysIdGenerator  idManager.IdGenerator
	TranIdGenerator idManager.IdGenerator
	run.BaseWorker
	TransMsg TransMessage
	Mutex    sync.RWMutex
	MsgDone  bool
}

func (t *SystemParams) CheckMsgAndSetDown() bool {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	if t.MsgDone != false {
		return false //已经处理，不再处理
	}
	t.MsgDone = true
	return true
}

func (t SystemParams) CheckMsg() bool {
	t.Mutex.RLock()
	defer t.Mutex.RUnlock()
	return t.MsgDone
}
