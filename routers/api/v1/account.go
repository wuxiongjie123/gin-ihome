package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ihome_gin/models"
	"ihome_gin/pkg/e"
	"regexp"
)

type RegReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	SmsCode  string `json:"sms_code"`
}

type loginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// 用户注册
func Register(c *gin.Context) {
	var req RegReq
	err := c.BindJSON(&req)
	if err != nil {
		SendResp(c, e.RECODE_PARAMERR, nil)
		return
	}
	if req.Mobile == "" || req.Password == "" || req.SmsCode == "" {
		SendResp(c, e.RECODE_PARAMERR, nil)
		return
	}
	if ok, _ := regexp.MatchString(`1[345678]\d{9}`, req.Mobile); !ok {
		SendResp(c, e.RECODE_PARAMERR, nil)
		return
	}

	// TODO 获取短信验证码,查看是否过期 使用redis保存验证码 并设置设置时间

	// 查询用户是否存在
	var acc *models.Account
	acc, err = models.SelectAccount(req.Mobile)
	if err!=nil{
		SendResp(c, e.RECODE_DBERR, nil)
		return
	}
	if acc.Mobile != "" && acc.Mobile == req.Mobile {
		SendResp(c,e.RECODE_MOBILEERR,nil)
		return
	}

	// 创建新的用户 加密密码
	pwd := models.PassWordMd5(req.Password)
	user := models.Account{
		Mobile:       req.Mobile,
		PasswordHash: pwd,
	}
	if err := user.CreateAccount(); err != nil {
		SendResp(c, e.RECODE_DBERR, err)
		return
	}

	// 设置session
	session := sessions.Default(c)
	session.Set("user_id", acc.UserId)
	if acc.Name == "" {
		session.Set("user_name", user.Mobile)
	} else {
		session.Set("user_name", user.Name)
	}
	session.Set("mobile", user.Mobile)
	_ = session.Save()

	SendResp(c, e.RECODE_OK, req.Mobile)
}

// 用户登录
func Login(c *gin.Context) {
	var req loginReq
	err := c.BindJSON(&req)
	if err != nil {
		SendResp(c, e.RECODE_PARAMERR, nil)
		return
	}
	if req.Mobile == "" || req.Password == "" {
		SendResp(c, e.RECODE_PARAMERR, nil)
		return
	}

	// 检查手机
	if ok, _ := regexp.MatchString(`1[345678]\d{9}`, req.Mobile); !ok {
		SendResp(c, e.RECODE_PARAMERR, nil)
		return
	}

	// TODO 判断用户的错误次数,从redis中获取错误次数,大于5次返回登录过于频繁

	// 查询数据库,判断用户信息与密码
	acc, err := models.SelectAccount(req.Mobile)
	if err != nil {
		SendResp(c, e.RECODE_DBERR, nil)
		return
	}

	// 保存的是hash密码,应该解密后对比
	pwd := models.PassWordMd5(req.Password)
	if acc.PasswordHash != pwd {
		SendResp(c, e.RECODE_LOGINERR, nil)
		return
	}

	SendResp(c, e.RECODE_OK, nil)
}
