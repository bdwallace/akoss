package response



type ConfFile struct {
	ConfId				int
	ConfFileName 		string		`json:"conf_file_name"`
	ConfInfo 			string 		`json:"conf_info"`
	ConfRelation
}

type ConfRelation struct {
	ProjectId 			int
	PlatformId 			int
	ServiceId 			int
	HostId 				int
	ProjectName 		string
	PlatformName		string
	ServiceName			string
	ServiceClass		string
	HostIp				string
}
