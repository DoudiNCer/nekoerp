package service

import (
	"github.com/gin-gonic/gin"
	"nekoerp/pkg/service"
	"nekoerp/pkg/service/middleware"
	orm "nekoerp/pkg/storage/mysql"
	"nekoerp/pkg/storage/mysql/tables"
	"net/http"
)

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Role  string `json:"role"`
	Token string `json:"token"`
}

func UserLogin(ctx *gin.Context) {
	var req UserLoginReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "参数解析错误：" + err.Error()})
		return
	}
	var user tables.User
	orm.Db.First(&user, "account = ? && password = ? && blocked = false", req.Username, req.Password)
	if user.Id == 0 {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "用户名、密码错误或账户已禁用"})
		return
	}
	var role tables.Role
	orm.Db.First(&role, user.Role)
	token, err := middleware.GenerateJWT(middleware.Payload{
		Id:      user.Id,
		Account: user.Account,
		Role:    user.Role,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "服务器执行错误"})
		return
	}
	ctx.JSON(http.StatusOK, service.Response{Code: 0, Message: "请求正常", Data: UserLoginResp{
		Role:  role.Name,
		Token: token,
	}})
}

func UserList(ctx *gin.Context) {
	req := service.ListSearchReq{
		Token:   ctx.Query("token"),
		Keyword: ctx.Query("keyword"),
	}
	payload, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	if payload.Role > 1 {
		ctx.JSON(http.StatusForbidden, service.Response{Code: -1, Message: "权限不足"})
		return
	}
	var total int64
	orm.Db.Model(tables.User{}).Count(&total)
	var results []service.ListUserResp
	page := orm.Db.Table("user").Select("`user`.id, `user`.account, `user`.created_at, `user`.updated_at, `user`.blocked, role.`name` AS \"role\"").
		Joins("LEFT JOIN role ON `user`.role = role.id")
	if req.Keyword != "" {
		page = page.Where("`user`.account LIKE ? ", "%"+req.Keyword+"%")
	}
	page.Scan(&results)
	ctx.JSON(http.StatusOK, service.ResponseList{Total: uint(total), List: results})
}

func BlockUser(ctx *gin.Context) {
	var req service.OperateUserReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "参数解析错误：" + err.Error()})
		return
	}
	payload, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	if payload.Role > 1 {
		ctx.JSON(http.StatusForbidden, service.Response{Code: -1, Message: "权限不足"})
		return
	}
	var user tables.User
	orm.Db.First(&user, req.Id)
	if user.Account == "" {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "用户不存在"})
		return
	}
	user.Blocked = true
	orm.Db.Save(user)
	ctx.JSON(http.StatusOK, service.Response{
		Code:    0,
		Message: "请求正常",
	})
}

func UnBlockUser(ctx *gin.Context) {
	var req service.OperateUserReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "参数解析错误：" + err.Error()})
		return
	}
	payload, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	if payload.Role > 1 {
		ctx.JSON(http.StatusForbidden, service.Response{Code: -1, Message: "权限不足"})
		return
	}
	var user tables.User
	orm.Db.First(&user, req.Id)
	if user.Account == "" {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "用户不存在"})
		return
	}
	user.Blocked = false
	orm.Db.Save(user)
	ctx.JSON(http.StatusOK, service.Response{
		Code:    0,
		Message: "请求正常",
	})
}

func EditUser(ctx *gin.Context) {
	var req service.OperateUserReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "参数解析错误：" + err.Error()})
		return
	}
	payload, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	if payload.Role > 1 {
		ctx.JSON(http.StatusForbidden, service.Response{Code: -1, Message: "权限不足"})
		return
	}
	var user tables.User
	orm.Db.First(&user, req.Id)
	if user.Account == "" {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "用户不存在"})
		return
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Account != "" {
		user.Account = req.Account
	}
	if req.Role > 0 && req.Role < 3 {
		user.Role = req.Role
	}
	orm.Db.Save(user)
	ctx.JSON(http.StatusOK, service.Response{
		Code:    0,
		Message: "请求正常",
	})
}

func AddUser(ctx *gin.Context) {
	var req service.OperateUserReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "参数解析错误：" + err.Error()})
		return
	}
	payload, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	if payload.Role > 1 {
		ctx.JSON(http.StatusForbidden, service.Response{Code: -1, Message: "权限不足"})
		return
	}
	if req.Role < 0 || req.Role > 2 || req.Account == "" || req.Password == "" {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "参数不足"})
		return
	}
	var user tables.User
	user.Account = req.Account
	user.Role = req.Role
	user.Password = req.Password
	orm.Db.Save(user)
	ctx.JSON(http.StatusOK, service.Response{
		Code:    0,
		Message: "请求正常",
	})
}
