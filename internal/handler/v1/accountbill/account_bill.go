// Created on 2022/5/27.
// @author tony
// email xmgtony@gmail.com
// description 账目清单handler

package accountbill

import (
	"apiserver-gin/internal/middleware"
	"apiserver-gin/internal/model"
	"apiserver-gin/internal/service"
	"apiserver-gin/pkg/response"
	"apiserver-gin/pkg/xerrors"
	"apiserver-gin/pkg/xerrors/ecode"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// Handler 账目清单handler，从分层来讲，这里已经是最外层，只要提供实例在router中使用
// 所以这里没有定义接口，而是直接使用struct来组织多个handler func
type Handler struct {
	accountBillServ service.AccountBillService
}

func NewAccountBillHandler(_accountBillServ service.AccountBillService) *Handler {
	return &Handler{
		accountBillServ: _accountBillServ,
	}
}

func (abh *Handler) AddAccountBill() gin.HandlerFunc {
	return func(c *gin.Context) {
		addAccountBillReq := model.AddAccountBillReq{}
		if err := c.ShouldBindJSON(&addAccountBillReq); err != nil {
			response.JSON(c, xerrors.WithCode(ecode.ValidateErr, err.Error()), nil)
			return
		}
		uid := middleware.GetUserId(c)
		amd, err := decimal.NewFromString(addAccountBillReq.Amount)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.ValidateErr, "金额填写错误"), nil)
			return
		}
		if amd.IsNegative() {
			response.JSON(c, xerrors.WithCode(ecode.ValidateErr, "金额须填写大于0的数字"), nil)
			return
		}
		// 数据库存储的单位为分，所以要*100
		amount := amd.Mul(decimal.NewFromInt32(100)).IntPart()
		// 组织model信息
		accountBill := addAccountBillReq.ToAccountBill(uint64(uid), uint(amount))
		err = abh.accountBillServ.Save(c, &accountBill)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.RecordCreateErr, "保存账目清单信息失败"), nil)
			return
		}
		response.JSON(c, nil, nil)
	}
}

func (abh *Handler) GetAccountBillList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用auth中间件的接口，一定能解析出来，否则中间件会响应错误
		uid := middleware.GetUserId(c)
		bills, err := abh.accountBillServ.SelectListByUserId(c, uid)
		if err != nil {
			response.JSON(c, xerrors.Wrap(err, ecode.NotFoundErr, "查询错误，未找到记录"), nil)
			return
		}
		respBills := make([]model.AccountBillResp, 0)
		for _, bill := range bills {
			respBills = append(respBills, bill.ToAccountBillResp())
		}
		response.JSON(c, nil, respBills)
		return
	}
}
