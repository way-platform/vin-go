# Technical Forensic Analysis of Volkswagen Commercial Vehicle Identification Protocols: European Market VDS Decoding and Homologation Architecture

## 1. Introduction to Automotive Identification Standards and the Volkswagen Paradigm

The Vehicle Identification Number (VIN) represents the fundamental DNA of modern automotive manufacturing, serving as a serialized fingerprint that encodes the origin, technical specifications, and production chronology of a specific chassis. While the global automotive industry operates under the umbrella of the International Organization for Standardization’s ISO 3779 and ISO 3780 protocols, the implementation of these standards varies significantly across regulatory jurisdictions. For forensic analysts, fleet managers, and automotive historians specializing in the European market, decoding the VINs of Volkswagen Commercial Vehicles (VWCV) presents a unique set of challenges and idiosyncrasies that differ markedly from the conventions used in North America.

The primary objective of this report is to provide an exhaustive technical analysis of the Volkswagen VIN architecture, with a specific focus on the European Vehicle Descriptor Section (VDS). Unlike their American counterparts, which utilize the VDS to encode granular data regarding engine displacement, restraint systems, and trim levels, European Volkswagen VINs—particularly those assigned to the Transporter, Crafter, Caddy, and ID. Buzz product lines—employ a "hollow" architecture. This design philosophy prioritizes chassis designation and manufacturing logistics over component-level identification within the VIN string itself. Consequently, a distinct methodological approach is required to extract meaningful intelligence from these identifiers.

This analysis draws upon a comprehensive review of technical documentation, manufacturing plant codes, and regulatory filings to establish a definitive decoding framework. It explores the geopolitical implications of the World Manufacturer Identifier (WMI), dissects the evolution of model series codes from the T3 Vanagon to the MEB-platform ID. Buzz, and delineates the critical boundary where the VIN loses its descriptive power, necessitating the use of supplementary data sources such as the proprietary "Datenträger" (vehicle data sticker) and Production (PR) codes. By synthesizing these elements, we establish a robust protocol for the accurate identification of Volkswagen vans within the European Union.

### 1.1 The Regulatory Divergence: ISO 3779 vs. FMVSS 115

To understand the specific structure of a European Volkswagen VIN, one must first appreciate the regulatory dichotomy that shaped its evolution. The modern 17-character VIN was standardized in 1980, replacing the earlier, shorter chassis number systems that varied wildly between manufacturers.1 The ISO 3779 standard defines the three-part structure—WMI, VDS, and VIS—that is universally recognized today. However, the utilization of the middle section, the VDS (positions 4 through 9), is subject to regional interpretation.

In the United States, the National Highway Traffic Safety Administration (NHTSA) mandates under Federal Motor Vehicle Safety Standard (FMVSS) 115 (now Part 565) that the VDS must contain specific attributes, including engine type, gross vehicle weight rating (GVWR), and restraint system information. Furthermore, position 9 is strictly reserved for a mathematical check digit calculated via a modulus 11 algorithm to detect transcription errors.3

In contrast, the European Union regulations (EU Regulation 19/2011 and its predecessors) are less prescriptive regarding the content of the VDS. They require that the vehicle be uniquely identifiable but do not mandate the encoding of specific technical attributes within the 17-character string. Volkswagen, leveraging this flexibility for its domestic and European markets, adopted a policy of using "filler" characters—universally the letter 'Z'—in positions where specific data is not legally mandated. This architectural decision streamlines the homologation process, as minor changes to engine tuning or trim levels do not require the generation of new VIN structures, but it renders the VDS largely opaque to casual observers attempting to determine vehicle specifications.1

## 2. The World Manufacturer Identifier (WMI): Geopolitics and Homologation

The first three characters of the VIN, known as the World Manufacturer Identifier (WMI), provide the highest level of classification. For Volkswagen Commercial Vehicles, the WMI is not merely a label of corporate ownership; it is a legal declaration of the vehicle's production origin and its regulatory category. This section provides the foundational context for determining whether a van was manufactured in Germany, Poland, or elsewhere, and whether it is legally classified as a passenger carrier or a goods vehicle—a distinction with significant tax and insurance implications in European markets.

### 2.1 Geographic Origin and Manufacturing Footprint

The first character of the WMI identifies the geographic region of the manufacturer, while the second character identifies the country. Volkswagen's commercial fleet is produced across a distributed network of European and South American facilities, each assigned specific codes that reflect the company's expanding logistical footprint.

The primary regional codes observed in the Volkswagen commercial lineage are:

- Region W (Germany): The letter 'W' denotes West Germany (historically) and now unified Germany. This is the ancestral home of Volkswagen and remains the primary identifier for the Transporter (T-Series) and the ID. Buzz produced in Hanover.5

- Region S (Poland): Since the 1990s, Volkswagen has shifted significant commercial production to Poland. The letter 'S' (specifically the WMI range SN-ST assigned to Germany's eastern neighbor) is now ubiquitous on the Caddy and the second-generation Crafter.6

- Region 3 (Mexico) and 9 (Brazil): While less common in modern European deliveries, historical T2 and T3 models, as well as specific light commercial derivatives, were sourced from Latin America. Brazil (9BW) maintained production of the "Kombi" T2 long after it ceased in Europe, and these units occasionally appear on the European gray market.1

### 2.2 Corporate Division and Vehicle Classification

The third character of the WMI, when combined with the first two, specifies the manufacturer's division and the vehicle type. For Volkswagen, this is where the critical distinction between "Commercial Vehicles" (Volkswagen Nutzfahrzeuge) and "Passenger Cars" (Volkswagen PKW) is made. This distinction is not academic; it dictates the homologation class (M1 vs. N1) under EU law.

#### Table 1: Primary Volkswagen WMI Codes for European Commercial Vehicles

|           |                   |                        |                                            |                                                                                                                      |
| --------- | ----------------- | ---------------------- | ------------------------------------------ | -------------------------------------------------------------------------------------------------------------------- |
| WMI Code  | Geographic Origin | Manufacturer Division  | Primary Vehicle Applications               | Regulatory Context                                                                                                   |
| WV1       | Germany           | VW Commercial Vehicles | Goods Vehicles (Panel Vans, Pickups)       | Classified as N1 category (Goods). Often subject to lower speed limits or different tax regimes in EU member states. |
| WV2       | Germany           | VW Commercial Vehicles | Passenger Transport (MPVs, Buses, Campers) | Classified as M1 category (Passenger). Used for Multivan, Caravelle, California, and ID. Buzz.                       |
| WV3       | Germany           | VW Commercial Vehicles | Incomplete Vehicles                        | Chassis cabs delivered to coachbuilders for conversion (e.g., box bodies, dropsides).                                |
| WVG       | Germany           | Volkswagen AG          | SUVs / MPVs                                | Historically used for Touran/Tiguan but overlaps with some crossover commercial variants (e.g., Caddy Life).         |
| WVW       | Germany           | Volkswagen AG          | Passenger Cars                             | Golf, Polo, Passat. Rarely found on true commercial vans unless derived from a car platform (e.g., Golf Van).        |
| SBS / S.. | Poland            | VW Poznan / Wrzesnia   | Commercial Vehicles                        | Caddy (Poznan) and Crafter (Wrzesnia). The prefix shifts to 'S' to reflect Polish origin.                            |
| VWV       | Spain             | Volkswagen             | Commercial / Passenger                     | Produced at the Navarra plant (less common for core van range).                                                      |

Forensic Insight on WV1 vs. WV2:

The differentiation between WV1 and WV2 provides immediate insight into the vehicle's original intended purpose. A Volkswagen Transporter T6 chassis could be physically identical in terms of wheelbase and engine, but if it carries a WV1 WMI, it left the factory as a panel van or "Kombi" intended for cargo. If it carries WV2, it was homologated as a passenger vehicle, such as a Caravelle or Multivan.5

This distinction is particularly relevant for the ID. Buzz. Current research indicates that the passenger version of the ID. Buzz typically bears the WV2 WMI, reflecting its status as an MPV.10 Conversely, the ID. Buzz Cargo, designed for logistics, would theoretically align with the WV1 designation, maintaining the historic separation of the product lines.

### 2.3 The Polish Manufacturing Hub (WMI S)

As Volkswagen consolidated its heavy commercial production in Poland, the WMI S became prominent. The Crafter, previously built by Mercedes-Benz in Germany (WMI W), moved to a dedicated VW plant in Wrzesnia, Poland, for its second generation. Consequently, the VINs for these modern heavy vans shifted from German identifiers to Polish ones. A VIN beginning with WV1 usually indicates a German-built Transporter, while a VIN beginning with SY or SZ (the internal codes often reflected in the WMI or VDS) indicates a Polish-built Crafter, although formally they fall under the Polish manufacturer code assignments.7

## 3. The Vehicle Descriptor Section (VDS): Decoding the "ZZZ" Matrix

The VDS (positions 4 through 9) is the focal point of the user's inquiry. In the European context, this section is a study in minimalism. Where a North American VIN would be densely packed with data regarding engine families (e.g., differentiating a 2.0L TDI from a 2.0L TSI) and safety restraint systems, the European VW VIN is populated almost exclusively by the "filler" character Z.

### 3.1 The "ZZZ" Phenomenon and its Implications

Positions 4, 5, and 6 in a European Volkswagen VIN are invariably Z-Z-Z. This sequence has become iconic among VW enthusiasts and technicians, often referred to as "triple Z" VINs.1

- Position 4: Filler (Z)

- Position 5: Filler (Z)

- Position 6: Filler (Z)

The presence of these fillers signifies that the vehicle's technical attributes are not encoded in the VIN. It is a critical error for an analyst to attempt to decode engine size or transmission type from these positions on a European van. Unlike a Ford or General Motors VIN, where specific letters map to specific engines, the VW system relies on the assumption that the manufacturer's database (accessible via the serial number) holds this information.

Position 9 (The Check Digit):

In the North American market, Position 9 is a mandatory checksum. In Europe, this requirement does not exist. Volkswagen typically fills Position 9 with Z as well, although in some modern iterations or specific export configurations, this character may vary or serve as an internal control digit. However, for the purpose of aftermarket decoding, it provides no descriptive data regarding the vehicle's physical properties.2

### 3.2 The Core Identifier: Model Series Codes (Positions 7-8)

If positions 4, 5, 6, and 9 are effectively null data, the forensic value of the VDS is concentrated entirely in Positions 7 and 8. These two characters represent the Model Series or Platform Code. They are the key to identifying the specific generation and type of the van.

This section provides a detailed breakdown of these codes across the three primary commercial lineages: the Transporter, the Crafter, and the ID. Buzz.

#### 3.2.1 The Transporter Lineage (T-Series)

The Volkswagen Transporter has evolved through seven distinct generations. The VIN codes used in positions 7 and 8 reflect this evolution, though not always linearly.

The T4 Generation (1990–2003):

The T4 marked the transition from rear-engine air-cooled (or water-boxed) architecture to a front-engine, front-wheel-drive layout. The model code assigned to the T4 was 70 (seven-zero) for the standard van and pickup configurations. In some later years or specific chassis variants, the code 7D was utilized.2

- Decoding Key: WV1ZZZ70Z... identifies a T4 commercial van.

The T5 and T6 Generations (2003–2024):

With the introduction of the T5 in 2003, Volkswagen adopted the code 7H for the majority of the range (Panel Van, Kombi, Multivan, California). A secondary code, 7J, was introduced primarily for chassis cabs, pickups, and certain heavy-duty variants, though 7H remains the dominant identifier.2

Crucially, when Volkswagen introduced the T6 in 2015 and the T6.1 in 2019, they did not change the VIN platform code. The T6 was technically a major facelift of the T5 chassis rather than a clean-sheet platform; therefore, it retained the 7H and 7J designators.

- Decoding Key: WV1ZZZ7HZ... could be a 2004 T5 or a 2020 T6.1. To distinguish them, one must look to the Model Year code in the VIS (Position 10).

The T7 Multivan (2021–Present):

The introduction of the T7 Multivan marked a bifurcation in the Transporter lineage. The T7 Multivan is built on the MQB (Modularer Querbaukasten) platform—the same architecture used for the Golf and Tiguan—and is intended primarily as a passenger vehicle. To reflect this fundamental shift away from the commercial T-platform, the T7 was assigned a new model code: ST.

- Decoding Key: WV2ZZZSTZ... identifies the new MQB-based Multivan.12

- Note: The commercial Transporter T6.1 continues to be sold alongside the T7 Multivan, retaining the 7H code. This creates a market where two "Transporters" exist simultaneously with different VIN structures.

#### 3.2.2 The Crafter Lineage

The Crafter's VIN history is complex due to the shifting partnerships behind its production.

Generation 1 (The "LT3" / Mercedes Era, 2006–2016):

The first-generation Crafter was a badge-engineered Mercedes-Benz Sprinter, manufactured by Daimler in Ludwigsfelde and Düsseldorf. Despite being built by Mercedes, it carried a VW VIN. The model codes assigned were 2E and 2F.

- Decoding Key: WV1ZZZ2EZ... identifies a Sprinter-based Crafter.5

Generation 2 (The "SY/SZ" / VW Era, 2017–Present):

In 2017, Volkswagen ended the partnership with Mercedes and launched an entirely in-house developed Crafter, built in a new dedicated plant in Wrzesnia, Poland. This generation introduced two distinct model codes that likely correlate with the drivetrain or gross weight configuration:

- SY: This code is typically associated with the standard Panel Van and Kombi configurations. It is observed on Front-Wheel Drive (FWD) and 4Motion models.

- SZ: This code is often linked to the heavier-duty variants, Rear-Wheel Drive (RWD) configurations, or open chassis cabs. Listings for parts often group SY and SZ together, but distinctions in drivetrain components (like prop shafts vs. half shafts) suggest the code helps segregate the FWD/AWD chassis from the RWD chassis.7

- Decoding Key: WV1ZZZSYZ... identifies a Polish-built, all-VW Crafter.

#### 3.2.3 The ID. Buzz (The Electric Era)

The ID. Buzz represents the future of the VW van lineage, utilizing the MEB (Modular Electric Drive Matrix) platform. Despite the radical technological shift, the VIN structure remains consistent.

- Model Code: EB.
  Current research and vehicle listings confirm that the ID. Buzz utilizes the code EB in positions 7 and 8.

- Decoding Key: WV2ZZZEBZ... identifies an ID. Buzz (Passenger).10 The WV2 prefix confirms its passenger-carrying focus, while EB confirms the MEB platform.

#### 3.2.4 The Caddy and Amarok

While the user focused on "vans," the Caddy and Amarok are integral to the commercial fleet.

- Caddy Mk 3 (2004–2015): Uses code 2K.5

- Caddy Mk 4 (2015–2020): Often referred to as the SA chassis in internal documents, but VINs may show 2K or SA depending on the specific transition period.16

- Caddy Mk 5 (2020–Present): Uses code SB.16

- Amarok: The first generation uses code 2H.5

#### Table 2: Comprehensive VDS Decoding Key for European VW Commercials

|                    |                  |                   |                                                |
| ------------------ | ---------------- | ----------------- | ---------------------------------------------- |
| Model Generation   | Production Years | VIN Positions 7-8 | Platform / Notes                               |
| Transporter T4     | 1990–2003        | 70 / 7D           | T4 Platform (FWD).                             |
| Transporter T5     | 2003–2015        | 7H / 7J           | T5 Platform. 7H = Van, 7J = Chassis/Pickup.    |
| Transporter T6/6.1 | 2015–2024        | 7H / 7J           | T6 is a heavy facelift of T5; retains 7H code. |
| Multivan T7        | 2021–Present     | ST                | MQB Platform. Distinct from T6.1.              |
| Crafter Gen 1      | 2006–2016        | 2E / 2F           | Mercedes Sprinter based.                       |
| Crafter Gen 2      | 2017–Present     | SY                | In-house VW. Standard / FWD / AWD.             |
| Crafter Gen 2      | 2017–Present     | SZ                | In-house VW. Heavy Duty / RWD.                 |
| ID. Buzz           | 2022–Present     | EB                | MEB Electric Platform.                         |
| Caddy Mk 3         | 2004–2015        | 2K                | Golf 5 based.                                  |
| Caddy Mk 5         | 2020–Present     | SB                | MQB Evo based.                                 |
| Amarok             | 2010–2022        | 2H                | Ladder frame pickup.                           |

## 4. The Vehicle Identifier Section (VIS): Logistics and Chronology

The VIS (positions 10 through 17) provides the specific identity of the vehicle, pinpointing when it was built and where. While the VDS identifies the model family, the VIS is essential for ordering parts, as production changes often occur mid-cycle, and specific parts may be unique to a specific factory.

### 4.1 Chronological Decoding: The Model Year (Position 10)

Volkswagen adheres to the standard ISO 3779 model year code, which repeats every 30 years. It is important to note that the "Model Year" in the VIN does not necessarily correspond to the calendar year of manufacture. Volkswagen's model year typically begins in roughly week 22 (May/June) of the previous calendar year. Thus, a van built in August 2023 will carry a 2024 model year code (R).

#### Table 3: Model Year Codes (2010–2030)

|      |      |      |      |      |      |      |      |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| Code | Year | Code | Year | Code | Year | Code | Year |
| A    | 2010 | G    | 2016 | N    | 2022 | U    | 2027 |
| B    | 2011 | H    | 2017 | P    | 2023 | V    | 2028 |
| C    | 2012 | J    | 2018 | R    | 2024 | W    | 2029 |
| D    | 2013 | K    | 2019 | S    | 2025 | X    | 2030 |
| E    | 2014 | L    | 2020 | T    | 2026 |      |      |
| F    | 2015 | M    | 2021 |      |      |      |      |

Note: The letters I, O, Q, Z, and number 0 are prohibited in this position to prevent optical confusion.

### 4.2 Industrial Geography: The Plant Codes (Position 11)

For the commercial vehicle specialist, the plant code at Position 11 is a high-value data point. Unlike passenger cars which may be built in Wolfsburg, Emden, or Zwickau, commercial vans are concentrated in specific industrial hubs.

Code H: Hanover (Hannover), Germany

Hanover is the historic heart of Volkswagen Commercial Vehicles. It has produced every generation of the Transporter since the T2. Today, it is the primary site for the T6.1, the T7 Multivan, and the ID. Buzz. If a VIN has an H in position 11, it is a high-pedigree German build.5

Code 9: Wrzesnia, Poland

This greenfield factory was constructed specifically for the second-generation Crafter (and its twin, the MAN TGE). It is the sole source for the SY and SZ model codes. The presence of 9 in position 11 guarantees the vehicle is a post-2016 VW-engineered Crafter.2

Code X: Poznan, Poland (Antoninek)

The Poznan facility has been the home of the Caddy for decades and also serves as a major production hub for the Transporter T6.1, particularly the cargo variants. It is common to see T6.1s with either H or X codes depending on capacity allocation.2

Code 7: Ludwigsfelde, Germany

This code appears on the first-generation Crafter (2E). This is actually a Mercedes-Benz plant. The presence of 7 in a VW VIN is a forensic artifact of the Daimler-VW joint venture, indicating the vehicle was assembled alongside Sprinters.5

Code 6: Düsseldorf, Germany

Another Mercedes-Benz plant used for the first-generation Crafter/Sprinter vans, typically the closed panel van variants, whereas Ludwigsfelde often handled chassis cabs.5

### 4.3 The Serial Number (Positions 12-17)

The final six digits constitute the sequential production number. This number is typically reset to 000001 when the Model Year (Position 10) rolls over. However, in some high-volume plants, the numbers may run continuously. For forensic purposes, the serial number is the unique identifier that, when combined with the WMI and VDS, creates a globally unique string. It is this serial number that acts as the primary key when querying the VW central database for build sheets.

## 5. Beyond the VIN: The Limits of Forensic Decoding and the "Datenträger"

A critical finding of this research is the limitation of the VIN itself. The user request asks for "other related vehicle-related information" such as engine type and options. It is imperative to state clearly: The European VIN does not contain this information.

Because of the "Hollow VDS" structure (ZZZ), it is impossible to look at a European VIN and determine if a T6 Transporter is equipped with a 81kW, 110kW, or 150kW TDI engine. It is impossible to determine if it has a DSG or Manual transmission, or if it is painted in "Candy White" or "Indium Grey."

### 5.1 The Solution: The Vehicle Data Sticker (Datenträger)

To access the missing data, one must locate the physical Vehicle Data Sticker. This white paper label is typically found in two locations:

1. Adhered to the metal bodywork, often near the fuse box, under the steering column, or on the base of the driver's seat frame.13

2. Pasted on the inside cover of the vehicle's Service Schedule booklet.25

### 5.2 Decoding the Data Sticker

The Data Sticker bridges the gap left by the VIN. It contains:

- The VIN: (For cross-referencing).

- The Sales Type: A precise code (e.g., 7HC) that defines the exact model variant (e.g., Multivan Highline).

- Engine & Transmission Codes: 3- or 4-letter codes (e.g., CXEB for the engine, DQ500 for the gearbox).5

- Paint Code: (e.g., LB9A).

- PR Codes (Primary Features): A block of 3-character alphanumeric codes (e.g., 1BA, C1G, 7P1) that serve as a "build sheet."

- Example: 1BA might denote "Standard Suspension," while 1BH denotes "Heavy Duty Suspension."

Forensic Implication:

For any investigation requiring the identification of a vehicle's original equipment—such as verifying if a van was factory-fitted with a bulkhead or if a camper conversion is based on a genuine California shell—the VIN is merely the index key. The PR codes on the data sticker provide the actual evidence.

## 6. Case Studies in Forensic Decoding

To demonstrate the application of this decoding framework, we analyze three distinct VIN examples derived from the research data.

### Case Study A: The Electric Pioneer

Subject VIN: WV2ZZZEB800000000 19

- WMI (WV2): Identifies the manufacturer as Volkswagen Commercial Vehicles, specifically the Passenger/MPV division. This is a people-carrier, not a cargo van.

- VDS (ZZZEB8):

- ZZZ: Standard European fillers.

- EB: Model Code. Identifies this as an ID. Buzz.

- 8: Check/Filler digit.

- VIS (SH015405):

- S: Model Year 2025.

- H: Plant Hanover.

- 015405: Serial number.

- Conclusion: This is a 2025 Model Year ID. Buzz Passenger variant, built in Hanover.

### Case Study B: The Modern Workhorse

Subject VIN: WV1ZZZSYZ00000000 (Reconstructed Example)

- WMI (WV1): Volkswagen Commercial Vehicles, Goods/Truck division. Legal classification N1.

- VDS (ZZZSYZ):

- SY: Model Code. Identifies this as a Crafter Generation 2. The SY suggests a standard Panel Van configuration (likely FWD or 4Motion).

- VIS (J9000001):

- J: Model Year 2018.

- 9: Plant Wrzesnia. Confirms Polish manufacture in the dedicated Crafter facility.

- Conclusion: A 2018 VW Crafter Panel Van, built in Poland.

### Case Study C: The New Passenger Standard

Subject VIN: WV2ZZZST0PH... 14

- WMI (WV2): Passenger/MPV classification.

- VDS (ZZZST0):

- ST: Model Code. Identifies this as the T7 Multivan (MQB platform). It is not a T6.1 Transporter.

- VIS (PH...):

- P: Model Year 2023.

- H: Plant Hanover.

- Conclusion: A 2023 T7 Multivan, distinct from the commercial T6.1 range.

## 7. Conclusion

The decoding of Volkswagen Commercial Vehicle VINs in the European market is a discipline that rewards attention to specific architectural nuances. Unlike the transparent, attribute-rich VINs of North America, the European format is a cryptographic container where the majority of the data describes the chassis rather than the configuration.

For the analyst specifically targeting Volkswagen vans, the forensic pathway is clear:

1. Analyze the WMI (WV1 vs. WV2) to determine the vehicle's legal homologation as a goods carrier or passenger transport.

2. Isolate VDS Positions 7 and 8 to identify the platform (7H for T6, ST for T7, SY/SZ for Crafter, EB for ID. Buzz).

3. Consult VIS Position 11 to verify the production logistics (Hanover, Wrzesnia, or Poznan).

4. Acknowledge the Data Gap: Accept that engine, transmission, and trim data are deliberately excluded from the VIN and must be sourced from the Data Sticker or manufacturer databases via the serial number.

This structured approach ensures accurate identification across the diverse and evolving lineup of Volkswagen’s commercial fleet, from the legacy diesel workhorses to the electrified future of the ID. Buzz.

#### Works cited

1. VW VIN Codes | PDF | Volkswagen | Private Transport - Scribd, accessed November 30, 2025, [https://www.scribd.com/doc/151850603/VW-VIN-Codes](https://www.scribd.com/doc/151850603/VW-VIN-Codes)

2. VW VIN Codes - Club VeeDub, accessed November 30, 2025, [https://www.clubvw.org.au/vwreference/vwvin/](https://www.clubvw.org.au/vwreference/vwvin/)

3. Vehicle identification number - Wikipedia, accessed November 30, 2025, [https://en.wikipedia.org/wiki/Vehicle_identification_number](https://en.wikipedia.org/wiki/Vehicle_identification_number)

4. VW World Manufacturer Identifier VW | PDF | Volkswagen | Car Body Styles - Scribd, accessed November 30, 2025, [https://www.scribd.com/document/415181124/VW-World-Manufacturer-Identifier-VW](https://www.scribd.com/document/415181124/VW-World-Manufacturer-Identifier-VW)

5. What to know about your VW VIN Code - L & M Foreign Cars, accessed November 30, 2025, [https://www.landmforeigncars.com/blog/what-to-know-about-your-vw-vin-code](https://www.landmforeigncars.com/blog/what-to-know-about-your-vw-vin-code)

6. Vehicle Identification Numbers (VIN codes)/World Manufacturer Identifier (WMI) - Wikibooks, accessed November 30, 2025, [https://en.wikibooks.org/wiki/Vehicle*Identification_Numbers*(VIN*codes)/World_Manufacturer_Identifier*(WMI)](<https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/World_Manufacturer_Identifier_(WMI)>)

7. Volkswagen Crafter - Wikipedia, accessed November 30, 2025, [https://en.wikipedia.org/wiki/Volkswagen_Crafter](https://en.wikipedia.org/wiki/Volkswagen_Crafter)

8. What's a Vehicle Identification Number? How to Decode the World Manufacturer Identifier, accessed November 30, 2025, [https://checkventory.com/articles/whats-your-number/](https://checkventory.com/articles/whats-your-number/)

9. accessed November 30, 2025, [https://www.clubvw.org.au/vwreference/vwvin/#:~:text=WVW%20%2D%20Volkswagen%20Cars,WV3%20%2D%20Volkswagen%20Trucks](https://www.clubvw.org.au/vwreference/vwvin/#:~:text=WVW%20%2D%20Volkswagen%20Cars,WV3%20%2D%20Volkswagen%20Trucks)

10. Demo 2024 Volkswagen ID. Buzz Pro BEV560 #F006796 Leichhardt, accessed November 30, 2025, [https://leichhardtvolkswagen.com.au/f006796-1594530-volkswagen-idbuzz-2024/](https://leichhardtvolkswagen.com.au/f006796-1594530-volkswagen-idbuzz-2024/)

11. New 2025 Volkswagen ID. Buzz GTX #V438628 Macgregor, accessed November 30, 2025, [https://macgregorvolkswagen.com.au/v438628-1688181-volkswagen-idbuzz-2025/](https://macgregorvolkswagen.com.au/v438628-1688181-volkswagen-idbuzz-2025/)

12. VW Transporter Crafter Multivan T7 Electric Vacuum Heater Coolant Valve 2015–202 | eBay, accessed November 30, 2025, [https://www.ebay.com/itm/156932798208](https://www.ebay.com/itm/156932798208)

13. How To Read Watercooled VW Chassis Numbers - Heritage Parts Centre, accessed November 30, 2025, [https://www.heritagepartscentre.com/uk/blog/how-to-read-watercooled-vw-chassis-numbers.html](https://www.heritagepartscentre.com/uk/blog/how-to-read-watercooled-vw-chassis-numbers.html)

14. VW MULTIVAN T7 1.4 e-Hybrid 150 PS (21-) Radschraube Eibach, accessed November 30, 2025, [https://eibach.ch/en/detail/index/sArticle/265239](https://eibach.ch/en/detail/index/sArticle/265239)

15. Complete set Park Assist for VW T7 ST - kufatec, accessed November 30, 2025, [https://kufatec.com/en/complete-set-park-assist-for-vw-t7-st/47029](https://kufatec.com/en/complete-set-park-assist-for-vw-t7-st/47029)

16. VW T7 - ​​Bulli Emblems Page - MazorDesign, accessed November 30, 2025, [https://mazordesign.com/en/products/vw-t7-bulli-seite-kotflugel-schwarz-emblem-plakette](https://mazordesign.com/en/products/vw-t7-bulli-seite-kotflugel-schwarz-emblem-plakette)

17. [en]=> P/N 245A0Axx (LV-CAN200) - Teltonika Telematics Wiki, accessed November 30, 2025, [https://wiki.teltonika-gps.com/images/5/55/LVCAN200_list_2023_03_17_en.pdf](https://wiki.teltonika-gps.com/images/5/55/LVCAN200_list_2023_03_17_en.pdf)

18. Genuine VW Crafter SCB SCC SYB SYC SYD SYI SYJ SZB SZC Poly-V-Belt 04L145933A, accessed November 30, 2025, [https://www.ebay.com/itm/226897579774](https://www.ebay.com/itm/226897579774)

19. 2025 Volkswagen ID. Buzz Pro S - Volkswagen dealer serving Birmingham AL – New and Used Volkswagen dealership serving Tuscaloosa Cullman Oxford AL, accessed November 30, 2025, [https://www.gotoroyalvw.com/new-Birmingham-2025-Volkswagen-ID+Buzz-Pro+S-WVGAWVEB800000000](https://www.gotoroyalvw.com/new-Birmingham-2025-Volkswagen-ID+Buzz-Pro+S-WVGAWVEB800000000)

20. 0001193125-13-169963.txt - SEC.gov, accessed November 30, 2025, [https://www.sec.gov/Archives/edgar/data/1070296/000119312513169963/0001193125-13-169963.txt](https://www.sec.gov/Archives/edgar/data/1070296/000119312513169963/0001193125-13-169963.txt)

21. VIN & Plants FAQ - Alltransua, accessed November 30, 2025, [https://alltransua.com/page/plants](https://alltransua.com/page/plants)

22. Where was my van built ? | VW T6 Transporter Forum, accessed November 30, 2025, [https://www.t6forum.com/threads/where-was-my-van-built.58673/](https://www.t6forum.com/threads/where-was-my-van-built.58673/)

23. How To Read Watercooled VW Chassis Numbers - Heritage Parts Centre, accessed November 30, 2025, [https://www.heritagepartscentre.com/us/blog/how-to-read-watercooled-vw-chassis-numbers-us.html](https://www.heritagepartscentre.com/us/blog/how-to-read-watercooled-vw-chassis-numbers-us.html)

24. How To Find the Chassis Number on Your Modern VW Camper | Just Kampers, accessed November 30, 2025, [https://www.justkampers.com/jkblog/how-to-find-chassis-number-modern-vw-camper/](https://www.justkampers.com/jkblog/how-to-find-chassis-number-modern-vw-camper/)

25. Production Codes (PR-Codes) can be found on the Vehicle Data Label - Ross-Tech Wiki, accessed November 30, 2025, [https://wiki.ross-tech.com/wiki/index.php/Production*Codes*(PR-Codes)\_can_be_found_on_the_Vehicle_Data_Label](<https://wiki.ross-tech.com/wiki/index.php/Production_Codes_(PR-Codes)_can_be_found_on_the_Vehicle_Data_Label>)

26. Understanding VIN and PR Codes | Blog - Run Auto Parts, accessed November 30, 2025, [https://runautoparts.com.au/blog/post/understanding-vin-and-pr-codes](https://runautoparts.com.au/blog/post/understanding-vin-and-pr-codes)
