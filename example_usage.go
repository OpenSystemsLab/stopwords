package main

import (
	"fmt"
	"log"

	"github.com/OpenSystemsLab/stopwords"
)

func main() {
	// Example 1: Register a single language
	fmt.Println("=== Example 1: Single Language Registration ===")

	// Register only English - saves memory by not loading all languages
	err := stopwords.RegisterLanguage("en")
	if err != nil {
		log.Fatal(err)
	}

	// Check if words are stopwords
	fmt.Printf("Is 'the' a stopword in English? %t\n", stopwords.IsStopWord("en", "the"))
	fmt.Printf("Is 'car' a stopword in English? %t\n", stopwords.IsStopWord("en", "car"))

	// Try unregistered language - returns false
	fmt.Printf("Is 'le' a stopword in French (not loaded)? %t\n", stopwords.IsStopWord("fr", "le"))

	fmt.Println()

	// Example 2: Register multiple languages
	fmt.Println("=== Example 2: Multiple Language Registration ===")

	// Register multiple languages at once
	err = stopwords.RegisterLanguages("fr", "es", "de")
	if err != nil {
		log.Fatal(err)
	}

	// Now French works
	fmt.Printf("Is 'le' a stopword in French? %t\n", stopwords.IsStopWord("fr", "le"))
	fmt.Printf("Is 'el' a stopword in Spanish? %t\n", stopwords.IsStopWord("es", "el"))
	fmt.Printf("Is 'der' a stopword in German? %t\n", stopwords.IsStopWord("de", "der"))

	fmt.Println()

	// Example 3: Check loaded languages
	fmt.Println("=== Example 3: Language Management ===")

	loaded := stopwords.LoadedLanguages()
	fmt.Printf("Currently loaded languages: %v\n", loaded)

	fmt.Printf("Is English loaded? %t\n", stopwords.IsLanguageLoaded("en"))
	fmt.Printf("Is Japanese loaded? %t\n", stopwords.IsLanguageLoaded("ja"))

	fmt.Println()

	// Example 4: Get supported languages
	fmt.Println("=== Example 4: Supported Languages ===")

	supported := stopwords.GetSupportedLanguages()
	fmt.Printf("Total supported languages: %d\n", len(supported))
	fmt.Printf("First 10 supported languages: %v\n", supported[:10])

	fmt.Println()

	// Example 5: Memory management
	fmt.Println("=== Example 5: Memory Management ===")

	// Unregister a language to free memory
	stopwords.UnregisterLanguage("de")
	fmt.Printf("Is German still loaded after unregistering? %t\n", stopwords.IsLanguageLoaded("de"))

	// Clear all loaded languages
	stopwords.Clear()
	loaded = stopwords.LoadedLanguages()
	fmt.Printf("Languages after clearing: %v\n", loaded)

	fmt.Println()

	// Example 6: Using custom registry
	fmt.Println("=== Example 6: Custom Registry ===")

	// Create a custom registry for isolated use
	customRegistry := stopwords.NewRegistry()

	// Load language only in custom registry
	err = customRegistry.RegisterLanguage("ja")
	if err != nil {
		log.Fatal(err)
	}

	// Japanese works in custom registry but not in default
	fmt.Printf("Japanese in custom registry: %t\n", customRegistry.IsStopWord("ja", "の"))
	fmt.Printf("Japanese in default registry: %t\n", stopwords.IsStopWord("ja", "の"))

	fmt.Println()

	// Example 7: Error handling
	fmt.Println("=== Example 7: Error Handling ===")

	// Try to register unsupported language
	err = stopwords.RegisterLanguage("unsupported")
	if err != nil {
		fmt.Printf("Expected error for unsupported language: %v\n", err)
	}
}
