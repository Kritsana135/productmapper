# Product Mapper

A Go library for processing and cleaning product orders with support for complementary items and price diffusion.

## Overview

Product Mapper is a Go package that provides functionality to process and clean product orders from various platforms. It handles:
- Product ID extraction and parsing
- Price diffusion across product components
- Complementary item management
- Order cleaning and standardization

## Installation

```bash
go get github.com/Kritsana135/productmapper
```

## Usage

```go
import "github.com/Kritsana135/productmapper"

// Create input orders
orders := []productmapper.InputOrder{
    {
        No:                1,
        PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX*2",
        Qty:               2,
        UnitPrice:         100.00,
        TotalPrice:        200.00,
    },
}

// Define complementary items if needed
complementaryItems := []productmapper.ComplementaryItem{
    // Add complementary items here
}

// Clean and process orders
cleanedOrders, err := productmapper.CleanOrder(context.Background(), orders, complementaryItems)
if err != nil {
    // Handle error
}
```

## Features

- **Order Cleaning**: Transforms platform-specific product IDs into standardized format
- **Price Diffusion**: Distributes prices across product components
- **Complementary Items**: Handles additional items that should be included with orders
- **Comprehensive Testing**: Includes extensive test coverage for all functionality

## Project Structure

- `productmapper.go`: Core functionality for order processing
- `extractor.go`: Product ID extraction and parsing
- `diffuseprice.go`: Price diffusion logic
- `complementary.go`: Complementary item handling
- `*_test.go`: Test files for each component

## Dependencies

- `github.com/stretchr/testify`: Testing utilities
- `github.com/elliotchance/orderedmap`: Ordered map implementation
- `gopkg.in/yaml.v3`: YAML processing
