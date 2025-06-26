package repositiory

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"wallet/internal/model"
)

type Repository interface {
	CreateWallet(ctx context.Context, wallet *model.Wallet) error
	GetBalance(ctx context.Context, walletID int) (*model.Wallet, error)
	UpdateBalance(ctx context.Context, wallet *model.Wallet) error
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

func (r *RepositoryDB) CreateWallet(ctx context.Context, wallet *model.Wallet) error {

	query := `INSERT INTO wallet (walletId, operationType, amount) VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, wallet.WalletID, wallet.OperationType, wallet.Amount)
	if err != nil {
		log.Println("Ошибка создания кошелька: %v", err)
		return err
	}
	return nil
}

func (r *RepositoryDB) GetBalance(ctx context.Context, walletID int) (*model.Wallet, error) {

	query := `SELECT operationType, amount FROM wallet WHERE walletId = $1`

	row := r.db.QueryRowContext(ctx, query, walletID)

	var operationType string
	var amount int64

	err := row.Scan(&operationType, &amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Кошелек с ID %d не найден", walletID)
		}
		return nil, err
	}

	return &model.Wallet{
		WalletID:      walletID,
		OperationType: operationType,
		Amount:        amount,
	}, nil
}

func (r *RepositoryDB) UpdateBalance(ctx context.Context, wallet *model.Wallet) error {

	query := "UPDATE wallet SET operationType = $1, amount = $2 WHERE walletId = $3"

	_, err := r.db.ExecContext(ctx, query, wallet.OperationType, wallet.Amount, wallet.WalletID)
	if err != nil {
		return fmt.Errorf("Не удалось обновить баланс: %w", err)
	}

	return nil

}

func (r *RepositoryDB) DeleteWallet(ctx context.Context, walletID int) error {
	query := "DELETE FROM wallet WHERE walletId = $1"

	_, err := r.db.ExecContext(ctx, query, walletID)
	if err != nil {
		return fmt.Errorf("не удалось удалить кошелёк: %w", err)
	}

	return nil
}
