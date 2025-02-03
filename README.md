# go-freed

A command-line utility and Go module for reading, parsing, and logging FreeD protocol data to files.

## Overview
go-freed provides both a library and command-line tool for working with FreeD protocol data, commonly used in virtual production and broadcast environments for camera tracking. The project is functional and actively used.

## Features
- FreeD protocol parsing and interpretation
- File-based data logging
- Command-line interface
- Reusable Go module for FreeD handling
- Real-time data monitoring

## Usage

### As a Command-Line Tool
```bash
go-freed [options] output_file
```

### As a Go Module
```go
import "github.com/ZeroSpace-Studios/go-freed/freed"

// Create a new FreeD listener
freed := freed.New()

// Start listening for FreeD data
data := freed.Listen()
```

## Technical Details
The module handles:
- FreeD packet parsing
- Data validation
- File I/O operations
- Network communication

## Data Format
FreeD protocol includes:
- Camera position (X, Y, Z)
- Camera rotation (Pan, Tilt, Roll)
- Lens data (Zoom, Focus)
- Additional metadata

## Installation
```bash
go get github.com/ZeroSpace-Studios/go-freed
```

## Development
The project is structured as:
```
freed/
├── freed.go  # Core FreeD handling
└── types.go  # Data structures
```
