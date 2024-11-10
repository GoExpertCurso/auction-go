package auction_usercase

import (
	"context"
	"time"

	"github.com/GoExpertCurso/auction-go/internal/entity/auction_entity"
	"github.com/GoExpertCurso/auction-go/internal/entity/bid_entity"
	"github.com/GoExpertCurso/auction-go/internal/internal_error"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition" binding:"oneof=0 1 2"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductCondition int
type AuctionStatus int

func NewAuctionUseCase(
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface,
	bidRepositoryInterface bid_entity.BidRepositoryInterface) AuctionUsecaseInterface {
	return &AuctionUsecase{
		auctionRepositoryInterface: auctionRepositoryInterface,
		bidRepositoryInterface:     bidRepositoryInterface,
	}
}

type AuctionUsecaseInterface interface {
	CreateAuction(ctx context.Context, auctionInputDTO AuctionInputDTO) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)
}

type AuctionUsecase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidRepositoryInterface
}

// CreateAuction implements AuctionUsecaseInterface.
func (au *AuctionUsecase) CreateAuction(ctx context.Context, auctionInputDTO AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInputDTO.ProductName,
		auctionInputDTO.Category,
		auctionInputDTO.Description,
		auction_entity.ProductCondition(auctionInputDTO.Condition)
	)
	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}
	return nil
}