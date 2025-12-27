package services

import (
	"context"

	pb "github.com/tigerbig/go-loan-car/generated/grpc-add-brand/proto"
	"github.com/tigerbig/go-loan-car/internal/models"
	"gorm.io/gorm"
)

type BrandService struct {
	pb.UnimplementedBrandServiceServer
	DB *gorm.DB
}

func NewBrandServer(db *gorm.DB) *BrandService {
	return &BrandService{DB: db}
}

func (s *BrandService) CreateBrand(ctx context.Context, req *pb.RequestBrandData) (*pb.ResponseBrandData, error) {
	brand := models.Brand{
		BrandName:   req.BrandName,
		Country:     req.Country,
		Description: req.Description,
	}

	if err := s.DB.Create(&brand).Error; err != nil {
		return nil, err
	}

	return &pb.ResponseBrandData{
		Id:          int32(brand.ID),
		BrandName:   brand.BrandName,
		Country:     brand.Country,
		Description: brand.Description,
		Message:     "Created Brand Success",
	}, nil
}

func (s *BrandService) GetBrand(ctx context.Context, req *pb.RequestGetBrand) (*pb.ResponseBrand, error) {
	var brand models.Brand

	if err := s.DB.First(&brand, req.Id).Error; err != nil {
		return nil, err
	}

	return &pb.ResponseBrand{
		Id:          int32(brand.ID),
		BrandName:   brand.BrandName,
		Country:     brand.Country,
		Description: brand.Description,
		Message:     "Found brand!!",
	}, nil
}

func (s *BrandService) GetBrands(ctx context.Context, req *pb.RequestGetBrans) (*pb.ResponseBrandList, error) {
	var brands []models.Brand

	if err := s.DB.Find(&brands).Error; err != nil {
		return nil, err
	}

	respItems := make([]*pb.ResponseBrand, 0, len(brands))
	for _, b := range brands {
		respItems = append(respItems, &pb.ResponseBrand{
			Id:          int32(b.ID),
			BrandName:   b.BrandName,
			Country:     b.Country,
			Description: b.Description,
			Message:     "Found Brands",
		})
	}

	return &pb.ResponseBrandList{
		Brands:  respItems,
		Message: "Get found Brands completed",
	}, nil
}
