package services

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/tigerbig/go-loan-car/generated/grpc-brand-models/proto"
	"github.com/tigerbig/go-loan-car/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ModelService struct {
	pb.UnimplementedModelServiceServer
	DB *gorm.DB
}

func NewModelBrandService(db *gorm.DB) *ModelService {
	return &ModelService{DB: db}
}

func (s *ModelService) CreateBrandModel(ctx context.Context, req *pb.RequestBrandModel) (*pb.ResponseBrandModel, error) {
	brandModels := models.ModelsCar{
		BrandID:         uint(req.BrandId),
		ModelCode:       req.ModelCode,
		ModelName:       req.ModelName,
		YearStart:       req.YearStart,
		YearEnd:         req.YearEnd,
		BodyType:        req.BodyType,
		EngineType:      req.EngineType,
		EngineSize:      req.EngineSize,
		FuelConsumption: req.FuelConsumption,
		Transmission:    req.Transmission,
		Generation:      req.Generation,
	}

	// validate data is match require
	if len(req.ModelName) <= 1 {
		return nil, status.Errorf(codes.InvalidArgument, "model name must be at least 2 characters")
	}

	if err := s.DB.Create(&brandModels).Error; err != nil {
		return nil, err
	}

	return &pb.ResponseBrandModel{
		Id:              int32(brandModels.ID),
		BrandId:         int32(brandModels.BrandID),
		ModelCode:       brandModels.ModelCode,
		ModelName:       brandModels.ModelName,
		YearStart:       brandModels.YearStart,
		YearEnd:         brandModels.YearEnd,
		BodyType:        brandModels.BodyType,
		EngineType:      brandModels.EngineType,
		EngineSize:      brandModels.EngineSize,
		FuelConsumption: brandModels.FuelConsumption,
		Transmission:    brandModels.Transmission,
		Generation:      brandModels.Generation,
		Message:         "Create model success",
	}, nil
}

func (s *ModelService) GetModelsBrand(ctx context.Context, req *pb.RequestBrandModelSingle) (*pb.ResponseBrandModelSingle, error) {
	var modelsBrand models.ModelsCar

	if err := s.DB.First(&modelsBrand, req.Id).Error; err != nil {
		return nil, err
	}

	return &pb.ResponseBrandModelSingle{
		Id:              int32(modelsBrand.ID),
		BrandId:         int32(modelsBrand.BrandID),
		ModelName:       modelsBrand.ModelName,
		BodyType:        modelsBrand.BodyType,
		EngineType:      modelsBrand.EngineType,
		EngineSize:      modelsBrand.EngineSize,
		FuelConsumption: modelsBrand.FuelConsumption,
		Transmission:    modelsBrand.Transmission,
		Generation:      modelsBrand.Generation,
		Message:         "Model brand found",
	}, nil
}

func (s *ModelService) GetModelsWithBrand(ctx context.Context, req *pb.RequestModelWithBrand) (*pb.ResponseModelWithBrand, error) {

	var brand models.Brand
	findBrandExist := s.DB.First(&brand, req.BrandId)
	// if findBrandExist.Error != nil {
	// 	fmt.Printf("Error: Brand has not fond: %v\n", findBrandExist.Error)
	// 	return nil, status.Errorf(codes.Internal, "Brand is not exist: %v", findBrandExist.Error)
	// }
	if findBrandExist.Error != nil {
		if errors.Is(findBrandExist.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(
				codes.NotFound,
				"brand with id %d not found",
				req.BrandId,
			)
		}
	}

	var modelBrand []models.ModelsCar
	result := s.DB.Where("brand_id = ?", req.BrandId).Find(&modelBrand)
	if result.Error != nil {
		fmt.Printf("❌ ERROR: Database query failed: %v\n", result.Error)
		return nil, status.Errorf(codes.Internal, "database query failed: %v", result.Error)
	}
	fmt.Printf("✅ Query OK: found %d records\n", len(modelBrand))

	protoModels := make([]*pb.CarModels, 0, len(modelBrand))
	for i, listModel := range modelBrand {
		fmt.Printf("  Converting model %d: Code=%s, Name=%s\n", i+1, listModel.ModelCode, listModel.ModelName)
		protoModels = append(protoModels, &pb.CarModels{
			ModelCode: listModel.ModelCode,
			ModelName: listModel.ModelName,
		})
	}
	fmt.Printf("%+v\n", protoModels)
	fmt.Println("brand_id:", req.BrandId)

	return &pb.ResponseModelWithBrand{
		Model: protoModels,
	}, nil
}
