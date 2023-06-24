package alertEmail

import (
	"fmt"
	"net/smtp"
)

func EmailInfoOfBackendService(unhealthyServices []map[string]string)(emailInfo string){

	unhealthyServicesNumber := len(unhealthyServices)
	if unhealthyServicesNumber > 0 {
		for _, itme := range unhealthyServices{
			// edit email
			emailInfo += "服务容器异常："
			emailInfo += EditAlertEmail(itme)
			emailInfo += "\n"
		}
	}
	return
}

func EditAlertEmail(item map[string]string)(mailInfo string){
	mailInfo = fmt.Sprintf("project_name: %s\nservice_name: %s\nhost_id: %s\npublic_ip: %s\nprivate_ip: %s\n",item["project_name"],item["service_name"],item["host_id"],item["ip_pub"],item["ip"])

/*
	mailInfo += fmt.Sprintf("project_name: %s",item["project_name"])
	mailInfo += fmt.Sprintf("service_name: %s", item["service_name"])
	mailInfo += fmt.Sprintf("host_id: %s", item["host_id"])
	mailInfo += fmt.Sprintf("public_ip: %s", item["ip_pub"])
	mailInfo += fmt.Sprintf("private_ip: %s", item["ip"])
*/

	return
}

func SendEmail(emailInfo string)error{

	auth := smtp.PlainAuth("", "dominic@sixthnet.com", "dominic@123456", "smtp.mail.us-east-1.awsapps.com")

	to := []string{"smtp.mail.us-east-1.awsapps.com"}
	//msg := []byte
	info := fmt.Sprintf("To: dominic@sixthnet.com\r\nSubject: 服务容器状态报警\r\n\r\n%s\r\n",emailInfo)
	msg := []byte(info)
	if err := smtp.SendMail("dominic@sixthnet.com", auth, "dominic@sixthnet.com", to, msg); err != nil{
		return err
	}


	return nil
}
