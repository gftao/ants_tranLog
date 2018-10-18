package digest

import (
	"net/http"
	"tranLog/models"
	"io/ioutil"
	"encoding/json"
	"golib/modules/run"
	"golib/gerror"
	"sync"
	"golib/modules/gormdb"
	"github.com/jinzhu/gorm"
	"tranLog/godefs"
	"tranLog/models/sutil"
)

type TranLogTask struct {
	W        http.ResponseWriter
	R        *http.Request
	TransMsg models.TransMessage
	run.BaseWorker
	Wg       sync.WaitGroup
}

func (t *TranLogTask) DoWork() {
	t.Info("开始处理请求")
	gerr := t.getParamsInit()
	if gerr != nil {
		t.rejectMsg(gerr.GetErrorCode(), gerr.GetErrorString(), gerr)
		return
	}
	t.SetCldOrderId(t.TransMsg.MsgBody.Order_id)
	gerr = t.insertDb()
	if gerr != nil {
		t.rejectMsg(gerr.GetErrorCode(), gerr.GetErrorString(), gerr)
		return
	}
	t.sendMsgToClient()
	t.Wg.Done()
}

func (t *TranLogTask) getParamsInit() gerror.IError {
	body, err := ioutil.ReadAll(t.R.Body)
	if err != nil {
		t.Error("读取请求报文失败", err)
		t.W.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	t.Info("get Request msg:", string(body))
	err = json.Unmarshal([]byte(body), &t.TransMsg)
	if err != nil {
		t.Error("decode json body fail", err)
		return gerror.New(51005, "96", err, "decode json body fail")
	}
	t.Infof("get msg_body:%s", t.TransMsg.Msg_body)
	err = json.Unmarshal([]byte(t.TransMsg.Msg_body), &t.TransMsg.MsgBody)
	if err != nil {
		t.Error("decode json body fail", err)
		return gerror.New(51006, "96", err, "decode json body fail")
	}
	t.Infof("get msg_body:%+v", t.TransMsg.MsgBody)
	return nil
}

func (t *TranLogTask) insertDb() gerror.IError {
	dbc := gormdb.GetInstance()
	dbc.SetLogger(t.GetEntry())
	msgBody := t.TransMsg.MsgBody
	dbtl := models.DbTranLog{}
	if len(msgBody.Tran_cd) != 4 {
		return gerror.New(51001, "98", nil, "交易码非法："+t.TransMsg.MsgBody.Sys_order_id)
	}
	if msgBody.Tran_cd[3:] == "1" {
		dbtl.InitByTransMsg(t.TransMsg)
		err := dbc.Create(&dbtl).Error
		if err != nil {
			return gerror.New(51003, "96", err, "记录流水失败："+t.TransMsg.MsgBody.Cld_order_id)
		}
	} else if msgBody.Tran_cd[3:] == "2" {
		tx := dbc.Begin()
		err := t.UpdRspTranInfo(tx, &dbtl)

		if err != nil {
			tx.Rollback()
			return gerror.New(51001, "96", err, "更新源交易流水失败："+t.TransMsg.MsgBody.Cld_order_id)
		}
		tx.Commit()
	}
	return nil
}
func (t *TranLogTask) sendMsgToClient() {
	TransMsg := &models.TransMessage{}
	TransMsg.MsgBody = &models.TransParams{}
	TransMsg.MsgBody.Tran_cd = t.TransMsg.MsgBody.Tran_cd
	TransMsg.MsgBody.Resp_cd = "00"
	TransMsg.MsgBody.Resp_msg = "SUCCESS"
	TransMsg.SetMsgBody()

	t.W.Header().Set("Content-Type", "application/json;charset=utf-8")
	slen, err := t.W.Write([]byte(TransMsg.ToString()))
	if err != nil {
		t.Error("发送报文失败", err)
		return
	}

	t.Info("发送应答成功:", slen, t.TransMsg.ToString())

	return
}

func (t *TranLogTask) rejectMsg(resp_cd, resp_msg string, e error) {
	t.Info("交易失败:"+resp_msg, e)
	if t.TransMsg.MsgBody == nil {
		t.TransMsg.MsgBody = &models.TransParams{}
	}
	t.TransMsg.MsgBody.Resp_cd = resp_cd
	t.TransMsg.MsgBody.Resp_msg = resp_msg
	if t.TransMsg.MsgBody.Tran_cd == "" {
		t.Warn("交易码为空，请求不处理")
		t.W.WriteHeader(http.StatusForbidden)
		return
	}
	//t.TransMsg.MsgBody.Tran_cd = t.TransMsg.MsgBody.Tran_cd[:3] + "2"
	t.TransMsg.SetMsgBody()

	t.W.Header().Set("Content-Type", "application/json;charset=utf-8")
	slen, err := t.W.Write([]byte(t.TransMsg.ToString()))
	if err != nil {
		t.Error("发送报文失败", err)
		return
	}

	t.Info("发送应答成功:", slen, t.TransMsg.ToString())

	return
}

func (t *TranLogTask) UpdRspTranInfo(tx *gorm.DB, dbtl *models.DbTranLog) gerror.IError {
	err := tx.Set("gorm:query_option", "FOR UPDATE").Where(" order_id = ? ",
		t.TransMsg.MsgBody.Order_id).Find(&dbtl).Error
	if err != nil {
		return gerror.New(80050, godefs.TRN_SYS_ERROR, nil, "查询交易流水失败："+t.TransMsg.MsgBody.Cld_order_id)
	}
	if dbtl.Resp_cd != "" {
		return gerror.New(80060, godefs.TRN_SYS_ERROR, nil, "交易状态异常：" + t.TransMsg.MsgBody.Cld_order_id+
			t.TransMsg.MsgBody.Resp_cd)
	}

	dbtl.Resp_cd = t.TransMsg.MsgBody.Resp_cd
	if dbtl.Resp_cd == godefs.TRN_TIME_OUT {
		t.Info("交易超时，不更新流水状态")
		return nil
	}

	if t.TransMsg.MsgBody.Ext_print_info != "" {
		Print := &models.PrintInfo{}
		err = json.Unmarshal([]byte(t.TransMsg.MsgBody.Ext_print_info), Print)
		dbtl.DesAcqInsId = Print.Des_term__cd //转换的终端号
		dbtl.DesMchntCd = Print.Des_mchnt_cd  //转换的商户号
	}
	if t.TransMsg.MsgBody.Gift_info != "" {
		git := &models.Gift_info{}
		err = json.Unmarshal([]byte(t.TransMsg.MsgBody.Gift_info), git)
		if err == nil {
			dbtl.Ext_fld1 = git.Real_amt
			t.TransMsg.MsgBody.Tran_amt = git.Real_amt
			dbtl.Ext_fld2 = git.Prob_amt
			dbtl.Ext_fld3 = git.Prob_status
			dbtl.Ext_fld6 = git.Prob_sucess_info
		} else {
			t.Debug("Gift_info:", err)
		}
	}
	if t.TransMsg.MsgBody.Foreign_info != nil {
		foreign := &models.ForeignInfo{}
		err = json.Unmarshal([]byte(t.TransMsg.MsgBody.ForeignInfo), foreign)
		if err == nil {
			dbtl.Ext_fld10 = foreign.Foreign_amt
			dbtl.Ext_fld11 = foreign.Foreign_rate
			dbtl.Ext_fld12 = foreign.Foreign_curr_cd
			dbtl.Ext_fld13 = foreign.Foreign_flag
		} else {
			t.Debug("Foreign_info:", err)
		}
	}
	switch t.TransMsg.MsgBody.Tran_cd {

	case "2012", "3012", "2132", "3132":
		if t.TransMsg.MsgBody.Resp_cd == godefs.TRN_SUCCESS {
			dbt := models.DbTranLog{}
			err := tx.Set("gorm:query_option", "FOR UPDATE").Where(" sys_order_id = ? ",
				t.TransMsg.MsgBody.Orig_sys_order_id).Find(&dbt).Error
			if err != nil {
				return gerror.New(80050, godefs.TRN_SYS_ERROR, nil, "查询源交易流水失败："+t.TransMsg.MsgBody.Cld_order_id)
			}
			if t.TransMsg.MsgBody.Tran_amt == dbt.Tran_amt {
				dbt.Trans_st = sutil.ReplaceAtIndex(dbt.Trans_st, godefs.PRTA_STAT_FLG, godefs.PRT_IDX)
			} else {
				dbt.Trans_st = sutil.ReplaceAtIndex(dbt.Trans_st, godefs.PRTP_STAT_FLG, godefs.PRT_IDX)
			}
			err = tx.Model(&dbt).Updates(dbt).Error
			if err != nil {
				return gerror.New(80070, godefs.TRN_SYS_ERROR, nil, "更新源交易结果失败："+t.TransMsg.MsgBody.Cld_order_id)
			}
		}
	}
	dbtl.UpdateTransSt()
	err = tx.Model(&dbtl).Updates(models.DbTranLog{Resp_cd: t.TransMsg.MsgBody.Resp_cd,
		Trans_st: dbtl.Trans_st,
		Sys_order_id: t.TransMsg.MsgBody.Sys_order_id,
		Resp_msg: t.TransMsg.MsgBody.Resp_msg,
		Iss_ins_id_cd: t.TransMsg.MsgBody.Iss_ins_id_cd,
		DesAcqInsId: dbtl.DesAcqInsId,
		DesMchntCd: dbtl.DesMchntCd,
		Be_order_id: t.TransMsg.MsgBody.Tran_order_id,
		Pre_auth_id: t.TransMsg.MsgBody.Pre_auth_id,
		Pri_acct_no: t.TransMsg.MsgBody.Pri_acct_no,
		CustomExtInfo: models.CustomExtInfo{Ext_fld1: dbtl.Ext_fld1,
			Ext_fld2: dbtl.Ext_fld2,
			Ext_fld3: dbtl.Ext_fld3,
			Ext_fld4: t.TransMsg.MsgBody.Card_tp,
			Ext_fld5: t.TransMsg.MsgBody.Ref_no,
			Ext_fld6: dbtl.Ext_fld6,
			Ext_fld10: dbtl.Ext_fld10,
			Ext_fld11: dbtl.Ext_fld11,
			Ext_fld12: dbtl.Ext_fld12,
			Ext_fld13: dbtl.Ext_fld13},
	}).Error
	if err != nil {
		return gerror.New(80070, godefs.TRN_SYS_ERROR, nil, "更新交易结果失败："+t.TransMsg.MsgBody.Cld_order_id)
	}
	return nil
}
