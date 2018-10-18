package models

type AppSourceInfo struct {
	App_source_type         string `json:"app_source_type,omitempty"`
	App_source_code         string `json:"app_source_code,omitempty"`
	App_source_url          string `json:"app_source_url,omitempty"`
	App_source_version_code int    `json:"app_source_version_code,omitempty"`
	App_source_version_name string `json:"app_source_version_name,omitempty"`
	App_source_md5          string `json:"app_source_md5,omitempty"`
}
