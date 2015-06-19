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
	
	//TODO:
	//1. 首先检查session中是否有chest， 若无则返回出错信息: 宝箱为空.
	//2. 若session中存在chest， 则将form信息，及chest信息一并存入数据库相应表中.
	//3. 写入数据库成功后， 清除session中的chest， 然后返回提示成功信息.
	//4. 若写入数据库失败， 则不清除session中的chest（以供用户重试）然后返回出错信息.
	r.ParseForm()
	var res struct {
			
		Id      string `json:"id"`
		ErrCode int 	`json:"errcode"`
		ErrMsg  string `json:"errmsg"`	
		
	}
	res.Id = "10001" //订单号
	res.ErrCode = 0;
	res.ErrMsg  = "成功啦！，法宝结缘成功！"
		
	encoded, err := json.Marshal(&res);
	if err != nil {
			
		fmt.Fprintf(w, "{id:%s,errcode:%d,errmsg:%s}", res.Id, 1, "失败啦！法宝没能结缘成功!")
		return
			
	}
		
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(encoded)
	
}
func D7_Apply(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	//TODO:
	//1. 首先通过手机号，与身份证号在数据库中查找， 用户是否已报过名了，且开七时间为未来时间。
	//2. 若报过了， 则提示：不可重报.
	//3. 若未报过名， 则写入数据库， 然后返回成功提示信息.
	fmt.Fprintf(w, "%v", r.Form)
}
func Add2TreasureChest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	log.Printf("Add2TreasureChest: %s", ps.ByName("id"))
	//TODO: 
	//1.首先检查此法宝是否已在chest中， 若在， 则返回错误信息，提示法宝早已在宝箱中，不可再重复添加.
	//2.若为新法宝，则存入session中，以备后用.
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
func GetFBaoDetailInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	fbao_id := r.URL.Query().Get("fbao_id")
	log.Printf("fbao_id: %s", fbao_id)
	
	//---
	type FBAO_INFO struct {
		
		Id   string `json:"id"`
		Name string `json:"name"`
		Desc string `json:"desc,omitempty"`
		ImageNames []string `json:"imagenames,omitempty"`
		
	}
	
	var info FBAO_INFO
	info.Id = fbao_id
	info.Name = "念佛成佛"
	info.Desc = "念佛是因， 成佛是果，不可思议！信愿行三者不可缺一也！若真是笃信，当即放下娑婆万有，志求西方，得生彼国！则此生不枉过也!!!"
	info.ImageNames = []string{"1.jpg","2.jpg"}
	
	encoded, err := json.Marshal(&info)
	if err != nil {
		
		fmt.Fprintf(w, "{}")
		
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(encoded)
	
}
func GetFBaoList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	//第几页.
	page := r.URL.Query().Get("page")
	//每一页的法宝数.
	count := r.URL.Query().Get("count")
	
	log.Printf("Page: %d, Count: %d", page, count)
	
	//--
	type FBAO struct {
		
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	
	fbaoSlice := make([]*FBAO, 8)
	
	for i, _ := range fbaoSlice {
		
		fbaoSlice[i] = &FBAO{Id:"1001", Name: "念佛成佛"}
		
	}
	
	encoded, err := json.Marshal(&fbaoSlice);
	if err != nil {
			
		fmt.Fprintf(w, "[]")
		return
			
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(encoded)
	
}
//从百宝箱中取出法宝列表,引接口返回的信息只是用于展示.
func GetFBaoListFromChest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
		//--
	type FBAO struct {
		
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	
	fbaoSlice := make([]*FBAO, 8)
	
	for i, _ := range fbaoSlice {
		
		fbaoSlice[i] = &FBAO{Id:"1001", Name: "念佛成佛"}
		
	}
	
	encoded, err := json.Marshal(&fbaoSlice);
	if err != nil {
			
		fmt.Fprintf(w, "[]")
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
	
	//check if we had openid.
	var openid string
	session, _ := SNs.SessionStart(w,r)
	defer session.SessionRelease(w)
	iId := session.Get("openid")
	
	if iId == nil {
	
		//get code.
		r.ParseForm()
		code  := r.FormValue("code")
		//state := r.FormValue("state")
	
		//get access token.
		oauthConfig := oauth2web.NewOAuth2Config(config.AppId, config.AppSecret, config.WebHostUrl + "/fbao_entry", "snsapi_base")
		oClient := &oauth2web.Client{OAuth2Config:oauthConfig}
		oClient.ExchangeOAuth2AccessTokenByCode(code)
		info, _ := oClient.UserInfo("zh_CN")
		
		if info != nil && len(info.OpenId) > 0 {
			
			openid = info.OpenId
			//store openid to session.
			session.Set("openid", info.OpenId)
		
		}
		
	
	} else {
		
		openid = iId.(string)
			
	}
		
	//redirect to fbao_list page.
	http.Redirect(w, r, targetUrl + "?openid=" + openid, http.StatusFound)
	
}

//微网站 end
