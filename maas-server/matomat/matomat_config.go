package matomat

type Config struct {
	AllowCreditDebt   bool
	ItemNameMinLength int
	ItemNameMaxLength int
}

func NewConfig(allowCreditDebt bool, itemNameMinLength int, itemNameMaxLength int) *Config {
	return &Config{AllowCreditDebt: allowCreditDebt, ItemNameMinLength: itemNameMinLength, ItemNameMaxLength: itemNameMaxLength}
}
