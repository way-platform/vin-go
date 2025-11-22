# Global Vehicle Identification Taxonomy: Architecture, Regulatory Compliance, and Data Normalization for World Manufacturer Identifiers

## 1. The Historical and Regulatory Genesis of Vehicle Identification

The standardization of the automobile's identity is not merely a bureaucratic convenience but the foundational bedrock of global automotive commerce, safety regulation, and supply chain logistics. To construct a robust, reliable, and exhaustive lookup table for World Manufacturer Identifier (WMI) tags, one must first deconstruct the historical and regulatory layers that have sedimented over the past seven decades to form the current global Vehicle Identification Number (VIN) system. The journey from chaotic, proprietary serial numbers to the structured ISO 3780 standard reveals why modern data normalization is so fraught with complexity.

### 1.1 The Pre-Standardization Era (1954–1980) and the Need for Harmonization

Prior to 1981, the automotive landscape was a fragmented ecosystem of proprietary identification schemas. Manufacturers utilized distinct, non-interoperable serial number formats that often functioned more as internal production control codes than as global identifiers. VINs were first utilized in the United States as early as 1954, yet for nearly three decades, there was no accepted standard.1 During this period, different manufacturers—and indeed, different divisions within the same corporate entity—employed divergent formats. A General Motors vehicle from 1965 might utilize a VIN structure entirely alien to a Ford vehicle of the same vintage, and arguably more problematic for modern data architects, these codes were often little more than sequential serial numbers lacking descriptive attributes.1

The lack of standardization created significant friction for law enforcement, theft recovery agencies, and the nascent insurance telematics industry. The confusion reached a critical mass in the mid-20th century as the global trade of vehicles expanded. The United States government, recognizing the necessity for a unified tracking mechanism for safety recalls and theft prevention, began mandating a 13-character VIN starting in January 1966.1 However, this initial mandate lacked the rigorous structural definitions required for a truly global system. It was not until the early 1980s that the modern 17-character architecture was solidified, driven by the need to uniquely identify vehicles across international borders for a period of 30 years.2

### 1.2 The ISO 3779 and ISO 3780 Frameworks

The modern VIN system is governed principally by two International Organization for Standardization (ISO) protocols: ISO 3779, which dictates the content and structure of the VIN, and ISO 3780, which specifically governs the World Manufacturer Identifier (WMI) code.1 These standards provide the skeleton upon which all national regulations are built.

ISO 3780 defines the WMI as the first section of the VIN, consisting of three characters that designate the manufacturer of the vehicle. The core objective of this standard is ensuring uniqueness. When utilized in conjunction with the remaining sections of the VIN—the Vehicle Descriptor Section (VDS) and the Vehicle Identifier Section (VIS)—the WMI ensures that no two vehicles manufactured within a 30-year window share the same identifier.2

The architectural rigidity of ISO 3780 is strict regarding character sets. To prevent optical character recognition (OCR) errors and human misinterpretation, the Arabic numerals 0–9 and Roman letters A–Z are permitted, with the explicit exclusion of the letters I, O, and Q.2 This exclusion is a critical validation rule for any data ingestion pipeline; any WMI or VIN containing these characters is, by definition, invalid or fraudulent.

### 1.3 The Role of SAE International as the Global Clearinghouse

While ISO provides the theoretical framework, the operational management of WMI allocation is handled by SAE International (formerly the Society of Automotive Engineers). SAE acts under contract with the National Highway Traffic Safety Administration (NHTSA) and ISO to assign selected portions of the VIN, specifically the WMI.4

This centralized coordination is vital for avoiding duplication. SAE J1044, the SAE Recommended Practice, establishes the procedure for the issuance and assignment of WMIs on a uniform basis.5 The standard is intended to be used in conjunction with other SAE reports (J853, J187, J272) to assist agencies in identifying a vehicle's point of origin.5 For data architects, this means that while government databases like vPIC are the primary public source of data, the master registry is technically proprietary to SAE. This distinction influences data acquisition strategies, necessitating reliance on regulatory filings (CFR 49 Part 565 in the US) rather than direct access to the SAE master database.7

### 1.4 Regional Regulatory Divergence

Although ISO 3780 implies a unified global system, regional implementation varies, creating "dialects" of the VIN language that a comprehensive lookup table must translate.

- North America (NHTSA): The United States operates under 49 CFR Part 565, which mandates the Check Digit in the 9th position. This is a mathematical checksum used to validate the entire VIN.8 North American WMIs are strictly enforced, and the definitions of "manufacturer" are tied to safety certification liabilities.

- Europe (EU/KBA): European implementations, such as those overseen by the Kraftfahrt-Bundesamt (KBA) in Germany or NSAI in Ireland, generally follow ISO 3779 but do not mandate the Check Digit in position 9.9 This means a validation algorithm designed for US VINs will reject valid European VINs. Furthermore, European registration documents often separate the manufacturer code (HSN) from the vehicle type (TSN), creating a parallel taxonomy that must be mapped to the WMI.10

- China (GB Standards): China has adopted its own national standards, GB 16735 and GB 16737, which align closely with ISO but impose specific requirements for New Energy Vehicles (NEVs) and joint venture identification.11

- Japan (MLIT): The Japanese domestic market (JDM) frequently utilizes a Chassis Code (Frame Number) system that does not conform to the 17-character ISO VIN standard at all.13 This represents a fundamental schema deviation for any global VIN decoder.

The "manufacturer" in the VIN system is a regulatory construct. It identifies the entity accepting liability for the vehicle's compliance with safety standards. This is why a re-manufacturer or a final-stage manufacturer (like a limousine builder) will assign their own WMI, superseding the WMI of the chassis provider.14

---

## 2. The Anatomy of ISO 3780 and The WMI Structure

To build a reliable lookup table, one must understand the positional logic of the WMI. It is not a random string; it is a hierarchical code.

### 2.1 Position 1: Geographic Region Allocation

The first character of the WMI identifies the broad geographic area of manufacture. This assignment is static and strictly partitioned by SAE to prevent overlap.

Table 1: ISO 3780 Geographic Region Codes

|            |                                |                                                             |        |
| ---------- | ------------------------------ | ----------------------------------------------------------- | ------ |
| Code Range | Geographic Region              | Notable Assignments                                         | Source |
| 1, 4, 5    | North America (USA)            | Divided to accommodate high volume of US manufacturers.     | 1      |
| 2          | North America (Canada)         | Exclusively Canada.                                         | 1      |
| 3          | North America (Mexico/Central) | Mexico (3A-3W), Costa Rica (3X-37), Cayman Islands (38-39). | 1      |
| 6, 7       | Oceania                        | Australia (6), New Zealand (7).                             | 1      |
| 8, 9       | South America                  | Argentina (8A-8E), Brazil (9A-9E, 93-99), Colombia (9F-9K). | 1      |
| J - R      | Asia                           | Japan (J), Korea (K), China (L), India (M), Indonesia (M).  | 14     |
| S - Z      | Europe                         | UK (S), Germany (W), France (V), Italy (Z), Sweden (Y).     | 14     |
| A - H      | Africa                         | South Africa (AA-AH).                                       | 14     |

Insight for Data Modeling: The first character serves as the primary shard key for any global VIN database. It immediately narrows the search space. For instance, any VIN starting with J is legally required to be a Japanese-origin manufacturer entity, though not necessarily manufactured in Japan if it's a transplant factory (though usually, J is reserved for domestic production, while transplant factories use the local region code, e.g., Toyota Kentucky uses 4 17).

### 2.2 Position 2: Country Code Assignment

The second character, in combination with the first, designates the specific country.

- Germany: W (Region) + A-0 (Country Range). Common German codes include W (West Germany historical) and S (East Germany historical, though now unified). W covers the vast majority of German production.14

- UK: S (Region) + A-M (Country Range). Hence SAJ is Jaguar, SAL is Land Rover.14

- Shared Codes: Note that 1, 4, and 5 are all USA. Therefore, the second character in a US VIN does not define a sub-region but rather expands the namespace for manufacturers. 1G (GM), 1F (Ford), 1C (Chrysler) utilize the second digit to denote the major corporate entity, whereas smaller nations might use it to define the country itself (e.g., KL = Korea 15).

### 2.3 Position 3: Manufacturer Identifier

The third character identifies the specific manufacturer. However, the allocation logic here is where simple key-value mapping fails.

- Volume Manufacturers: Large manufacturers are assigned a unique third character. For example, 1G1 is Chevrolet Passenger Cars. 1G defines GM US, and 1 defines the Chevrolet division.14

- Corporate Divisions: The third character often distinguishes vehicle type or division. For General Motors:

- 1G1 = Chevrolet Cars

- 1GC = Chevrolet Trucks

- 1G2 = Pontiac Cars

- 1G6 = Cadillac

- 1G8 = Saturn 18

- Shared Third Characters: In some jurisdictions or smaller production runs, the third character might be rotated or reused over decades, necessitating a check against the Model Year (10th digit of VIN) to ensure accurate decoding.

---

## 3. The Low Volume Manufacturer (LVM) Paradigm

The most critical architectural exception in WMI normalization is the handling of Low Volume Manufacturers (LVM). Failing to implement LVM logic correctly is the single most common failure point in amateur VIN decoders.

### 3.1 The "9" Code Anomaly

According to ISO 3780 and US 49 CFR Part 565.13, manufacturers producing fewer than 1,000 vehicles of a given type per year are assigned a WMI where the third character is always 9.2

- Standard WMI: 3 Characters (Positions 1-3). Uniquely identifies manufacturer.

- LVM WMI: 3 Characters (Positions 1-3) PLUS Positions 12, 13, and 14.

When the third character is 9, the first three characters (e.g., 1A9) do not identify a unique manufacturer. 1A9 simply signifies "A low-volume manufacturer located in the United States." To identify the specific entity, the decoder must look at positions 12, 13, and 14 of the VIS (Vehicle Identifier Section).

### 3.2 LVM Data Structure and Lookups

The specific manufacturer identifier for an LVM is the combination of the WMI (Pos 1-3) and the Manufacturer ID (Pos 12-14).

- Example: A hypothetical small trailer builder might have the WMI 19U. This is shared by hundreds of builders.

- Builder A: 19U......001...

- Builder B: 19U......002...

- Regulatory Burden: The "Replica Vehicle" amendment to 49 CFR Part 566 and 565 in 2022 expanded the LVM program to allow for replica cars (e.g., a new company building 1960s Cobras) to be registered.20 This has led to an influx of new LVM registrations.

Implications for Lookup Table Schema:

Your database cannot use the 3-character WMI as a unique primary key. The schema must support a "Extended WMI" concept.

- Column wmi_prefix: The first 3 characters.

- Column lvm_suffix: The 12th-14th characters (nullable, populated only if wmi_prefix == '9').

- Validation Logic: IF wmi == '9' THEN LOOKUP(wmi_prefix + vin[11:14]) ELSE LOOKUP(wmi_prefix).

### 3.3 High-Value LVMs

This is not just for obscure trailer companies. Many high-prestige hypercar manufacturers started or operate as LVMs.

- Koenigsegg: Early models used LVM coding logic before securing a standard WMI.

- Shelby American: 123 (Pos 12-14).

- Saleen: 19U...
  Missing this logic means your table will fail to identify some of the most expensive assets in the automotive market.

---

## 4. North American Regulatory Framework (NHTSA & vPIC)

For the North American market, the National Highway Traffic Safety Administration (NHTSA) is the supreme regulatory body. Their database, vPIC, is the Rosetta Stone for VIN decoding in this region.

### 4.1 The vPIC Ecosystem

The Product Information Catalog and Vehicle Listing (vPIC) is a consolidated platform presenting data collected from manufacturers under 49 CFR Parts 551–595.7 Manufacturers are legally required to submit their VIN decoding information to NHTSA before selling vehicles in the US.

- Data Availability: NHTSA makes this data available via a public API (vpic.nhtsa.dot.gov/api/) and bulk downloads (CSV and SQL backups).7

- The "565" Submittals: The core data comes from "Part 565" submissions, where manufacturers define their VIN schemas.22 This means the vPIC database is not just a list of cars; it is a list of decoding rules provided by the automakers themselves.

### 4.2 Analyzing the vPIC Data Model

The vPIC database schema is extensive. For WMI normalization, the key tables and columns are:

- WMI Table: Contains the WMI code and links to the Manufacturer table.

- Manufacturer Table: Contains the Mfr_ID, Mfr_Name, and Mfr_CommonName.

- Make Table: Contains the consumer-facing brand names (e.g., "Honda", "Acura").

API Endpoints for Data Gathering:

- GetWMIsForManufacturer: Returns all WMIs assigned to a specific manufacturer ID. Useful for mapping corporate groups.22

- DecodeVin / DecodeVinValues: The "Batch Decode" capability is essential for validating your lookup table against the "ground truth" of the NHTSA engine.22

- WMI Response Example: A query for "Honda" might return JHM (Honda Motor Co Japan), 1HG (Honda of America), 2HG (Honda Canada), etc..23

### 4.3 Error Handling and Validation

The vPIC system utilizes specific error codes that your ingestion engine must recognize 24:

- Code 0: VIN decoded clean. Check Digit correct.

- Code 1: Check Digit failed validation.

- Code 10: Off-road vehicle warning (indicates a non-compliant VIN or a vehicle not under NHTSA jurisdiction).

- Code 400: Invalid characters present (I, O, Q).

Integrating these error codes into your lookup logic allows the system to flag "dirty" data before attempting to map it to a manufacturer.

---

## 5. European and Asian Identification Frameworks

While vPIC covers North America, a global lookup table must integrate data from jurisdictions with different reporting requirements.

### 5.1 Europe: KBA and the HSN/TSN System

In Germany, the Kraftfahrt-Bundesamt (KBA) utilizes a system that runs parallel to the VIN: the Herstellerschlüsselnummer (HSN) and Typschlüsselnummer (TSN).10

- HSN: A 4-digit manufacturer code (e.g., 0005 for BMW).

- WMI Correlation: KBA documentation explicitly links HSNs to WMIs. For example, HSN 0005 correlates to WMIs WBA, WBS (BMW M), and WBM (BMW Motorrad).10

- Relevance: This dataset is invaluable for resolving "Multi-Stage Manufacturers" common in Europe. Companies like Alpina are legally distinct manufacturers in Germany. An Alpina B7 has a specific HSN (7656) and a specific WMI (WAP), distinct from the BMW chassis it is based on. A US-centric decoder might incorrectly identify it as a BMW based on visual similarity, but the WMI WAP correctly points to Alpina Burkard Bovensiepen GmbH.

### 5.2 China: GB 16735 and Joint Ventures

The Chinese market operates under GB 16735-2019 (Vehicle Identification Number) and GB 16737 (WMI Code).11

- The "L" Code: Chinese WMIs predominantly start with L.

- The Joint Venture (JV) Complexity: Due to ownership regulations, foreign brands operate as JVs.

- SAIC Volkswagen: WMI LSV.12

- FAW-Volkswagen: WMI LFV.

- NEV Tracking: The 2019 update to GB 16735 integrated requirements for New Energy Vehicle (NEV) battery tracing, linking the VIN to the battery coding standard GB/T 34014.25 This provides a mechanism to track battery lifecycle from the VIN, a feature unique to the Chinese regulatory environment.

- Data Source: The "Notice on Road Motor Vehicle Manufacturing Enterprises and Products" issued by the Ministry of Industry and Information Technology (MIIT) is the authoritative source for these assignments.

### 5.3 Japan: The Chassis Code (Frame Number) Divergence

Japan presents a unique architectural challenge. Domestic vehicles (JDM) often utilize a Frame Number format (e.g., DAA-ZWR80G-1234567) rather than a 17-char VIN.13

- Structure:

- Emission Code: DAA (indicates compliance with specific emission standards).

- Model Code: ZWR80G (Manufacturer's model designation, e.g., Toyota Voxy).

- Serial: 1234567.

- Export vs. Domestic: Japanese vehicles built for export do receive standard ISO VINs starting with J (e.g., JT1 for Toyota).

- Implication: A global lookup table must have a "Mode Switch." If the input is < 17 characters and follows the AAA-AAAAAA pattern, it must bypass the WMI lookup and query a Japanese Model Code Table instead. This table maps ZWR80G -> Toyota Voxy.

- Recall Data as Source: Recall notices from MLIT (e.g., the Daihatsu Pixis Joy recall) are excellent sources for mapping Model Codes to specific production runs and chassis number ranges.26

---

## 6. Corporate Taxonomy and WMI Allocation: The "Big Three" (US)

We now move to the specific WMI mappings required for the lookup table. We begin with the US "Big Three," whose corporate lineages require careful handling of defunct brands.

### 6.1 General Motors (GM)

GM utilizes the third character of the WMI to designate division, a practice that creates a clean but extensive list of codes.

Table 2: General Motors WMI Allocation

|     |                   |               |         |                                        |
| --- | ----------------- | ------------- | ------- | -------------------------------------- |
| WMI | Division / Brand  | Vehicle Type  | Region  | Notes                                  |
| 1G1 | Chevrolet         | Passenger Car | USA     |                                        |
| 1GC | Chevrolet         | Truck         | USA     |                                        |
| 1G2 | Pontiac           | Passenger Car | USA     | Defunct Brand (Map to GM Parent)       |
| 1G3 | Oldsmobile        | Passenger Car | USA     | Defunct Brand                          |
| 1G4 | Buick             | Passenger Car | USA     |                                        |
| 1G6 | Cadillac          | Passenger Car | USA     |                                        |
| 1G8 | Saturn            | Passenger Car | USA     | Defunct Brand                          |
| 1GM | Pontiac / Holden  | Passenger Car | USA/Aus | Shared allocation for specific imports |
| 2G1 | Chevrolet         | Passenger Car | Canada  |                                        |
| 3G1 | Chevrolet         | Passenger Car | Mexico  |                                        |
| KL1 | GM Korea (Daewoo) | Car           | Korea   | Formerly Daewoo, now Chevy             |
| LSG | SAIC-GM           | Car           | China   | Joint Venture                          |

Taxonomy Note: Pontiac, Oldsmobile, and Saturn are legally "General Motors" entities. The lookup table should maintain them as distinct Brands (brand_id) but link them to the GENERAL_MOTORS Parent Group (group_id). This preserves historical accuracy while allowing group-level aggregation.18

### 6.2 Ford Motor Company

Ford's WMI structure is relatively stable, with less division-level granularity in the WMI compared to GM.

Table 3: Ford Motor Company WMI Allocation

|     |                 |         |                                  |
| --- | --------------- | ------- | -------------------------------- |
| WMI | Entity          | Region  | Notes                            |
| 1FA | Ford Motor Co   | USA     | Ford Passenger Car               |
| 1FT | Ford Motor Co   | USA     | Ford Truck                       |
| 1FM | Ford Motor Co   | USA     | Ford MPV/SUV                     |
| 1L1 | Lincoln         | USA     | Lincoln Passenger Car            |
| 2FA | Ford Canada     | Canada  |                                  |
| 3FA | Ford Mexico     | Mexico  |                                  |
| NM0 | Ford Turkey     | Turkey  | Ford Otosan (Transit production) |
| WFO | Ford Werke GmbH | Germany | European Ford models             |
| SFA | Ford UK         | UK      | Legacy / Specialized             |

### 6.3 Stellantis (The Merger Complexity)

Stellantis represents the ultimate data normalization challenge. It is an aggregation of FCA (Fiat Chrysler) and PSA (Peugeot Citroen). The lookup table must map legacy codes from three nations (US, Italy, France) to a single parent.

Table 4: Stellantis WMI Mapping

|     |                |                |              |
| --- | -------------- | -------------- | ------------ |
| WMI | Legacy Entity  | Current Brand  | Parent Group |
| 1C3 | Chrysler Corp  | Chrysler       | STELLANTIS   |
| 1C4 | Chrysler Corp  | Chrysler (MPV) | STELLANTIS   |
| 1D3 | Dodge          | Dodge          | STELLANTIS   |
| 1J4 | Jeep           | Jeep           | STELLANTIS   |
| 1C6 | RAM            | RAM            | STELLANTIS   |
| ZFA | Fiat SpA       | Fiat           | STELLANTIS   |
| ZAR | Alfa Romeo     | Alfa Romeo     | STELLANTIS   |
| ZAM | Maserati       | Maserati       | STELLANTIS   |
| VF3 | Peugeot        | Peugeot        | STELLANTIS   |
| VF7 | Citroën        | Citroën        | STELLANTIS   |
| VR1 | DS Automobiles | DS             | STELLANTIS   |
| W0L | Adam Opel AG   | Opel           | STELLANTIS   |

Critical Insight: The WMI W0L (Opel) was historically owned by General Motors. In 2017, it was sold to PSA, and subsequently became Stellantis. The WMI did not change. A VIN decoder must simply map W0L to OPEL. The group_id relationship is what changed over time. A sophisticated database might temporally version this relationship (e.g., "If Model Year < 2017, Parent=GM; Else Parent=Stellantis").

---

## 7. Corporate Taxonomy and WMI Allocation: European Conglomerates

European manufacturers often distinguish between "Passenger Car" and "Commercial/Utility" divisions within the WMI, and rely heavily on regional manufacturing codes.

### 7.1 The Volkswagen Group (VAG)

VAG is a master of platform sharing, yet maintains distinct WMIs for its brands, which simplifies decoding compared to US "Badge Engineering."

Table 5: Volkswagen Group WMI Allocation

|     |             |                        |                            |
| --- | ----------- | ---------------------- | -------------------------- |
| WMI | Brand       | Entity                 | Region                     |
| WVW | Volkswagen  | VW AG (Cars)           | Germany                    |
| WV1 | Volkswagen  | VW Commercial          | Germany                    |
| WV2 | Volkswagen  | VW Bus/Van             | Germany                    |
| WAU | Audi        | Audi AG                | Germany                    |
| TRU | Audi        | Audi Hungaria          | Hungary (TT/Q3 production) |
| WUA | Audi Sport  | Quattro GmbH           | Germany (RS Models)        |
| WP0 | Porsche     | Porsche AG             | Germany (Sports Cars)      |
| WP1 | Porsche     | Porsche AG             | Germany (SUVs - Cayenne)   |
| TMB | Škoda       | Škoda Auto             | Czech Republic             |
| VSS | SEAT        | SEAT S.A.              | Spain                      |
| ZHW | Lamborghini | Automobili Lamborghini | Italy                      |
| W1V | Bugatti     | Bugatti Rimac          | Germany/France             |

Note on Bugatti: Historical Bugatti (EB110 era) used ZA9. Modern Bugatti (Veyron/Chiron) uses VF9 (France) or W1V. The newest entity, Bugatti-Rimac, introduces new complexity that may yield new codes.9

### 7.2 BMW Group

BMW separates its US manufacturing (Spartanburg) distinctly from Munich, a critical distinction for trade compliance and tariff calculation.

- WBA: BMW AG (Munich - Core Passenger Cars).

- WBS: BMW M GmbH (Motorsport vehicles - M3, M5). Note: M-Performance models like M340i often use WBA; WBS is reserved for full "M" cars.

- 4US: BMW Manufacturing Co (Spartanburg, SC - USA).

- 5UX: BMW USA (SUVs like X3, X5, X7).

- 5YM: BMW USA (M-Class SUVs like X5M).

- WMW: MINI (Germany/UK).

- SCA: Rolls-Royce Motor Cars (Goodwood, UK).

### 7.3 Mercedes-Benz Group (Daimler)

- WDB: Mercedes-Benz (Legacy/Standard).

- WDD: Mercedes-Benz (Daimler AG - Modern passenger cars).

- WDF: Mercedes-Benz (Commercial/Trucks).

- 4JG: Mercedes-Benz USA (Tuscaloosa - GLE/GLS production).

- W1K: Mercedes-Benz (A-Class/Compact).

- VSA: Mercedes-Benz Spain (V-Class vans).

---

## 8. Corporate Taxonomy and WMI Allocation: Asian Giants

Asian manufacturers exhibit a strong bifurcation between "Domestic" and "Transplant" manufacturing codes.

### 8.1 Toyota Group

Toyota is the world's largest automaker and has a massive WMI footprint.

Table 6: Toyota Group WMI Allocation

|          |        |              |                               |
| -------- | ------ | ------------ | ----------------------------- |
| WMI      | Brand  | Region       | Notes                         |
| JT1, JT2 | Toyota | Japan        | Cars, MPVs                    |
| JT5      | Toyota | Japan        | Trucks, SUVs                  |
| JTH      | Lexus  | Japan        | Lexus Passenger Cars          |
| JTJ      | Lexus  | Japan        | Lexus SUVs                    |
| 4T1      | Toyota | USA          | KY Production (Camry)         |
| 4T3      | Lexus  | USA          | KY/TX Production (RX/ES)      |
| 5TD      | Toyota | USA          | TX Production (Tundra/Tacoma) |
| 2T1      | Toyota | Canada       | RAV4 / Corolla                |
| MR0      | Toyota | Thailand     | Hilux / Fortuner              |
| AHT      | Toyota | South Africa | Hilux                         |

### 8.2 Hyundai Motor Group

Hyundai and Kia share platforms and engines but maintain strictly separate WMI identities.

- Hyundai:

- KMH: Hyundai Motor Co (Korea - Car)

- KM8: Hyundai Motor Co (Korea - SUV)

- 5NP: Hyundai Motor Mfg Alabama (USA)

- MAL: Hyundai India

- TMA: Hyundai Czech Republic

- Kia:

- KNA: Kia Motors Corp (Korea - Car)

- KND: Kia Motors Corp (Korea - MPV/SUV)

- 5XX: Kia Motors Mfg Georgia (USA)

- U5Y: Kia Slovakia

### 8.3 Chinese New Energy Vehicles (NEVs)

The rise of Chinese EV brands introduces new codes that are essential for a modern lookup table.

- BYD: LGX (BYD Auto).

- NIO: LNZ.

- Tesla China: LRW (Gigafactory Shanghai).

- Polestar: LPS (Polestar is owned by Geely/Volvo but has Chinese production codes).

---

## 9. Data Engineering and Lookup Table Architecture

To move from lists to a functional system, we must architect the data model. This section addresses the "build a complete lookup table ourselves" requirement.

### 9.1 Schema Design for Relational Mapping

A flat CSV is insufficient. The schema must support the one-to-many relationships inherent in the taxonomy.

Table A: wmi_master (The Entry Point)

- wmi_code (PK, CHAR(3)): The 3-char lookup key.

- manufacturer_id (FK): Links to the manufacturers table.

- region_id (FK): Links to geographic region.

- is_lvm (BOOLEAN): Trigger for LVM logic.

- active (BOOLEAN): Status flag.

Table B: lvm_registry (The "9" Code Handler)

- id (PK): UUID.

- wmi_prefix (CHAR(3)): e.g., '19U'.

- serial_range_start (CHAR(3)): Pos 12-14 start.

- serial_range_end (CHAR(3)): Pos 12-14 end.

- manufacturer_id (FK): The specific small manufacturer.

Table C: manufacturers (The Brand Layer)

- manufacturer_id (PK): UUID.

- name: Brand Name (e.g., "Chevrolet").

- parent_group_id (FK): Links to corporate_groups.

Table D: corporate_groups (The Financial Layer)

- group_id (PK): UUID.

- name: e.g., "General Motors", "Stellantis".

- market_cap_tier: Useful for analytics (Large Cap vs Niche).

### 9.2 ETL and Ingestion Strategy

The data pipeline should be multi-sourced to ensure completeness.

1. Ingest vPIC: Use the NHTSA vPIC SQL backup (vPICList_lite.bak) as the base layer.22 Map vPIC.WMI to wmi_master.

2. Enrich with KBA: Parse the German KBA list for European manufacturers not in vPIC.10 Use HSN codes to verify manufacturer names.

3. Apply LVM Logic: Write a specific transformation script that isolates all WMIs where char == '9'. Move these to the lvm_registry table, populating the specific manufacturer names from the "Part 565" submissions in vPIC.

4. Sanitization: Run a filter to remove any entries containing I, O, or Q.

### 9.3 Handling "Zombie" Data

Historical data is messy. 1G2 (Pontiac) is valid for a 2009 vehicle but invalid for a 2025 vehicle.

- Recommendation: Do not delete defunct WMIs. Mark them active=FALSE or legacy=TRUE. The lookup table must return "Pontiac" for a 1G2 query, even if the brand no longer exists. The parent_group should remain General Motors.

---

## 10. Implementation Case Studies and Open Source Analysis

Several open-source projects have attempted this, offering lessons on what to do (and what to avoid).

### 10.1 Case Study: php-vin-decoder vs. node-vin-lite

- php-vin-decoder: Often relies on static arrays of regex patterns.28 This is brittle. If a manufacturer changes their WMI allocation (like Tesla moving from 5YJ to adding 7G2), the code breaks.

- node-vin-lite: Uses a partial WMI list (approx 1,000 codes).29 This is insufficient for a commercial-grade application. The SAE database contains over 33,000 codes.

- Lesson: Hard-coding is failure. The system must be database-driven with an update capability (e.g., nightly sync with vPIC API).

### 10.2 Case Study: vininfo (Python)

The vininfo library allows for standalone execution and checksum verification.30

- Strength: It includes specific dictionaries for identifying brands like Nissan or Opel based on the WMI.

- Weakness: It requires manual updates to the dicts/wmi.py file to add new manufacturers.31

- Takeaway: Your implementation should decouple the code from the data. The WMI list should be a JSON or SQL file loaded at runtime, not hardcoded into the application logic.

### 10.3 The "Corgi" Library (TypeScript)

The corgi library 32 uses a compressed SQLite database of the NHTSA vPIC dataset.

- Architecture: It bundles the database with the library (~20MB). This provides offline capability and zero latency.

- Relevance: For high-performance applications (like scanning VINs on a mobile device), an offline SQLite approach is superior to a REST API call. Your lookup table architecture should support exporting a "Lite" version (SQLite/JSON) for edge deployment.

## 11. Conclusion

Constructing a "reliable and complete" WMI lookup table is a challenge of architectural taxonomy, not merely data entry. It requires a system that can simultaneously handle the rigid structure of ISO 3780, the flexible "9-code" logic of Low Volume Manufacturers, and the evolving corporate hierarchies of the global automotive industry.

By synthesizing the raw data from NHTSA vPIC with the regional specificity of KBA and GB standards, and organizing it into a relational model that separates Entity, Brand, and Group, one can build a definitive source of truth. This system will correctly identify a ZAR VIN as an Alfa Romeo (Brand) owned by Stellantis (Group), a 19U... VIN as a unique small trailer builder, and a J VIN as a Japanese export, providing the nuanced insight required for modern automotive data applications.

#### Works cited

1. Vehicle identification number - Wikipedia, accessed November 21, 2025, [https://en.wikipedia.org/wiki/Vehicle_identification_number](https://en.wikipedia.org/wiki/Vehicle_identification_number)

2. World Manufacturer Identifier | NSAI - National Standards Authority of Ireland, accessed November 21, 2025, [https://www.nsai.ie/certification/automotive/transport-schemes/world-manufacturer-identifier/](https://www.nsai.ie/certification/automotive/transport-schemes/world-manufacturer-identifier/)

3. ISO 3780 - iTeh Standards, accessed November 21, 2025, [https://cdn.standards.iteh.ai/samples/45844/8aa6bf9bd1ee4463aafb4135b5e74cfb/ISO-3780-2009.pdf](https://cdn.standards.iteh.ai/samples/45844/8aa6bf9bd1ee4463aafb4135b5e74cfb/ISO-3780-2009.pdf)

4. World Manufacturer Codes/Product Identification Numbers (WMC/PIN) Program Page | SAE Industry Technologies Consortia (SAE ITC), accessed November 21, 2025, [https://www.sae-itc.com/programs/wmc-pin](https://www.sae-itc.com/programs/wmc-pin)

5. J1044_202501 : World Manufacturer Identifier - SAE International, accessed November 21, 2025, [https://www.sae.org/standards/j1044_202501-world-manufacturer-identifier](https://www.sae.org/standards/j1044_202501-world-manufacturer-identifier)

6. J1044_202403 : World Manufacturer Identifier - SAE International, accessed November 21, 2025, [https://www.sae.org/standards/j1044_202403-world-manufacturer-identifier](https://www.sae.org/standards/j1044_202403-world-manufacturer-identifier)

7. NHTSA Product Information Catalog and Vehicle Listing - Home, accessed November 21, 2025, [https://vpic.nhtsa.dot.gov/](https://vpic.nhtsa.dot.gov/)

8. What's in a VIN? How to decode the vehicle identification number, your car's unique fingerprint | Clemson News, accessed November 21, 2025, [https://news.clemson.edu/whats-in-a-vin-how-to-decode-the-vehicle-identification-number-your-cars-unique-fingerprint/](https://news.clemson.edu/whats-in-a-vin-how-to-decode-the-vehicle-identification-number-your-cars-unique-fingerprint/)

9. VW World Manufacturer Identifier VW | PDF | Volkswagen | Car Body Styles - Scribd, accessed November 21, 2025, [https://www.scribd.com/document/415181124/VW-World-Manufacturer-Identifier-VW](https://www.scribd.com/document/415181124/VW-World-Manufacturer-Identifier-VW)

10. Verzeichnis der Hersteller List of manufacturers - Kraftfahrt-Bundesamt, accessed November 21, 2025, [https://www.kba.de/SharedDocs/Downloads/EN/SV/sv32_pdf_en.pdf?\_\_blob=publicationFile&v=2](https://www.kba.de/SharedDocs/Downloads/EN/SV/sv32_pdf_en.pdf?__blob=publicationFile&v=2)

11. China - National Standard of the P.R.C., Road Vehicles – Wor - TUV Rheinland, accessed November 21, 2025, [https://www.tuv.com/regulations-and-standards/en/china-national-standard-of-the-p-r-c-road-vehicles-world-manufacturer-identifier-wmi-code.html](https://www.tuv.com/regulations-and-standards/en/china-national-standard-of-the-p-r-c-road-vehicles-world-manufacturer-identifier-wmi-code.html)

12. GB 16735-2019 PDF English, accessed November 21, 2025, [https://www.chinesestandard.net/PDF.aspx/GB16735-2019](https://www.chinesestandard.net/PDF.aspx/GB16735-2019)

13. About Vehicle Model Codes - Provide Cars, accessed November 21, 2025, [https://providecars.co.jp/japan-used-car-guide-about-vehicle-model-codes/](https://providecars.co.jp/japan-used-car-guide-about-vehicle-model-codes/)

14. Vehicle Identification Numbers (VIN codes)/World Manufacturer Identifier (WMI) - Wikibooks, accessed November 21, 2025, [https://en.wikibooks.org/wiki/Vehicle*Identification_Numbers*(VIN*codes)/World_Manufacturer_Identifier*(WMI)](<https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/World_Manufacturer_Identifier_(WMI)>)

15. VIN Country Codes of Vehicles | EpicVIN Blog, accessed November 21, 2025, [https://epicvin.com/blog/vin-country-codes-of-the-vehicles](https://epicvin.com/blog/vin-country-codes-of-the-vehicles)

16. GEAR.gr, accessed November 21, 2025, [http://www.ship.gr/gear/vin.html](http://www.ship.gr/gear/vin.html)

17. Every Car Brand Toyota Owns In 2025 - SlashGear, accessed November 21, 2025, [https://www.slashgear.com/1944566/every-toyota-owned-car-brand/](https://www.slashgear.com/1944566/every-toyota-owned-car-brand/)

18. Automotive Parent Companies - North American Lubricants, accessed November 21, 2025, [https://nalube.com/automotive-parent-companies/](https://nalube.com/automotive-parent-companies/)

19. new manufacturers handbook - NHTSA's vPIC, accessed November 21, 2025, [https://vpic.nhtsa.dot.gov/ManufacturerHandbook.pdf?d=251025](https://vpic.nhtsa.dot.gov/ManufacturerHandbook.pdf?d=251025)

20. Vehicle Identification Number (VIN) Requirements; Manufacturer Identification; Certification; Replica Motor Vehicles; Importation of Vehicles and Equipment Subject to Federal Safety, Bumper, and Theft Prevention Standards - Federal Register, accessed November 21, 2025, [https://www.federalregister.gov/documents/2022/03/09/2022-04030/vehicle-identification-number-vin-requirements-manufacturer-identification-certification-replica](https://www.federalregister.gov/documents/2022/03/09/2022-04030/vehicle-identification-number-vin-requirements-manufacturer-identification-certification-replica)

21. NHTSA Product Information Catalog and Vehicle Listing (vPIC) - MID, accessed November 21, 2025, [https://catalog.data.gov/dataset/nhtsa-product-information-catalog-and-vehicle-listing-vpic-mid](https://catalog.data.gov/dataset/nhtsa-product-information-catalog-and-vehicle-listing-vpic-mid)

22. Vehicle API - NHTSA's vPIC - Department of Transportation, accessed November 21, 2025, [https://vpic.nhtsa.dot.gov/api/](https://vpic.nhtsa.dot.gov/api/)

23. https://vpic.nhtsa.dot.gov/api/vehicles/GetWMIsForManufacturer/hon?format=xml, accessed November 21, 2025, [https://vpic.nhtsa.dot.gov/api/vehicles/GetWMIsForManufacturer/hon?format=xml](https://vpic.nhtsa.dot.gov/api/vehicles/GetWMIsForManufacturer/hon?format=xml)

24. Product Information Catalog and Vehicle Listing (vPIC) Analytical User's Manual 2020 - CrashStats - NHTSA, accessed November 21, 2025, [https://crashstats.nhtsa.dot.gov/Api/Public/Publication/813252](https://crashstats.nhtsa.dot.gov/Api/Public/Publication/813252)

25. Key technology and application analysis of quick coding for recovery of retired energy vehicle battery - Sci-Hub, accessed November 21, 2025, [https://2024.sci-hub.se/8301/1c6fc82225205bce9d6ed34be16d18f4/yu2021.pdf](https://2024.sci-hub.se/8301/1c6fc82225205bce9d6ed34be16d18f4/yu2021.pdf)

26. Recall Notification | Corporate | Global Newsroom | Toyota Motor Corporation Official Global Website, accessed November 21, 2025, [https://global.toyota/en/newsroom/corporate/40353033.html](https://global.toyota/en/newsroom/corporate/40353033.html)

27. The 15 Corporations That Create Most Cars: A Family Tree of Automotive Makers, accessed November 21, 2025, [https://alansfactoryoutlet.com/infographics/the-15-corporations-that-create-most-cars-a-family-tree-of-automotive-makers/](https://alansfactoryoutlet.com/infographics/the-15-corporations-that-create-most-cars-a-family-tree-of-automotive-makers/)

28. TNT VIN Decoder, accessed November 21, 2025, [https://tntautoparts.org/index.php/vin-decoder](https://tntautoparts.org/index.php/vin-decoder)

29. ApelSYN/node-vin-lite: VIN (Vehicle Identification Number) Checker Lite - GitHub, accessed November 21, 2025, [https://github.com/ApelSYN/node-vin-lite](https://github.com/ApelSYN/node-vin-lite)

30. vininfo - PyPI, accessed November 21, 2025, [https://pypi.org/project/vininfo/](https://pypi.org/project/vininfo/)

31. accessed January 1, 1970, [https://github.com/idlesign/vininfo/blob/master/vininfo/dicts/wmi.py](https://github.com/idlesign/vininfo/blob/master/vininfo/dicts/wmi.py)

32. cardog-ai/corgi: A TypeScript library for decoding and validating Vehicle Identification Numbers (VINs) using a customized VPIC database. - GitHub, accessed November 21, 2025, [https://github.com/cardog-ai/corgi](https://github.com/cardog-ai/corgi)
