# vPIC Manufacturers Tool Plan

## Objective
Create a CLI tool (`vpic-manufacturers`) that transforms the relational vPIC database (provided as CSV files) into a stream of `Manufacturer` Protocol Buffer messages, serialized as JSONL. This allows the vPIC data to be consumed by the wider `vin-go` ecosystem.

## Input Data
The tool accepts the following vPIC CSV tables:
- `Wmi.csv`: Core registry of World Manufacturer Identifiers.
- `Manufacturer.csv`: Manufacturer names and IDs.
- `Country.csv`: Country mappings.
- `Make.csv`: Brand (Make) names.
- `ManufacturerMake.csv`: Link table between Manufacturers and Makes.
- `WmiMake.csv`: Link table between specific WMIs and Makes.
- `MakeModel.csv`: Link table between Makes and Models.
- `Model.csv`: Model names.

## Architecture

### 1. Data Loading & Indexing
We will load the CSVs into memory and build indices to allow for efficient $O(1)$ lookups during the transformation process.

**Indices:**
- **Manufacturers:** `map[ManufacturerID]ManufacturerRecord`
  - Used to resolve `vpic_id` to `display_name`.
- **ManufacturerMakes:** `map[ManufacturerID][]MakeID`
  - Aggregates all Makes (Brands) associated with a Manufacturer.
- **WmiMakes:** `map[WmiID][]MakeID`
  - Aggregates Makes specifically tied to a WMI record.
- **MakeModels:** `map[MakeID][]ModelID`
  - Aggregates all Models produced under a specific Make.

### 2. Transformation Logic
We iterate through the `Wmi` table. Each valid record (valid WMI string) becomes a candidate for a `Manufacturer` message.

**For each WMI Record:**
1.  **WMI Parsing:**
    - Extract `wmi1` (chars 0-2).
    - Extract `wmi2` (chars 3-5, if present).
    - Create a unique key (`wmi1` + `wmi2`) to deduplicate entries if necessary (though vPIC `Wmi` table IDs are unique, multiple IDs might map to the same WMI code).

2.  **Manufacturer Details:**
    - **`vpic_id`**: Taken from `Wmi.ManufacturerId`.
    - **`display_name`**: Looked up in the `Manufacturers` index using `vpic_id`.

3.  **Country Resolution:**
    - **`country`**: Look up `Wmi.CountryId` using `internal/vpic.ResolveCountry`.

4.  **Vehicle Type Resolution:**
    - **`vehicle_types`**: Look up `Wmi.VehicleTypeId` using `internal/vpic.ResolveVehicleType`.

5.  **Brand (Make) Resolution:**
    - We collect a set of distinct `MakeID`s from:
        1. `Wmi.MakeId` (Direct column).
        2. `WmiMakes` index (Specific to this WMI).
        3. `ManufacturerMakes` index (All makes for this Manufacturer - *Broad Approach*).
    - For each `MakeID`, resolve to a `Brand` enum using `internal/vpic.ResolveBrand`.

6.  **Model Resolution:**
    - For every resolved `MakeID`, look up associated `ModelID`s using the `MakeModels` index.
    - For each `ModelID`, resolve to a `Model` enum using `internal/vpic.ResolveModel`.

### 3. Output
- The resulting `Manufacturer` message is marshaled to a single line of JSON using `protojson`.
- Output is written to `stdout` or a specified file.

## Helper Functions
We leverage the existing and newly created helpers in `@internal/vpic`:
- `ResolveCountry(int32)`
- `ResolveBrand(int32)`
- `ResolveVehicleType(int32)`
- `ResolveModel(int32)`
