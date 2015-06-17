// httpserver.go
package controller

import (
	

	"log"
	"fmt"
	"sfs/config"
	"net/http"
	"encoding/json"
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
		
	replyText := wx.ReplyText("http://http://webapp.jinliangyu_weinxin_dev.tunnel.mobi/hello", r)
	w.Write([]byte(replyText))

}
func HandleTextMsg(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	q7_OAuthConfig := oauth2web.NewOAuth2Config(wx.GetAppId(), wx.GetAppSecret(), config.WebHostUrl + "/q7_entry", "snsapi_base")
	q7_url := q7_OAuthConfig.AuthCodeURL("q7_entry")
	replyText := wx.ReplyText(q7_url, r)
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
	//generate auth url.
	q7_OAuthConfig := oauth2web.NewOAuth2Config(wx.GetAppId(), wx.GetAppSecret(), config.WebHostUrl + "/q7_entry", "snsapi_base")
	q7_url := q7_OAuthConfig.AuthCodeURL("q7_entry")
	menu.Buttons[0].Url	 = q7_url
	//---
	menu.Buttons[1].Name = "论坛"
	menu.Buttons[1].Type = mp.ButtonTypeView
	menu.Buttons[1].Url  = "https://mp.weixin.qq.com"
	//---
	menu.Buttons[2].Name = "结缘"
	menu.Buttons[2].SubButtons = make([]mp.MenuButton, 2)
	menu.Buttons[2].SubButtons[0].Name = "结缘法宝"
	menu.Buttons[2].SubButtons[0].Type = mp.ButtonTypeView
	fbao_OAuthConfig := oauth2web.NewOAuth2Config(wx.GetAppId(), wx.GetAppSecret(), config.WebHostUrl + "/fbao_entry", "snsapi_base")
	fbao_url := fbao_OAuthConfig.AuthCodeURL("fbao_entry")
	menu.Buttons[2].SubButtons[0].Url	= fbao_url
	//---
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
func D7_Apply(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	fmt.Fprintf(w, "%v", r.Form)
}
func Add2TreasureChest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	log.Printf("Add2TreasureChest: %s", ps.ByName("id"))
	var res struct {
			
		Id      string `json:"id"`
		ErrCode int 	`json:"errcode"`
		ErrMsg  string `json:"errmsg"`	
		
	}
	res.Id = ps.ByName("id");
	res.ErrCode = 0;
	res.ErrMsg  = "成功啦！，法宝已收入百宝箱中！"
		
	encoded, err := json.Marshal(&res);
	if err != nil {
			
		fmt.Fprintf(w, "{id:%s,errcode:%d,errmsg:%s}", res.Id, 1, "失败啦！没能加入百宝箱!")
		return
			
	}
		
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(encoded)
	
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, _ := SNs.SessionStart(w,r)
	defer session.SessionRelease(w)
	
	openid := session.Get("openid")
	
	fmt.Fprintf(w, "%s", openid)
	
}
func FBaoEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	redirect2targetWithOpenId(w, r, config.WebHostUrl + "/static/html/fbao_list.html")
	
}
func Q7Entry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	redirect2targetWithOpenId(w, r, config.WebHostUrl + "/static/html/q7_list.html")
	
}
func redirect2targetWithOpenId(w http.ResponseWriter, r *http.Request, targetUrl string) {
	
	//get code.
	r.ParseForm()
	code  := r.FormValue("code")
	//state := r.FormValue("state")
	log.Print("code: " + code)
	
	//get access token.
	oauthConfig := oauth2web.NewOAuth2Config(config.AppId, config.AppSecret, config.WebHostUrl + "/fbao_entry", "snsapi_base")
	oClient := &oauth2web.Client{OAuth2Config:oauthConfig}
	oClient.ExchangeOAuth2AccessTokenByCode(code)
	info, _ := oClient.UserInfo("zh_CN")
	
	//store openid to session.
	session, _ := SNs.SessionStart(w,r)
	defer session.SessionRelease(w)
	session.Set("openid", info.OpenId)
	
	log.Print("openid: " + info.OpenId)
	
	//redirect to fbao_list page.
	http.Redirect(w, r, targetUrl + "?openid=" + info.OpenId, http.StatusFound)
	
}

//微网站 end
