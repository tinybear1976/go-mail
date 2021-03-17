package mail

import (
	"testing"
)

func TestSendMailOK(t *testing.T) {
	sender := MailSender{
		SenderMail:   "hhyt_devops@163.com",
		AuthPassword: "ATOLOIEHZIOAWTBC",
		SmtpServer:   "smtp.163.com:465",
	}
	body := MailBody{
		To:      "908247013@qq.com",
		Cc:      "",
		Subject: "test mail测试邮件",
		Text:    "2021年3月17日。go测试邮件",
		Attas: []MailAtta{{
			FileName: "20200309_bak.zip",
			FullPath: "/Users/wang/Downloads/20200309_bak.zip",
		}, {
			FileName: "RedisDesktopManager2020079.zip",
			FullPath: "/Users/wang/Downloads/RedisDesktopManager2020079.zip",
		}},
	}
	if err := SendMail(sender, body); err != nil {
		t.Error(err.Error())
	}
}
