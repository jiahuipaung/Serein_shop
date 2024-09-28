package service

import (
	"context"
	"strconv"

	// "fmt"
	"mime/multipart"
	"sync"

	"serein/consts"
	"serein/types"

	conf "serein/config"

	// "serein/pkg/utils/ctl"

	"serein/pkg/utils/log"
	"serein/repository/db/dao"
	"serein/repository/db/model"

	util "serein/pkg/utils/upload"
)

var ProductIns *ProductSrv
var ProductOnce sync.Once

type ProductSrv struct {
}

func GetProductSrv() *ProductSrv {
	ProductOnce.Do(func() {
		ProductIns = &ProductSrv{}
	})
	return ProductIns
}

// ProductCreate 创建商品
func (s *ProductSrv) ProductCreate(ctx context.Context, files []*multipart.FileHeader, req *types.ProductCreateReq) (resp interface{}, err error) {
	// 以第一张图作为封面
	tmp, _ := files[0].Open()
	var path string
	if conf.Config.System.UploadModel == consts.UploadModelLocal {
		path, err = util.ProductUploadToLocalStatic(tmp, req.Name)
	}

	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	product := &model.Product{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Info:       req.Info,
		ImgPath:    path,
		Price:      req.Price,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		tmp, _ := file.Open()
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			path, err = util.ProductUploadToLocalStatic(tmp, req.Name+num)
		}
		if err != nil {
			log.LogrusObj.Error(err)
			return
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = dao.NewProductImgDaoByDB(productDao.DB).CreateProductImg(productImg)
		if err != nil {
			log.LogrusObj.Error(err)
			return
		}
		wg.Done()
	}

	wg.Wait()

	return
}

func (s *ProductSrv) ProductList(ctx context.Context, req *types.ProductListReq) (resp interface{}, err error) {
	var total int64
	condition := make(map[string]interface{})
	if req.CategoryID != 0 {
		condition["category_id"] = req.CategoryID
	}

	productDao := dao.NewProductDao(ctx)
	// 从数据库中获取所有符合条件的商品信息
	products, _ := productDao.ListProductByCondition(condition, req.BasePage)
	total, err = productDao.CountProductByCondition(condition)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	pRespList := make([]*types.ProductResp, 0)
	for _, p := range products {
		pResp := &types.ProductResp{
			ID:         p.ID,
			Name:       p.Name,
			CategoryID: p.CategoryID,
			Title:      p.Title,
			Info:       p.Info,
			ImgPath:    p.ImgPath,
			Price:      p.Price,
		}
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
		}
		pRespList = append(pRespList, pResp)
	}

	resp = &types.DataListResp{
		Item:  pRespList,
		Total: total,
	}
	return
}

func (s *ProductSrv) ProductShow(ctx context.Context, req *types.ProductShowReq) (resp interface{}, err error) {
	p, err := dao.NewProductDao(ctx).ShowProductById(req.ID)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	pResp := &types.ProductResp{
		ID:         req.ID,
		Name:       p.Name,
		CategoryID: p.CategoryID,
		Title:      p.Title,
		Info:       p.Info,
		ImgPath:    p.ImgPath,
		Price:      p.Price,
	}

	if conf.Config.System.UploadModel == consts.UploadModelLocal {
		pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
	}

	resp = pResp
	return
}

func (s *ProductSrv) ProductSearch(ctx context.Context, req *types.ProductSearchReq) (resp interface{}, err error) {
	products, count, err := dao.NewProductDao(ctx).SearchProduct(req.Info, req.BasePage)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	pRespList := make([]*types.ProductResp, 0)
	for _, p := range products {
		pResp := &types.ProductResp{
			ID:         p.ID,
			Name:       p.Name,
			CategoryID: p.CategoryID,
			Title:      p.Title,
			Info:       p.Info,
			ImgPath:    p.ImgPath,
			Price:      p.Price,
		}
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
		}
		pRespList = append(pRespList, pResp)
	}

	resp = &types.DataListResp{
		Item:  pRespList,
		Total: count,
	}

	return
}

func (s *ProductSrv) ProductDelete(ctx context.Context, req *types.ProductDeleteReq) (resp interface{}, err error) {
	err = dao.NewProductDao(ctx).DeleteProduct(req.ID)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	return
}
