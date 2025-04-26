# Go Password Generator

A flexible and secure random password generator implemented in Go.

## Features

- Customizable password length
- Configurable character set
- Secure random generation using `math/rand` with time-based seeding
- Prevents consecutive duplicate characters
- Simple command-line interface

## Installation

```bash
go get github.com/softwareplace/go-password
```

## Usage

### Basic Usage

```bash
go run cmd/main.go
```

Generates a 24-character password using the default character set.

### Custom Length

```bash
go run cmd/main.go -length=32
```

Generates a 32-character password.

### Custom Character Set

```bash
go run cmd/main.go -chars="ABCDEF123!@#"
```

Generates a password using only the specified characters.

### Help

```bash
go run cmd/main.go -help
```

Displays usage information.

## API Documentation

### `Generator` Struct

```go
type Generator struct {
    Length int    // Length of the password to generate
    Chars  string // Character set to use for generation
}
```

### Methods

#### `SetLength(length int) *Generator`
Sets the password length.

#### `SetChars(chars string) *Generator`
Sets the character set to use.

#### `Generate() string`
Generates and returns the password.

### Helper Functions

#### `New() *Generator`
Creates a new Generator instance with default values.

#### `Default() *Generator`
Creates a new Generator instance configured via command-line flags.

## Default Character Set

The default character set includes:
- Lowercase letters: `a-z`
- Uppercase letters: `A-Z`
- Numbers: `0-9`
- Special characters: `!@#$%^&*?`

## Examples

### Programmatic Usage

```go
package main

import (
    "github.com/softwareplace/go-password/pkg/str"
)

func main() {
    // Using defaults
    password := str.Default().Generate()
    
    // Custom configuration
    customPass := str.New().
        SetLength(16).
        SetChars("ABCD1234").
        Generate()
}
```

### Command Line Examples

1. Generate a 12-character password:
   ```bash
   go run cmd/main.go -length=12
   ```

2. Generate a password with only numbers:
   ```bash
   go run cmd/main.go -chars="0123456789"
   ```

3. Generate a complex 20-character password:
   ```bash
   go run cmd/main.go -length=20 -chars="abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!@#$%^&*?"
   ```

## Security Notes

- The generator uses Go's `math/rand` package seeded with the current time
- For cryptographic applications, consider using `crypto/rand` instead
- The implementation prevents consecutive duplicate characters
- Minimum recommended length is 12 characters for basic security


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
