# Forensic Analysis of Scania Vehicle Identification Architectures: Decoding Protocols, Modular Engineering Inference, and Global Manufacturing Traceability

## 1. Introduction to Commercial Vehicle Identification Standards

The global logistics and heavy transport sector operates on a foundation of precision engineering and rigorous identification standards. At the core of this operational framework lies the Vehicle Identification Number (VIN), a unique 17-character alphanumeric sequence that serves as the definitive fingerprint for every motor vehicle produced worldwide. For Scania AB, a premier Swedish manufacturer of heavy trucks, buses, and industrial engines, the VIN is not merely a regulatory requirement; it is a complex cryptographic index that encodes the company’s sophisticated modular production philosophy. Unlike mass-market passenger vehicle manufacturers that produce fixed trim levels, Scania utilizes a highly granular modular system where vehicles are tailored to specific transport tasks through a vast library of interchangeable components. Consequently, the Scania VIN functions as a primary data key, unlocking the specific "Individual Chassis Specification" (ICS) necessary for maintenance, regulatory compliance, and lifecycle management.

This report provides an exhaustive technical examination of the Scania VIN architecture. It dissects the three principal sections of the code—the World Manufacturer Identifier (WMI), the Vehicle Descriptor Section (VDS), and the Vehicle Identifier Section (VIS)—to provide a comprehensive decoding guide for industry professionals. Furthermore, it explores the nuances of inferring commercial model designations from raw chassis data, distinguishing between vehicle generations such as the PGRT range and the Next Generation (NTG) models, and analyzing the geopolitical manufacturing footprint encoded within the string. By synthesizing data from regulatory standards, manufacturer bodybuilder instructions, and forensic decoding databases, this analysis offers a definitive reference for understanding the digital identity of Scania vehicles.

### 1.1 The Regulatory Framework: ISO 3779 and Global Harmonization

The structure of the modern VIN is governed by the International Organization for Standardization (ISO) under standard 3779, with additional guidelines provided by ISO 4030 regarding the location and attachment of the VIN plate.1 This standardization was adopted in 1981 to replace the fragmented and proprietary numbering systems used by manufacturers between 1954 and 1980.2 The primary objective of ISO 3779 is to ensure that every vehicle manufactured globally can be uniquely identified for a period of 30 years, facilitating theft recovery, safety recalls, and international commerce.2

For heavy-duty truck manufacturers like Scania, adherence to ISO 3779 presents unique challenges. The standard was originally designed with a bias toward fixed-model passenger cars (e.g., a "Ford Mustang"). However, a Scania truck is defined by its duty class (e.g., long-haulage vs. construction), chassis adaptation, and axle configuration, rather than a static model name. Therefore, Scania utilizes the permissible flexibility within the Vehicle Descriptor Section (VDS) to encode the architectural attributes of the chassis.1

The VIN is segmented into three distinct functional blocks:

1. World Manufacturer Identifier (WMI): Characters 1–3, designating the manufacturer and country of origin.1

2. Vehicle Descriptor Section (VDS): Characters 4–9, describing the vehicle's general attributes such as model series, engine type, and chassis configuration.5

3. Vehicle Identifier Section (VIS): Characters 10–17, providing specific production information including model year, plant of manufacture, and the sequential serial number.5

This report will systematically analyze each of these sections in the context of Scania’s engineering lexicon.

---

## 2. The World Manufacturer Identifier (WMI): Geopolitical and Corporate Origins

The first three characters of the VIN constitute the World Manufacturer Identifier (WMI). This section acts as the highest level of classification, assigning the vehicle to a specific manufacturer and a specific country of final assembly. For a global conglomerate like Scania, which operates production facilities across Europe, South America, and Asia, the WMI provides critical insight into the geopolitical origin of the chassis and the supply chain network involved in its construction.

### 2.1 The Logic of Geographic Allocation

The WMI system is hierarchical. The first character identifies the geographic region, the second identifies the country within that region, and the third identifies the specific manufacturer.5 Scania’s manufacturing footprint is primarily concentrated in Europe and South America, which is reflected in the dominance of specific regional codes in its VIN population.

- Region Code 'S–Z' (Europe): The letter range S through Z is assigned to Europe. Within this range, the specific code 'Y' is assigned to Sweden.7 This forms the basis of the most common Scania WMI, reflecting the company's headquarters and primary production base in Södertälje.

- Region Code '9' (South America): The numeric code '9' is assigned to South America. Within this region, specific ranges are allocated to Brazil, Scania's second-largest manufacturing hub.3

- Region Code '3' (North America/Mexico): The numeric code '3' is assigned to Mexico, where Scania has established assembly operations for specialized markets.7

### 2.2 Primary Scania WMI Codes and Corporate Divisions

Research identifies several distinct WMI codes associated with Scania’s global operations. These codes not only indicate the country of origin but often differentiate between vehicle types (e.g., trucks vs. buses) or specific subsidiaries.

#### Table 1: Detailed Analysis of Scania World Manufacturer Identifiers (WMI)

|          |                   |             |                            |                                                                                                                                                                             |
| -------- | ----------------- | ----------- | -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| WMI Code | Geographic Region | Country     | Manufacturer / Division    | Operational Context                                                                                                                                                         |
| YS2      | Europe            | Sweden      | Scania CV AB (Södertälje)  | The primary code for Scania trucks and commercial vehicles assembled in Sweden. This is the most ubiquitous WMI found in global markets.7                                   |
| YS4      | Europe            | Sweden      | Scania Bus (Katrineholm)   | Historically linked to Scania’s bus and bus chassis production in Katrineholm. While production has consolidated, this code appears on older bus chassis units (pre-2002).7 |
| 9BS      | South America     | Brazil      | Scania Latin America Ltda. | Designated for trucks and buses manufactured at the São Bernardo do Campo facility. Brazil serves as a major export hub for Latin America and occasionally other markets.8  |
| 3AX      | North America     | Mexico      | Scania Mexico              | Assigned to Scania trucks assembled in Mexico.7                                                                                                                             |
| 3BE      | North America     | Mexico      | Scania Mexico              | Assigned specifically to Scania buses assembled in Mexico.7                                                                                                                 |
| XLER     | Europe            | Netherlands | Scania Production Zwolle   | Found in VIN databases linked to trucks assembled at the Zwolle plant, utilizing the Dutch country code 'X'.10                                                              |
| XLEP     | Europe            | Netherlands | Scania Production Meppel   | Linked to component manufacturing or specialized assembly in Meppel, Netherlands.10                                                                                         |

### 2.3 Manufacturing Logistics and Supply Chain Implications

The variation in WMI codes reflects Scania’s "Global Production System." While Södertälje (Sweden) is the heart of the operation, the Zwolle (Netherlands) plant is actually the largest assembly facility for Scania in the world.9 Vehicles assembled in Zwolle may carry a Swedish WMI (YS2) in many instances to maintain brand consistency, as the "Manufacturer of Record" remains Scania CV AB in Sweden. However, the existence of XLER and XLEP codes indicates that for certain regulatory or market-specific registrations, the Dutch origin is explicitly encoded.10

Similarly, the 9BS code for Brazil is significant for supply chain analysts. Scania Latin America Ltda. operates with a high degree of localization. A vehicle with a 9BS WMI likely contains engine blocks or chassis components cast and machined in Brazil (São Bernardo do Campo) rather than imported from Sweden, although Scania maintains rigid global quality standards to ensure parts are functionally identical.9 This distinction is vital for forensics and parts ordering, as legacy vehicles from Brazilian production may have utilized different sub-suppliers for non-critical components compared to their European counterparts.

### 2.4 Low Volume and Incomplete Vehicles

In accordance with ISO 3779, manufacturers producing fewer than 500 or 1,000 vehicles per year (depending on the jurisdiction) must use the digit '9' in the third position of the WMI.4 When this occurs, the 12th, 13th, and 14th characters of the VIN (located in the VIS) are used to uniquely identify the specific manufacturer.

While Scania is a high-volume manufacturer, this "Low Volume" logic allows for the identification of specialized converters or bodybuilders who may purchase a Scania chassis and re-manufacture it significantly enough to issue a new VIN. However, for standard Scania trucks (P, G, R, S series), the WMI will almost invariably be one of the high-volume codes listed in Table 1 (e.g., YS2, 9BS).

---

## 3. The Vehicle Descriptor Section (VDS): Decoding the Engineering DNA

The Vehicle Descriptor Section (VDS), occupying positions 4 through 9 of the VIN, is the most technically dense and architecturally significant portion of the identifier. For a modular manufacturer like Scania, the VDS does not simply point to a marketing model name (like "Civic" or "Corolla"). Instead, it encodes the fundamental engineering architecture of the vehicle, including the cab series, the chassis adaptation, the wheel configuration, and the engine family. Decrypting this section provides the "genetic code" of the truck.

### 3.1 Position 4: The Model Series (Cab Architecture)

The fourth character of the VIN is the primary indicator of the vehicle's series. In Scania's nomenclature, the "Series" is defined principally by the cab mounting height and the engine tunnel configuration. This character allows for the immediate differentiation between urban distribution trucks, long-haul tractors, and heavy construction vehicles.

#### 3.1.1 Analysis of Truck Series Codes

- P-Series (Code 'P'): The P-series features a low forward-control cab. The cab is mounted lower on the chassis to facilitate frequent entry and exit, making it ideal for regional distribution, construction, and municipal services. The engine tunnel is prominent inside the cab due to the low floor height.4

- G-Series (Code 'G'): The G-series represents a medium forward-control cab. It offers a balance between the compactness of the P-series and the spaciousness of the R-series. It is widely used in construction and regional haulage where a sleeper cab is required but extreme height is a disadvantage.4

- R-Series (Code 'R'): Historically the flagship of the Scania range, the R-series features a high forward-control cab. Prior to the Next Generation, the R-series had a small engine tunnel. It is designed for long-haul transport.4

- S-Series (Code 'S'): Introduced with the Next Generation (NTG) in 2016, the S-series is the new flagship. Its defining feature is a completely flat floor with no engine tunnel. The presence of the 'S' code in Position 4 of the VIN is the definitive marker of a top-tier NTG vehicle.11

- L-Series (Code 'L'): A specialized low-entry cab introduced in 2017 for urban environments. It features a kneeling function and a cab that sits ahead of the front axle in a low position, often used for refuse collection.11

- T-Series (Code 'T'): The "Torpedo" or bonneted cab. Production of the T-series ceased in 2005. A VIN containing a 'T' in Position 4 identifies a legacy vehicle where the engine is mounted in front of the cab rather than underneath it.11

#### 3.1.2 Analysis of Bus and Coach Series Codes

- K-Series (Code 'K'): The K-series designates a chassis with a longitudinally mounted engine, located centrally or at the rear. This is the standard platform for intercity coaches and touring buses.4

- N-Series (Code 'N'): The N-series designates a chassis with a transversely mounted engine at the rear. This configuration is compact and allows for a low floor throughout the entire length of the bus, making it standard for city transit applications.4

- F-Series (Code 'F'): Designates a front-engine chassis. These are typically ruggedized chassis used in developing markets or for specific applications like school buses.

Inferential Insight: The transition from the PGRT generation to the NTG generation introduced the S and L codes. Therefore, any VIN with an 'S' or 'L' in Position 4 can be immediately identified as a Next Generation vehicle (post-2016). Conversely, 'P', 'G', and 'R' codes appear in both generations, requiring further cross-referencing with the Model Year (Position 10) to determine the generation.

### 3.2 Position 5: Cab Type and Chassis Adaptation

Position 5 modifies the broad Series designation found in Position 4. It provides granular detail regarding the specific cab variant (e.g., Day Cab, Sleeper Cab) or the chassis adaptation (e.g., Tractor vs. Rigid). This position is critical for bodybuilders as it defines the physical envelope of the driver's compartment.

#### 3.2.1 Decoding Cab Types and Duty Classes

While Scania uses detailed internal codes like "CR19" (Cab R-series, 19dm length), the VIN compresses this into a single character in Position 5. Research indicates a strong correlation between this character and the Duty Class or Chassis Height.

- Chassis Adaptation Codes:

- A: Articulated (Tractor unit). Designed to pull semi-trailers.8

- B: Basic (Rigid truck). Designed to carry a fixed body (e.g., box, tank, tipper).14

- Chassis Height and Duty Class:

- M (Medium): Standard duty class for normal road conditions.14

- H (Heavy): High chassis height, often associated with construction or off-road applications to provide ground clearance.14

- E (Extra Heavy): Severe duty or heavy haulage.

- L (Low): Low chassis height for volume transport (maximizing cargo height within legal limits).15

- N (Normal): Standard chassis height.14

Forensic Application: A VIN containing a code mapping to 'H' in this section, combined with a 'P' series code (Position 4), strongly suggests a construction tipper or mixer. A code mapping to 'L' with an 'S' series code suggests a long-haul volume tractor.

### 3.3 Positions 6 and 7: Drive Configuration and Axle Arrangement

The combination of characters in positions 6 and 7 encodes the vehicle's wheel configuration. This is one of the most vital specifications for regulatory compliance (axle weight limits) and functional capability. Scania offers a massive array of configurations, from standard 4x2 tractors to complex 8x4\*4 tridems.

#### 3.3.1 Common Axle Configurations

- 4x2: Two axles, one driven (the rear). The standard configuration for European long-haul tractors.16

- 6x2: Three axles, one driven. Extremely common in Europe for distribution and heavy haulage. This configuration often includes a "Tag Axle" (behind the drive axle) or a "Pusher Axle" (in front of the drive axle) to distribute weight.15

- 6x4: Three axles, two driven (the rear bogie). Standard for construction vehicles requiring traction on loose surfaces.8

- 8x4: Four axles, two driven. Used for heavy tippers, concrete mixers, and crane trucks.16

- 6x2*4: A specific Scania notation often found in technical descriptions, denoting a 6x2 vehicle with a steered tag axle (*4 indicates steering on the tag).18

Decoding Logic: While the VIN may not explicitly spell out "6x2", the alphanumeric code in Positions 6 and 7 maps directly to these configurations in the Scania manufacturing database. For example, specific sequences are reserved for "3-axle rigid" vs "2-axle tractor." Decoding databases utilize these positions to populate the "Axle Configuration" field.8

### 3.4 Position 8: The Engine Family and Security Variance

Position 8 is highly variable and its interpretation depends heavily on the market (European Union vs. North America) and the specific era of production.

#### 3.4.1 Engine Family Designation (Primary EU Usage)

In many Scania VINs, particularly those for the European market, Position 8 identifies the Engine Family.3 It is crucial to understand that this code identifies the hardware platform, not the software calibration.

- The Engine Hardware: The code distinguishes between engine blocks, such as the DC09 (9-liter inline-5), DC13 (13-liter inline-6), or the legendary DC16 (16-liter V8).8

- The Software Nuance: A single hardware engine, such as the DC13, can be calibrated for various power outputs (e.g., 410 hp, 450 hp, 500 hp, 540 hp). The VIN Position 8 code generally identifies the hardware (DC13) but not the specific horsepower rating.

- Implication: You can infer from Position 8 that a truck has a 13-liter engine, but you cannot definitively state it is a "450 hp" model without accessing the specific chassis build record or the Type Designation plate.20

#### 3.4.2 Safety and Restraint Systems (North American Influence)

In North American markets, Position 8 is frequently mandatorily designated for safety restraint systems (e.g., airbags, seatbelts) to comply with NHTSA regulations.6 This creates a divergence in VIN logic. A Scania truck built for Mexico (WMI 3AX) might utilize Position 8 for safety codes, whereas a Swedish-built unit (WMI YS2) uses it for engine codes.

Forensic Insight: The presence of a V8 engine is a major value driver in the used truck market. If Position 8 codes for the V8 engine family, the vehicle commands a premium. This makes Position 8 a critical check for buyers verifying that a truck is a genuine V8 and not a re-badged inline-6.

### 3.5 Position 9: The Check Digit and Integrity Verification

Position 9 acts as a data integrity layer. Its function varies strictly by geography.

#### 3.5.1 North America: The Mandatory Checksum

For vehicles destined for North American markets (USA, Canada, Mexico), Position 9 is a mandatory Check Digit. It is the result of a mathematical algorithm (modulus 11) applied to the weighted values of all other 16 characters in the VIN.22

- Values: The result is always a number (0–9) or the letter 'X'.

- Purpose: It allows computer systems to immediately flag invalid VINs caused by typos or fraud.

#### 3.5.2 Europe: Optional Usage

In Europe, the Check Digit is not legally mandatory. However, many manufacturers, including Scania, increasingly adopt the North American checksum logic for their global VINs to ensure compatibility with international databases.3 In instances where the checksum is not used, Position 9 may be used for internal coding, though standard Scania practice leans toward the checksum or a filler character.

---

## 4. The Vehicle Identifier Section (VIS): Production Logistics and Traceability

The VIS (Positions 10-17) transitions the identifier from describing the type of vehicle to identifying the specific physical unit. This section is the sequential log of the manufacturing process.

### 4.1 Position 10: Model Year Encoding

Scania adheres to the ISO 3779 standard for Model Year encoding. This character indicates the model year of the vehicle, which may differ slightly from the calendar year of production (e.g., a vehicle built in late 2023 may be a 2024 model year). The code follows a 30-year cycle.

#### Table 2: Scania Model Year Codes (2010–2030)

|      |      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- | ---- |
| Code | Year | Code | Year | Code | Year |
| A    | 2010 | H    | 2017 | P    | 2023 |
| B    | 2011 | J    | 2018 | R    | 2024 |
| C    | 2012 | K    | 2019 | S    | 2025 |
| D    | 2013 | L    | 2020 | T    | 2026 |
| E    | 2014 | M    | 2021 | V    | 2027 |
| F    | 2015 | N    | 2022 | W    | 2028 |
| G    | 2016 |      |      |      |      |

Note: The letters I, O, Q, U, Z, and the digit 0 are strictly excluded to prevent confusion with numeric characters.24

Generation Gap Analysis: The Model Year code is the primary filter for distinguishing vehicle generations.

- Code 'G' (2016): This year marks the pivot point. A VIN with Year Code 'G' could be one of the last PGRT series or one of the very first NTG series.

- Code 'S' (2025): Indicates current production.

### 4.2 Position 11: Plant of Manufacture

Position 11 identifies the specific factory where the vehicle underwent final assembly. This code is crucial for tracing manufacturing quality and supply chain origins.

#### Table 3: Identified Scania Assembly Plant Codes

|             |                      |                          |                                                                                                                       |
| ----------- | -------------------- | ------------------------ | --------------------------------------------------------------------------------------------------------------------- |
| Code        | Location             | Facility                 | Significance                                                                                                          |
| S (or var.) | Södertälje, Sweden   | Scania Global HQ         | The historic home of Scania. Produces complex chassis and houses the R&D center.7                                     |
| R (or var.) | Zwolle, Netherlands  | Scania Production Zwolle | The largest Scania assembly plant globally. Handles a massive volume of European and export production.4              |
| 9 (or A)    | Angers, France       | Scania Production Angers | A major hub for French and Southern European markets. Explicitly linked to code '9' in some technical documentation.9 |
| B (or var.) | São Bernardo, Brazil | Scania Latin America     | The central hub for South American production.9                                                                       |

Note: Plant codes can be subject to internal rotation or specific market variations. The codes listed are derived from available decoder snippets and historical data.

### 4.3 Positions 12–17: Sequential Production Number

The final six characters (Positions 12–17) constitute the sequential serial number. This is a unique number assigned to the chassis as it moves along the production line.

- Recall Precision: Safety recalls are almost always issued based on a range of sequential numbers (e.g., "All R-series trucks with serials 200500 through 200900").18

- Component Matching: While the VIN identifies the chassis, the Engine Serial Number is stamped separately on the block. However, the VIN allows a dealer to access the SWS database and retrieve the original Engine Serial Number associated with that specific chassis sequence to verify matching numbers.29

---

## 5. Inferring Vehicle Model and Specifications: A Forensic Methodology

One of the most complex tasks for analysts is inferring the commercial model name (e.g., "Scania R 450") solely from the VIN. Unlike passenger cars, where the VDS often spells out the model, Scania requires a deductive algorithm.

### 5.1 The Inference Algorithm

To reconstruct the model name, an analyst must synthesize data from three separate VIN sections.

Step 1: Determine the Series (Position 4)

- If Character = 'S' $\rightarrow$ S-Series (NTG Flagship).

- If Character = 'R' $\rightarrow$ R-Series (High Cab).

- If Character = 'G' $\rightarrow$ G-Series (Medium Cab).

- If Character = 'P' $\rightarrow$ P-Series (Low Cab).

- If Character = 'L' $\rightarrow$ L-Series (Low Entry).

Step 2: Determine the Generation (Position 4 + Position 10)

- If Series is 'R' AND Year is < 2016 ('G') $\rightarrow$ PGR Generation.

- If Series is 'R' AND Year is > 2017 ('H') $\rightarrow$ Likely NTG Generation.

- If Series is 'S' $\rightarrow$ Always NTG Generation (S-series did not exist in PGR).

Step 3: Estimate Engine Power (Position 8)

- Decode Position 8 to find the Engine Family (e.g., 13-Liter vs. 16-Liter).

- Inference: If the code indicates a V8 (16-liter), the model number will be high (R 520, R 580, R 650, R 730/770). If the code indicates a 9-liter, the model number will be low (P 280, P 320).

- Limitation: The VIN does not distinguish between an R 410 and an R 450 if they share the same DC13 hardware. This specific number is defined by the ECU software and fuel map. To obtain the exact horsepower rating, one must consult the Type Designation Plate or the SWS database using the VIN.8

### 5.2 Case Study: Decoding a Hypothetical VIN

Consider the VIN: YS2R4x2...

1. WMI (YS2): Manufactured in Sweden.

2. Series (R): R-Series Cab.

3. Chassis (4x2): Two-axle tractor unit (Standard Long Haul).

4. Inferred Model: Scania R-Series Tractor (e.g., R 450).

Consider the VIN: YS2S6x2...

1. WMI (YS2): Manufactured in Sweden.

2. Series (S): S-Series Cab (Flat Floor, Next Gen).

3. Chassis (6x2): Three-axle rigid or tractor (likely with tag axle).

4. Inferred Model: Scania S-Series (e.g., S 500 or S 650).

---

## 6. Distinguishing Generations: PGR vs. Next Generation (NTG)

The transition from the "PGRT" range (2004–2018) to the "Next Generation" (NTG) introduced in 2016 represented a complete re-engineering of the Scania modular system. Distinguishing these generations via VIN is a common requirement for parts supply and valuation.

### 6.1 The "S" Factor

The introduction of the S-Series is the most reliable VIN-based indicator of the NTG platform. The S-cab, with its completely flat floor, was a new addition to the lineup. Therefore, any VIN containing an 'S' in Position 4 is definitively an NTG vehicle.

### 6.2 The Overlapping "R" Series

The R-series exists in both the old (PGR) and new (NTG) generations.

- PGR R-Series: Features a small engine tunnel.

- NTG R-Series: Redesigned cab, different aerodynamics, but still retains a small engine tunnel (unlike the S-series).

- Differentiation: Analysts must look at the Model Year (Position 10). A 2015 'R' is PGR. A 2019 'R' is NTG. For the transition years (2016-2017), the Sequential Number (Positions 12-17) will show a distinct break in the sequence where the new line began.11

---

## 7. The Type Designation System: The Shadow of the VIN

While the VIN is the key, the Type Designation is the value. Scania trucks carry a secondary identification string, often found on the B-pillar or door jamb, which provides the granular data that the VIN compresses.

Example Type Designation: R 580 LA6x2MNA

- R: Cab Series (Matches VIN Pos 4).

- 580: Power Rating (The missing link in the VIN).

- L: Chassis Class (Long Haulage).

- A: Chassis Adaptation (Articulated/Tractor).

- 6x2: Wheel Configuration (Matches VIN Pos 6-7).

- M: Duty Class (Medium).

- N: Chassis Height (Normal).

- A: Suspension (Air).

The VIN leads you to the chassis record in the database, which contains this Type Designation string. This string confirms the exact horsepower and suspension setup, which the VIN only implies.17

---

## 8. Bodybuilder Instructions and Regulatory Compliance

For third-party bodybuilders (companies that add cranes, mixers, or bus bodies to Scania chassis), the VIN is a critical tool for regulatory reporting.

### 8.1 SCIP Database and REACH

Under EU regulations, manufacturers must report "Substances of Concern in Products" (SCIP) to the European Chemicals Agency (ECHA). Scania uses the VIN as the primary identifier for these submissions. Bodybuilders referencing the VIN can access Scania's SCIP dossiers to ensure their finished vehicle compliance.29

### 8.2 Type Approval

The VIN is the link to the Whole Vehicle Type Approval (WVTA). When a bodybuilder completes a vehicle (e.g., building a coach on a K-series chassis), they must issue a second stage Manufacturer's Record. The Scania VIN remains the primary chassis identifier, linking the finished bus back to the original incomplete vehicle certificate.32

---

## 9. Conclusion

The Scania Vehicle Identification Number is a triumph of modular data engineering. It successfully compresses the immense complexity of a custom-specified heavy truck into a standardized 17-character string. By mastering the decoding of the WMI (Origin), VDS (Architecture), and VIS (Production), stakeholders can reconstruct the vehicle's lineage with high precision.

Key Takeaways for Decoding:

1. WMI dictates Origin: Distinguish between Swedish (YS2) and Brazilian (9BS) production for accurate parts sourcing.

2. Position 4 dictates Generation: Look for the 'S' code to confirm Next Generation (NTG) status.

3. Position 8 dictates Capability: Use the engine family code to distinguish between inline-6 and V8 assets.

4. Position 10 dictates Age: Verify the model year against the generation timeline (PGR vs NTG).

5. The VIN is the Key, not the Map: For exact horsepower and software configurations, use the VIN to retrieve the "Type Designation" or "Individual Chassis Specification."

In an industry driven by uptime and precision, the ability to fluently read this digital language is indispensable for fleet managers, service technicians, and forensic analysts alike.

---

### Table 4: Master Summary of Scania VIN Logic

|         |          |               |                   |                                                                |
| ------- | -------- | ------------- | ----------------- | -------------------------------------------------------------- |
| Section | Position | Name          | Decoding Function | Key Codes                                                      |
| WMI     | 1–3      | World Mfr. ID | Origin & Mfr.     | YS2 (Sweden), 9BS (Brazil), 3AX (Mexico), XLER (Netherlands)   |
| VDS     | 4        | Series        | Cab Family        | P, G, R (PGR/NTG), S (NTG), T (Bonnet), K (Bus), N (Bus)       |
| VDS     | 5        | Cab/Chassis   | Adaptation        | A (Tractor), B (Rigid), M/H (Duty Class)                       |
| VDS     | 6–7      | Configuration | Axles             | Encodes 4x2, 6x2, 6x4, 8x4 (Market dependent coding)           |
| VDS     | 8        | Variable      | Engine/Safety     | Engine Family (EU) or Safety System (NA). Identifies V8 vs I6. |
| VDS     | 9        | Check Digit   | Security          | Calculated checksum (0-9, X) or Factory code filler.           |
| VIS     | 10       | Model Year    | Age               | G=2016, H=2017, J=2018... S=2025                               |
| VIS     | 11       | Plant Code    | Assembly          | S (Södertälje), R (Zwolle), 9 (Angers), B (São Bernardo)       |
| VIS     | 12–17    | Serial Number | Production        | Unique sequential number for tracking and recalls.             |

Report compiled by Senior Automotive Compliance and Forensic Data Analyst.

November 29, 2025

#### Works cited

1. Heavy-Duty Truck VIN Decoder - Fullbay, accessed November 29, 2025, [https://www.fullbay.com/tools/vin/](https://www.fullbay.com/tools/vin/)

2. What's a Vehicle Identification Number? How to Decode the World Manufacturer Identifier, accessed November 29, 2025, [https://checkventory.com/articles/whats-your-number/](https://checkventory.com/articles/whats-your-number/)

3. Vehicle identification number - Wikipedia, accessed November 29, 2025, [https://en.wikipedia.org/wiki/Vehicle_identification_number](https://en.wikipedia.org/wiki/Vehicle_identification_number)

4. Get Scania VIN History Report | Scania Vindecoder, accessed November 29, 2025, [https://vindecoder.eu/scania](https://vindecoder.eu/scania)

5. What is a vehicle identification number (VIN)? - SPV, accessed November 29, 2025, [https://spv-vehicle.com/industrial-news/what-is-a-vehicle-identification-number-vin-329.html](https://spv-vehicle.com/industrial-news/what-is-a-vehicle-identification-number-vin-329.html)

6. VIN Decoder | AutoZone, accessed November 29, 2025, [https://www.autozone.com/vin-decoder](https://www.autozone.com/vin-decoder)

7. Vehicle Identification Numbers (VIN codes)/World Manufacturer Identifier (WMI) - Wikibooks, accessed November 29, 2025, [https://en.wikibooks.org/wiki/Vehicle*Identification_Numbers*(VIN*codes)/World_Manufacturer_Identifier*(WMI)](<https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/World_Manufacturer_Identifier_(WMI)>)

8. Chassis serial number, accessed November 29, 2025, [https://www.seminuevos-scania.pe/wp-content/uploads/2020/10/Ficha-T%C3%A9cnica-ASZ-733.pdf](https://www.seminuevos-scania.pe/wp-content/uploads/2020/10/Ficha-T%C3%A9cnica-ASZ-733.pdf)

9. Scania AB | PDF | Motor Vehicle - Scribd, accessed November 29, 2025, [https://www.scribd.com/document/894604558/Scania-AB](https://www.scribd.com/document/894604558/Scania-AB)

10. Scania VIN Decoder, accessed November 29, 2025, [https://www.freevindecoder.eu/make/scania](https://www.freevindecoder.eu/make/scania)

11. Scania PRT-range - Wikipedia, accessed November 29, 2025, [https://en.wikipedia.org/wiki/Scania_PRT-range](https://en.wikipedia.org/wiki/Scania_PRT-range)

12. The difference between the next-Generation Scania R-Series and the S-Series - Reddit, accessed November 29, 2025, [https://www.reddit.com/r/trucksim/comments/7dobm8/the_difference_between_the_nextgeneration_scania/](https://www.reddit.com/r/trucksim/comments/7dobm8/the_difference_between_the_nextgeneration_scania/)

13. Great buses and coaches start here - Scania, accessed November 29, 2025, [https://www.scania.com/content/dam/scanianoe/market/fr/products-and-services/buses-and-coaches/notre-gamme/chassis-serie-k/scania-serie-k.pdf](https://www.scania.com/content/dam/scanianoe/market/fr/products-and-services/buses-and-coaches/notre-gamme/chassis-serie-k/scania-serie-k.pdf)

14. Scania Off-Road Training Guide | PDF | Transmission (Mechanics) | Axle - Scribd, accessed November 29, 2025, [https://www.scribd.com/document/533249238/Pengenalan-Product-Scania](https://www.scribd.com/document/533249238/Pengenalan-Product-Scania)

15. Truck type Dimensions Weights Engine Fuel SCR Propulsion batteries Scania ICS (Individual Chassis Specification), accessed November 29, 2025, [https://bodybuilder.scania.com/content/dam/bodybuilder/tbb-files/japan/quick_reference_guide_jp/R420_B6x2LB_R17N_325S+230S.pdf](https://bodybuilder.scania.com/content/dam/bodybuilder/tbb-files/japan/quick_reference_guide_jp/R420_B6x2LB_R17N_325S+230S.pdf)

16. 4x2, 6x2, 8x4 configurations: vehicle types - BAS World, accessed November 29, 2025, [https://www.basworld.com/content/4x2-6x2-8x4-configurations-vehicle-types](https://www.basworld.com/content/4x2-6x2-8x4-configurations-vehicle-types)

17. Scania Type Designation | PDF | Truck | Axle - Scribd, accessed November 29, 2025, [https://www.scribd.com/document/329503186/Scania-Type-Designation](https://www.scribd.com/document/329503186/Scania-Type-Designation)

18. Chassis Chassis type VIN code 2135251 G 360 B6x2LB YS2G6X20000000000 2137658 G 360 B6x2LB 5603427 G 360 B6x2LB YS2G6X20000000000 - Scania, accessed November 29, 2025, [https://www.scania.com/content/dam/www/market/jp/about-scania/recall-campaign/chassis_list/GAI3962_CHASSISLIST.html](https://www.scania.com/content/dam/www/market/jp/about-scania/recall-campaign/chassis_list/GAI3962_CHASSISLIST.html)

19. Revision to General Motors' Vehicle Identification Number decoding for 2015 Model Year - NHTSA, accessed November 29, 2025, [https://www.nhtsa.gov/es/filebrowser/download/222336](https://www.nhtsa.gov/es/filebrowser/download/222336)

20. The 8th Eighth Digit in the VIN Vehicle Identification Number Indicates Engine - YouTube, accessed November 29, 2025, [https://www.youtube.com/watch?v=XDTaAUK5-4I](https://www.youtube.com/watch?v=XDTaAUK5-4I)

21. Scania VIN decoder - VinDocs, accessed November 29, 2025, [https://vindocs.com/us/scania-vin-decoder](https://vindocs.com/us/scania-vin-decoder)

22. What's in a VIN? How to decode the vehicle identification number, your car's unique fingerprint | Clemson News, accessed November 29, 2025, [https://news.clemson.edu/whats-in-a-vin-how-to-decode-the-vehicle-identification-number-your-cars-unique-fingerprint/](https://news.clemson.edu/whats-in-a-vin-how-to-decode-the-vehicle-identification-number-your-cars-unique-fingerprint/)

23. Check Digit Calculator for VINs - CJ Pony Parts, accessed November 29, 2025, [https://www.cjponyparts.com/resources/check-digit-calculator](https://www.cjponyparts.com/resources/check-digit-calculator)

24. VIN-to-Year Chart - ALLDATA, accessed November 29, 2025, [https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart](https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart)

25. 10th Digit of VIN to Vehicle Model Year Chart - Diminished Value Georgia, accessed November 29, 2025, [https://diminishedvalueofgeorgia.com/wp-content/uploads/10th-Digit-of-VIN-to-Vehicle-Model-Year-Chart.pdf](https://diminishedvalueofgeorgia.com/wp-content/uploads/10th-Digit-of-VIN-to-Vehicle-Model-Year-Chart.pdf)

26. The Vehicle Identification Number (VIN) - NISR - National Institute of Safety Research, accessed November 29, 2025, [https://cmvid.nisrinc.com/cmv_id/VIN_Visor.pdf](https://cmvid.nisrinc.com/cmv_id/VIN_Visor.pdf)

27. List of Volkswagen Group factories - Wikipedia, accessed November 29, 2025, [https://en.wikipedia.org/wiki/List_of_Volkswagen_Group_factories](https://en.wikipedia.org/wiki/List_of_Volkswagen_Group_factories)

28. Information package | Scania Group, accessed November 29, 2025, [https://www.scania.com/group/en/home/products-and-services/services/rmi/information-package.html](https://www.scania.com/group/en/home/products-and-services/services/rmi/information-package.html)

29. Frequently Asked Questions - Scania SURE, accessed November 29, 2025, [https://sure.scania.com/faq](https://sure.scania.com/faq)

30. Scania Designacion | PDF | Axle | Truck - Scribd, accessed November 29, 2025, [https://www.scribd.com/document/421869792/scania-designacion](https://www.scribd.com/document/421869792/scania-designacion)

31. SCIP - Scania Bodybuilder Portal, accessed November 29, 2025, [https://bodybuilder.scania.com/buses/en/type-approval/database-for-information-on-substances-of-concern-in-products--s.html](https://bodybuilder.scania.com/buses/en/type-approval/database-for-information-on-substances-of-concern-in-products--s.html)

32. Type Approval Support - Scania Bodybuilder Portal, accessed November 29, 2025, [https://bodybuilder.scania.com/trucks/en/local-information/uk/type-approval-support.html](https://bodybuilder.scania.com/trucks/en/local-information/uk/type-approval-support.html)

33. Legal Requirements and certificates - Scania Bodybuilder Portal, accessed November 29, 2025, [https://bodybuilder.scania.com/content/dam/bodybuilder/bbb-files/type-approval/Legal_Requirements_and_certificates_KC_series_v2.pdf](https://bodybuilder.scania.com/content/dam/bodybuilder/bbb-files/type-approval/Legal_Requirements_and_certificates_KC_series_v2.pdf)
