package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          int64     `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Picture     string    `json:"picture"`
	Price       float32   `json:"price"`
	Stock       int64     `json:"stock"`
	Sale        int64     `json:"sale"`
	PublicState int64     `json:"public_state"`
	LockStock   int64     `json:"lock_stock"`
	CategoryId  int64     `json:"category_id"`
	BrandId     int64     `json:"brand_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	RealStock   int64     `gorm:"-" json:"quantity"`
	Category    Category  `gorm:"foreignKey:category_id" json:"category"`
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

func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	p.RealStock = p.Stock - p.LockStock
	return nil
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
	result := db.WithContext(ctx).Delete(&Product{ID: id})
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
