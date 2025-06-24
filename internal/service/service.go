package service

import (
	"context"
	"wallet/internal/repositiory"
)

type WalletService interface {
	CreateWallet(ctx context.Context, req *Wallet) (*Wallet, error)
	PostBalance(ctx context.Context, walletID int) (*Wallet, error)
	GetBalance(ctx context.Context, walletID int) (*CheckBalance, error)
	UpdateBalance(ctx context.Context, walletID int) (*Wallet, error)
	DeleteWallet(ctx context.Context, walletID int) error
}

type CheckBalance struct {
	Amount int64 `json:"amount"`
}

type Wallet struct {
	WalletID      int    `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int64  `json:"amount"`
}

type WalletServiceImpl struct {
	repo repositiory.Repository
}

func NewWalletService(repo repositiory.Repository) *WalletServiceImpl {
	return &WalletServiceImpl{
		repo: repo,
	}
}

func (s *WalletServiceImpl) CreateWallet(ctx context.Context, req *Wallet) (*Wallet, error) {

}

func (s *WalletServiceImpl) PostBalance(ctx context.Context, walletID int) (*Wallet, error) {

}

func (s *WalletServiceImpl) GetBalance(ctx context.Context, walletID int) (*CheckBalance, error) {

}

func (s *WalletServiceImpl) UpdateBalance(ctx context.Context, walletID int) (*Wallet, error) {

}

func (s *WalletServiceImpl) DeleteWallet(ctx context.Context, walletID int) error {

}
