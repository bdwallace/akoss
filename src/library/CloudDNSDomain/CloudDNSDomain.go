package CloudDNSDomain

import (
	"fmt"
	"models/request"
)



func DomainSearch(searchReq *request.DomainForSearchRequest) (resp interface{}, err error){

	if searchReq.Src == "" {
		return nil, fmt.Errorf("error:  DomainAdapterSearchOrBackUpLocal  ::  src or dest is empty")
	}

	if searchReq.Src == "ALi"{
		// domain search from ali domain
		resp, err = SearchAliDomain(searchReq)
		if err != nil{
			return nil, err
		}
	}

	if searchReq.Src == "TenCent" {
		// domain search from tencent domain
		resp, err = SearchTenCentDomain(searchReq)
		if err != nil{
			return nil, err
		}
	}

	if searchReq.Src == "CloudFare"{
		// domain search from cloudfare domain

	}

	return resp, nil

}



func DomainBackUpToLocal(backupReq *request.DomainForBackUpTolocalRequest) (err error) {

	if backupReq.Src == "" {
		return fmt.Errorf("error:  DomainAdapterSearchOrBackUpLocal  ::  src or dest is empty")
	}

	if backupReq.Src == "ALi"{
		// domain backup to local ali domain
		err = BackUpOrUpdateToLocalAliDomain(backupReq)
		if err != nil{
			return
		}
	}

	if backupReq.Src == "TenCent" {
		// domain backup to local tencent domain
		err = BackUpOrUpdateToLocalTenCentDomain(backupReq)
		if err != nil{
			return err
		}

	}

	if backupReq.Src == "CloudFare"{
		// domain backup to local cloudfare domain

	}

	return nil
}



func DomainImportToCloud(importReq *request.DomainForImportOtherCloudRequest) (resp interface{}, err error){

	if importReq.Dest == "ALi"{
		return ImportToAliDomain(importReq)
	}
	if importReq.Dest == "TenCent"{
		return ImportToTenCentDomain(importReq)


	}
	if importReq.Dest == "CloudFare"{

	}

	return
}


