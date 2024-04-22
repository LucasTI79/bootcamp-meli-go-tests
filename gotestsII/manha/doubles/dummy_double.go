package doubles

type DummySearchEngine struct{}

func (d DummySearchEngine) SearchByName(name string) string {
	return ""
}

func (d DummySearchEngine) SearchByPhone(phone string) string {
	return ""
}

func (d DummySearchEngine) AddEntry(name, phone string) error {
	return nil
}
