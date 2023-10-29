# Start.gg GraphQL Interface

This project is a Go interface for [start.gg's GraphQL API](https://developer.start.gg/docs/intro).

# Getting Started

Create your start.gg API key. [Guide](https://developer.start.gg/docs/authentication)

For an example of creating and using this package, see `main.go`

To run main.go in this project, place your API key into a `.env` file.

```
api_key=your_key_here
```

# Use Cases

This library will provide functions for turning pre-defined queries into Go structs.

Higher level functions will combine many queries to create reports for players, tournaments, or games.

