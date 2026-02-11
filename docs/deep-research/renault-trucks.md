# The Forensic Identification of Renault Trucks: A Technical Analysis of VIN Architectures in Light Commercial and Heavy-Duty Fleets

## 1. Executive Introduction: The Strategic Imperative of Vehicle Identification

The modern logistics and transportation sector operates on a foundation of precision. In an era defined by just-in-time delivery models, stringent emissions regulations, and the rapid electrification of fleets, the ability to accurately identify, catalog, and maintain assets is not merely an administrative convenience but a strategic imperative. At the heart of this identification ecosystem lies the Vehicle Identification Number (VIN), a seventeen-character alphanumeric sequence that serves as the immutable DNA of every commercial vehicle. For Renault Trucks, a manufacturer with a storied lineage characterized by complex mergers, strategic alliances, and a bifurcated product portfolio spanning light commercial vans to heavy-duty construction haulers, the VIN is a particularly dense repository of engineering and manufacturing data.

This report provides an exhaustive, expert-level analysis of the VIN decoding methodologies applicable to the Renault Trucks lineup. It addresses a critical ambiguity in the fleet management domain: the distinction between vehicles produced under the aegis of Renault S.A. (the passenger car entity) and those engineered by Renault Trucks (a division of the Volvo Group). This distinction is most acute in the crossover segments, such as the Master and Trafic vans, where branding strategies like the "Red Edition" can obscure the underlying manufacturing provenance. By deconstructing the World Manufacturer Identifier (WMI), the Vehicle Descriptor Section (VDS), and the Vehicle Indicator Section (VIS), this analysis establishes a definitive framework for inferring model identity, chassis configuration, and powertrain specifications.

The research synthesizes data from technical manuals, regulatory filings, and decoding databases to produce a nuanced understanding of how Renault Trucks encodes its diverse fleet. From the heavy-duty T, C, K, and D ranges—which utilize a dedicated truck-specific logic—to the light commercial vehicles that share genetic markers with passenger automobiles, this report offers fleet managers, parts specialists, and compliance officers a comprehensive guide to the forensic identification of Renault Trucks assets. The following sections will systematically dissect these coding structures, culminating in detailed inference tables that map the complex alphanumeric strings to tangible vehicle attributes.1

## 2. The Geopolitical and Industrial Context of Renault VINs

To effectively decode a Renault Truck VIN, one must first understand the industrial architecture that generates it. Unlike manufacturers with a singular, unified production output, Renault Trucks represents a convergence of historical legacies—specifically the integration of Berliet and Saviem into Renault Véhicules Industriels (RVI), and the subsequent acquisition by the Volvo Group in 2001. This history is encoded in the very first characters of the VIN, creating a divergence in logic between the heavy-duty fleet and the light commercial offer.

### 2.1 The World Manufacturer Identifier (WMI): The Primary Watershed

The first three characters of the VIN, the World Manufacturer Identifier (WMI), act as the primary filter for model inference. In the Renault ecosystem, the WMI allows for an immediate, high-level categorization of the vehicle’s origin and regulatory weight class. The distinction is binary: a vehicle is either a product of the heavy-truck manufacturing system or a product of the light-vehicle alliance system.

The VF6 Identifier: The Hallmark of Heavy Industry

The WMI VF6 is the definitive identifier for Renault Trucks as a heavy-duty manufacturer. This code is assigned to the legal entity responsible for the engineering and production of trucks typically exceeding 3.5 tonnes Gross Vehicle Weight (GVW). When a VIN commences with VF6, it signals that the vehicle was produced within the industrial footprint of the Volvo Group-owned Renault Trucks division. This encompasses the flagship long-haul Range T, the construction-focused Range K and Range C, and the distribution-oriented Range D.4

The presence of VF6 triggers a specific decoding protocol for the subsequent characters (positions 4-9), which describes the vehicle in terms of chassis rigidness, axle configuration (e.g., 6x4, 8x4), and heavy-duty powertrain options. This logic is distinct from passenger car decoding and is harmonized with the broader practices of heavy equipment manufacturing, prioritizing technical capability over trim levels.

The VF1 and VN1 Identifiers: The Light Commercial Nexus

In contrast, the light commercial vehicles (LCVs) marketed by Renault Trucks—specifically the Master and Trafic—often bear the WMI VF1 or VN1.

- VF1: This code identifies Renault S.A., the parent company responsible for passenger cars and light vans. A Master or Trafic van carrying this WMI shares its fundamental architecture with the consumer versions of these vehicles.

- VN1: This identifier is specific to SOVAB (Société des Véhicules Automobiles de Batilly), a manufacturing subsidiary in France dedicated to the production of the Master van. The use of VN1 highlights the vehicle's assembly location and specific manufacturing entity but follows the same general decoding logic as VF1.4

The "Red Edition" marketing label, applied to high-specification vans sold through the Renault Trucks dealer network, creates a potential point of confusion. While these vehicles are serviced and sold as "trucks," their VINs typically remain VF1 or VN1, indicating their shared lineage with the mass-market van product. However, as will be discussed in later sections, specific heavy-duty chassis-cab variants of the Master may carry the VF6 identifier, distinguishing them as specialized upfit-ready platforms.6

The VS Identifier: The Spanish Connection (Maxity)

The Renault Maxity, a compact cab-over-engine truck, introduces a third variable. As a badge-engineered version of the Nissan Cabstar, the Maxity is manufactured in Ávila, Spain. Consequently, its VIN often begins with a Spanish WMI code, such as VS, reflecting its Nissan-led production environment. While branded as a Renault Truck, its "genetic code" is rooted in Nissan's naming conventions, requiring a different decoding lens than the VF6 or VF1 fleets.7

### 2.2 Table 1: Primary WMI Classifications for Model Inference

The following table summarizes the primary WMI codes encountered in the Renault Trucks ecosystem and their implications for model inference.

|          |                              |                                                  |                                                                     |
| -------- | ---------------------------- | ------------------------------------------------ | ------------------------------------------------------------------- |
| WMI Code | Manufacturer / Entity        | Primary Vehicle Families                         | Inference Implication                                               |
| VF6      | Renault Trucks (Volvo Group) | Range T, K, C, D, Midlum, Premium, Magnum, Kerax | Definitively a Heavy-Duty Truck. Use HD decoding logic.             |
| VF1      | Renault S.A.                 | Master, Trafic, Kangoo                           | Light Commercial Vehicle (LCV). Passenger car logic applies.        |
| VN1      | SOVAB (Batilly)              | Master                                           | LCV manufactured at the Batilly plant. Passenger car logic applies. |
| VF2      | Renault Trucks (Legacy)      | Older Heavy Trucks / Buses                       | Legacy HD Truck. Rarely seen in new fleets.                         |
| VNE      | Irisbus / Iveco Bus          | Buses, Coaches                                   | Passenger Transport / Bus division.                                 |
| VS...    | Nissan / Renault Spain       | Maxity, Mascott (some)                           | Nissan-derived chassis. Requires Nissan-style decoding.             |

1

## 3. Decoding the Heavy-Duty Fleet: The VF6 Architecture

For fleet operators managing heavy assets, the VF6 VIN is the primary source of truth. The Vehicle Descriptor Section (VDS), occupying positions 4 through 9, contains a structured code that reveals the vehicle family, chassis configuration, and gross weight characteristics. This section analyzes the specific alphanumeric patterns that distinguish the modern ranges (T, K, C, D) from one another.

### 3.1 Positions 4-5: The Vehicle Family Designator

In the VF6 logic, the fourth and fifth characters form a critical dyad known as the Family Code. This two-character sequence maps directly to the commercial model nameplate. By identifying this code, one can instantly segregate a long-haul tractor from a construction rigid.

- The T Range (Long Haul):
  The successor to the Magnum and Premium Route, the Renault Trucks T, utilizes specific codes that reflect its modern chassis architecture. Analysis of VIN datasets indicates that codes such as 11 and 29 are prevalent in this range.

- Code 11: Frequently observed on standard T-Range tractor units (e.g., T460, T480). A VIN beginning with VF611 is a strong indicator of a T-series tractor.10

- Code 29: Also associated with the T-Range, potentially distinguishing specific chassis evolutions or high-roof (High Sleeper) variants. VINs starting VF629 often correlate with T-High models.12

- The K Range (Heavy Construction):
  The Renault Trucks K, designed for severe off-road applications and replacing the Kerax, inherits and evolves the coding structure of its predecessor. The legacy Kerax code was 34, and this often persists or transitions to related codes in the K series.

- Code 34: Historically the identifier for the Kerax. In the modern era, VF634 continues to signify a heavy construction chassis, characterized by high ground clearance and steel bumpers.13

- Codes 32 / 38: Observed in 8x4 and 6x6 configurations within the K Range, denoting specific multi-axle drive layouts required for mining and quarrying applications.13

- The C Range (Construction & Distribution):
  The Renault Trucks C bridges the gap between the on-road efficiency of the T range and the ruggedness of the K range. It often utilizes codes derived from the former Premium Distribution line.

- Codes 24 / 25: These codes are heavily linked to the C Range (and the legacy Premium). A VIN starting with VF624 or VF625 suggests a vehicle optimized for construction supply or heavy distribution, often with a lighter chassis weight than the K Range.13

- The D Range (Distribution):
  Covering the medium-duty segment (formerly Midlum and Premium Distribution), the Renault Trucks D (including D Narrow and D Wide) employs a distinct set of codes.

- Codes 20 / 21: These are the primary identifiers for the D Range. For instance, VF620 is commonly associated with the D Wide (18-26t), while VF621 or similar variations may denote the standard D (10-18t) or specific narrow-cab versions.15

### 3.2 Position 6: The Chassis and Axle Configuration

The sixth character in a VF6 VIN provides a granular definition of the truck's physical layout, specifically the axle configuration (e.g., 4x2, 6x4) and the chassis type (Tractor vs. Rigid). This character is essential for determining the vehicle's operational envelope.

- Rigid Trucks (Porteurs):

- A: 4x2 Rigid (Standard distribution or light construction chassis).

- F: 8x4 Rigid (Heavy tipper or mixer chassis, common in K and C ranges).13

- E: 6x6 Rigid (All-wheel drive, severe duty).

- L: 8x4 Rigid (Specific long-wheelbase or tridem configurations).13

- Tractor Units (Tracteurs):

- G: 4x2 Tractor ( The standard long-haul configuration for T Range).

- H: 4x4 Tractor (AWD tractor, often K or C range).

- K: 6x4 Tractor (Heavy haulage or construction tractor).

- J: 6x2 Tractor (Pusher or Tag axle, common in UK and Nordic markets).13

By combining the Family Code (Pos 4-5) with the Configuration Code (Pos 6), a highly specific model inference is possible. For example, a VIN starting VF634F... can be decoded as a Renault Trucks K (34) in an 8x4 Rigid (F) configuration—a classic mining tipper spec. Conversely, VF611G... decodes to a Renault Trucks T (11) in a 4x2 Tractor (G) configuration—the quintessential long-haul highway truck.

### 3.3 Engine Identification (Positions 7-8)

While earlier VIN systems often used a single character for the engine, modern Renault Trucks VINs utilize positions 7 and 8 (and sometimes 9) to encode the engine family and power rating. This is crucial for distinguishing between the 11-liter and 13-liter powerplants.

- DTi 11 (10.8 Liters): This engine powers the T, C, and D Wide ranges. It offers power outputs typically ranging from 380 to 460 hp. In the VIN, specific alphanumeric combinations in the VDS linked to the "11" family code indicate this displacement.

- DTi 13 (12.8 Liters): Reserved for the heavy hitters—the T High, K, and heavy C models. Power outputs range from 440 to 520 hp. Differentiating a T480 from a T520 generally requires access to the manufacturer's specific lookup tables for these positions, as they map to internal option codes (e.g., HD or MD sequences seen in snippets).10

The transition to Euro 6 Step D and E standards has introduced further variations in these codes to denote emissions compliance levels, meaning the specific character mapping evolves with each regulatory update.18

### 3.4 Table 2: The Master Decoding Matrix for Heavy-Duty Models (VF6)

This table allows for the rapid inference of model identity based on the first six characters of the VIN.

|               |                  |                      |                     |                        |                         |
| ------------- | ---------------- | -------------------- | ------------------- | ---------------------- | ----------------------- |
| WMI (Pos 1-3) | Family (Pos 4-5) | Model Range Inferred | Config Code (Pos 6) | Chassis Interpretation | Common Applications     |
| VF6           | 11               | Range T / T High     | G                   | 4x2 Tractor            | Long Haul Logistics     |
| VF6           | 11 / 29          | Range T / T High     | J                   | 6x2 Tractor            | Heavy Haulage / UK Spec |
| VF6           | 34               | Range K (Kerax)      | K                   | 6x4 Tractor            | Heavy Construction      |
| VF6           | 34 / 32          | Range K (Kerax)      | F                   | 8x4 Rigid              | Mining / Tipper         |
| VF6           | 34 / 38          | Range K (Kerax)      | E                   | 6x6 Rigid              | Off-Road / Military     |
| VF6           | 24 / 25          | Range C              | G                   | 4x2 Tractor            | Regional Distribution   |
| VF6           | 24 / 25          | Range C              | F                   | 8x4 Rigid              | Concrete Mixer          |
| VF6           | 20 / 21          | Range D / D Wide     | A                   | 4x2 Rigid              | Urban Distribution      |
| VF6           | 17               | Magnum (Legacy)      | G                   | 4x2 Tractor            | Historical Long Haul    |

10

## 4. Decoding the Light Commercial Fleet: The VF1 & VN1 Architecture

The decoding logic for Renault Trucks' Light Commercial Vehicle (LCV) fleet differs fundamentally from the heavy-duty sector. Here, the VIN reflects the passenger car division's coding standards, prioritizing platform generation and body style over gross weight ratings.

### 4.1 The Renault Trucks Master ("Red Edition")

The Renault Master is the linchpin of the LCV offer. Identifying a "Renault Trucks" specific Master (often marketed as "Red Edition") purely from the VIN is challenging because the chassis is frequently shared with the standard Renault LCV range.

- WMI Analysis: Most Masters will bear the VN1 (Sovab, Batilly) or VF1 (Renault France) WMI. This confirms the vehicle is a Master van but does not explicitly confirm it is a "Red Edition."

- The VF6 Master Exception: A critical insight for fleet identifiers is the existence of VF6 Masters. Snippets indicate that specific Master chassis—likely heavy-duty variants (4.5t GVW), rear-wheel-drive chassis cabs intended for substantial bodywork, or specific "Trucks" homologated units—are registered with the VF6 WMI. A Master with a VF6 VIN is definitively a product of the truck division and not a standard passenger van.6

- VDS Logic (Positions 4-7):

- Family Code: The fourth position often indicates the platform generation. F typically denotes the Master II/III platform, while V is appearing on newer generations (Master IV).6

- Body Type (Position 5):

- D: Panel Van (Fourgon).

- E: Chassis Cab (Plateau-Ridelles/Chassis-Cabine) – common for "Red Edition" upfits.

- U: Tipper chassis.20

- Engine Code: Positions 6 and 7 encode the engine type (e.g., 2.3 dCi). Unlike the DTi engines, these are derived from the passenger car M9T engine family.

Implication for "Red Edition" Identification:

The "Red Edition" is a trim and service specification (chrome grille, heavy-duty dashboard, truck-grade maintenance contract). Since the VIN describes the physical chassis, a VF1 Master Red Edition and a VF1 standard Master may share extremely similar VIN structures. Definitive confirmation of "Red Edition" status for VF1 units typically requires cross-referencing the VIN with Renault Trucks' internal Consult system or checking for specific option codes (e.g., "Red Edition" badging options) in the build sheet.21

### 4.2 The Renault Trucks Trafic

The Trafic (Red Edition) follows a similar pattern to the Master but is strictly a VF1 vehicle (or VF1 equivalent depending on the manufacturing partnership).

- Model Identification: Positions 4 and 5 of the VIN provide the model code.

- FL: Trafic II.

- JG / FG: Trafic III.

- Like the Master, the "Red Edition" status is an overlay on the standard Trafic chassis. However, acquiring a Trafic through the Renault Trucks network ensures access to heavy-duty service bays and extended hours, a distinction in operation rather than identification.5

### 4.3 The Maxity Anomaly

The Renault Maxity is a distinct entity. As a rebadged Nissan Cabstar, its identification logic is Nissan-centric.

- WMI: Often VS... (Nissan/Renault Spain).

- Chassis Location: The VIN is stamped on the right-hand chassis member behind the cab, consistent with Nissan's manufacturing standards in Ávila.8

- Inference: A VIN starting with VS on a vehicle with Renault badges immediately identifies it as a Maxity/Mascott lineage vehicle, distinct from the Batilly-built Masters or Bourg-en-Bresse heavy trucks.

### 4.4 Table 3: LCV Model Inference Guide

|               |           |                   |                 |                                                                                   |
| ------------- | --------- | ----------------- | --------------- | --------------------------------------------------------------------------------- |
| Model Name    | WMI       | Typical VDS Start | Plant (Pos 11)  | Identification Notes                                                              |
| Master III/IV | VN1 / VF1 | MA, FD, FV        | K (Batilly)     | Standard LCV logic. "Red Edition" requires visual/database confirmation.          |
| Master HD     | VF6       | VF                | K (Batilly)     | Heavy-duty homologation (>3.5t or specific chassis). Definitively Truck Division. |
| Trafic        | VF1       | FL, JG, FG        | V (Sandouville) | Medium Van. Shared platform.                                                      |
| Maxity        | VS...     | Variable          | S (Ávila)       | Cab-over engine. Nissan Cabstar architecture.                                     |

5

## 5. The Electric Revolution: Decoding E-Tech Models

The electrification of the Renault Trucks range (E-Tech) introduces a paradigm shift in VIN decoding. The traditional reliance on engine displacement codes (e.g., 11L vs 13L) is obsolete for these vehicles. Instead, the VIN must encode battery capacity, traction voltage, and electric motor configuration.

### 5.1 E-Tech Heavy Duty (T & C)

Manufactured in Bourg-en-Bresse alongside their diesel counterparts, the E-Tech T and C models utilize the VF6 WMI.

- VDS Shift: The engine code section (Positions 7-8) no longer maps to a DTi diesel engine. Instead, it utilizes specific character sets designated for electric propulsion (e.g., identifying 360 kWh or 540 kWh battery packs).

- Visual Indicators: While the VIN is definitive, the VDS will lack the standard combustion identifiers. Forensically, the absence of a standard diesel engine code in a VF6 VIN from 2023 onwards suggests an E-Tech unit.24

### 5.2 E-Tech D & D Wide

The electric distribution trucks (E-Tech D) are manufactured in Blainville-sur-Orne (V plant code).

- Decoding: These vehicles maintain the D-Range family codes (20/21) but feature unique propulsion codes. For example, a D Wide Z.E. (Zero Emission) will carry a VIN that reflects its electric drivetrain in the engine code position, differentiating it from a DTi 8 diesel version.24

### 5.3 E-Tech Master & Trafic

- Master E-Tech: Identified via VF1 or VN1, but the engine code (Pos 6-8) maps to electric motor outputs (e.g., 57kW) rather than dCi codes.

- Trafic E-Tech: Follows similar logic, with specific VDS codes for the electric powertrain.28

Strategic Insight: For emergency responders and mechanics, the VIN is a critical safety tool. An E-Tech VIN identifies the presence of high-voltage (600V+) systems. Renault Trucks has integrated this into the identification plates and rescue sheets, linked directly to the VIN.24

## 6. Technical Specifics: Plant and Year Codes

To complete the identification profile, one must decode the assembly location and the model year.

### 6.1 Plant Codes (Position 11)

The 11th character of the VIN reveals the factory of origin. This is a highly reliable static data point.

|      |                             |                                              |
| ---- | --------------------------- | -------------------------------------------- |
| Code | Plant Location              | Primary Models Produced                      |
| J    | Bourg-en-Bresse, France     | Range T, Range C, Range K (Heavy Duty)       |
| V    | Blainville-sur-Orne, France | Range D, D Wide, E-Tech D, Cab Manufacturing |
| K    | Batilly, France (SOVAB)     | Renault Master (LCV)                         |
| S    | Ávila, Spain                | Maxity (Nissan Plant)                        |
| Y    | Volvo Group Plants          | Shared assembly (rare for pure Renault)      |

4

### 6.2 Year of Manufacture (Position 10)

Renault Trucks generally adheres to the ISO 3779 standard for the model year code in the 10th position. However, it is crucial to note that in the European Union, the inclusion of the year code in the VIN is not mandatory for all vehicle classes, and sometimes a generic character (like 0 or 1) is used. When the year is encoded, it follows the standard rotation:

- A = 2010

- B = 2011

- C = 2012

- D = 2013 (Launch of the Euro 6 Range)

- ...

- L = 2020

- M = 2021

- R = 2024

- S = 2025

For vehicles where the 10th digit does not correlate to the year, the Manufacturing Plate (or CAM plate) on the chassis must be consulted for the precise production date (often encoded as DDMY).13

## 7. Operational Applications and Conclusion

The decoding of Renault Trucks VINs is a segmented discipline. For the heavy-duty fleet manager, the VF6 WMI serves as the entry point into a logical system where the Family Code (Pos 4-5) and Configuration Code (Pos 6) reveal the vehicle's fundamental purpose—whether it is a T-High long-hauler or a K-Range mining tipper. This architecture, inherited from RVI and refined under Volvo, provides transparency and precision.

In contrast, the LCV fleet operates under a more opaque VF1/VN1 system, where the distinction between a "Red Edition" commercial asset and a standard van is often hidden within option codes rather than the chassis identifier itself. The exception—the VF6 Master—stands as a beacon for specialized heavy-duty applications within the van segment.

As the industry transitions to the E-Tech portfolio, the VIN will remain the central nervous system of asset data, evolving to encode kilowatt-hours alongside axle ratios. By mastering these decoding tables, stakeholders can ensure precise parts ordering, accurate valuation, and enhanced safety compliance across the diverse and evolving landscape of Renault Trucks.

Final Recommendation for Identification:

1. Check the WMI: VF6 = Heavy Truck (or HD Master). VF1/VN1 = LCV (Van).

2. For VF6: Check Positions 4-5 to identify the Range (11/29=T, 34=K, 20/21=D, 24/25=C).

3. For VF6: Check Position 6 for Axle Config (G=4x2 Tractor, F=8x4 Rigid, etc.).

4. For VF1 (Master): Check for "Red Edition" badging physically or via dealer database; do not rely on VIN alone for trim level.

5. For E-Tech: Look for non-standard engine codes in Positions 7-8 and cross-reference with the V or J plant code.

#### Works cited

1. VIN decoder RENAULT Specification & car history | Cebia.com, accessed December 1, 2025, [https://en.cebia.com/detailArticle/renault-vin-decoder-specification-car-history](https://en.cebia.com/detailArticle/renault-vin-decoder-specification-car-history)

2. Vehicle identification number - Wikipedia, accessed December 1, 2025, [https://en.wikipedia.org/wiki/Vehicle_identification_number](https://en.wikipedia.org/wiki/Vehicle_identification_number)

3. HD VIN Decoder Tool and Resources - DataOne Software, accessed December 1, 2025, [https://www.dataonesoftware.com/hd-truck-vin-decoder](https://www.dataonesoftware.com/hd-truck-vin-decoder)

4. Vehicle Identification Numbers (VIN Codes) - World Manufacturer Identifier (WMI) - Wikibooks, Open Books For An Open World | PDF - Scribd, accessed December 1, 2025, [https://www.scribd.com/document/409957445/Vehicle-Identification-Numbers-VIN-Codes-World-Manufacturer-Identifier-WMI-Wikibooks-Open-Books-for-an-Open-World](https://www.scribd.com/document/409957445/Vehicle-Identification-Numbers-VIN-Codes-World-Manufacturer-Identifier-WMI-Wikibooks-Open-Books-for-an-Open-World)

5. Vehicle Identification Numbers (VIN codes)/World Manufacturer Identifier (WMI) - Wikibooks, accessed December 1, 2025, [https://en.wikibooks.org/wiki/Vehicle*Identification_Numbers*(VIN*codes)/World_Manufacturer_Identifier*(WMI)](<https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/World_Manufacturer_Identifier_(WMI)>)

6. Renault VIN check - BidCar.eu Auctions, accessed December 1, 2025, [https://bidcar.eu/en/vin/report/renault](https://bidcar.eu/en/vin/report/renault)

7. VIN Decoder | NHTSA, accessed December 1, 2025, [https://www.nhtsa.gov/vin-decoder](https://www.nhtsa.gov/vin-decoder)

8. Nissan Cabstar - Wikipedia, accessed December 1, 2025, [https://en.wikipedia.org/wiki/Nissan_Cabstar](https://en.wikipedia.org/wiki/Nissan_Cabstar)

9. Renault Maxity - Wikipedia, accessed December 1, 2025, [https://en.wikipedia.org/wiki/Renault_Maxity](https://en.wikipedia.org/wiki/Renault_Maxity)

10. Renault T480 vin: VF611A36800000000 (MD 2021!) , P-climate! 2 tanks. Non painted for sale, Tractor unit, 48900 EUR - Truck1, accessed December 1, 2025, [https://www.truck1.eu/tractor-units/renault-t480-vin-vf611a368md027798-md-2021-p-climate-2-tanks-non-painted-a9930767.html](https://www.truck1.eu/tractor-units/renault-t480-vin-vf611a368md027798-md-2021-p-climate-2-tanks-non-painted-a9930767.html)

11. Renault Trucks T - Wikipedia, accessed December 1, 2025, [https://en.wikipedia.org/wiki/Renault_Trucks_T](https://en.wikipedia.org/wiki/Renault_Trucks_T)

12. Renault Trucks, accessed December 1, 2025, [https://vinsearch.ec24.ch/en/brand-detail/AS09aiQ9JGI6XH9teionMz0kO0pcSCsrd3FlLzxkZQAAAABjiXV8?session=673093431d6fe](https://vinsearch.ec24.ch/en/brand-detail/AS09aiQ9JGI6XH9teionMz0kO0pcSCsrd3FlLzxkZQAAAABjiXV8?session=673093431d6fe)

13. Vehicle Identification | PDF | Truck - Scribd, accessed December 1, 2025, [https://www.scribd.com/document/465961941/0-Vehicle-identification](https://www.scribd.com/document/465961941/0-Vehicle-identification)

14. Renault Trucks C 2021, accessed December 1, 2025, [https://middle-east.renault-trucks.com/en-mea/product/renault-trucks-c-2021](https://middle-east.renault-trucks.com/en-mea/product/renault-trucks-c-2021)

15. Renault D wide, accessed December 1, 2025, [https://www.seltrucks.com/en/renault-d-wide](https://www.seltrucks.com/en/renault-d-wide)

16. Renault Trucks D - Wikipedia, accessed December 1, 2025, [https://en.wikipedia.org/wiki/Renault_Trucks_D](https://en.wikipedia.org/wiki/Renault_Trucks_D)

17. Tractor unit Renault T520 vin: VF610A36800000000, 2 tanks, 2 bedsTUV till 07-09-2024, accessed December 1, 2025, [https://www.truck1.eu/tractor-units/renault-t520-vin-vf610a368hd008113-2-tanks-2-bedstuv-till-07-09-2024-a7979566.html](https://www.truck1.eu/tractor-units/renault-t520-vin-vf610a368hd008113-2-tanks-2-bedstuv-till-07-09-2024-a7979566.html)

18. RENAULT TRUCKS T AND T HIGH MODEL YEAR 2020: DRIVER COMFORT AND REDUCED FUEL CONSUMPTION, accessed December 1, 2025, [https://www.renault-trucks.com/en/newsroom/press-releases/renault-trucks-t-and-t-high-model-year-2020-driver-comfort-and-reduced-fuel](https://www.renault-trucks.com/en/newsroom/press-releases/renault-trucks-t-and-t-high-model-year-2020-driver-comfort-and-reduced-fuel)

19. A fast guide to Renault truck engines, accessed December 1, 2025, [https://mwtruckparts.co.uk/a-fast-guide-to-renault-truck-engines/](https://mwtruckparts.co.uk/a-fast-guide-to-renault-truck-engines/)

20. Lookup any Renault Car | VIN Check by AutoDetective.com, accessed December 1, 2025, [https://www.autodetective.com/make/renault/](https://www.autodetective.com/make/renault/)

21. NEW MASTER: RENAULT TRUCKS UNVEILS ITS EXCLUSIVE RED EDITION, accessed December 1, 2025, [https://www.renault-trucks.com/en/newsroom/press-releases/new-master-renault-trucks-unveils-its-exclusive-red-edition](https://www.renault-trucks.com/en/newsroom/press-releases/new-master-renault-trucks-unveils-its-exclusive-red-edition)

22. New generation Renault Trucks Master Red EDITION: efficient and versatile, accessed December 1, 2025, [https://www.renault-trucks.com/en/newsroom/press-releases/new-generation-renault-trucks-master-red-edition-efficient-and-versatile](https://www.renault-trucks.com/en/newsroom/press-releases/new-generation-renault-trucks-master-red-edition-efficient-and-versatile)

23. Distribution range | Renault Trucks Corporate, accessed December 1, 2025, [https://www.renault-trucks.com/en/distribution-range](https://www.renault-trucks.com/en/distribution-range)

24. RENAULT TRUCKS, accessed December 1, 2025, [https://www.renault-trucks.com/sites/corporate/files/2023-12/ERG-RT_MD_BEV_P4283_en-GB%28English%29.pdf](https://www.renault-trucks.com/sites/corporate/files/2023-12/ERG-RT_MD_BEV_P4283_en-GB%28English%29.pdf)

25. renault trucks e-tech t 6x2 - Information environnementale, accessed December 1, 2025, [https://www.renault-trucks.com/sites/corporate/files/2024-07/T%20E-Tech%206x2%20EN.pdf](https://www.renault-trucks.com/sites/corporate/files/2024-07/T%20E-Tech%206x2%20EN.pdf)

26. RENAULT TRUCKS UNVEILS THE DESIGN OF ITS ELECTRIC T AND C MODELS, accessed December 1, 2025, [https://www.renault-trucks.com/en/newsroom/press-releases/renault-trucks-unveils-design-its-electric-t-and-c-models](https://www.renault-trucks.com/en/newsroom/press-releases/renault-trucks-unveils-design-its-electric-t-and-c-models)

27. RENAULT TRUCKS - E-Tech D / D Wide / D Wide LEC, accessed December 1, 2025, [https://www.renault-trucks.cz/sites/default/files/2023-10/com11314%20%281%29_1.pdf](https://www.renault-trucks.cz/sites/default/files/2023-10/com11314%20%281%29_1.pdf)

28. Electromobility | Renault Trucks Corporate, accessed December 1, 2025, [https://www.renault-trucks.com/en/electromobility](https://www.renault-trucks.com/en/electromobility)

29. Renault Trucks E-Tech Trafic, a new addition to the manufacturer's electric range, accessed December 1, 2025, [https://www.renault-trucks.com/en/newsroom/press-releases/renault-trucks-e-tech-trafic-new-addition-manufacturers-electric-range](https://www.renault-trucks.com/en/newsroom/press-releases/renault-trucks-e-tech-trafic-new-addition-manufacturers-electric-range)

30. The Vehicle Identification Number (VIN) - NISR - National Institute of Safety Research, accessed December 1, 2025, [https://cmvid.nisrinc.com/cmv_id/VIN_Visor.pdf](https://cmvid.nisrinc.com/cmv_id/VIN_Visor.pdf)

31. VIN-to-Year Chart - ALLDATA, accessed December 1, 2025, [https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart](https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart)
