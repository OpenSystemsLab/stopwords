# Stopwords for Golang ![Last release](https://img.shields.io/github/release/OpenSystemsLab/stopwords.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/OpenSystemsLab/stopwords)](https://goreportcard.com/report/github.com/OpenSystemsLab/stopwords)

| Branch | Status                                                                                                                                                | Coverage                                                                                                                                         |
| ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| master | [![Go](https://github.com/OpenSystemsLab/stopwords/actions/workflows/go.yml/badge.svg)](https://github.com/OpenSystemsLab/stopwords/actions/workflows/go.yml) | [![Coveralls](https://img.shields.io/coveralls/OpenSystemsLab/stopwords/master.svg)](https://coveralls.io/github/OpenSystemsLab/stopwords?branch=master) |

```sh
go get -u github.com/OpenSystemsLab/stopwords
```

## Quick Start

```go
import (
    "fmt"

    "github.com/OpenSystemsLab/stopwords"
)

func main() {
    // Register a language first
    stopwords.RegisterLanguage("fr")
    
    // Check if a word is a stopword
    fmt.Print(stopwords.IsStopWord("fr", "avec")) // true
}
```

## Registry Feature

The library now includes a **Registry** system that allows you to load only the languages you need, saving memory and improving performance. This is especially useful when working with multiple languages or in memory-constrained environments.

### Basic Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/OpenSystemsLab/stopwords"
)

func main() {
    // Register a single language
    err := stopwords.RegisterLanguage("en")
    if err != nil {
        log.Fatal(err)
    }

    // Check stopwords
    fmt.Printf("Is 'the' a stopword? %t\n", stopwords.IsStopWord("en", "the"))
    fmt.Printf("Is 'car' a stopword? %t\n", stopwords.IsStopWord("en", "car"))
}
```

### Multiple Languages

```go
// Register multiple languages at once
err := stopwords.RegisterLanguages("en", "fr", "es", "de")
if err != nil {
    log.Fatal(err)
}

// Now you can check stopwords in any registered language
fmt.Printf("English: %t\n", stopwords.IsStopWord("en", "the"))
fmt.Printf("French: %t\n", stopwords.IsStopWord("fr", "le"))
fmt.Printf("Spanish: %t\n", stopwords.IsStopWord("es", "el"))
fmt.Printf("German: %t\n", stopwords.IsStopWord("de", "der"))
```

### Memory Management

```go
// Check which languages are currently loaded
loaded := stopwords.LoadedLanguages()
fmt.Printf("Loaded languages: %v\n", loaded)

// Check if a specific language is loaded
fmt.Printf("Is English loaded? %t\n", stopwords.IsLanguageLoaded("en"))

// Unregister a language to free memory
stopwords.UnregisterLanguage("de")

// Clear all loaded languages
stopwords.Clear()
```

### Custom Registry

For isolated use cases, you can create your own registry:

```go
// Create a custom registry
customRegistry := stopwords.NewRegistry()

// Load languages only in this registry
err := customRegistry.RegisterLanguage("ja")
if err != nil {
    log.Fatal(err)
}

// Use the custom registry
fmt.Printf("Japanese stopword: %t\n", customRegistry.IsStopWord("ja", "の"))

// The default registry remains unaffected
fmt.Printf("Japanese in default registry: %t\n", stopwords.IsStopWord("ja", "の"))
```

### Supported Languages

Get a list of all supported languages:

```go
supported := stopwords.GetSupportedLanguages()
fmt.Printf("Total supported languages: %d\n", len(supported))
fmt.Printf("Languages: %v\n", supported)
```

The library supports **60+ languages** including English, French, Spanish, German, Japanese, Chinese, Arabic, and many more.

## API Reference

### Registry Methods

- `RegisterLanguage(lang string) error` - Load a single language
- `RegisterLanguages(langs ...string) error` - Load multiple languages
- `IsStopWord(lang, word string) bool` - Check if a word is a stopword
- `IsLanguageLoaded(lang string) bool` - Check if a language is loaded
- `LoadedLanguages() []string` - Get list of loaded languages
- `UnregisterLanguage(lang string)` - Remove a language from memory
- `Clear()` - Remove all languages from memory
- `GetSupportedLanguages() []string` - Get all supported language codes

### Benefits of the Registry System

1. **Memory Efficient**: Only load the languages you need
2. **Thread Safe**: All operations are protected by read/write mutexes
3. **Flexible**: Use the default registry or create custom ones
4. **Error Handling**: Proper error handling for unsupported languages
5. **Performance**: Faster startup and reduced memory footprint

## License

stopwords is licensed under [the MIT license](LICENSE.md).
