package internal


import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"github.com/lishaodong/MazeDemo/controllers"
)
type Launcher struct {
	APP_VER string
}
func NewLauncher() (*Launcher) {
	return &Launcher{
		APP_VER:"0.1.1.0227",
	}
}

func (l *Launcher) Run(){
	beego.Info(beego.AppName, l.APP_VER)
	SearchController  := &controllers.SearchController{}

	// Register routers.
	beego.Router("/", &controllers.AppController{},"get:Get;post:Get")

	beego.Router("/empty", &controllers.AppController{},"get:Empty")
// Indicate AppController.Join method to handle POST requests.
	beego.Router("/create", &controllers.CreateController{},"post:Get")

	beego.Router("/login", &controllers.CreateController{},"post:Join")

	beego.Router("/login", &controllers.CreateController{},"get:Get")

	beego.Router("/search", SearchController,"post:Get")



// Long polling.
beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
beego.Router("/lp/post", &controllers.LongPollingController{})
beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

// WebSocket.
beego.Router("/ws", &controllers.WebSocketController{})
beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

// Register template functions.
beego.AddFuncMap("i18n", i18n.Tr)

beego.Run()
}
