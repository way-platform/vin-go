# Global Automotive Identification Standards: A Technical Treatise on the Ford of Europe (WF0) VIN Architecture

## 1. Introduction: The Geopolitics of Vehicle Identification

The standardized identification of motor vehicles is not merely a bureaucratic requirement; it is the foundational bedrock of the global automotive ecosystem. It facilitates everything from supply chain logistics and forensic accident reconstruction to international trade compliance and safety recall management. At the heart of this system lies the Vehicle Identification Number (VIN), a 17-character alphanumeric string that serves as the machine-readable DNA of a modern automobile. While the International Organization for Standardization (ISO) established the ISO 3779 standard to harmonize these identifiers globally, the implementation of this standard varies significantly across regional jurisdictions.

Nowhere is this divergence more technically profound than in the internal coding architectures of the Ford Motor Company. As one of the world's few truly trans-continental manufacturers, Ford operates under two distinct, albeit related, identification paradigms: the North American standard, governed by the Code of Federal Regulations (CFR) Title 49 Part 565, and the European standard, historically rooted in Council Directive 76/114/EEC and current EU regulations. For the automotive forensic analyst, fleet manager, or parts specialist, the failure to distinguish between these two architectures is the single most common source of decoding errors.

This report presents an exhaustive technical analysis of the Ford of Europe VIN format, specifically the WF0 World Manufacturer Identifier structure. Unlike its North American counterpart, which prioritizes safety equipment and engine displacement within the VIN string, the European Ford VIN is a masterclass in assembly logistics and body configuration coding. Through a granular deconstruction of the Vehicle Descriptor Section (VDS) and the Vehicle Identifier Section (VIS), this document will elucidate the cryptic logic governing positions 4 through 9, identify the mechanisms for determining assembly origin and production chronology, and provide the definitive methodology for decoding the identity of European Ford vehicles.

The analysis draws upon technical service data, homologation documents, and fleet identification guides to construct a unified theory of Ford’s European identification strategy. It addresses the critical "blind spots" in the European format—specifically the absence of engine codes and check digits—and offers robust methodologies for overcoming these limitations in professional practice.

---

## 2. The Regulatory and Corporate Framework

To decode the VIN, one must first understand the regulatory environment that shapes it. The 17-character VIN is divided into three standardized zones: the World Manufacturer Identifier (WMI), the Vehicle Descriptor Section (VDS), and the Vehicle Identifier Section (VIS). While the length is fixed, the content is fluid, determined by the interaction between international standards and regional legal requirements.

### 2.1 The ISO 3779 vs. FMVSS 115 Divide

In North America, the National Highway Traffic Safety Administration (NHTSA) enforces a rigid interpretation of the VDS (Positions 4-8). Manufacturers must encode specific safety data (restraint systems) and technical specifications (engine type) to comply with Federal Motor Vehicle Safety Standards (FMVSS). Furthermore, the 9th position is mandated as a mathematical check digit—a modulus 11 calculation designed to detect transcription errors.1

In Europe, the regulatory touch is lighter regarding the specific content of the VDS. EU regulations mandate the presence of a VIN and its uniqueness over a 30-year period, but they grant manufacturers significant latitude in how they utilize the descriptor section. Ford of Europe utilizes this flexibility to prioritize logistical data over technical specifications. Consequently, the "Check Digit" at Position 9 is absent, replaced by a Model Code, and the "Engine Code" at Position 8 is replaced by an Assembly Plant identifier.4 This architectural inversion means that a decoder algorithm designed for a US Ford F-150 will return catastrophic errors when applied to a German-built Ford Focus.

### 2.2 The 'WF0' Paradigm: Legal Entity vs. Manufacturing Location

The first three characters of the VIN, the World Manufacturer Identifier (WMI), ostensibly indicate the country of origin and the manufacturer. For Ford of Europe, the dominant WMI is WF0.

Decoding WF0:

- W: Germany (Geographic Zone: Europe).

- F: Ford-Werke GmbH (Manufacturer).

- 0: Passenger Vehicle / Legal Entity Designation.1

It is imperative to understand that WF0 does not necessarily guarantee that the vehicle was assembled in Germany. In the era of the European Economic Community and later the EU, WF0 serves as the designator for Ford of Europe as a legal homologation entity headquartered in Cologne, Germany. A Ford Fiesta assembled in the Valencia plant in Spain may still carry a WF0 VIN because the vehicle is type-approved under the German subsidiary's authority. This contrasts with other regional codes like VS6 (Ford Spain) or SFA (Ford UK), which were more prevalent in the pre-consolidation era or are used for specific local production runs.5

For the analyst, WF0 acts as a regional umbrella. It signals that the vehicle adheres to the "European" decoding logic. If a VIN begins with 1FA (USA), 3FA (Mexico), or NM0 (Turkey), different decoding tables apply. The focus of this report is the WF0 architecture, which covers the vast majority of passenger cars (Fiesta, Focus, Mondeo, Kuga, Puma) sold in the European market.6

Table 1: Comparative WMI Designators within the Ford Ecosystem

|     |         |                     |                                                    |
| --- | ------- | ------------------- | -------------------------------------------------- |
| WMI | Region  | Manufacturer Entity | Primary Usage                                      |
| WF0 | Germany | Ford-Werke GmbH     | Standard EU Passenger Cars (Focus, Fiesta, Kuga) 1 |
| WF1 | Germany | Ford-Werke (Merkur) | Historical export models (XR4Ti) 7                 |
| NM0 | Turkey  | Ford Otosan         | Transit Commercial Range (Custom, Courier) 4       |
| VS6 | Spain   | Ford España S.A.    | Valencia production (often superseded by WF0) 5    |
| SFA | UK      | Ford Motor Co. Ltd. | Historical Dagenham/Halewood production 5          |
| X9F | Russia  | Ford Sollers        | Russian market Focus/Mondeo 5                      |
| 1FA | USA     | Ford Motor Company  | North American Passenger Cars (Mustang) 3          |
| 3FM | Mexico  | Ford Motor Company  | North American MPVs (Mach-E, Bronco Sport) 8       |

The distinction between WF0 and NM0 is particularly relevant for commercial fleets, as the Ford Transit Connect and Transit Custom are often manufactured in Turkey (NM0) but may be registered or imported under logistics chains that blur these lines. However, the internal VDS logic remains largely consistent across the European footprint.4

---

## 3. The Vehicle Descriptor Section (VDS): Anatomy of an Identity

The Vehicle Descriptor Section, occupying positions 4 through 9, is the technical core of the VIN. In the Ford of Europe architecture, this section is used to define the vehicle's physical form (Body), its corporate lineage (Source), its assembly point (Plant), and its model line. This differs radically from the US standard which uses this space for Safety Restraints, Series, and Engine.

### 3.1 Position 4: The Body Configuration Architecture

In the North American 1FA format, Position 4 is legally mandated to encode the Restraint System (e.g., airbags, seatbelts) and GVWR class.3 In the WF0 format, Position 4 is liberated from this requirement and is utilized to designate the Body Type.

This character defines the fundamental geometry of the vehicle—number of doors, rear configuration (hatch vs. saloon), and roofline. It is the first filter in identifying the vehicle's silhouette.

Common Position 4 Codes (Passenger Vehicles):

- A: 5-Door Hatchback (Standard configuration for models like the Escort, early Focus).

- B: 3-Door Hatchback / Saloon (Common on Fiesta and Focus).

- C: 2-Door Coupe (Rare, seen on models like the Ford Cougar or specific Fiesta ST trims).

- D: Estate / Wagon / 5-Door (Ubiquitous on Focus Wagon and Mondeo Estate).

- E: 3-Door Hatchback (Alternative code depending on generation).

- F: 4-Door Saloon (Sedan) (Standard for Mondeo and Focus Sedans).

- G: Multi-Purpose Vehicle (MPV) / Crossover (Used for C-Max, Kuga, S-Max).

- T: Convertible (Focus CC, StreetKa).

- W: 3-Door Commercial or Van derivative.

5

The Transit Correlation (Position 4 & 10):

A unique quirk exists within the commercial vehicle lineage (Transit). For many Transit generations, the Body Type code in Position 4 is often repeated or correlated with Position 10. While Position 10 is usually the Model Year in US VINs, in European Transits, it serves as a structural echo of the body style. For example, a "Double Cab Short Wheelbase" might code as 'M' in Position 4 and 'M' in Position 10. This redundancy helps confirm the configuration in the absence of a check digit.4

- G: 5-Seater Van SWB

- H: Kombi SWB

- J: Kombi LWB

- K: 5-Seater Van LWB

- M: Double Cab SWB

- T: Closed Van SWB

4

### 3.2 Positions 5 and 6: The "XX" Constant and Filler Logic

Perhaps the most perplexing feature for analysts transitioning from American to European VIN decoding is the content of positions 5 and 6. In a US VIN, these positions are data-rich, encoding the specific Series (trim level) and Chassis Type.3

In the WF0 format, positions 5 and 6 are almost invariably populated with the characters XX.

The Logic of Null Data:

The ISO 3779 standard requires the VDS to be six characters long. However, Ford of Europe’s internal coding system allows the Model Code (Position 9) and the Body Code (Position 4) to carry sufficient weight to identify the vehicle shell. The trim level (e.g., Titanium, ST-Line, Vignale) is not encoded in the VIN string. Since the trim level is determined by the "Build Sheet" and not the chassis manufacturing process, Ford chooses not to hardcode it into the chassis number.

Therefore, XX serves as a standardized filler to satisfy the character count requirement without conveying variable data. This is observed across the spectrum, from the Ford Ka to the Ford Mondeo. Deviations from XX are extremely rare and typically denote specialized homologation series, export-specific batches, or non-standard chassis supplied to third-party coachbuilders.4

Forensic Insight: If a European Ford VIN presents distinct characters in positions 5 and 6 (e.g., WF0A12...), it warrants immediate scrutiny. It may indicate a grey-market import, a mis-stamped chassis, or a specialized vehicle (like an armored variant or emergency service chassis) that falls outside standard passenger car production lines.

### 3.3 Position 7: The Product Source Lineage

Position 7 provides a glimpse into the corporate geography of Ford of Europe. It designates the Product Source Company or the affiliate responsible for the production line. This is distinct from the physical assembly plant (Position 8); rather, it indicates the manufacturing authority or legacy division.

Table 2: Position 7 Product Source Codes

|      |                              |                                                                                                                                   |
| ---- | ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| Code | Corporate Affiliate / Source | Historical Context                                                                                                                |
| G    | Ford-Werke AG (Germany)      | The default code for most modern European production (Focus, Fiesta, Kuga). Reflects the centralized management from Cologne.5    |
| B    | Ford Britain                 | A legacy code tracing back to Ford UK's independence (Dagenham/Halewood). Rarely seen on new vehicles but common on classics.5    |
| W    | Ford Spain                   | Vehicles managed under the Spanish subsidiary, typically Valencia output.5                                                        |
| T    | Ford Turkey (Otosan)         | Exclusively used for the Transit commercial range managed by the Ford Otosan joint venture.4                                      |
| S    | Mazda (Japan)                | A relic of the Ford-Mazda partnership. Found on vehicles like the Ford Fiesta (Mk4/5) which shared platforms with the Mazda 121.5 |
| P    | AutoEuropa (Portugal)        | Specific to the Ford Galaxy/VW Sharan joint venture era in Palmela, Portugal.5                                                    |

In the modern "One Ford" era, the code G has become ubiquitous, representing the consolidated European operations. Even vehicles built in Romania or Spain often carry the 'G' source code, reinforcing the centralization of VIN assignment under the German HQ.5

### 3.4 Position 8: The Assembly Plant – The Critical Divergence

This position represents the single most significant architectural difference between US and EU VINs.

- US Standard: Assembly Plant is Position 11.

- EU Standard: Assembly Plant is Position 8.

Misinterpreting this position is the primary cause of errors in origin tracing. In the WF0 format, Position 8 identifies the physical factory where the vehicle was welded and painted.

Table 3: Comprehensive Ford Europe Assembly Plant Codes (Position 8)

|      |                              |                                                        |                                    |
| ---- | ---------------------------- | ------------------------------------------------------ | ---------------------------------- |
| Code | Plant Location               | Primary Models (Historical & Current)                  | Status                             |
| G    | Cologne (Köln), Germany      | Fiesta, Puma (original), Scorpio, Capri                | Active (Transitioning to EV) 5     |
| C    | Saarlouis, Germany           | Focus, C-Max, Kuga (Mk1), Escort                       | Active (Scheduled for wind-down) 5 |
| P    | Valencia (Almussafes), Spain | Kuga (Current), Mondeo, Galaxy, S-Max, Transit Connect | Active 5                           |
| B    | Genk, Belgium                | Mondeo, S-Max, Galaxy                                  | Closed (2014) 5                    |
| A    | Dagenham, UK                 | Fiesta, Sierra, Cortina                                | Closed (2002) 5                    |
| D    | Southampton, UK              | Transit                                                | Closed (2013) 5                    |
| R    | Craiova, Romania             | Puma (New), EcoSport, Transit Courier                  | Active 5                           |
| T    | Kocaeli, Turkey              | Transit, Transit Custom                                | Active (Ford Otosan) 4             |
| S    | St. Petersburg, Russia       | Focus, Mondeo                                          | Suspended/Closed 5                 |
| K    | Rheine (Karmann), Germany    | Escort Cabriolet                                       | Closed 5                           |
| E    | Cork, Ireland                | Sierra, Cortina                                        | Closed (1984) 5                    |

Analytical Scenario: The Mondeo Shift

An analyst tracking the history of the Ford Mondeo can observe the industrial history of Ford through Position 8.

- Mondeos from 2007–2014 carry the code B (Genk, Belgium).

- Following the closure of the Genk plant in 2014, the code shifts to P (Valencia, Spain) for the 2015+ model years.
  This transition provides a forensic marker for verifying the model generation and ensuring the VIN corresponds to the correct manufacturing window.5

### 3.5 Position 9: The Model Identifier (No Check Digit)

In the US 1FA format, Position 9 is the Check Digit. In the WF0 format, Position 9 is the Model Code. Because Ford of Europe does not utilize a check digit, the VIN lacks an intrinsic mathematical self-validation mechanism. This makes the system prone to data entry errors (e.g., typing an 'S' instead of a '5'), which will not be flagged by standard validation algorithms used in DMS (Dealer Management Systems) software designed for US VINs.14

Position 9 encodes the specific model family. Note that these codes are often reused over decades for different generations of the same nameplate.

Table 4: Ford Europe Model Codes (Position 9)

|      |                          |                                                                                       |
| ---- | ------------------------ | ------------------------------------------------------------------------------------- |
| Code | Model Identity           | Notes                                                                                 |
| A    | Escort / Orion           | Also used for early Focus development mules.5                                         |
| B    | Mondeo / Sierra / Cougar | The lineage of mid-size D-segment cars.5                                              |
| C    | Focus / C-Max            | The primary C-segment identifier.5                                                    |
| D    | Focus                    | Often used for specific body variations (Wagon) or later generations.5                |
| E    | Puma (1997) / Capri      | The compact sports coupe lineage.5                                                    |
| F    | Fiesta / Transit         | The B-segment hatch. Also used for some Transit variants (context dependent on WMI).5 |
| G    | Scorpio / Granada        | Historical executive saloons.5                                                        |
| J    | Fiesta / Fusion          | Used during the Mk5/Mk6 crossover era.15                                              |
| M    | Kuga                     | The C-segment SUV (Escape equivalent).15                                              |
| R    | Ka                       | The A-segment city car.15                                                             |
| S    | Galaxy                   | The large MPV.15                                                                      |
| W    | Galaxy / S-Max           | Shared platform MPVs.15                                                               |
| X    | Transit                  | Large commercial variants.                                                            |

Contextual Decoding Required:

The code F illustrates the necessity of context.

- If WMI = WF0 (Passenger) and Pos 4 = A (5-door hatch) and Pos 9 = F -> Ford Fiesta.

- If WMI = NM0 (Commercial) and Pos 4 = T (Van) and Pos 9 = F -> Ford Transit variant.
  The model code cannot be interpreted in isolation; it requires the VDS context.5

---

## 4. The Vehicle Identifier Section (VIS): Chronology and Serialization

The VIS (Positions 10-17) allows for the precise dating and serialization of the unit. Here, the deviation from the US standard is absolute and is the most frequent source of confusion regarding the "Year" of a European Ford.

### 4.1 Position 10: The Phantom Year

In almost every other VIN standard globally, Position 10 indicates the Model Year (e.g., 'X' = 1999, 'Y' = 2000). In the WF0 format, Position 10 is NOT the Model Year.

Instead, Position 10 is largely a structural placeholder or a repeater of the Body Type code (Position 4).

- In Transits, Pos 10 often duplicates Pos 4 (e.g., ...TTFD... where Pos 4 is D and Pos 10 is D).

- In Passenger cars, it may be a constant based on the model line (e.g., '1' or 'A') but holds no chronological value.

Attempting to read Position 10 as a year will result in nonsensical dates (e.g., interpreting a 'D' in Position 10 as 2013, when the vehicle is actually a 2005 model).4

### 4.2 Position 11: The True Year of Manufacture

Ford of Europe places the Year of Manufacture code in Position 11.

- US VIN: Pos 11 = Plant.

- EU VIN: Pos 11 = Year.

The coding schema follows a repeating alphabet cycle. Unlike the ISO standard which cycles every 30 years, Ford's internal cycle has historically repeated more frequently or utilized specific subsets of the alphabet.

Table 5: Ford Europe Year of Manufacture Codes (Position 11)

|      |      |      |      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| Code | Year | Code | Year | Code | Year | Code | Year |
| Y    | 2000 | 5    | 2005 | A    | 2010 | F    | 2015 |
| 1    | 2001 | 6    | 2006 | B    | 2011 | G    | 2016 |
| 2    | 2002 | 7    | 2007 | C    | 2012 | H    | 2017 |
| 3    | 2003 | 8    | 2008 | D    | 2013 | J    | 2018 |
| 4    | 2004 | 9    | 2009 | E    | 2014 | K    | 2019 |
| L    | 2020 | M    | 2021 | N    | 2022 | P    | 2023 |
| R    | 2024 | S    | 2025 |      |      |      |      |

Note: Letters I, O, Q, U, Z are typically omitted to prevent confusion with numbers..4

### 4.3 Position 12: The Monthly Rotation Cipher

To determine the precise build date, Position 11 must be paired with Position 12, which encodes the Month of Manufacture.

Ford uses a proprietary rotational cipher for months. The code for "January" is not fixed; it changes based on the year. This ensures that the combination of Year (Pos 11) and Month (Pos 12) is unique within a 4-year cycle, preventing overlaps.

The Rotation Mechanism:

The sequence of month codes typically follows the pattern: C - K - D - E - L - Y - S - T - J - U - M - P - B - R - A - G.

This sequence is mapped against the calendar months, shifting each year.

Table 6: Representative Month Decoding Matrix

This table demonstrates the rotation logic. Exact alignment relies on the specific "Start Year" of the cycle.

|       |                    |                    |                    |                    |
| ----- | ------------------ | ------------------ | ------------------ | ------------------ |
| Month | Year 1 (e.g. 2024) | Year 2 (e.g. 2025) | Year 3 (e.g. 2026) | Year 4 (e.g. 2027) |
| Jan   | L                  | C                  | B                  | J                  |
| Feb   | Y                  | K                  | R                  | U                  |
| Mar   | S                  | D                  | A                  | M                  |
| Apr   | T                  | E                  | G                  | P                  |
| May   | J                  | L                  | C                  | B                  |
| Jun   | U                  | Y                  | K                  | R                  |
| Jul   | M                  | S                  | D                  | A                  |
| Aug   | P                  | T                  | E                  | G                  |
| Sep   | B                  | J                  | L                  | C                  |
| Oct   | R                  | U                  | Y                  | K                  |
| Nov   | A                  | M                  | S                  | D                  |
| Dec   | G                  | P                  | T                  | E                  |

Interpretation: If Position 11 is L (2020) and Position 12 is J, one consults the 2020 column. If J corresponds to May in that year's rotation, the build date is May 2020. This system effectively creates a hash of the production date, requiring specific manufacturer tables to decode accurately.4

---

## 5. Global Conflicts and Identifying the "Hybrid" VIN

A major challenge in modern decoding arises when Ford utilizes global platforms, leading to "hybrid" VIN structures that blend US and EU logic.

### 5.1 The Mustang Mach-E Case Study

The Ford Mustang Mach-E presents a unique forensic anomaly. Physically manufactured in Cuautitlán, Mexico, North American units carry the WMI 3FM. However, units destined for the European market are homologated under Ford-Werke GmbH and carry the WMI WF0.8

Decoding the EU Mach-E (WF0):

- WMI: WF0 (Germany/Europe Entity).

- Pos 4: Body Type (e.g., T).

- Pos 5-7: Unlike standard EU VINs which use 'XX', the Mach-E often utilizes these positions for Line/Series codes (e.g., K1E), adopting a US-style descriptor within the EU framework to align with global EV platform standards.20

- Pos 8: Motor/Drive Type (e.g., M, 7, S) rather than the strict Plant Code used in legacy ICE vehicles.20

- Pos 11: Assembly Plant. In this specific global platform instance, the Plant Code (e.g., M for Cuautitlán) may appear in Position 11, mimicking the US standard, or in Position 11 as the Year, depending on the specific batch homologation.21

Conclusion: The Mach-E represents a transition point where the rigid separation between 1FA and WF0 logic is blurring due to the "One Ford" global strategy. Forensics on these vehicles require checking the specific homologation plate (B-Pillar) rather than relying solely on legacy VIN tables.

### 5.2 The "Missing" Engine Code

A critical deficiency in the WF0 format is the lack of engine data.

- US VIN: Pos 8 = Engine (e.g., 'F' = 5.0L V8).

- EU VIN: Pos 8 = Plant.

Where is the Engine Code?

For European Fords, the engine code is not in the VIN. It is a 4-character alpha-numeric code (e.g., G8DA, M1DA) stamped physically on the engine block or printed on the vehicle identification plate (sticker) located on the B-pillar or lock carrier.

To determine the engine from the VIN remotely, one must access the Ford ETIS (Electronic Technical Information System) database, which maps the VIN serial number to the specific "As-Built" data. There is no algorithm to derive the engine solely from the 17 characters.22

---

## 6. Comprehensive Decoding Table (WF0 Standard)

The following table synthesizes the analysis into a lookup tool for the standard European Ford VIN.

|              |                |                                                                                                                                                                                |
| ------------ | -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| VIN Position | Description    | Decoding Logic / Typical Values                                                                                                                                                |
| 1-3          | WMI            | WF0 = Ford-Werke GmbH (Standard Passenger)<br><br> <br><br>NM0 = Ford Otosan (Commercial)<br><br> <br><br>VS6 = Ford Spain                                                     |
| 4            | Body Type      | A/B = Hatchback<br><br> <br><br>F = Sedan<br><br> <br><br>D = Estate<br><br> <br><br>G = MPV/SUV (Kuga/C-Max)<br><br> <br><br>W = Van                                          |
| 5-6          | Filler         | XX (Standard Constant)                                                                                                                                                         |
| 7            | Product Source | G = Germany (Ford-Werke)<br><br> <br><br>W = Spain<br><br> <br><br>T = Turkey<br><br> <br><br>B = UK (Legacy)                                                                  |
| 8            | Assembly Plant | C = Saarlouis (Focus)<br><br> <br><br>G = Cologne (Fiesta)<br><br> <br><br>P = Valencia (Kuga/Mondeo)<br><br> <br><br>R = Craiova (Puma)<br><br> <br><br>T = Kocaeli (Transit) |
| 9            | Model Code     | C = Focus<br><br> <br><br>F = Fiesta<br><br> <br><br>M = Kuga<br><br> <br><br>B = Mondeo<br><br> <br><br>S = Galaxy                                                            |
| 10           | Body/Filler    | Often repeats Position 4 (e.g., A or D). NOT THE YEAR.                                                                                                                         |
| 11           | Year           | L=2020, M=2021, N=2022, P=2023, R=2024 16                                                                                                                                      |
| 12           | Month          | Rotational Cipher (See Table 6). Requires Year context.                                                                                                                        |
| 13-17        | Serial         | Sequential production number.                                                                                                                                                  |

---

## 7. Conclusion

The decoding of the Ford WF0 VIN is a distinct discipline from its North American equivalent. It requires the analyst to discard the expectations of FMVSS-compliant structures—specifically the check digit and the engine code—and embrace a system rooted in logistical hierarchy.

The VDS (Positions 4-9) functions as a logistical map: it tells us the form of the body (Pos 4), the managing subsidiary (Pos 7), the physical factory (Pos 8), and the model line (Pos 9). The VIS completes this picture with a time-stamp (Pos 11-12) and a serial identity.

For the professional, the key markers of a valid European Ford VIN are the presence of XX in positions 5-6, the absence of a check digit validation, and the location of the Plant Code in Position 8. Understanding these nuances is essential for accurate asset identification, preventing the conflation of manufacturing origin with legal homologation, and navigating the complex history of Ford’s European operations.

#### Works cited

1. Ford VIN Decoder - Lookup and check Ford VIN Number and Get Vehicle History - Vininspect.com, accessed November 29, 2025, [https://vininspect.com/vin/ford](https://vininspect.com/vin/ford)

2. 2006 Model Year Vehicle Identification Number Codes - Revision #4 - NHTSA, accessed November 29, 2025, [https://www.nhtsa.gov/filebrowser/download/186861](https://www.nhtsa.gov/filebrowser/download/186861)

3. A Guide to Decode VIN Numbers in Ford Vehicles, accessed November 29, 2025, [https://highlandford.com/blog/a-guide-to-decode-vin-numbers-in-ford-vehicles/](https://highlandford.com/blog/a-guide-to-decode-vin-numbers-in-ford-vehicles/)

4. Wf0 D XX TT F D 8 A 52617: Decoding Ford Transit Chassis Number - Year 2002 To 2009, accessed November 29, 2025, [https://www.scribd.com/document/522475983/ford-transit-vincodes](https://www.scribd.com/document/522475983/ford-transit-vincodes)

5. Vehicle Identification Numbers (VIN codes)/Ford/VIN Codes - Wikibooks, open books for an open world, accessed November 29, 2025, [https://en.wikibooks.org/wiki/Vehicle*Identification_Numbers*(VIN_codes)/Ford/VIN_Codes](<https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/Ford/VIN_Codes>)

6. Vehicle Identification Numbers (VIN codes)/World Manufacturer Identifier (WMI) - Wikibooks, accessed November 29, 2025, [https://en.wikibooks.org/wiki/Vehicle*Identification_Numbers*(VIN*codes)/World_Manufacturer_Identifier*(WMI)](<https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/World_Manufacturer_Identifier_(WMI)>)

7. Vehicle identification number - Wikipedia, accessed November 29, 2025, [https://en.wikipedia.org/wiki/Vehicle_identification_number](https://en.wikipedia.org/wiki/Vehicle_identification_number)

8. VIN starting with a W - Mach E Forum, accessed November 29, 2025, [https://www.macheforum.com/site/threads/vin-starting-with-a-w.3307/](https://www.macheforum.com/site/threads/vin-starting-with-a-w.3307/)

9. VIN Lookup & Decoder - Ford Pro, accessed November 29, 2025, [https://www.fordpro.com/en-us/fleet-vehicles/vin-decoder-and-guides/](https://www.fordpro.com/en-us/fleet-vehicles/vin-decoder-and-guides/)

10. accessed November 29, 2025, [https://www.scribd.com/document/522475983/ford-transit-vincodes#:~:text=Position%201%2D3%20indicate%20the,13%2D17%20the%20sequential%20number.](https://www.scribd.com/document/522475983/ford-transit-vincodes#:~:text=Position%201%2D3%20indicate%20the,13%2D17%20the%20sequential%20number.)

11. In. Plant Code Liste - Ford, accessed November 29, 2025, [https://www.suppcomm.ford.com/europe/Documents/GlobalPlantList.xls](https://www.suppcomm.ford.com/europe/Documents/GlobalPlantList.xls)

12. 2022 VIN GUIDE | Ford Pro, accessed November 29, 2025, [https://content.fordpro.com/content/dam/fordpro/us/en-us/pdf/fleet-vehicles/vin-lookup-and-guides/2022-vin-guide.pdf](https://content.fordpro.com/content/dam/fordpro/us/en-us/pdf/fleet-vehicles/vin-lookup-and-guides/2022-vin-guide.pdf)

13. VIN: FORD - Carvin-Info, accessed November 29, 2025, [https://carvin-info.com/marka-avto/ford/](https://carvin-info.com/marka-avto/ford/)

14. Position 1 The very first letter or number of the VIN tells you in what region of the world your vehicle was made. Match the let, accessed November 29, 2025, [http://dpefuel.com/wp-content/uploads/2018/06/VIN-DECODER.pdf](http://dpefuel.com/wp-content/uploads/2018/06/VIN-DECODER.pdf)

15. Vehicle Identification Numbers Vin Codes Ford Vin Codes Wikibooks Open Books For An Open World Compress - Scribd, accessed November 29, 2025, [https://www.scribd.com/document/812413223/Vehicle-Identification-Numbers-Vin-Codes-Ford-Vin-Codes-Wikibooks-Open-Books-for-an-Open-World-Compress](https://www.scribd.com/document/812413223/Vehicle-Identification-Numbers-Vin-Codes-Ford-Vin-Codes-Wikibooks-Open-Books-for-an-Open-World-Compress)

16. Check / Decode Ford Car VIN Number & Manufacturing Date - V3Cars, accessed November 29, 2025, [https://www.v3cars.com/car-guide/how-to-find-ford-vehicle-manufacturing-date-with-vin-number](https://www.v3cars.com/car-guide/how-to-find-ford-vehicle-manufacturing-date-with-vin-number)

17. VIN-to-Year Chart - ALLDATA, accessed November 29, 2025, [https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart](https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart)

18. 2024 MY VIN Attachment - Ford Pro, accessed November 29, 2025, [https://content.fordpro.com/content/dam/fordpro/ca/en-ca/pdf/fleet-vehicles/vin-lookup-and-guides/2024-VIN-Guide.pdf](https://content.fordpro.com/content/dam/fordpro/ca/en-ca/pdf/fleet-vehicles/vin-lookup-and-guides/2024-VIN-Guide.pdf)

19. Ford Manufacture Dates - Burton Power, accessed November 29, 2025, [https://www.burtonpower.com/tuning-guides/tuning-guide-pages/ford-manufacture-dates.html](https://www.burtonpower.com/tuning-guides/tuning-guide-pages/ford-manufacture-dates.html)

20. 2021 Mach-E Mustang VIN Decoder | Page 3 | MachEforum, accessed November 29, 2025, [https://www.macheforum.com/site/threads/2021-mach-e-mustang-vin-decoder.651/page-3](https://www.macheforum.com/site/threads/2021-mach-e-mustang-vin-decoder.651/page-3)

21. VIN does not match manufacturer country | MachEforum, accessed November 29, 2025, [https://www.macheforum.com/site/threads/vin-does-not-match-manufacturer-country.18838/](https://www.macheforum.com/site/threads/vin-does-not-match-manufacturer-country.18838/)

22. Engine number - Vehicle identification, accessed November 29, 2025, [https://www.fordservicecontent.com/Ford_Content/vdirsnet/OwnerManual/Home/Content?variantid=1789&languageCode=en&countryCode=USA&moidRef=G539686&Uid=G1076512&ProcUid=G563122&userMarket=GBR&div=f&vFilteringEnabled=False&buildtype=web](https://www.fordservicecontent.com/Ford_Content/vdirsnet/OwnerManual/Home/Content?variantid=1789&languageCode=en&countryCode=USA&moidRef=G539686&Uid=G1076512&ProcUid=G563122&userMarket=GBR&div=f&vFilteringEnabled=False&buildtype=web)

23. Ford Tourneo Courier, Transit Courier (2014-2020) - Service Manual - Wiring Diagrams - Owners Manual | PDF | Rechargeable Battery | Electric Generator - Scribd, accessed November 29, 2025, [https://www.scribd.com/document/658370642/Ford-Tourneo-Courier-Transit-Courier-2014-2020-Service-Manual-Wiring-Diagrams-Owners-Manual](https://www.scribd.com/document/658370642/Ford-Tourneo-Courier-Transit-Courier-2014-2020-Service-Manual-Wiring-Diagrams-Owners-Manual)
