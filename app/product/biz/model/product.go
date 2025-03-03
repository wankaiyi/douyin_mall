package model

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Product struct {
	Base
	Name        string  `gorm:"not null;type:varchar(100)" json:"name"`
	Description string  `gorm:"not null;type:text" json:"description"`
	Picture     string  `gorm:"not null;type:varchar(255)" json:"picture"`
	Price       float32 `gorm:"not null;type:decimal(10,2)" json:"price"`
	Stock       int64   `gorm:"not null;type:int(32)" json:"stock"`
	Sale        int64   `gorm:"not null;type:int(32)" json:"sale"`
	PublicState int64   `gorm:"not null;type:int(32)" json:"public_state"`
	LockStock   int64   `gorm:"not null;type:int(32)" json:"lock_stock"`
	CategoryId  int64   `gorm:"not null;index:idx_category_id;not null;type:int(32)" json:"category_id"`
}
type ProductWithCategory struct {
	ProductId          int64   `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductPrice       float32 `json:"product_price"`
	ProductDescription string  `json:"product_description"`
	ProductStock       int64   `json:"product_stock"`
	ProductLockStock   int64   `json:"product_lock_stock"`
	ProductSale        int64   `json:"product_sale"`
	ProductPublicState int64   `json:"product_public_state"`
	ProductPicture     string  `json:"product_picture"`
	CategoryID         int64   `json:"category_id"`
	CategoryName       string  `json:"category_name"`
	RealStock          int64   `gorm:"-" json:"quantity"`
}

func (p *Product) TableName() string {
	return "tb_product"
}

func (p *ProductWithCategory) AfterFind(tx *gorm.DB) (err error) {
	p.RealStock = p.ProductStock - p.ProductLockStock
	return nil
}

func SelectProduct(db *gorm.DB, ctx context.Context, id int64) (product Product, err error) {
	product = Product{}
	result := db.WithContext(ctx).Where("id=?", id).First(&product)
	err = result.Error
	return product, err
}
func SelectProductAll(db *gorm.DB, ctx context.Context, index int64, total int64) (product []Product, err error) {
	product = []Product{}
	result := db.WithContext(ctx).Model(&Product{}).Where("id%? = ?", total, index).Find(&product)
	err = result.Error
	return product, err
}

func SelectProductAllWithoutCondition(db *gorm.DB, ctx context.Context) (products []ProductWithCategory, err error) {
	result := db.WithContext(ctx).Model(&Product{}).
		Select("tb_product.id as product_id,tb_product.name as product_name,tb_product.price as product_price,tb_product.description as product_description,tb_product.stock as product_stock,tb_product.lock_stock as product_lock_stock,tb_product.sale as product_sale,tb_product.public_state as product_public_state,tb_product.picture as product_picture,tb_category.id as category_id,tb_category.name as category_name").
		Joins("left join tb_category on tb_product.category_id=tb_category.id").
		Scan(&products)
	err = result.Error
	return products, err
}

func SelectProductList(db *gorm.DB, ctx context.Context, ids []int64) (products []ProductWithCategory, err error) {
	result := db.WithContext(ctx).Model(&Product{}).
		Where("tb_product.id IN ?", ids).
		Select("tb_product.id as product_id,tb_product.name as product_name,tb_product.price as product_price,tb_product.description as product_description,tb_product.stock as product_stock,tb_product.lock_stock as product_lock_stock,tb_product.sale as product_sale,tb_product.public_state as product_public_state,tb_product.picture as product_picture,tb_category.id as category_id,tb_category.name as category_name").
		Joins("left join tb_category on tb_product.category_id=tb_category.id").
		Scan(&products)
	err = result.Error
	return products, err
}

func UpdateProduct(db *gorm.DB, ctx context.Context, product *Product) (err error) {
	result := db.WithContext(ctx).Updates(&product)
	err = result.Error
	return err
}
func DeleteProduct(db *gorm.DB, ctx context.Context, id int64) (err error) {
	result := db.WithContext(ctx).Delete(&Product{Base: Base{ID: id}})
	err = result.Error
	return
}
func CreateProduct(db *gorm.DB, ctx context.Context, product *Product) (err error) {
	result := db.WithContext(ctx).Create(product)
	err = result.Error
	return err
}

func UpdateLockStock(db *gorm.DB, ctx context.Context, productQuantityMap map[int64]int64) (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		for id, quantity := range productQuantityMap {
			result := db.WithContext(ctx).
				Model(&Product{}).
				Where("id=?", id).
				Where("stock >= lock_stock+ ?", quantity).
				Update("lock_stock", gorm.Expr("lock_stock + ?", quantity))
			err = result.Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func PushToRedisBaseInfo(ctx context.Context, product Product, client *redis.Client, key string) (err error) {
	//改用lua脚本实现
	luaScript := `
		local key = KEYS[1]
		local id = ARGV[1]
		local name = ARGV[2]
		local description = ARGV[3]
		local picture = ARGV[4]
		local price = ARGV[5]
		local sale = ARGV[6]
		local publish_status = ARGV[7]
		
		if redis.call("EXISTS", key) == 0 then
			redis.call("HSET", key, "id", id, "name", name, "description", description, "picture", picture, "price", price ,'sale',sale,'publish_status',publish_status)
			return 1
		else
			return 0
		end
`
	keys := []string{key}
	args := []interface{}{product.ID, product.Name, product.Description, product.Picture, product.Price, product.Sale, product.PublicState}
	result, err := client.Eval(ctx, luaScript, keys, args).Result()
	if err != nil {
		return err
	}
	if result.(int64) == 1 {
		return nil
	} else {
		return errors.New("product already exists")
	}
}
func PushToRedisStock(ctx context.Context, product Product, client *redis.Client, key string) (err error) {
	//改用lua脚本实现
	luaScript := `
		local key = KEYS[1]
		local stock = ARGV[1]
		local lock_stock = ARGV[2]
		
		if redis.call("EXISTS", key) == 0 then
			redis.call("HSET", key, "stock", stock, "lock_stock", lock_stock)
			return 1
		else
			return 0
		end
`
	keys := []string{key}
	args := []interface{}{product.Stock, product.LockStock}
	result, err := client.Eval(ctx, luaScript, keys, args).Result()
	if err != nil {
		return err
	}
	if result.(int64) == 1 {
		return nil
	} else {
		return errors.New("product stock already exists")
	}
}

var (
	BaseInfoPattern     = "product:base:"
	BaseInfoLockPattern = "product:base:lock"
	StockPattern        = "product:stock:"
	StockLockPattern    = "product:stock:lock:"
)

func BaseInfoKey(ctx context.Context, id int64) string {
	return BaseInfoPattern + strconv.FormatInt(int64(id), 10)
}

func BaseInfoLockKey(ctx context.Context, id int64) string {
	return BaseInfoLockPattern + strconv.FormatInt(int64(id), 10)
}

func StockKey(ctx context.Context, id int64) string {
	return StockPattern + strconv.FormatInt(int64(id), 10)
}

func StockLockKey(ctx context.Context, id int64) string {
	return StockLockPattern + strconv.FormatInt(int64(id), 10)
}

// SafeDeleteLock 安全地删除锁
func SafeDeleteLock(ctx context.Context, client *redis.Client, key string, value string) (err error) {
	luaScript := `
		local key = KEYS[1]
		local value = ARGV[1]
		if redis.call("GET", key) == value then
			return redis.call("DEL", key)
		else
			return 0
		end
`
	keys := []string{key}
	args := []interface{}{value}
	result, err := client.Eval(ctx, luaScript, keys, args).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	if result.(int64) == 1 {
		return nil
	} else {
		return errors.New("锁不存在或过期")
	}
}
func SetLock(ctx context.Context, client *redis.Client, key string, value string) (lock bool, err error) {
	nx := client.SetNX(ctx, key, value, 2*time.Second)
	return nx.Val(), nx.Err()
}
