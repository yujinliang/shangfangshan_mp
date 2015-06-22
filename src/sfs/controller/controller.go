// httpserver.go
package controller

import (
	
	//"io"
	//"os"
	"bufio"
	"log"
	"fmt"
	"strconv"
	"strings"
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
	
	//TODO:保存至关注者列表，用于群发消息.
	replyText := wx.ReplyText("订阅事件!", r)
	w.Write([]byte(replyText))
}
func HandleUnSubscribeEvent(wx *mp.WeiXin, w http.ResponseWriter, r *request.WeiXinRequest, timestamp, nonce string) {
	
	//TODO:从关注者列表中删除.
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
//menu
func CreateMenu(wx *mp.WeiXin) {
	

	
	menu := &mp.Menu{make([]mp.MenuButton,3)}
	menu.Buttons[0].Name = "我要打七"
	menu.Buttons[0].Type = mp.ButtonTypeView
	menu.Buttons[0].Url	 = config.WebHostUrl + "/static/html/q7_list.html"
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
	menu.Buttons[2].SubButtons[1].Url	= config.WebHostUrl + "/static/html/contact_info.html"
	
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
	
	//1.是否已登录。
//	if checkAuth(w, r) == false {
		
//		http.Redirect(w, r, config.WebHostUrl + "/static/html/admin_login.html", http.StatusFound)
		
//	}
	//2.检查参数. 
	r.ParseMultipartForm(10 << 20)
	msg_type := r.FormValue("mass_type_q7sendmassmsg")
	if msg_type == config.WeiXinMassMsgNews {//news type.
		
		title := r.FormValue("title_mass_tuwen")
		content := r.FormValue("text_mass_tuwen")
		if len(title) <= 0 || len(content) <= 0 {
			
			fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",1 ,"No title or content!")
			return
			
		}
		
		author := r.FormValue("author_mass_tuwen")
		digest := r.FormValue("digest_mass_tuwen")
		content_source_url := r.FormValue("sourceurl_mass_tuwen")	
		show_pic := r.FormValue("showcoverpic_mass_tuwen")
		is_to_all := r.FormValue("toall_mass_tuwen")
	
		//upload file handle block.
		file, handler, err := r.FormFile("thumbpic_mass_tuwen")
		if err != nil {
		
			fmt.Fprintf(w, "{errcode:%d, errmsg:%s}", 1, err)
			return
		
		}
		defer file.Close()
	
		//3.通过微信素材管理上传thumb, 成功后获得mediaId, thumb, image
		if WX != nil {
		
			mediaId, err := WX.UploadTmpMedia("image", strings.ToLower(handler.Filename), bufio.NewReader(file))
			if err != nil {
			
				fmt.Fprintf(w, "{errcode:%d, errmsg:%s}", 2, err)
				return
			
			}
			//4. 上传图文消息.
			var news mp.MPNews;
			news.Title = title
			news.ThumbMediaId = mediaId
			news.Author = author
			news.Digest = digest
			news.Content = content
			news.ContentSourceUrl = content_source_url
			picN, err := strconv.Atoi(show_pic)
			if err != nil {
			
				picN = 1
			
			}
			news.ShowCoverPic = int8(picN)
			materialId, err := WX.UploadNews([]mp.MPNews{news})
			if err != nil {
				
				fmt.Fprintf(w, "{errcode:%d, errmsg:%s}", 2, err)
				return
				
			}
			
			if len(materialId) <= 0 {
				
				fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",2 ,"No MaterialId")
				return
				
			}
			
			is_to_all_bool, err := strconv.ParseBool(is_to_all)
			if err != nil {
				
				is_to_all_bool = false
				
			}
			msgid, err := WX.SendNewsByGroupID(config.WeiXinDefaultUserGroupId, materialId, is_to_all_bool)
			if err != nil {
				
				fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",2 ,err)
				return
				
			}
			
			fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",0 ,"Success: " + msgid)
			return
			
		}
		
		fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",1 ,http.StatusText(http.StatusInternalServerError))
		return
		

	} else { //just text message.
		
		content := r.FormValue("text_mass_tuwen")
		if len(content) <= 0 {
			
			fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",1 ,"No Content!")
			return
			
		}
		
		if WX != nil {
			
			is_to_all := r.FormValue("toall_mass_tuwen")
			is_to_all_bool, err := strconv.ParseBool(is_to_all)
			if err != nil {
				
				is_to_all_bool = false
				
			}
			
			msgid, err := WX.SendTextByGroupID(config.WeiXinDefaultUserGroupId, content, is_to_all_bool)
			if err != nil {
				
				fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",2 ,err)
				return
				
			}
			
			fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",0 ,"Success: " + msgid)
			return
			
		}
		
		fmt.Fprintf(w, "{errcode:%d, errmsg:%s}",1 ,http.StatusText(http.StatusInternalServerError))
		return
		
	}

}
func JieYuanFABAO_Order(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	//TODO:
	//1. 首先检查session中是否有chest， 若无则返回出错信息: 宝箱为空.
	//2. 若session中存在chest， 则将form信息，及chest信息一并存入数据库相应表中.
	//3. 写入数据库成功后， 清除session中的chest， 然后返回提示成功信息.
	//4. 若写入数据库失败， 则不清除session中的chest（以供用户重试）然后返回出错信息.
	r.ParseForm()
	fmt.Printf("q7openid: %v", r.FormValue("q7openid"))
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
		Tip  string	 `json:"tip,omitempty"`
		ImageNames []string `json:"imagenames,omitempty"`
		
	}
	
	var info FBAO_INFO
	info.Id = fbao_id
	info.Name = "念佛成佛"
	info.Desc = "念佛是因， 成佛是果，不可思议！信愿行三者不可缺一也！若真是笃信，当即放下娑婆万有，志求西方，得生彼国！则此生不枉过也!!!"
	info.Tip  = "一定要恭敬法宝，不可污损，不可丢弃，不可放在杂物处，不可与世间书放在一起，不可放在卧室及污秽之处，总之见法宝如见佛，必需恭敬之！阅读时要着装正式大方，洗手洗脸，刷牙，端身正坐，双手捧着读诵!"
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
func GetQ7List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	//--
	type Q7Type struct {
		
		Id   string `json:"id"`
		Name string `json:"name"`
		
	}
	
	q7slice := make([]*Q7Type, 10)
	
	for i, _ := range q7slice {
		
		q7slice[i] = &Q7Type{Id:"1001", Name: "基础七"}
		
	}
	
	encoded, err := json.Marshal(&q7slice);
	if err != nil {
			
		fmt.Fprintf(w, "[]")
		return
			
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(encoded)
	
}
func GetQ7DetailInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	q7_id := r.URL.Query().Get("q7_id")	
	//---
	type Q7_INFO struct {
		
		Id   		string `json:"id"`
		Name 		string `json:"name"`
		Desc 		string `json:"desc,omitempty"`
		Q7LimitTip	string `json:"q7limit,omitempty"`
		Q7Plan		string `json:"q7plan"`
		ImageNames 	[]string `json:"imagenames,omitempty"`
		EnrollWay 	string `json:"enrollway"`
		
	}
	
	var info Q7_INFO
	info.Id = q7_id
	info.Name = "基础七"
	info.Desc = "学习佛法六部曲，吃素戒杀，拜忏，诵经，放生，日行一善，念佛；使身心皆得大利益,坚持长久修行，必然当生得见阿弥陀佛，得生西方极乐世界!"
	info.Q7LimitTip = "不接受生活不能自理者，有心脑血管等疾病者需家人陪同方可参加，并签署自愿免责协议."
	info.ImageNames = []string{"1.jpg","2.jpg"}
	info.EnrollWay = "光瑞师兄:13712348908"
	info.Q7Plan = "每月1号，9号，19号开七."
	
	encoded, err := json.Marshal(&info)
	if err != nil {
		
		fmt.Fprintf(w, "{}")
		
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(encoded)
	
}

func FBaoEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	redirect2targetWithOpenId(w, r, config.WebHostUrl + "/static/html/fbao_list.html")
	
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
//admin start
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	//1.检查输入参数
	r.ParseForm()
	userName := r.FormValue("user_name")
	userPwd	 := r.FormValue("user_pwd")
	if len(userName) <= 0 || len(userPwd) <= 0 {
		
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		
	}
	
	//2.到session中查找用户信息，找到则之前已登录， 否则为新登录.

	if checkAuth(w, r) == true {
		
		http.Redirect(w, r, config.WebHostUrl + "/static/html/admin.html", http.StatusFound)
		
	} else {
		
		//3. TODO:去数据库中，查找此user_name 并验证其密码,以及是否被disable.
		if userName == "yu" {
			
			//将管理员信息存入session.
			session, _ := SNs.SessionStart(w,r)
			defer session.SessionRelease(w)
			session.Set("user_name", userName)
			session.Set("admin_level", 0)
			//--
			http.Redirect(w, r, config.WebHostUrl + "/static/html/admin.html", http.StatusFound)
			
		} else {
			
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			
		}
	}
}
func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	
	session, _ := SNs.SessionStart(w,r)
	defer session.SessionRelease(w)
	user_name_in_session := session.Get("user_name")
	if user_name_in_session != nil {
		
		return true
		
	}
		
	return false
		
}
//admin end
//微网站 end
