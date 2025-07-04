package dtos

import (
	"api-money-management/internal/models"
	"time"
)

type RequestWallet struct {
	Name    string  `json:"name" validate:"required"`
	Balance float64 `json:"balance" validate:"number,gte=0"`
}

type WalletResponse struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WalletListResponse struct {
	Wallets []WalletResponse `json:"wallets"`
}

func ToWalletResponse(wallet *models.Wallet) *WalletResponse {
	return &WalletResponse{
		ID:        wallet.ID,
		UserId:    wallet.UserId,
		Name:      wallet.Name,
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}
}
func ToWalletListResponse(wallets []models.Wallet) *WalletListResponse {
	var walletResponse []WalletResponse
	for _, wallet := range wallets {
		walletResponse = append(walletResponse, *ToWalletResponse(&wallet))
	}
	return &WalletListResponse{
		Wallets: walletResponse,
	}
}
