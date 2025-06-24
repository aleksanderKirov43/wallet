package repositiory

import (
	"context"
	"database/sql"
	"wallet/internal/service"
)

type Repository interface {
	CreateWallet(ctx context.Context, wallet *service.Wallet) error
	GetBalance(ctx context.Context, walletID int) (*service.Wallet, error)
	UpdateBalance(ctx context.Context, wallet *service.Wallet) error
	DeleteWallet(ctx context.Context, walletID int) error
}

type RepositoryDB struct {
	db *sql.DB
}

func NewRepositoryDB(db *sql.DB) *RepositoryDB {
	return &RepositoryDB{
		db: db,
	}
}

func (r *RepositoryDB) CreateWallet(ctx context.Context, wallet *service.Wallet) error {

}

func (r *RepositoryDB) GetBalance(ctx context.Context, walletID int) (*service.Wallet, error) {

}

func (r *RepositoryDB) UpdateBalance(ctx context.Context, wallet *service.Wallet) error {

}

func (r *RepositoryDB) DeleteWallet(ctx context.Context, walletID int) error {

}
