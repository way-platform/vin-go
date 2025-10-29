# vin-go

[![Go Reference](https://pkg.go.dev/badge/github.com/way-platform/vin-go.svg)](https://pkg.go.dev/github.com/way-platform/vin-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/way-platform/vin-go)](https://goreportcard.com/report/github.com/way-platform/vin-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library and CLI tool for decoding Vehicle Identification Numbers (VINs) based on the ISO 3779 standard.

## Features

- **VIN Validation**: Validates VIN format, character set, and check digit
- **Complete Decoding**: Extracts all standardized VIN components:
  - World Manufacturer Identifier (WMI) - manufacturer, country, geographic area
  - Vehicle Descriptor Section (VDS) - manufacturer-specific codes and check digit
  - Vehicle Identifier Section (VIS) - model year, plant code, serial number
- **Extensive Database**: Includes 8000+ manufacturer codes (WMI)
- **CLI Tool**: Command-line interface for easy VIN decoding

## Installation

### As a Library

```bash
go get github.com/way-platform/vin-go
```

### As a CLI Tool

```bash
go install github.com/way-platform/vin-go/cmd/vin@latest
```

Or download prebuilt binaries from the [Releases](https://github.com/way-platform/vin-go/releases) page.

## Usage

### Library Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/way-platform/vin-go"
)

func main() {
    vinStr := "1HGBH41JXMN109186"

    decoded, err := vin.Decode(&vinStr)
    if err != nil {
        log.Fatalf("Failed to decode VIN: %v", err)
    }

    fmt.Printf("Manufacturer: %s\n", decoded.WMI.Manufacturer)
    fmt.Printf("Country: %s\n", decoded.WMI.Country)
    fmt.Printf("Model Year: %s\n", decoded.VIS.ModelYear)
    fmt.Printf("Serial Number: %s\n", decoded.VIS.SerialNumber)
}
```

### CLI Usage

Decode a VIN and display the results as JSON:

```bash
vin decode 1HGBH41JXMN109186
```

Output:

```json
{
  "vin": "1HGBH41JXMN109186",
  "wmi": {
    "code": "1HG",
    "manufacturer": "Honda",
    "country": "United States",
    "geographic_area": "North America"
  },
  "vds": {
    "manufacturerSpecific": "BH41J",
    "checkDigit": "X",
    "checkDigitValid": true
  },
  "vis": {
    "modelYear": "2021",
    "plantCode": "M",
    "serialNumber": "109186"
  }
}
```

## VIN Structure

A VIN consists of 17 characters divided into three sections:

1. **World Manufacturer Identifier (WMI)** - Characters 1-3

   - Identifies the manufacturer and country of origin
   - First character indicates geographic area

2. **Vehicle Descriptor Section (VDS)** - Characters 4-9

   - Characters 4-8: Manufacturer-specific vehicle attributes
   - Character 9: Check digit for VIN validation

3. **Vehicle Identifier Section (VIS)** - Characters 10-17
   - Character 10: Model year
   - Character 11: Manufacturing plant code
   - Characters 12-17: Sequential serial number

## Development

### Prerequisites

- Go 1.24 or later
- [buf](https://buf.build/docs/installation) (automatically installed via `go install` when needed)
- [Mage](https://magefile.org/) (optional, for build automation)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/way-platform/vin-go.git
cd vin-go

# Generate code and build (run from project root with -d tools)
go tool mage build

# The binary will be in bin/vin
./bin/vin decode 1HGBH41JXMN109186
```

### Available Mage Targets

```bash
go tool mage -l
```

## Data Sources

- **ISO 3779**: International standard for VIN structure
- **WMI Database**: Comprehensive database of World Manufacturer Identifiers
- **Model Year Codes**: Standard 30-year cycle of model year codes

## License

MIT License - see [LICENSE](LICENSE) for details.

## Contributing

Contributions are welcome!
