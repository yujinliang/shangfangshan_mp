// httpserver.go
package controller

import (
	

	"log"
	"fmt"
	"sfs/config"
	"net/http"
	"github.com/yujinliang/wechat/mp"
	"github.com/yujinliang/wechat/mp/request"
	"github.com/yujinliang/wechat/mp/oauth2web"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/session"
	
)

//just define global vars.
var (
	
	WX *mp.WeiXin = nil
	SNs *session.Manager
	
)
func HandleVoiceMsg(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
		
	//test oauth2
	oauthConfig := oauth2web.NewOAuth2Config(wx.GetAppId(), wx.GetAppSecret(), "http://webapp.jinliangyu_weinxin_dev.tunnel.mobi/showuserinfo", "snsapi_userinfo")
	oauthUrl := oauthConfig.AuthCodeURL("testOauth2")
	//oClient := &oauth2web.Client{OAuth2Config:oauthConfig}
	//oClient.CheckAccessTokenValid()
	
	replyText := wx.ReplyText(oauthUrl, r)
	w.Write([]byte(replyText))

}
func HandleTextMsg(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("文本消息!", r)
	//data, _ := wx.MakeEncryptResponse([]byte(replyText), timestamp, nonce)
	w.Write([]byte(replyText))
	
	//send custom message
	//wx.PostText(r.FromUserName, "我是客服消息， 你好！", "")
	
}
func HandleImgeMsg(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("图片消息!", r)
	w.Write([]byte(replyText))
}
func HandleVideoMsg(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("视频消息!", r)
	w.Write([]byte(replyText))
}
func HandleLocationMsg(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("位置消息!", r)
	w.Write([]byte(replyText))
}
func HandleLinkMsg(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("link消息!", r)
	w.Write([]byte(replyText))
}
func HandleSubscribeEvent(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("订阅事件!", r)
	w.Write([]byte(replyText))
}
func HandleUnSubscribeEvent(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("取消订阅事件!", r)
	w.Write([]byte(replyText))
}
func HandleScanEvent(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("扫二维码事件！", r)
	w.Write([]byte(replyText))
}
func HandleLocationEvent(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("位置事件!", r)
	w.Write([]byte(replyText))
}
func HandleMenuClickEvent(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	var info string
	
	switch r.EventKey {
		
		case "da_7": {
			
			//返回地藏七种类列表， 点进详情可进入每一个“七”的详情页，并中有报名入口
			info = "1.基础七  <a href='http://webapp.jinliangyu_weinxin_dev.tunnel.mobi/static/html/qi_detail.html'>详情</a>\n\n2.老年七\n\n3.备孕七\n\n4.排毒七"
			
		}
		case "lianxidaochang": {
			
			//直接返回配置好的道场联系方式信息串.
			info = "小明师兄:13688886666\n\n小红师兄:18623458907"
			
		}
	}
	replyText := wx.ReplyText(info, r)
	w.Write([]byte(replyText))
	
}
func HandleMenuViewEvent(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	replyText := wx.ReplyText("打开网页事件!", r)
	w.Write([]byte(replyText))
}
//test menu
func CreateMenu(wx *mp.WeiXin) {
	
	menu := &mp.Menu{make([]mp.MenuButton,3)}
	menu.Buttons[0].Name = "我要打七"
	menu.Buttons[0].Type = mp.ButtonTypeView
	menu.Buttons[0].Url	 = "http://webapp.jinliangyu_weinxin_dev.tunnel.mobi/static/html/q7_list.html"
	menu.Buttons[1].Name = "论坛"
	menu.Buttons[1].Type = mp.ButtonTypeView
	menu.Buttons[1].Url  = "https://mp.weixin.qq.com"
	menu.Buttons[2].Name = "结缘"
	menu.Buttons[2].SubButtons = make([]mp.MenuButton, 2)
	menu.Buttons[2].SubButtons[0].Name = "结缘法宝"
	menu.Buttons[2].SubButtons[0].Type = mp.ButtonTypeView
	menu.Buttons[2].SubButtons[0].Url	= "http://webapp.jinliangyu_weinxin_dev.tunnel.mobi/static/html/fbao_list.html"
	menu.Buttons[2].SubButtons[1].Name = "联系道场"
	menu.Buttons[2].SubButtons[1].Type = mp.ButtonTypeView
	menu.Buttons[2].SubButtons[1].Url	= "http://webapp.jinliangyu_weinxin_dev.tunnel.mobi/static/html/contact_info.html"
	
	err := wx.CreateMenu(menu)
	
	if err != nil {
		
		log.Println(err)
		
	}
}
//-----------

//微网站 start
//专门处理静态文件,如：.html, .jpg, .js等
func Static(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	filePath := config.StaticResourcesDir +  ps.ByName("filepath")
	//log.Printf("filePath:%s, urlPath:%s, host:%s", filePath, r.URL.Path, r.Host)
	http.ServeFile(w, r, filePath)
		

}
func GetCurrentId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
//	session := sessions.GetSession(r)
//	id_type , _ := session.Get("id_type").(string)
//	id      , _ := session.Get("id").(string)
//	fmt.Printf("%s, %s", id_type, id)
//	var res struct {
			
//		Id_type string `json:"id_type"`
//		Id      string `json:"id"`
			
//	}
//	res.Id = id;
//	res.Id_type = id_type
		
//	encoded, err := json.Marshal(&res);
//	if err != nil {
			
//		fmt.Fprintf(w, "{id_type:%s, id:%s}", "", "")
//		return
			
//	}
		
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Write(encoded)
				
}
//err_code: [0:成功， 1:我方服务器问题，2: 微信方问题]
func MassMsg2WeinXinUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	r.ParseForm()
	fmt.Print(r.FormValue("q7_text"))
	//---
	if WX != nil {
		
		msgid, err := WX.SendTextByGroupID("0", r.FormValue("q7_text"), false)
		if err != nil {
			
			fmt.Fprintf(w, "{msg_id:%s, err_code:%d, err_msg:%s}","" , 2, err)
			
		} else {
			
			fmt.Fprintf(w, "{msg_id:%s, err_code:%d, err_msg:%s}", msgid, 0, "SUCCESS")
			
		}
		
			
	} else {
		
		fmt.Fprintf(w, "{msg_id:%s, err_code:%d, err_msg:%s}","", 1, "FAILED")
		
	}

}
func JieYuanFABAO_Order(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	r.ParseForm()
	fmt.Fprintf(w, "%v", r.Form)
}
func JieYuanFABAO_Prepare_Order(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	r.ParseForm()
	fmt.Fprintf(w, "id : %s", ps.ByName("current_fbao_id"))
}
func D7_Apply(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	fmt.Fprintf(w, "%v", r.Form)
}
func Add2TreasureChest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	fmt.Fprintf(w, "id: %s", ps.ByName("current_fbao_id"))
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, _ := SNs.SessionStart(w,r)
	defer session.SessionRelease(w)
	
	iName := session.Get("name")
	if iName != nil {
		
		str, ok := iName.(string)
		if ok && str == ps.ByName("name") {
			
			fmt.Fprintf(w, "同主机，同名字，您又来了，%s", str)
			
		} else {
			
			session.Set("name", ps.ByName("name"))
			fmt.Fprintf(w, "新来的的名字, %s", ps.ByName("name"))
			
		}
		
	} else {
		
		session.Set("name", ps.ByName("name"))
		fmt.Fprintf(w, "新来的主机, %s", ps.ByName("name"))
		
	}
}
func ShowUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	r.ParseForm()
	code  := r.FormValue("code")
	state := r.FormValue("state")
	fmt.Fprintf(w, "code: %s, state: %s\n", code, state)
	//get access token.
	oauthConfig := oauth2web.NewOAuth2Config(config.AppId, config.AppSecret, "http://webapp.jinliangyu_weinxin_dev.tunnel.mobi/showuserinfo", "snsapi_userinfo")
	oClient := &oauth2web.Client{OAuth2Config:oauthConfig}
	oClient.ExchangeOAuth2AccessTokenByCode(code)
	info, _ := oClient.UserInfo("zh_CN")
	fmt.Fprintf(w, "openid:%s, nickname:%s, sex:%s, city:%s, province:%s, country:%s,UnionId:%s, HeadImageURL:%s, Privilege:%v", info.OpenId,info.Nickname, info.Sex, info.City, info.Province, info.Country, info.UnionId, info.HeadImageURL,info.Privilege )
	
}

//微网站 end
