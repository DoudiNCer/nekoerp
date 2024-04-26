package service

import (
	"github.com/gin-gonic/gin"
	"nekoerp/pkg/service"
	"nekoerp/pkg/service/middleware"
	orm "nekoerp/pkg/storage/mysql"
	"nekoerp/pkg/storage/mysql/tables"
	"net/http"
)

func ListStorage(ctx *gin.Context) {
	req := service.ListSearchReq{
		Token:   ctx.Query("token"),
		Keyword: ctx.Query("keyword"),
	}
	_, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	var total int64
	orm.Db.Model(tables.Storage{}).Count(&total)
	var results []service.ListStorageResp
	page := orm.Db.Table("storage").Select("`storage`.id, `storage`.`count`, goods.`name`, goods.price").Joins("LEFT JOIN goods ON `storage`.type = goods.id")
	if req.Keyword != "" {
		page = page.Where("goods.`name` LIKE ?", "%"+req.Keyword+"%")
	}
	page.Scan(&results)
	ctx.JSON(http.StatusOK, service.ResponseList{Total: uint(total), List: results})
}

func ListDi(ctx *gin.Context) {
	req := service.ListSearchReq{
		Token:   ctx.Query("token"),
		Keyword: ctx.Query("keyword"),
	}
	_, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	var total int64
	orm.Db.Model(tables.Di{}).Count(&total)
	var results []service.ListDiTiaoResp
	page := orm.Db.Table("di").Select("di.id, di.`count`, di.created_at, goods.`name`, goods.price, `user`.account").Joins("LEFT JOIN goods ON di.type = goods.id").
		Joins("LEFT JOIN `user` ON di.operator = `user`.id")
	if req.Keyword != "" {
		page = page.Where("goods.`name` LIKE ? OR `user`.account LIKE ? ", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	page.Scan(&results)
	ctx.JSON(http.StatusOK, service.ResponseList{Total: uint(total), List: results})
}

func ListTiao(ctx *gin.Context) {
	req := service.ListSearchReq{
		Token:   ctx.Query("token"),
		Keyword: ctx.Query("keyword"),
	}
	_, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	var total int64
	orm.Db.Model(tables.Tiao{}).Count(&total)
	var results []service.ListDiTiaoResp
	page := orm.Db.Table("tiao").Select("tiao.id, tiao.`count`, tiao.created_at, goods.`name`, goods.price, `user`.account").Joins("LEFT JOIN goods ON tiao.type = goods.id").
		Joins("LEFT JOIN `user` ON tiao.operator = `user`.id")
	if req.Keyword != "" {
		page = page.Where("goods.`name` LIKE ? OR `user`.account LIKE ? ", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	page.Scan(&results)
	ctx.JSON(http.StatusOK, service.ResponseList{Total: uint(total), List: results})
}

func ListGoods(ctx *gin.Context) {
	req := service.ListSearchReq{
		Token:   ctx.Query("token"),
		Keyword: ctx.Query("keyword"),
	}
	_, err := middleware.ParseJwt(req.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, service.Response{Code: -1, Message: "Token解析错误：" + err.Error()})
		return
	}
	var total int64
	orm.Db.Model(tables.Goods{}).Count(&total)
	var results []service.ListGoodsResp
	page := orm.Db.Table("goods")
	if req.Keyword != "" {
		page = page.Where("goods.`name` LIKE", "%"+req.Keyword+"%")
	}
	page.Scan(&results)
	ctx.JSON(http.StatusOK, service.ResponseList{Total: uint(total), List: results})
}

func AddGoods(ctx *gin.Context) {
	var req service.AddGoodsReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Message: "参数解析错误：" + err.Error()})
		return
	}
	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Data: "参数校验错误"})
	}
	var goods tables.Goods
	goods.Name = req.Name
	goods.Price = req.Price
	orm.Db.Save(&goods)
	ctx.JSON(http.StatusOK, service.Response{
		Code:    0,
		Message: "请求正常",
	})
}

func Di(ctx *gin.Context) {
	var req service.DiTiaoReq
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
	if req.Type <= 0 || req.Count <= 0 {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Data: "参数校验错误"})
	}
	var goods tables.Goods
	orm.Db.First(&goods, req.Type)
	if goods.Name == "" {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Data: "参数校验错误"})
		return
	}
	var storage tables.Storage
	orm.Db.First(&storage, "type = ?", req.Type)
	if storage.Type == 0 {
		storage.Type = req.Type
	}
	storage.Count += req.Count
	orm.Db.Save(&storage)
	di := tables.Di{
		Type:     req.Type,
		Count:    req.Count,
		Operator: payload.Id,
	}
	orm.Db.Save(&di)
	ctx.JSON(http.StatusOK, service.Response{
		Code:    0,
		Message: "请求正常",
	})
}

func Tiao(ctx *gin.Context) {
	var req service.DiTiaoReq
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
	if req.Type <= 0 || req.Count <= 0 {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Data: "参数校验错误"})
	}
	var goods tables.Goods
	orm.Db.First(&goods, req.Type)
	if goods.Name == "" {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Data: "参数校验错误"})
		return
	}
	var storage tables.Storage
	orm.Db.First(&storage, "type = ?", req.Type)
	if storage.Type == 0 || storage.Count < req.Count {
		ctx.JSON(http.StatusBadRequest, service.Response{Code: -1, Data: "仓库数量不足，出库失败"})
		return
	}
	storage.Count -= req.Count
	orm.Db.Save(&storage)
	tiao := tables.Tiao{
		Type:     req.Type,
		Count:    req.Count,
		Operator: payload.Id,
	}
	orm.Db.Save(&tiao)
	ctx.JSON(http.StatusOK, service.Response{
		Code:    0,
		Message: "请求正常",
	})
}
