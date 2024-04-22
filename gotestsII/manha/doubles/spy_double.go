package doubles

type SpySearchEngine struct {
	SearchByPhoneWasCalled bool
}

func (s *SpySearchEngine) SearchByName(name string) string {
	return ""
}

func (s *SpySearchEngine) SearchByPhone(phone string) string {
	s.SearchByPhoneWasCalled = true
	return ""
}

func (s *SpySearchEngine) AddEntry(name, phone string) error {
	return nil
}
