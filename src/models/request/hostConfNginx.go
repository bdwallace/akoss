package request

import "library/components"

type ProjectConfNginx struct {
	ProjectId       	int         `json:"project_id"`
	ProjectAlias     	string      `json:"project_alias"`
	Platforms			[]*PlatformConfNginx
}

type PlatformConfNginx struct {
	PlatformId      	int         `json:"platform_id"`
	PlatformName     	string      `json:"platform_name"`
	Services			[]*ServiceConfNginx
}

type ServiceConfNginx struct {
	ServiceId	    	int         `json:"service_id"`
	ServiceName     	string      `json:"service_name"`
	ServiceClass 		string		`json:"service_class"`
	Hosts				[]*HostConfNginx
	BaseDocker	 		*components.BaseDocker
}

type HostConfNginx struct {
	HostId 				int			`json:"host_id"`
	HostName			string		`json:"host_name"`
	UseIp				string		`json:"use_ip"`
	NginxConfs			[]*ConfFile
}

type ConfFile struct {
	FilePath			string		`json:"conf_dir"`
	FileDir				string		`json:"conf_dir"`
	FileName 			string		`json:"file_name"`
}
