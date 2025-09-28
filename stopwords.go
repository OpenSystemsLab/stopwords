package stopwords

import (
	"sync"

	"github.com/OpenSystemsLab/stopwords/data"
)

// Registry manages loaded languages to avoid loading all languages into memory
type Registry struct {
	mu        sync.RWMutex
	languages map[string]map[string]struct{}
}

// DefaultRegistry is the global registry instance
var DefaultRegistry = NewRegistry()

// NewRegistry creates a new language registry
func NewRegistry() *Registry {
	return &Registry{
		languages: make(map[string]map[string]struct{}),
	}
}

// RegisterLanguage loads and registers a specific language
func (r *Registry) RegisterLanguage(lang string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.languages[lang]; exists {
		return nil // already loaded
	}

	langData, err := data.LoadLanguage(lang)
	if err != nil {
		return err
	}

	r.languages[lang] = langData
	return nil
}

// RegisterLanguages loads and registers multiple languages
func (r *Registry) RegisterLanguages(langs ...string) error {
	for _, lang := range langs {
		if err := r.RegisterLanguage(lang); err != nil {
			return err
		}
	}
	return nil
}

// IsStopWord checks if a word is a stopword in the given language
func (r *Registry) IsStopWord(lang, word string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	langData, exists := r.languages[lang]
	if !exists {
		return false // language not loaded
	}

	_, ok := langData[word]
	return ok
}

// IsLanguageLoaded checks if a language is currently loaded
func (r *Registry) IsLanguageLoaded(lang string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.languages[lang]
	return exists
}

// LoadedLanguages returns a list of currently loaded languages
func (r *Registry) LoadedLanguages() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	langs := make([]string, 0, len(r.languages))
	for lang := range r.languages {
		langs = append(langs, lang)
	}
	return langs
}

// UnregisterLanguage removes a language from memory
func (r *Registry) UnregisterLanguage(lang string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.languages, lang)
}

// Clear removes all loaded languages from memory
func (r *Registry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.languages = make(map[string]map[string]struct{})
}

// Convenience functions for the default registry

// RegisterLanguage loads and registers a specific language in the default registry
func RegisterLanguage(lang string) error {
	return DefaultRegistry.RegisterLanguage(lang)
}

// RegisterLanguages loads and registers multiple languages in the default registry
func RegisterLanguages(langs ...string) error {
	return DefaultRegistry.RegisterLanguages(langs...)
}

// IsStopWord checks if a word is a stopword in the given language using the default registry
func IsStopWord(lang, word string) bool {
	return DefaultRegistry.IsStopWord(lang, word)
}

// IsLanguageLoaded checks if a language is currently loaded in the default registry
func IsLanguageLoaded(lang string) bool {
	return DefaultRegistry.IsLanguageLoaded(lang)
}

// LoadedLanguages returns a list of currently loaded languages from the default registry
func LoadedLanguages() []string {
	return DefaultRegistry.LoadedLanguages()
}

// UnregisterLanguage removes a language from memory in the default registry
func UnregisterLanguage(lang string) {
	DefaultRegistry.UnregisterLanguage(lang)
}

// Clear removes all loaded languages from memory in the default registry
func Clear() {
	DefaultRegistry.Clear()
}

// GetSupportedLanguages returns a list of all supported languages
func GetSupportedLanguages() []string {
	return data.GetSupportedLanguages()
}
