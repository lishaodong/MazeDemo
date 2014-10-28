package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lishaodong/MazeDemo/models"
)

type SearchController struct {
	baseController // Embed to use methods that are implemented in baseController.
	IP string
}

func (this *SearchController) Get() {
	beego.Trace("In search, model IP:",models.IP)
	if ((models.IP!="")) {
		//TODO:open another page

		this.Redirect("http://"+ models.IP +":8081/login",302)
		beego.Trace(models.IP)
	}else{
		this.Redirect("/empty",302)
	}
}


