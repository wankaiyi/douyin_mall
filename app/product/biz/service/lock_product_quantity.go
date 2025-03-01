package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type LockProductQuantityService struct {
	ctx context.Context
} // NewLockProductQuantityService new LockProductQuantityService
func NewLockProductQuantityService(ctx context.Context) *LockProductQuantityService {
	return &LockProductQuantityService{ctx: ctx}
}

// Run create note info
func (s *LockProductQuantityService) Run(req *product.ProductLockQuantityRequest) (resp *product.ProductLockQuantityResponse, err error) {
	originProducts := req.Products
	var ids = make([]int64, 0)
	var productQuantityMap = make(map[int64]int64)
	for _, pro := range originProducts {
		ids = append(ids, pro.Id)
		productQuantityMap[pro.Id] = pro.Quantity
	}
	productList, err := model.SelectProductList(mysql.DB, context.Background(), ids)
	//确定当前库存是否足够
	canLock := true
	var lowStockProductId int64

	for _, pro := range productList {
		//如果真实库存小于下单的数量，则库存锁定失败
		if pro.RealStock < productQuantityMap[pro.ProductId] {
			canLock = false
			lowStockProductId = pro.ProductId
			break
		}
	}
	//如果库存锁定失败，则返回失败信息
	if !canLock {
		klog.CtxInfof(s.ctx, "商品库存不足，无法锁定库存，productId：%v, quantity：%v", lowStockProductId, productQuantityMap[lowStockProductId])
		return &product.ProductLockQuantityResponse{
			StatusCode: 6014,
			StatusMsg:  constant.GetMsg(6014),
		}, nil
	}
	//如果库存锁定成功，则更新库存信息
	err = model.UpdateLockStock(mysql.DB, context.Background(), productQuantityMap)
	if err != nil {
		klog.CtxErrorf(s.ctx, "更新库存失败，原因：%v", err)
		return nil, err
	}
	return &product.ProductLockQuantityResponse{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}, nil
}
