package i18n

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/text/language"
)

// Translator interface for the internationalization system
type Translator interface {
	Translate(key string) string
	Translatef(key string, args ...interface{}) string
	SetLanguage(lang string) error
	GetLanguage() string
	GetAvailableLanguages() []string
}

// i18n implements the internationalization system
type i18n struct {
	mu           sync.RWMutex
	translations map[string]map[string]string
	currentLang  string
}

// New creates a new instance of the internationalization system
// It automatically detects the system language and falls back to English
func New() (Translator, error) {
	i := &i18n{
		translations: make(map[string]map[string]string),
	}

	// Detect system language
	systemLang := detectSystemLanguage()

	// Try to load the system language
	if err := i.loadLanguage(systemLang); err == nil {
		i.currentLang = systemLang
		log.Printf("System language detected: %s", systemLang)
	} else {
		// Fallback to English
		if err := i.loadLanguage("en"); err != nil {
			return nil, fmt.Errorf("failed to load fallback language (en): %v", err)
		}
		i.currentLang = "en"
		log.Printf("Using fallback language: en")
	}

	return i, nil
}

// NewWithLanguage creates a new instance with a specific language
func NewWithLanguage(lang string) (Translator, error) {
	i := &i18n{
		translations: make(map[string]map[string]string),
	}

	if err := i.loadLanguage(lang); err != nil {
		// Fallback to English if requested language is not available
		if err := i.loadLanguage("en"); err != nil {
			return nil, fmt.Errorf("failed to load languages: %v", err)
		}
		i.currentLang = "en"
	} else {
		i.currentLang = lang
	}

	return i, nil
}

// detectSystemLanguage detects the operating system's language
func detectSystemLanguage() string {
	// Check environment variables (Linux/macOS)
	langEnv := os.Getenv("LANG")
	if langEnv != "" {
		if strings.Contains(langEnv, "pt") {
			return "pt"
		} else if strings.Contains(langEnv, "es") {
			return "es"
		} else if strings.Contains(langEnv, "fr") {
			return "fr"
		} else if strings.Contains(langEnv, "de") {
			return "de"
		} else if strings.Contains(langEnv, "it") {
			return "it"
		} else if strings.Contains(langEnv, "ru") {
			return "ru"
		} else if strings.Contains(langEnv, "zh") {
			return "zh"
		} else if strings.Contains(langEnv, "ja") {
			return "ja"
		} else if strings.Contains(langEnv, "ko") {
			return "ko"
		}
	}

	// Check other common environment variables
	for _, envVar := range []string{"LC_ALL", "LC_MESSAGES", "LANGUAGE"} {
		envValue := os.Getenv(envVar)
		if envValue != "" {
			if strings.Contains(envValue, "pt") {
				return "pt"
			} else if strings.Contains(envValue, "es") {
				return "es"
			} else if strings.Contains(envValue, "fr") {
				return "fr"
			} else if strings.Contains(envValue, "de") {
				return "de"
			} else if strings.Contains(envValue, "it") {
				return "it"
			} else if strings.Contains(envValue, "ru") {
				return "ru"
			} else if strings.Contains(envValue, "zh") {
				return "zh"
			} else if strings.Contains(envValue, "ja") {
				return "ja"
			} else if strings.Contains(envValue, "ko") {
				return "ko"
			}
		}
	}

	// Use golang.org/x/text/language for more precise detection
	if langEnv != "" {
		tag, err := language.Parse(langEnv)
		if err == nil {
			base, _ := tag.Base()

			switch base.String() {
			case "pt":
				return "pt"
			case "es":
				return "es"
			case "fr":
				return "fr"
			case "de":
				return "de"
			case "it":
				return "it"
			case "ru":
				return "ru"
			case "zh":
				return "zh"
			case "ja":
				return "ja"
			case "ko":
				return "ko"
			}
		}
	}

	// Default fallback
	return "en"
}

// loadLanguage loads translations from a JSON file
func (i *i18n) loadLanguage(lang string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Path to the translation file
	path := filepath.Join("assets", "i18n", lang+".json")

	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("language file not found: %s", path)
	}

	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Decode JSON
	var translations map[string]string
	if err := json.Unmarshal(data, &translations); err != nil {
		return err
	}

	i.translations[lang] = translations
	return nil
}

// SetLanguage sets the current language
func (i *i18n) SetLanguage(lang string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	// If language is already loaded, just set as current
	if _, exists := i.translations[lang]; exists {
		i.currentLang = lang
		return nil
	}

	// Try to load the new language
	if err := i.loadLanguage(lang); err != nil {
		return err
	}

	i.currentLang = lang
	return nil
}

// GetLanguage returns the current language
func (i *i18n) GetLanguage() string {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.currentLang
}

// GetAvailableLanguages returns the list of available languages
func (i *i18n) GetAvailableLanguages() []string {
	i.mu.RLock()
	defer i.mu.RUnlock()

	languages := make([]string, 0, len(i.translations))
	for lang := range i.translations {
		languages = append(languages, lang)
	}
	return languages
}

// Translate returns the translation for the provided key
func (i *i18n) Translate(key string) string {
	i.mu.RLock()
	defer i.mu.RUnlock()

	// Try current language first
	if translations, exists := i.translations[i.currentLang]; exists {
		if text, found := translations[key]; found {
			return text
		}
	}

	// Fallback: try English if not found in current language
	if i.currentLang != "en" {
		if translations, exists := i.translations["en"]; exists {
			if text, found := translations[key]; found {
				return text
			}
		}
	}

	// Final fallback: return the key itself if no translation found
	return key
}

// Translatef returns a formatted translation
func (i *i18n) Translatef(key string, args ...interface{}) string {
	baseText := i.Translate(key)
	return fmt.Sprintf(baseText, args...)
}

// fallbackTranslator provides a minimal implementation for fallback scenarios
type fallbackTranslator struct{}

func (f *fallbackTranslator) Translate(key string) string {
	return key
}

func (f *fallbackTranslator) Translatef(key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}

func (f *fallbackTranslator) SetLanguage(lang string) error {
	return nil
}

func (f *fallbackTranslator) GetLanguage() string {
	return "en"
}

func (f *fallbackTranslator) GetAvailableLanguages() []string {
	return []string{"en"}
}
