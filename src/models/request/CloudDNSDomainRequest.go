package request

import (
	"models"
)

type DomainForSearchRequest struct {
	SrcKeyId 		string
	SrcKeySecret 	string
	Src 			string			// src from ali or tencent or cloudfare ...
	SearchDomain 	string

	PageSize 		int
	PageNum 		int
}
/*
type ALiDomain struct {
	KeyId  			string
	KeySecret 		string
	Domain			string

}
type TenCentDomain struct {

	Domain 			[]*DomainAdapterDomain
}
*/


type DomainForBackUpTolocalRequest struct {
	SrcKeyId 		string
	SrcKeySecret 	string
	Src 			string			// src from ali or tencent or cloudfare ...
}


type DomainForImportOtherCloudRequest struct {
	Src 			string			// src from ali or tencent or cloudfare ...
	Dest 			string			// src from ali or tencent or cloudfare ...
	//SrcKeyId 		string
	//SrcKeySecret 	string
	DestKeyId 		string
	DestKeySecret 	string
	AliDomain 		[]*models.AliDomain
	TenCentDomain	[]*models.TencentDomain

}