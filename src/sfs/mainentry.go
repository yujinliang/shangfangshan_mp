// httpserver.go
package main

import (
	
	"net/http"
	"sfs/config"
	"sfs/controller"
	"github.com/yujinliang/wechat/mp"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/session"
	
)

//global vars.
var (
	
	globalSessions *session.Manager
	
)
func init() {
	
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go globalSessions.GC()
	
}
//-----------
//切换至不同的域名 start
type HostSwitch map[string]http.Handler
// Implement the ServerHTTP method on our new type
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Check if a http.Handler is registered for the given host.
    // If yes, use it to handle the request.
    if handler := hs[r.Host]; handler != nil {
        handler.ServeHTTP(w, r)
    } else {
        // Handle host names for wich no handler is registered
        http.Error(w, r.Host, 403) // Or Redirect?
    }
}
//切换至不同的域名 end
func main() {
	
	wx := mp.New(config.Token, config.AppId, config.AppSecret ,config.EncodingAESKey, "")
	wx.HandleFunc(mp.MsgTypeText,  controller.HandleTextMsg)
	wx.HandleFunc(mp.MsgTypeVoice, controller.HandleVoiceMsg)
	wx.HandleFunc(mp.MsgTypeImage, controller.HandleImgeMsg)
	wx.HandleFunc(mp.MsgTypeVideo, controller.HandleVideoMsg)
	wx.HandleFunc(mp.MsgTypeLocation, controller.HandleLocationMsg)
	wx.HandleFunc(mp.MsgTypeLink, controller.HandleLinkMsg)
	//event.
	wx.HandleFunc(mp.GenHttpRouteKey(mp.MsgTypeEvent, mp.EventSubscribe), controller.HandleSubscribeEvent)
	wx.HandleFunc(mp.GenHttpRouteKey(mp.MsgTypeEvent, mp.EventUnsubscribe), controller.HandleUnSubscribeEvent)
	wx.HandleFunc(mp.GenHttpRouteKey(mp.MsgTypeEvent, mp.EventScan), controller.HandleScanEvent)
	wx.HandleFunc(mp.GenHttpRouteKey(mp.MsgTypeEvent, mp.EventLocation), controller.HandleLocationEvent)
	wx.HandleFunc(mp.GenHttpRouteKey(mp.MsgTypeEvent, mp.EventClick), controller.HandleMenuClickEvent)
	wx.HandleFunc(mp.GenHttpRouteKey(mp.MsgTypeEvent, mp.EventView), controller.HandleMenuViewEvent)
	
	//mux bind.
	controller.WX = wx
	controller.SNs = globalSessions
	router := httprouter.New()
	router.GET("/fbao_entry", controller.FBaoEntry)
	router.GET("/static/*filepath", controller.Static)
	router.POST("/admin/sendmass_msg", controller.MassMsg2WeinXinUser)
	router.POST("/do_forder", controller.JieYuanFABAO_Order)
	router.GET("/chest/:id",controller.Add2TreasureChest)
	router.GET("/getfbao_list", controller.GetFBaoList)
	router.GET("/getfbao_detail", controller.GetFBaoDetailInfo)
	router.GET("/getfbao_list_from_chest", controller.GetFBaoListFromChest)
	router.GET("/get_q7_list", controller.GetQ7List)
	router.GET("/get_q7_detail", controller.GetQ7DetailInfo)
	//admin---
	router.POST("/login", controller.Login)
	
	//mux chain.
	muxchain := make(HostSwitch)
	muxchain["wechat.jinliangyu_weinxin_dev.tunnel.mobi"] = wx
	muxchain["webapp.jinliangyu_weinxin_dev.tunnel.mobi"] = router
	
	n := negroni.Classic()
	n.UseHandler(muxchain)
	
	//create menu.
	controller.CreateMenu(wx)
	
	//launch weixin http server.
	n.Run(":8080")
	
}
