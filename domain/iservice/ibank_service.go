package iservice

type IBankService interface {
	CreateBank(name string, website string) error
}
