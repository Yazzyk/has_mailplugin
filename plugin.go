package mailplugin

import (
	"fmt"
	"github.com/drharryhe/has/common/herrors"
	"github.com/drharryhe/has/common/hlogger"
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/core"
	"github.com/jordan-wright/email"
	"io/ioutil"
	smtp2 "net/smtp"
	"strings"
)

type Plugin struct {
	core.BasePlugin

	conf       MailPlugin
}

var plugin = &Plugin{}

func New() *Plugin {
	return plugin
}

func (this *Plugin) Open(s core.IServer, ins core.IPlugin) *herrors.Error {
	if err := this.BasePlugin.Open(s, ins); err != nil {
		return err
	}
	return nil
}

func (this *Plugin) Config() core.IEntityConf {
	return &this.conf
}

func (this *Plugin) EntityStub() *core.EntityStub {
	return core.NewEntityStub(
		&core.EntityStubOptions{
			Owner: this,
		})
}

func (this *Plugin) Capability() htypes.Any {
	return this
}

func (this *Plugin) SendHTMLTmp(to, subject string, tmpStr map[string]string) {
	html, err := ioutil.ReadFile(this.conf.HTMLPath)
	if err != nil {
		hlogger.Error(err)
		return
	}
	htmlContent := string(html)
	if tmpStr != nil {
		for tmp, s := range tmpStr {
			htmlContent = strings.ReplaceAll(htmlContent, tmp, s)
		}
	}
	em := email.NewEmail()
	em.From = this.conf.FromMail
	em.To = []string{to}
	em.HTML = []byte(htmlContent)
	em.Subject = subject
	if err := em.Send(fmt.Sprintf("%s:%d", this.conf.SMTP.Server, this.conf.SMTP.Port), smtp2.PlainAuth("", this.conf.FromMail, this.conf.Pwd, this.conf.SMTP.Server)); err != nil {
		hlogger.Error(err)
		return
	}
}