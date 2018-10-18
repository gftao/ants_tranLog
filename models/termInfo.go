package models

type TermInfo struct {
	Ip_addr    string `json:"ip_addr,omitempty"`
	Gps_addr   string `json:"gps_addr,omitempty"`
	Term_prod  string `json:"term_prod,omitempty"`
	Term_model string `json:"term_model,omitempty"`
	Brand_ksn  string `json:"brand_ksn,omitempty"`
	Brand_sn   string `json:"brand_sn,omitempty"`
	Term_tp    string `json:"term_tp,omitempty"`
	Term_sn    string `json:"term_sn,omitempty"`
	Term_rand  string `json:"term_rand,omitempty"`
	Term_enc   string `json:"term_enc,omitempty"`
	Term_ver   string `json:"term_ver,omitempty"`
}

type Gift_info struct {
	Bal_num          string `json:"bal_num,omitempty"`
	Bal_total_num    string `json:"bal_total_num,omitempty"`
	Sum_num          string `json:"sum_num,omitempty"`
	Sum_total_num    string `json:"sum_total_num,omitempty"`
	Bal_amt          string `json:"bal_amt,omitempty"`
	Bal_total_amt    string `json:"bal_total_amt,omitempty"`
	Sum_amt          string `json:"sum_amt,omitempty"`
	Sum_total_amt    string `json:"sum_total_amt,omitempty"`
	Coupon_amt       string `json:"coupon_amt,omitempty"`
	Coupon_num       string `json:"coupon_num,omitempty"`
	Trans_amt        string `json:"trans_amt,omitempty"`
	Real_amt         string `json:"real_amt,omitempty"`
	Prob_amt         string `json:"prob_amt,omitempty"`
	Prob_sucess_info string `json:"prob_sucess_info,omitempty"`
	Prob_code        string `json:"prob_code,omitempty"`
	Prob_boss        string `json:"prob_boss,omitempty"`
	Prob_info        string `json:"prob_info,omitempty"`
	Prob_status      string `json:"prob_status,omitempty"`
	Ims_ext1         string `json:"ims_ext_1,omitempty"`
	Ims_ext2         string `json:"ims_ext_2,omitempty"`
}
