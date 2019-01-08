package matomat

type Config struct {
	CreditMin         int32
	CreditMax         int32
	ItemNameMinLength int
	ItemNameMaxLength int
}

func NewConfig(creditMin int32, creditMax int32, itemNameMinLength int, itemNameMaxLength int) *Config {
	return &Config{CreditMin: creditMin, CreditMax: creditMax, ItemNameMinLength: itemNameMinLength, ItemNameMaxLength: itemNameMaxLength}
}
