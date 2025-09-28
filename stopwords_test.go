package stopwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsStopWord(t *testing.T) {
	// Clear any previously loaded languages
	Clear()

	tt := []struct {
		lang string
		word string
		want bool
	}{
		{"fr", "au", true},
		{"fr", "aux", true},
		{"fr", "avec", true},
		{"fr", "ce", true},
		{"fr", "ces", true},
		{"fr", "Voiture", false},
		{"bad", "bad", false}, // This should return false since language won't be loaded
		{"en", "the", true},
		{"en", "car", false},
	}

	// Register required languages for testing
	err := RegisterLanguages("fr", "en")
	require.NoError(t, err)

	for _, tc := range tt {
		t.Run(tc.word, func(t *testing.T) {
			got := IsStopWord(tc.lang, tc.word)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRegisterLanguage(t *testing.T) {
	// Clear any previously loaded languages
	Clear()

	// Test that language is not loaded initially
	assert.False(t, IsLanguageLoaded("en"))

	// Register English
	err := RegisterLanguage("en")
	require.NoError(t, err)

	// Test that language is now loaded
	assert.True(t, IsLanguageLoaded("en"))

	// Test that stopwords work
	assert.True(t, IsStopWord("en", "the"))
	assert.False(t, IsStopWord("en", "car"))

	// Test registering invalid language
	err = RegisterLanguage("invalid")
	require.Error(t, err)
}

func TestRegisterLanguages(t *testing.T) {
	// Clear any previously loaded languages
	Clear()

	// Register multiple languages
	err := RegisterLanguages("en", "fr", "es")
	require.NoError(t, err)

	// Test all are loaded
	assert.True(t, IsLanguageLoaded("en"))
	assert.True(t, IsLanguageLoaded("fr"))
	assert.True(t, IsLanguageLoaded("es"))

	// Test loaded languages list
	loaded := LoadedLanguages()
	assert.Len(t, loaded, 3)
	assert.Contains(t, loaded, "en")
	assert.Contains(t, loaded, "fr")
	assert.Contains(t, loaded, "es")
}

func TestUnregisterLanguage(t *testing.T) {
	// Clear any previously loaded languages
	Clear()

	// Register a language
	err := RegisterLanguage("en")
	require.NoError(t, err)
	assert.True(t, IsLanguageLoaded("en"))

	// Unregister it
	UnregisterLanguage("en")
	assert.False(t, IsLanguageLoaded("en"))

	// Test that stopwords no longer work
	assert.False(t, IsStopWord("en", "the"))
}

func TestGetSupportedLanguages(t *testing.T) {
	supported := GetSupportedLanguages()

	// Should have many languages
	assert.Greater(t, len(supported), 50)

	// Should contain common languages
	assert.Contains(t, supported, "en")
	assert.Contains(t, supported, "fr")
	assert.Contains(t, supported, "es")
	assert.Contains(t, supported, "de")
}

func TestRegistryIsolation(t *testing.T) {
	// Test that different registries are isolated
	reg1 := NewRegistry()
	reg2 := NewRegistry()

	// Load language in first registry
	err := reg1.RegisterLanguage("en")
	require.NoError(t, err)

	// Should only be loaded in first registry
	assert.True(t, reg1.IsLanguageLoaded("en"))
	assert.False(t, reg2.IsLanguageLoaded("en"))

	// Should only work in first registry
	assert.True(t, reg1.IsStopWord("en", "the"))
	assert.False(t, reg2.IsStopWord("en", "the"))
}
