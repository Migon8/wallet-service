package service

import (
	"errors"
	"wallet-service/internal/model"
)

type WalletService interface {
	CreateWallet(ownerID int64) *model.Wallet
	GetWallet(id int64) (*model.Wallet, error)
	Deposit(id int64, amount int64) (*model.Wallet, error)
	Withdraw(id int64, amount int64) (*model.Wallet, error)
}

type walletService struct {
	storage map[int64]*model.Wallet
	nextID  int64
}

func NewWalletService() WalletService {
	return &walletService{
		storage: make(map[int64]*model.Wallet),
		nextID:  1,
	}
}

func (s *walletService) CreateWallet(ownerID int64) *model.Wallet {
	wallet := &model.Wallet{
		ID:      s.nextID,
		OwnerID: ownerID,
		Balance: 0,
	}
	s.storage[wallet.ID] = wallet
	s.nextID++
	return wallet
}

func (s *walletService) GetWallet(id int64) (*model.Wallet, error) {
	wallet, ok := s.storage[id]
	if !ok {
		return nil, errors.New("wallet not found")
	}
	return wallet, nil
}

func (s *walletService) Deposit(id int64, amount int64) (*model.Wallet, error) {
	wallet, err := s.GetWallet(id)
	if err != nil {
		return nil, err
	}
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	wallet.Balance += amount
	return wallet, nil
}

func (s *walletService) Withdraw(id int64, amount int64) (*model.Wallet, error) {
	wallet, err := s.GetWallet(id)
	if err != nil {
		return nil, err
	}
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	if wallet.Balance < amount {
		return nil, errors.New("insufficient funds")
	}
	wallet.Balance -= amount
	return wallet, nil
}
