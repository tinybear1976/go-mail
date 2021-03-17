# go-langpacks
**version: 0.9.1**



一个简易纯文本邮件发送程序，可以发送附件。

A simple text mail sending program, you can send attachments.



# Install

```bash
go get  github.com/tinybear1976/go-mail
```


# Example

```go
func TestSendMailOK(t *testing.T) {
	sender := MailSender{
		SenderMail:   "test@163.com",
		AuthPassword: "密码",
		SmtpServer:   "smtp.163.com:465",
	}
	body := MailBody{
		To:      "333@qq.com",
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
```

