package models

import (
	"crypto/rsa"
)

type KeyInfo struct {
	Key_down_flg   string `json:"key_down_flg,omitempty"`
	Session_key    string `json:"session_key,omitempty"`
	Term_pub_key   string `json:"term_pub_key,omitempty"`
	Server_pub_key string `json:"server_pub_key,omitempty"`
	Term_tmk       string `json:"term_tmk,omitempty"`
	Term_tmk_chk   string `json:"term_tmk_chk,omitempty"`
	Term_pik       string `json:"term_pik,omitempty"`
	Term_pik_chk   string `json:"term_pik_chk,omitempty"`
	Term_trk       string `json:"term_trk,omitempty"`
	Term_trk_chk   string `json:"term_trk_chk,omitempty"`
	Pboc_ic_key    string `json:"pboc_ic_key,omitempty"`
	Pboc_ic_param  string `json:"pboc_ic_param,omitempty"`
}

type KeyHandleInfo struct {
	Term_key     string
	TermPubKey   *rsa.PublicKey
	TermPriKey   *rsa.PrivateKey
	ServerPubKey *rsa.PublicKey
	ServerPriKey *rsa.PrivateKey
}
