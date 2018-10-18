package godefs

const DEFAULT = "default"

//const SYSTIMEFORMAT="20060102150405"
const SYSTIMEFORMAT = "20060102"
const SYSSEQLEN = 7 //系统流水号长

//应答码
const (
	TRN_SUCCESS            = "00"
	TRN_INVALID_MERCH      = "03" /* 无效商户 */
	TRN_PARTIAL_APPROVED   = "10"
	TRN_VIP_APPROVED       = "11"
	TRN_INVALID_AMT        = "13" /* 无效金额 */
	TRN_INVALID_PAN        = "14" /* 无效卡号 */
	TRN_INVALID_DEST_BIN   = "15" /* 无发卡行 */
	TRN_APPROVED_UPD_T3    = "16"
	TRN_ORIG_TRN_NOT_SUCC  = "12"
	TRN_ORIG_TRN_NOT_FUNC  = "22" /* 嫌疑交易 */
	TRN_ORIG_TRN_NOT_FOUND = "25" /* 找不到原始交易 */
	TRN_FORMAT_ERR         = "30" /* 格式错误 */
	TRN_DESINST_ERR        = "31" /* 找不到目标机构 */
	TRN_NOT_SUPPORT        = "40" /* 不支持交易 */
	TRN_NO_FALLBACK        = "45" /* 不允许降级交易*/
	TRN_AMT_LIMIT_ERR      = "61" /* 金额笔数超限 */
	TRN_INCORRECT_AMT      = "64" /* 原始金额错 */
	TRN_DEST_INS_DOWN      = "91"
	TRN_INST_ERROR         = "92" /* 机构错 */
	TRN_DUPL_TXN           = "94"
	TRN_SYS_ERROR          = "96" /* 系统故障 */
	TRN_INVALID_TERM       = "97" /* 无效终端 */
	TRN_TIME_OUT           = "98"
	TRN_TRANS_PIN_FAIL     = "99"
	TRN_INST_SINGOUT       = "C1" /* 未签到 */
	TRN_VERIFY_MAC_FAIL    = "A0"
	TRN_SUCC_FAULT_2       = "A2"
	TRN_SUCC_FAULT_4       = "A4"
	TRN_SUCC_FAULT_5       = "A5"
	TRN_SUCC_FAULT_6       = "A6"
	WAR_INVALID_PAN        = "B0" /* 风控黑卡 */
	WAR_BLACK_PAN          = "B1" /* 风控黑名单拒绝卡*/
	TRAN_ERR_PRODCD        = "H3" /* 上送产品码或业务代码有误*/
	TRN_SM_CLOSED          = "05" /*  订单已关闭 */
	TRN_SM_UNKOWN          = "C2" /*  订单状态未知 */
	TRN_SM_OTHERR          = "E0" /*  其它错误 */
	AGT_TRN_CUTOFF         = "73" /*  原交易已经日切 */
	AGT_TRN_OTHERR         = "E0" /*  其它错误 */
	TRN_REVSALED           = "J2" /*交易已冲正*/

	TRN_SYS_BUSY       = "J0" /* 系统忙，请稍后重试 */
	TRN_SYS_BUSY_INS   = "J3" /* 机构限制 系统忙，请稍后重试 */
	MGM_ACTIVE_INVALID = "J1" /* 非法的激活码*/
)


const (
	/*交易状态*/
	TRN_STAT_LEN     = 1
	TRN_STAT_NO_RSP  = '0'
	TRN_STAT_SUCC    = '1'
	TRN_STAT_FAIL    = '2'
	TRN_STAT_REJ     = '3'
	TRN_STAT_WAIT    = '4'
	TRN_STAT_TIMEOUT = '5'
	TRN_STAT_ZWXTREJ = '6'

	CAN_STAT_FLG  = '1' //已撤销
	REV_STAT_FLG  = '1' //已冲正
	PRTA_STAT_FLG = '2' //全额退货
	PRTP_STAT_FLG = '1' //部分退货
	COM_STAT_FLG  = '1' //已完成
	NOR_STAT_FLG  = '0'
	TRAN_ST_INIT  = "0000000000"

	LOC_IDX = 0 /*交易状态下标*/
	REV_IDX = 1 /*冲正标志下标*/
	CAN_IDX = 2 /*撤销标志下标*/
	PRT_IDX = 3 /*退货标志下标*/
	COM_IDX = 4 /*完成标志下标*/
)