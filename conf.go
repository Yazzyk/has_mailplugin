package mailplugin

import "github.com/drharryhe/has/core"

type MailPlugin struct {
	core.PluginConf

	FromMail string // 发送方邮件
	Pwd      string // 授权码
	HTMLPath string // html 地址
	SMTP     smtp   // smtp服务器
}

type smtp struct {
	Server string // 服务器地址
	Port   uint   // 端口
}
