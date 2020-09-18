package currency

type Repository interface {
	UpdateCache() (map[string]float32, error)
	GetByCurrency(currency string) (float32, error)
}
