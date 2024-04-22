package doubles

type StubSearchEngine struct{}

func (d StubSearchEngine) SearchByName(name string) string {
	return "12345678" // note que ao invés de retornar "", retornamos um valor concreto
}

func (d StubSearchEngine) SearchByPhone(phone string) string {
	return ""
}

func (d StubSearchEngine) AddEntry(name, phone string) error {
	return nil
}
