package service

import (
	"HiChat/common"
	"HiChat/dao"
	"HiChat/middlewear"
	"HiChat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"mseeage": "获取用户列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func LoginByNameAndPasseWord(ctx *gin.Context) {
	var name = ctx.PostForm("name")
	password := ctx.PostForm("password")
	data, err := dao.FindUserByName(name)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1, //0 表示成功， -1 表示失败
			"message": "登录失败",
		})
		return
	}
	ok := common.CheckPassWord(password, data.Salt, data.Password)
	if !ok {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "密码错误",
		})
		return
	}
	token, err := middlewear.GenerateToken(data.ID, "yk")
	if err != nil {
		zap.S().Info("生成token失败", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"tokens":  token,
		"userId":  data.ID,
	})
}

// @Summary	添加账号
// @Produce	json
// @Param		name	formData		string	true	"用户名"
// @Param		password		formData		string	true	"密码"
// @Param		Identity		formData		string	true	"重复密码"
// @Success	200			{object}	string	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/user/new [post]
func NewUser(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Request.FormValue("name")
	password := ctx.Request.FormValue("password")
	repassword := ctx.Request.FormValue("Identity")
	if user.Name == "" || password == "" || repassword == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "用户名或密码不能为空！",
			"data":    user,
		})
		return
	}
	_, err := dao.FindUser(user.Name)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户已注册",
			"data":    user,
		})
		return
	}

	if password != repassword {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "两次密码不一致！",
			"data":    user,
		})
		return
	}

	t := time.Now()
	salt := fmt.Sprintf("%d", rand.Int31())
	user.Password = common.SaltPassWord(password, salt)
	user.Salt = salt
	user.LoginTime = &t
	user.LoginOutTime = &t
	user.HeartBeatTime = &t
	dao.CreateUser(user)
	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功",
		"data":    user,
	})

}

func UpdateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(ctx.Request.FormValue("id"))
	if err != nil {
		zap.S().Info("类型转换失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "注销账号失败",
		})
		return
	}
	user.ID = uint(id)
	Name := ctx.Request.FormValue("name")
	PassWord := ctx.Request.FormValue("password")
	Email := ctx.Request.FormValue("email")
	Phone := ctx.Request.FormValue("phone")
	avatar := ctx.Request.FormValue("icon")
	gender := ctx.Request.FormValue("gender")
	if Name != "" {
		user.Name = Name
	}
	if PassWord != "" {
		salt := fmt.Sprintf("%d", rand.Int31())
		user.Salt = salt
		user.Password = common.SaltPassWord(PassWord, salt)
	}
	if Email != "" {
		user.Email = Email
	}
	if Phone != "" {
		user.Phone = Phone
	}
	if avatar != "" {
		user.Avatar = avatar
	}
	if gender != "" {
		user.Gender = gender
	}
	updateUser, err := dao.UpdateUser(user)
	if err != nil {
		zap.S().Info("更新用户失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "修改信息失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "修改成功",
		"data":    updateUser.Name,
	})
}

func DeleteUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(ctx.Request.FormValue("id"))
	if err != nil {
		zap.S().Info("类型转换失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "注销账号失败",
		})
		return
	}

	user.ID = uint(id)
	err = dao.DeleteUser(user)
	if err != nil {
		zap.S().Info("注销用户失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "注销账号失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "注销账号成功",
	})

}

func GroupList(ctx *gin.Context) {
	owner := ctx.PostForm("ownerId")
	ownerId, err := strconv.Atoi(owner)
	if err != nil {
		zap.S().Info("owner类型转换失败", err)
		return
	}

	if ownerId == 0 {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "您未登录",
		})
		return
	}

	rsp, err := dao.GetCommunityList(uint(ownerId))
	if err != nil {
		zap.S().Info("获取群列表失败", err)
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "你还没加入任何群聊",
		})
		return
	}

	common.RespOKList(ctx.Writer, rsp, len(*rsp))
}

func JoinGroup(ctx *gin.Context) {
	comInfo := ctx.PostForm("comId")
	if comInfo == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "群名称不能为空",
		})
		return
	}

	user := ctx.PostForm("userId")
	userId, err := strconv.Atoi(user)
	if err != nil {
		zap.S().Info("user类型转换失败")
	}
	if userId == 0 {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "你未登录",
		})
		return
	}

	code, err := dao.JoinCommunity(uint(userId), comInfo)
	if err != nil {
		HandleErr(code, ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "加群成功",
	})
}

func HandleErr(code int, ctx *gin.Context, err error) {
	switch code {
	case -1:
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": err.Error(),
		})
	case 0:
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "该好友已经存在",
		})
	case -2:
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "不能添加自己",
		})

	}
}

func NewGroup(ctx *gin.Context) {
	owner := ctx.PostForm("ownerId")
	ownerId, err := strconv.Atoi(owner)
	if err != nil {
		zap.S().Info("owner类型转换失败", err)
		return
	}

	ty := ctx.PostForm("cate")
	Type, err := strconv.Atoi(ty)
	if err != nil {
		zap.S().Info("ty类型转换失败", err)
		return
	}

	img := ctx.PostForm("icon")
	name := ctx.PostForm("name")
	desc := ctx.PostForm("desc")

	community := models.Community{}
	if ownerId == 0 {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "您未登录",
		})
		return
	}

	if name == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "群名称不能为空",
		})
		return
	}

	if img != "" {
		community.Image = img
	}
	if desc != "" {
		community.Desc = desc
	}

	community.Name = name
	community.Type = Type
	community.OwnerId = uint(ownerId)

	code, err := dao.CreateCommunity(community)
	if err != nil {
		HandleErr(code, ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "键群成功",
	})
}

// @Summary	添加好友
// @Produce	json
// @Param		targetName	body		string	true	"用户名"
// @Param		userId		body		string	true	"用户id"
// @Success	200			{object}	string	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/v1/relation/add [post]
func AddFriendByName(ctx *gin.Context) {
	user := ctx.PostForm("userId")
	userId, err := strconv.Atoi(user)
	if err != nil {
		zap.S().Info("类型转换失败", err)
		return
	}

	tar := ctx.PostForm("targetName")
	target, err := strconv.Atoi(tar)
	if err != nil {
		code, err := dao.AddFriendByName(uint(userId), tar)
		if err != nil {
			HandleErr(code, ctx, err)
			return
		}

	} else {
		code, err := dao.AddFriend(uint(userId), uint(target))
		if err != nil {
			HandleErr(code, ctx, err)
			return
		}
	}
	ctx.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "添加好友成功",
	})
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
