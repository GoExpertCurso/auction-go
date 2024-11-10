package bid_entity

import (
	"context"
	"time"

	"github.com/GoExpertCurso/auction-go/internal/internal_error"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

type BidRepositoryInterface interface {
	CreateBid(
		ctx context.Context,
		bidEntities []Bid) *internal_error.InternalError

	FindBidByAuctionId(
		ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}
