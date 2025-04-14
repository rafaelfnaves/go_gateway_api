package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByID(id string) (*Account, error)
	FindByAPIKey(api_key string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
