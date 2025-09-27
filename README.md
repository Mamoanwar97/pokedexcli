# Pokédex CLI

A command-line interface Pokédex application built in Go that allows you to explore Pokémon locations, catch Pokémon, and manage your collection using the [PokéAPI](https://pokeapi.co/).

## Features

- 🗺️ **Location Exploration**: Browse through different Pokémon locations
- 🔍 **Pokémon Discovery**: Explore locations to find available Pokémon
- ⚡ **Pokémon Catching**: Attempt to catch Pokémon with randomized success rates
- 📖 **Pokédex Management**: View and inspect your caught Pokémon collection
- 🚀 **Caching System**: Built-in cache for improved performance and reduced API calls
- 💬 **Interactive REPL**: User-friendly command-line interface

## Installation

### Prerequisites

- Go 1.24.4 or later

### Build from Source

1. Clone the repository:
```bash
git clone <repository-url>
cd pokedexcli
```

2. Build the application:
```bash
go build -o pokedexcli
```

3. Run the application:
```bash
./pokedexcli
```

## Usage

Once you start the application, you'll see the `Pokedex >` prompt. Here are the available commands:

### Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `help` | Displays help message with all available commands | `help` |
| `map` | Get the next page of locations | `map` |
| `mapb` | Get the previous page of locations | `mapb` |
| `explore <location>` | Explore a specific location to see available Pokémon | `explore canalave-city-area` |
| `catch <pokemon>` | Attempt to catch a Pokémon | `catch pikachu` |
| `inspect <pokemon>` | View details of a caught Pokémon | `inspect pikachu` |
| `pokedex` | Show all caught Pokémon | `pokedex` |
| `exit` | Exit the Pokédex | `exit` |

### Example Session

```
Pokedex > help

Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Get the next page of locations
mapb: Get the previous page of locations
exit: Exit the Pokedex
explore: Explore a location
catch: Catch a pokemon
inspect: Inspect a pokemon
pokedex: Show all caught pokemon

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
tentacool
tentacruel
staryu
...

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Base Experience: 112

Pokedex > pokedex
Pokedex:
- pikachu

Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Architecture

The project is organized into several key components:

### Core Files

- `main.go` - Application entry point and configuration setup
- `repl.go` - Read-Eval-Print Loop implementation and command routing
- `commands_*.go` - Individual command implementations

### Internal Packages

#### `internal/pokeapi/`
- **API Client**: Handles HTTP requests to the PokéAPI
- **Data Models**: Defines structures for API responses
- **Endpoints**: 
  - Location listing and exploration
  - Pokémon details retrieval

#### `internal/pokecache/`
- **Caching System**: Thread-safe cache implementation
- **Performance Optimization**: Reduces API calls and improves response times
- **Automatic Cleanup**: Background goroutine for cache expiration

### Key Features

#### Pokémon Catching Mechanism
The catching system uses a randomized algorithm based on the Pokémon's base experience:
- Higher base experience = harder to catch
- 50% base success rate with randomization
- Successful catches are stored in your personal Pokédex

#### Caching System
- **Duration**: 5-minute cache expiration
- **Thread-Safe**: Uses mutex locks for concurrent access
- **Automatic Cleanup**: Background process removes expired entries

## Configuration

The application is configured with the following defaults:
- **HTTP Timeout**: 5 seconds
- **Cache Duration**: 5 minutes
- **API Base URL**: `https://pokeapi.co/api/v2`

## Testing

Run the test suite:
```bash
go test ./...
```

The project includes tests for:
- Input cleaning and validation
- Cache functionality

## Development

### Project Structure
```
pokedexcli/
├── main.go                 # Application entry point
├── repl.go                 # REPL implementation
├── commands_*.go           # Command implementations
├── go.mod                  # Go module definition
├── internal/
│   ├── pokeapi/           # PokéAPI client and models
│   └── pokecache/         # Caching system
└── README.md              # This file
```

### Adding New Commands

1. Create a new `commands_<name>.go` file
2. Implement the command function with signature `func(cfg *config) error`
3. Add the command to the `getCommands()` map in `repl.go`

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Dependencies

- **Standard Library Only**: The project uses only Go's standard library
- **External API**: [PokéAPI](https://pokeapi.co/) for Pokémon data

## License

This project is built for educational purposes as part of the Boot.dev curriculum.

## Acknowledgments

- [PokéAPI](https://pokeapi.co/) for providing the Pokémon data
- [Boot.dev](https://boot.dev/) for the project inspiration and guidance

