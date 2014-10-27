package controllers

import "github.com/astaxie/beego"

type CreateController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *CreateController) Get() {
	this.TplNames = "login.html"
}

func (this *CreateController) Join() {

	beego.Trace("enter join")

	// Get form value.
	uname := this.GetString("uname")
	tech := this.GetString("tech")
	// Check valid.
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	switch tech {
	case "Long Polling", "长轮询":
		this.Redirect("/lp?uname="+uname, 302)
	case "WebSocket":
		this.Redirect("/ws?uname="+uname, 302)
	default:
		this.Redirect("/", 302)
	}

	// Usually put return after redirect.
	return
}
