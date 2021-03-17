package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/smtp"
	"strings"
	"time"
)

var (
	longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

/// 邮件附件结构：FileName主要用于存放于邮件之中，FullPath是附件文件的全路径+文件名，主要用于函数内搜索该文件并进行base64转码
type MailAtta struct {
	FileName string
	FullPath string
}

/// To接收人,如果接收人为多个，应该以 a@company.com;b@company.com 格式传入
type MailBody struct {
	To      string
	Cc      string
	Subject string
	Text    string
	Attas   []MailAtta
}

type MailSender struct {
	// 发送邮件所用的邮箱名。比如： hhyt_devops@163.com
	SenderMail string
	// 发送邮件使用的邮箱账号的授权码
	AuthPassword string
	// 发送邮件所使用的SMTP服务器及端口。比如：smtp.163.com:465
	SmtpServer string
}

func SendMail(sender MailSender, body MailBody) error {
	sendTo := strings.Split(body.To, ";")
	boundary := "derlegehedetinybear" //boundary 用于分割邮件内容，可自定义. 注意它的开始和结束格式
	mime := bytes.NewBuffer(nil)
	//设置邮件
	mime.WriteString(fmt.Sprintf("From: %s<%s>\r\nTo: %s\r\nCC: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\n",
		sender.SenderMail, sender.SenderMail, body.To, body.Cc, body.Subject))
	partid := string(randCharacter(8))
	mime.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s_%s\"\r\n", boundary, partid))
	mime.WriteString("Content-Type: multipart/alternative;\r\n")
	mime.WriteString("\r\n\r\nThis is a multi-part message in MIME format.\r\n\r\n")
	mime.WriteString(fmt.Sprintf("--=====%s_%s=====\r\n", boundary, partid))

	//邮件普通Text正文
	partid = string(randCharacter(8))
	mime.WriteString(fmt.Sprintf("--=====%s_%s=====\r\n", boundary, partid))
	mime.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	mime.WriteString(body.Text)
	mime.WriteString("\r\n")
	mime.WriteString(fmt.Sprintf("--=====%s_%s=====--\r\n", boundary, partid))
	// 附件
	for _, atta := range body.Attas {
		partid = string(randCharacter(8))
		mime.WriteString(fmt.Sprintf("\n--=====%s_%s=====\r\n", boundary, partid))
		mime.WriteString("Content-Type: application/octet-stream\r\n")
		mime.WriteString("Content-Transfer-Encoding: base64\r\n")
		mime.WriteString("Content-Disposition: attachment; filename=\"" + atta.FileName + "\"\r\n\r\n")
		//读取并编码文件内容
		attaData, err := ioutil.ReadFile(atta.FullPath)
		if err != nil {
			return err
		}
		b := make([]byte, base64.StdEncoding.EncodedLen(len(attaData)))
		base64.StdEncoding.Encode(b, attaData)
		mime.Write(b)
	}
	//邮件结束
	mime.WriteString(fmt.Sprintf("--=====%s_%s=====--\r\n\r\n", boundary, partid)) //最后一个位置
	fmt.Println(mime.String())
	//发送
	smtpHost, _, err := net.SplitHostPort(sender.SmtpServer)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", sender.SenderMail, sender.AuthPassword, smtpHost)
	return smtp.SendMail(sender.SmtpServer, auth, sender.SenderMail, sendTo, mime.Bytes())
}

func randCharacter(length int) []byte {
	if length <= 0 {
		return []byte{}
	}
	bytes := make([]byte, length)
	arc := uint8(0)
	if _, err := rand.Read(bytes[:]); err != nil {
		return []byte{}
	}
	for index, b := range bytes {
		arc = b & 61
		bytes[index] = longLetters[arc]
	}
	return bytes
}

func init() {
	rand.Seed(time.Now().Unix())
}
