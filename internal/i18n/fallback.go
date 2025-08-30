package i18n

import "fmt"

// FallbackTranslator provides a minimal implementation for fallback scenarios
type FallbackTranslator struct{}

func (f *FallbackTranslator) Translate(key string) string {
	return key
}

func (f *FallbackTranslator) Translatef(key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}

func (f *FallbackTranslator) SetLanguage(lang string) error {
	return nil
}

func (f *FallbackTranslator) GetLanguage() string {
	return "en"
}

func (f *FallbackTranslator) GetAvailableLanguages() []string {
	return []string{"en"}
}
