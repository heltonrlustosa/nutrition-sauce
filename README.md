# 🥗 Nutrition Sauce - Food & Nutrition Data Server for AI Agents

<div align="center">

An [MCP](https://modelcontextprotocol.io/) server written in Go that provides comprehensive food and nutrition data to AI agents.

[![Go Version](https://img.shields.io/badge/Go-1.20+-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)]()

</div>

---

## 📖 Overview

**Nutrition Sauce MCP** is an MCP server designed to provide AI agents with rich, structured nutritional information about foods and dishes. Whether your agent needs to calculate calorie content, analyze macros, or suggest healthy alternatives, this server acts as a reliable knowledge source for food and nutrition queries.

### What It Does

- ✅ Provides nutritional data (calories, proteins, carbs, fats, vitamins, minerals)
- ✅ Supports various food items from different cuisine types
- ✅ Compatible with the Model Context Protocol standard
- ✅ Built with Go for performance and reliability

---

## 🚀 Features

| Feature | Description |
|---------|-------------|
| **Nutritional Data** | Complete macronutrient and micronutrient information |
| **Food Search** | Search by name, category, or specific nutrients |
| **Recipe Support** | Nutritional breakdown for complete dishes |
| **Unit Conversion** | Flexible serving size and unit handling |
| **Structured Output** | MCP-compliant responses for easy agent parsing |

---

## 📦 Installation

### Prerequisites

- Go 1.20 or higher
- A compatible MCP client (e.g., [Smolagents](https://github.com/huggingface/smolagents))

### Quick Start

```bash
# Clone the repository
git clone https://github.com/your-org/nutrition-sauce.git
cd nutrition-sauce

# Build the server
go build -o nutrition-sauce-server main.go

# Run as standalone (optional)
./nutrition-sauce-server --config config/config.yaml
```

### Using with an MCP Client

If you're using a compatible AI agent framework:

```bash
# Configure your MCP client to point to this server
{
  "mcpServers": {
    "nutrition-sauce": {
      "command": "./nutrition-sauce-server",
      "config": {
        "apiEndpoint": "http://localhost:8080"
      }
    }
  }
}
```

---

## 🛠️ Usage Examples

### Example Query

**Agent Query:**  
> "What are the nutritional facts of Gorgonzola cheese?"

**Server Response:**
```json
{
  "food": {
    "name": "Gorgonzola cheese",
    "category": "dairy",
    "base_amount": 22.66,
  },
  "nutrition": {
    "calories": 239,
    "macronutrients": {
      "protein": 27.5,   // g
      "carbohydrates": 0,      // g
      "fat": 13.6           // g
    },
    "micronutrients": {
      "sodium": 94,       // mg
      "potassium": 256,    // mg
      "iron": 1.1          // mg
    }
  },
  "tags": ["high-protein", "low-carb"]
}
```

---

## ⚙️ Configuration

The server can be configured using a JSON config file:

```yaml
# config/config.yaml
server:
  host: "127.0.0.1"
  port: 8080
  debug: false

data_source:
  type: "local"
  path: "./data/nutrition_data.json"
```

---

## 📁 Project Structure

```
├── main.go                    # Application entry point
├── go.mod                     # Go module definition
│
├── cmd/                       # Entry points
│   └── nutrition-sauce/
│       └── main.go            # Server bootstrap with DI
│
├── internal/                  # Private application code
│   ├── domain/                # Domain layer - business logic & entities
│   │   ├── food/              # Food entity and domain interfaces
│   │   │   ├── food.go        # Food entity model
│   │   │   └── food_repo.go   # FoodRepository interface
│   │   └── nutrition/         # Nutrition calculation domain
│   │       └── nutrition_calc.go  # NutritionCalculator interface
│   │
│   ├── application/           # Use case layer (application logic)
│   │   └── usecases/
│   │       └── food_search_usecase.go
│   │
│   └── infrastructure/       # Implementation layer
│       ├── data/              # Data access implementations
│       │   └── json_repository.go  # JSON database abstraction
│       │
│       └── mcp/               # MCP protocol implementation
│           ├── server.go      # MCP server implementation
│           ├── transport.go   # HTTP/WebSocket transport
│           └── handler.go     # JSON request/response handlers
│
├── config/                    # Configuration files
│   ├── config.yaml            # Server configuration
│   └── .env.example           # Environment variables template
│
├── docs/                      # Documentation
│   ├── adr/                   # Architecture Decision Records
│   │   ├── 001-platform-and-architecture.md
│   │   └── 002-data-storage.md
│   │
│   └── assets/                # Diagrams and visual assets
│       └── diagrams/          # Architecture diagrams
│           ├── clean-arch.svg
│           ├── sequence-diagram.mmd
│           ├── component-diagram.mmd
│           └── system-overview.mmd
│
├── scripts/                   # Build/test utilities
│   ├── build.sh               # Build script
│   └── test.sh                # Test runner script
│
└── .gitignore                 # Git ignore rules
```

---

## 🔧 Development

### Setting up the development environment

```bash
# Install Go dependencies
go mod download

# Run with debugging enabled
go run -race main.go --debug

# Build for production
CGO_ENABLED=0 go build -ldflags="-s -w" -o nutrition-sauce-server
```

---

## 📚 API Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/food/search` | Search for food items |
| GET | `/api/v1/food/{id}` | Get detailed nutrition info |
| GET | `/api/v1/categories` | List all food categories |

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🔗 Resources

- [Model Context Protocol Documentation](https://modelcontextprotocol.io/)
- [Go Project Guidelines](https://go.dev/doc/projects)
- **Nutrition Sauce Platform** - Reference platform for nutrition data

---

## 📧 Support & Contact

For questions, issues, or feature requests:

- GitHub Issues: [Open an issue](https://github.com/your-org/nutrition-sauce/issues)

---

<div align="center">

**Made with ❤️ for AI developers**

</div>
