package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByID(id string) (*Account, error)
	FindByAPIKey(api_key string) (*Account, error)
	Update(account *Account) error
}
