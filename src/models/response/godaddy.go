package response


type RespGodaddy struct {
	CreatedAt						string			`json:"createdAt"`
	Domain							string			`json:"domain"`
	DomainId						int 			`json:"domainId"`
	ExpirationProtected				bool			`json:"expirationProtected"`
	Expires							string			`json:"expires"`
	ExposeWhois						bool			`json:"exposeWhois"`
	HoldRegistrar					bool			`json:"holdRegistrar"`
	Locked							bool			`json:"locked"`
	NameServers						string			`json:"nameServers"`
	Privacy							bool			`json:"privacy"`
	RenewAuto						bool			`json:"renewAuto"`
	RenewDeadline					string			`json:"renewDeadline"`
	Renewable						bool			`json:"renewable"`
	Status							string			`json:"status"`
	TransferAwayEligibleAt			string			`json:"transferAwayEligibleAt"`
	TransferProtected				bool			`json:"transferProtected"`
}