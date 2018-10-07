package models

type balance struct {
	Address string
	N_tx float64
	Total_received float64
	Total_sent float64
	Final_balance float64
}

func InitBalance() balance {
	return balance{}
}