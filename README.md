# Simple Slot

A modular and flexible slot machine engine implemented in Go. This project provides a clean architecture for creating and running different slot game engines with customizable behavior.

## Project Structure

```
simpleSlot/
├── src/                    # Source code
│   ├── bootstrap/          # Application bootstrap code
│   │   └── start.go
│   ├── engine/             # Game engines
│   │   ├── shared/         # Shared domain objects for all engines
│   │   │   └── domain/     # Core domain models (reels, symbols, etc.)
│   │   └── superwin/       # SuperWin engine implementation
│   │       ├── domain/     # Engine-specific domain logic
│   │       └── infrastructure/  # Engine implementation details
│   ├── platform/           # Platform-specific code
│   │   └── core/
│   │       └── application/  # Application services
│   └── shared/             # Shared components across the application
│       ├── domain/         # Core domain interfaces
│       │   └── ports/      # Interfaces for slot engines and RNG
│       └── infrastructure/ # Infrastructure implementations
│           └── rng/        # Random number generator implementation
```

## Architecture

The project follows clean architecture principles:

- **Domain Layer**: Contains core business logic and entities (symbols, reels, paytables)
- **Application Layer**: Orchestrates use cases with domain objects
- **Infrastructure/Adapter Layer**: Provides concrete implementations of interfaces
- **Interface/Ports Layer**: Handles external interactions (not fully implemented yet)

## Modularity and Flexibility

The system is designed to be modular and flexible:

1. **Engine Modularity**: Each slot game engine is isolated in its own package
2. **Dependency Injection**: Components are wired together via interfaces
3. **Shared Domain Objects**: Common slot elements are reusable across engines
4. **Domain-Driven Design**: Focus on slot game concepts as the core domain

## How to Add a New Engine

To create a new slot engine:

1. Create a new directory under `src/engine/` (e.g., `src/engine/mynewengine/`)
2. Create subdirectories for:
   - `domain/`: Domain logic specific to your engine
   - `infrastructure/`: Implementation details

3. Create the necessary adapters in `infrastructure/`

4. Register your engine in the application bootstrap if needed

## Running Tests

Run tests for the entire project:

```bash
cd src
go test ./...
```

Run tests with coverage report:

```bash
cd src
go test ./... -cover
```

Generate a detailed HTML coverage report:

```bash
cd src
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## RTP Testing

RTP (Return To Player) tests ensure the game's mathematical model behaves as expected:

```bash
cd src/engine/superwin/domain
go test -v -count=1 .
```

This test spins the reels millions of times to verify the actual RTP is close to the target RTP.

Adjust the number of spins, target RTP, and acceptable epsilon in the test for different levels of precision.

## Running the Project

To run the project:

```bash
cd src
go run main.go
```

This starts the application with the default configuration. You can also specify custom configuration through environment variables or command-line flags, depending on your implementation.

## Development Notes

- The codebase uses Go modules for dependency management
- Tests have been created for all objects in the `engine/shared/domain` package
- Current test coverage is approximately 97.5% for shared domain objects

## Architecture Decisions

1. **RNG Abstraction**: Random number generation is abstracted behind an interface to allow for different implementations (including deterministic ones for testing)
2. **Money Value Object**: A dedicated type for handling monetary values prevents floating-point issues
3. **Engine Isolation**: Each engine is isolated to allow for different game mechanics without affecting other engines
4. **Separate Domain and Infrastructure**: Clear separation of concerns between domain logic and implementation details

## Future Improvements

- Logging. Add proper logging with a proper logger.
- Error handling. There are some edges cases that are not handled properly.  
- Add HTTP or GPRC API for controlling the slot engines.
- Implement a web-based frontend for visualization.
- Add persistence for game state and player data.
- Expand the engine varieties with different game mechanics.
- Generalize the load of many engines. 
