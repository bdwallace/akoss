package components

import "models"

func GetServiceNacos(p * models.Project, s * models.Service){

	nacosList := make([]string,0)
	nacosList = append(nacosList, p.Nacos1)
	nacosList = append(nacosList, p.Nacos2)
	nacosList = append(nacosList, p.Nacos3)

	for i, data := range nacosList{
		if s.UseNacos == data{

			switch i {
			case 0:
				s.NacosUserName = p.Nacos1UserName
				s.NacosPwd = p.Nacos1Pwd
			case 1:
				s.NacosUserName = p.Nacos2UserName
				s.NacosPwd = p.Nacos2Pwd
			case 2:
				s.NacosUserName = p.Nacos2UserName
				s.NacosPwd = p.Nacos2Pwd
			}
		}
	}

	return
}