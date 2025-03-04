package service

import (
	"context"
	"crypto/md5"
	"douyin_mall/common/constant"
	product "douyin_mall/product/kitex_gen/product"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SelectProductService struct {
	ctx context.Context
} // NewSelectProductService new SelectProductService
func NewSelectProductService(ctx context.Context) *SelectProductService {
	return &SelectProductService{ctx: ctx}
}

// Run create note info
func (s *SelectProductService) Run(req *product.SelectProductReq) (resp *product.SelectProductResp, err error) {
	var searchIds []int64 = []int64{req.Id}
	var products = make([]*product.Product, 0)
	//根据id从缓存或者数据库钟获取数据
	//根据返回的数据确认是否有缺失数据，有的话把当前的id存进去
	var missingIds []int64
	//先判断redis是否存在数据，如果存在，则直接返回数据
	searchIdsBytes, err := sonic.Marshal(searchIds)
	if err != nil {
		return nil, err
	}
	harsher := md5.New()
	harsher.Write(searchIdsBytes)
	md5Bytes := harsher.Sum(nil)
	products, missingIds, err = GetCache(s.ctx, searchIds, md5Bytes)
	klog.CtxInfof(s.ctx, "products: %v,missingsIds:%v", products, missingIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "GetCache: missingsIds:%v,err:%v", missingIds, err)
		return nil, err
	}

	//如果不存在，则从数据库中获取数据，并存入redis
	missingProduct, err := GetMissingProduct(s.ctx, missingIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "GetMissingProduct: err:%v", err)
		return nil, err
	}
	products = append(products, missingProduct...)
	klog.CtxInfof(s.ctx, "搜索的products: %v", products)

	//根据商品id查询库存信息
	productStock, err := GetStock(s.ctx, searchIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取库存时, err: %v", err)
		return nil, err
	}
	for _, pro := range products {
		pro.Stock = productStock[pro.Id]
	}
	//products, err = model.SelectProductList(mysql.DB, s.ctx, req.Ids)
	//if err != nil {
	//	klog.CtxErrorf(s.ctx, "查询商品列表失败, error:%v", err)
	//	return &product.SelectProductListResp{
	//		StatusCode: 6003,
	//		StatusMsg:  constant.GetMsg(6003),
	//	}, nil
	//}
	var productList []*product.Product
	for i := range products {
		productList = append(productList, &product.Product{
			Id:            products[i].Id,
			Name:          products[i].Name,
			Description:   products[i].Description,
			Picture:       products[i].Picture,
			Price:         products[i].Price,
			CategoryName:  products[i].CategoryName,
			CategoryId:    products[i].CategoryId,
			Stock:         products[i].Stock,
			Sale:          products[i].Sale,
			PublishStatus: products[i].PublishStatus,
		})
	}
	if len(productList) == 0 {
		return &product.SelectProductResp{
			StatusCode: 6004,
			StatusMsg:  constant.GetMsg(6004),
		}, nil
	}
	pro := productList[0]
	return &product.SelectProductResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Product: &product.Product{
			Id:            pro.Id,
			Name:          pro.Name,
			Description:   pro.Description,
			Picture:       pro.Picture,
			Price:         pro.Price,
			Stock:         pro.Stock,
			Sale:          pro.Sale,
			PublishStatus: pro.PublishStatus,
			CategoryId:    pro.CategoryId,
		},
	}, nil
}
