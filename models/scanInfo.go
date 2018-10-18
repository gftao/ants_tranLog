package models

type QrCodeInfo struct {
	Auth_code    string `json:"auth_code,omitempty"`
	Qr_type      string `json:"qr_type,omitempty"`
	Time_out     string `json:"time_out,omitempty"`
	Buyer_user   string `json:"buyer_user,omitempty"`
	Ins_order_id string `json:"ins_order_id,omitempty"`
	Goods_id     string `json:"goods_id,omitempty"`
	Goods_nm     string `json:"goods_nm,omitempty"`
	Goods_num    string `json:"goods_num,omitempty"`
	Goods_price  string `json:"goods_price,omitempty"`
}
