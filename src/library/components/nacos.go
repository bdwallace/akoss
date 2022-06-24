package components

import "models"

func GetServiceNacos(p * models.Project, s * models.Service)(err error){

	nacosList := make([]string,0)
	if p.Nacos1 != "" && p.Nacos1UserName != "" && p.Nacos1Pwd != "" {
		nacosList = append(nacosList, p.Nacos1)
	}
	if p.Nacos2 != "" && p.Nacos2UserName != "" && p.Nacos2Pwd != "" {
		nacosList = append(nacosList, p.Nacos2)
	}
	if p.Nacos3 != "" && p.Nacos3UserName != "" && p.Nacos3Pwd != "" {
		nacosList = append(nacosList, p.Nacos3)
	}

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

	_, err = models.UpdateServiceById(s)

	return
}