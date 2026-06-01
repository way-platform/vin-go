# Traficom Fuel Tank / Battery Capacity Data

Source: [Finnish Transport and Communications Agency (Traficom) Open Data](https://tieto.traficom.fi/en/open-data#accordion-63420)

## Methodology

1. **Download** the full vehicle register CSV from Traficom open data (all vehicles registered in Finland).
2. **Filter** to N1 class only (light commercial vehicles, GVW <= 3500 kg).
3. **Deduplicate** by brand + model + fuel type, keeping the most common configuration per group based on registration count.
4. **Map** Traficom fields to vin-go enum names:
   - Brand names → `Brand` enum (e.g. "Volkswagen" → `VOLKSWAGEN`)
   - Model names → `Model` enum (e.g. "Transporter" → `TRANSPORTER`)
   - Traficom fuel codes → `FuelType` enum (01→`GASOLINE`, 02→`DIESEL`, 04→`ELECTRIC`, 13/38→`COMPRESSED_NATURAL_GAS`)
5. **Research** OEM fuel tank capacity and battery capacity for each unique configuration using:
   - ultimatespecs.com (primary source for ICE fuel tanks)
   - OEM PDF spec sheets (toyota.fi, citroen.fr, fiat.fi, ford-cdn)
   - Wikipedia (battery capacity for BEV/PHEV models)
6. **Verify** each entry has a fetchable source URL. Remove entries that cannot be independently verified.

## Output

`traficom.csv` — keyed on `brand,model,fuel_type` with columns:
- `fuel_tank_capacity_l` — fuel tank size in liters (0 for BEVs)
- `battery_capacity_kwh` — battery capacity in kWh (0 for pure ICE)
- `source` — URL where the spec was verified (not embedded in generated code)

The decoder returns all fuel type configurations for a brand+model via named fields on the Vehicle proto (`diesel`, `gasoline`, `electric`), allowing downstream consumers to select the matching entry based on their known fuel type.

## Regenerating

```bash
mise run traficom-gen-go
```

This reads `data/traficom/traficom.csv` and generates `traficom.gen.go` at the repo root.
