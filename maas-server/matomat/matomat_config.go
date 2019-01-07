package matomat

type Config struct {
	AllowCreditDebt   bool
	ItemNameMinLength uint32
	ItemNameMaxLength uint32
}

func NewConfig(allowCreditDebt bool, itemNameMinLength uint32, itemNameMaxLength uint32) *Config {
	return &Config{AllowCreditDebt: allowCreditDebt, ItemNameMinLength: itemNameMinLength, ItemNameMaxLength: itemNameMaxLength}
}
