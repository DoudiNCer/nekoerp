package service

import "time"

type ResponseList struct {
	Total uint `json:"total"`
	List  any  `json:"list"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ListSearchReq struct {
	Token    string `json:"token"`
	Keyword  string `json:"keyword"`
	PageNum  uint   `json:"pageNum"`
	PageSize uint   `json:"pageSize"`
}

type ListStorageResp struct {
	Id    uint    `json:"id"`
	Count uint    `json:"count"`
	Name  string  `json:"name"`
	Price float64 `gorm:"type:numeric(10,2);column:price" json:"price"`
}

type ListDiTiaoResp struct {
	Id        uint      `json:"id"`
	Count     uint      `json:"count"`
	Name      string    `json:"name"`
	Price     float64   `gorm:"type:numeric(10,2);column:price" json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	Account   string    `json:"account"`
}

type ListUserResp struct {
	Id        uint      `json:"id"`
	Account   string    `json:"account"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Blocked   bool      `json:"blocked"`
}

type OperateUserReq struct {
	Token    string `json:"token"`
	Id       uint   `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type ListGoodsResp struct {
	Id    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type AddGoodsReq struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type DiTiaoReq struct {
	Token string `json:"token"`
	Type  uint   `json:"type"`
	Count uint   `json:"count"`
}
