package service

import (
	"context"
	"errors"
	"serein/pkg/e"
	"serein/pkg/utils/ctl"
	"serein/repository/db/dao"
	"serein/types"
	"sync"

	util "serein/pkg/utils/log"
)

var CartSrvIns *CartSrv
var CartSrvOnce sync.Once

type CartSrv struct {
}

func GetCartSrv() *CartSrv {
	CartSrvOnce.Do(func() {
		CartSrvIns = &CartSrv{}
	})
	return CartSrvIns
}

// 创建购物车
func (s *CartSrv) CartCreate(ctx context.Context, req *types.CartCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	// 判断有无该产品
	_, err = dao.NewProductDao(ctx).GetProductById(req.ProductID)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	// 创建购物车
	cartDao := dao.NewCartDao(ctx)
	_, status, _ := cartDao.CreateCart(req.ProductID, u.Id)
	if status == e.ErrorProductMoreCart {
		err = errors.New(e.GetMsg(status))
		return
	}

	return
}

func (s *CartSrv) CartList(ctx context.Context, req *types.CartListReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	carts, err := dao.NewCartDao(ctx).ListCartByUserId(u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	resp = &types.DataListResp{
		Item:  carts,		// TODO 无分页，之后考虑要不要加
		Total: int64(len(carts)),
	}
	return
}

func (s *CartSrv) CartUpdate(ctx context.Context, req *types.UpdateCartServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	err = dao.NewCartDao(ctx).UpdateCartNumById(req.ID, u.Id, req.Num)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return
}

func (s *CartSrv) CartDelete(ctx context.Context, req *types.CartDeleteReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	err = dao.NewCartDao(ctx).DeleteCartById(req.Id, u.Id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	
	return
}
