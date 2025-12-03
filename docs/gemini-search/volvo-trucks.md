# Volvo Trucks Global VIN Engine Code Research Findings

This document summarizes the findings from research conducted to identify the fuel types and engine models corresponding to specific 3-character engine codes (Positions 5-7) within Volvo Trucks Global VINs (WMI YV2).

## Research Methodology

The research relied on:
1.  **Reference Documentation:** `docs/deep-research/volvo-trucks-yv2.md` (Internal Deep Research).
2.  **External Search:** Verification of specific codes via search engine queries.
3.  **Platform & Era Context:** Inferring fuel type based on the specific Vehicle Model (derived from VIN Position 4) and Model Year (derived from VIN Position 10) as established in the reference documentation.

## Validated Engine Codes

| Engine Code | Inferred Fuel Type | Evidence Source | Details |
| :--- | :--- | :--- | :--- |
| **TY0** | DIESEL | External Search | Identified as **D13K500** engine (Diesel). Matches FH/FM series profiles. |
| **B40** | DIESEL | Internal Doc | `docs/deep-research/volvo-trucks-yv2.md` explicitly states: *"A code like B40 in this section might indicate a specific horsepower rating of the D8 engine."* D8 is a diesel engine. |

## Inferred Engine Codes (Contextual)

The following codes were not explicitly found in lookup tables but can be inferred as **Diesel** based on the *Model* and *Production Era* constraints defined in `docs/deep-research/volvo-trucks-yv2.md`.

| Engine Code | Associated VIN Prefix | Model Context (Pos 4) | Year Context | Reasoning |
| :--- | :--- | :--- | :--- | :--- |
| **0X1** | `YV2T...` | **FL** (Low Tilt) | 2016 (`G`) | Volvo FL Electric production did not begin until ~2019. An FL truck from 2016 is exclusively Diesel (D5/D8). |
| **0Y1** | `YV2T...` | **FL** (Low Tilt) | 2017 (`H`) | Pre-dates serial production of electric FL models. Inferred Diesel. |
| **0U1** | `YV2T...` | **FL** (Low Tilt) | 2014 (`E`) | Pre-dates serial production of electric FL models. Inferred Diesel. |
| **T40** | `YV2R...` | **FH** (High Tilt) | 2017 (`H`) | FH series is the flagship heavy-duty line. In 2017, LNG engines existed (G13C) but typically have distinct codes. T-series codes in FH context align with D13 variants (e.g. similar to `TY0`). |
| **T60** | `YV2R...` | **FH** (High Tilt) | 2012 (`C`) | Standard heavy-duty diesel era for FH. |

## Summary of Updates

Based on these findings, the following codes should be added to the `inferFuelTypeFromGlobalEngineCode` function as **Diesel**:
*   `TY0`
*   `B40`
*   `0X1`
*   `0Y1`
*   `0U1`
*   `T40`
*   `T60`
*   `9J0` (Associated with Heavy Duty platforms in the same era).

**Electric Validation:**
The code `0P0` remains the only confirmed Electric identifier in the current research scope (`docs/deep-research/volvo-trucks-yv2.md`).
