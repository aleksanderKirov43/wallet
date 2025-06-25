package service

import (
	"context"
	"errors"
	"fmt"
	"wallet/internal/repositiory"
)

type WalletService interface {
	CreateWallet(ctx context.Context, req *Wallet) (*Wallet, error)
	PostBalance(ctx context.Context, req *Wallet) (*Wallet, error)
	GetBalance(ctx context.Context, walletID int) (*CheckBalance, error)
	UpdateBalance(ctx context.Context, req *Wallet) (*Wallet, error)
	DeleteWallet(ctx context.Context, walletID int) error
}

type CheckBalance struct {
	Amount int64 `json:"amount"`
}

type Wallet struct {
	WalletID      int    `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int64  `json:"amount,omitempty"`
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

	if req.OperationType != "DEPOSIT" {
		return nil, errors.New("Сначала внесите депозит")
	}

	err := s.repo.CreateWallet(ctx, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *WalletServiceImpl) PostBalance(ctx context.Context, req *Wallet) (*Wallet, error) {

	exist, err := s.repo.GetBalance(ctx, req.WalletID)
	if err != nil {
		return nil, fmt.Errorf("Кошелёк не найден: %w", err)
	}

	switch req.OperationType {
	case "DEPOSIT":
		exist.Amount += req.Amount
	case "WITHDRAW":
		return nil, errors.New("Метод не предназначен для вывода средств, выбирите : UpdateBalance")
	default:
		return nil, errors.New("Неизвестный тип операции")
	}

	err = s.repo.UpdateBalance(ctx, exist)
	if err != nil {
		return nil, fmt.Errorf("Ошибка пополнения депозита: %w", err)
	}

	return exist, nil
}

func (s *WalletServiceImpl) GetBalance(ctx context.Context, walletID int) (*CheckBalance, error) {

	exist, err := s.repo.GetBalance(ctx, walletID)
	if err != nil {
		return nil, fmt.Errorf("Кошелёк не найден: %w", err)
	}

	return &CheckBalance{Amount: exist.Amount}, nil
}

func (s *WalletServiceImpl) UpdateBalance(ctx context.Context, req *Wallet) (*Wallet, error) {

	exist, err := s.repo.GetBalance(ctx, req.WalletID)
	if err != nil {
		return nil, fmt.Errorf("Кошелёк не найден: %w", err)
	}

	switch req.OperationType {
	case "DEPOSIT":
		exist.Amount += req.Amount
	case "WITHDRAW":
		if exist.Amount < req.Amount {
			return nil, errors.New("Недостаточно средств для вывода")
		}
		exist.Amount -= req.Amount
	default:
		return nil, errors.New("Неизвестный тип операции")
	}

	err = s.repo.UpdateBalance(ctx, exist)
	if err != nil {
		return nil, fmt.Errorf("Ошибка обновления данных: %w", err)
	}

	return exist, nil

}

func (s *WalletServiceImpl) DeleteWallet(ctx context.Context, walletID int) error {
	err := s.repo.DeleteWallet(ctx, walletID)
	if err != nil {
		return fmt.Errorf("ошибка удаления кошелька: %w", err)
	}
	return nil
}
