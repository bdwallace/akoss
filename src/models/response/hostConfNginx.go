package response


type ConfRelation struct {
	ServiceId 			int					`orm:"column(id)"`
	ServiceName			string				`orm:"column(name)"`
	HostId				int					`orm:"column(t_host_id)"`
	HostUseIp			string				`orm:"column(use_ip)"`
	Confs 				[]*ConfFile
}


type ConfFile struct {
	ConfServiceId 		int			`json:"conf_service_id"`
	HostUseIp			string		`json:"host_use_ip"`
	ConfFileName 		string		`json:"conf_file_name"`
	ConfFullPath 		string		`json:"conf_full_path"`
	ConfInfo 			string 		`json:"conf_info"`
}

