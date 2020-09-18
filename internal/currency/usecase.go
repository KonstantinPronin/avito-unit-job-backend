package currency

type Usecase interface {
	Exchange(value float64, cur string) (float64, error)
}
