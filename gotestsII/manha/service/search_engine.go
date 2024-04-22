package service

type SearchEngine struct {
	engine Engine
}

type Engine interface {
	SearchByName(name string) string
	SearchByPhone(phone string) string
	AddEntry(name, phone string) error
}

func (s *SearchEngine) SearchByName(name string) string {
	if len(name) <= 3 {
		return ""
	}

	return s.engine.SearchByName(name)
}

func (s *SearchEngine) SearchByPhone(phone string) string {
	if len(phone) < 5 {
		return ""
	}

	return s.engine.SearchByPhone(phone)
}

func (s *SearchEngine) AddEntry(name, phone string) error {
	return s.engine.AddEntry(name, phone)
}

func (s *SearchEngine) GetVersion() string {
	return "1.0"
}

func NewEngine(search Engine) *SearchEngine {
	return &SearchEngine{
		engine: search,
	}
}
