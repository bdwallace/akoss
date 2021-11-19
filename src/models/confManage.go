package models

type ConfEdit struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Info      string `json:"info"`
	ServiceId string `json:"service_id"`
	IP        string `json:"host_ip"`
}
