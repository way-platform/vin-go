# The Taxonomy of Transport: A Definitive Technical Reference on Commercial Vehicle Identification, VIN Decoding Architectures, and Model Designation Systems for the European and Global Markets

## 1. Introduction: The Geopolitical and Engineering Architecture of Vehicle Identity

The commercial vehicle sector operates on a level of technical complexity that far exceeds the passenger car market. While a sedan is typically manufactured as a finished good with a static identity, a commercial vehicle—be it a heavy-duty tractor unit, a medium-duty distribution truck, or a light commercial van—is frequently a modular platform. It is an incomplete canvas upon which third-party bodybuilders, upfitters, and equipment manufacturers inscribe the final purpose of the machine. Consequently, the identity of a commercial vehicle is not a singular data point but a layered construct, anchored by the Vehicle Identification Number (VIN) assigned by the Original Equipment Manufacturer (OEM).

This report provides an exhaustive, expert-level analysis of the methodologies used to decode these identities. It focuses primarily on the Daimler Truck and Mercedes-Benz Vans product portfolio, the dominant force in the European commercial sector, while also providing a granular comparative analysis of its primary competitors: Volvo Trucks, Scania, MAN, DAF, Iveco, Renault Trucks, and Ford. The analysis is grounded in the tension between the two dominant global standards—ISO 3779/3780 in the European Union and FMVSS 565 in North America—and explores how these regulatory frameworks influence data retrieval, parts logistics, and the technical integration of bodywork.

### 1.1 The Regulatory Divergence: ISO 3779 versus FMVSS 565



At the foundational level, global vehicle identification is governed by a bifurcated regulatory landscape. Understanding this divergence is critical for any analyst, fleet manager, or parts supplier operating across borders, as it dictates which data points can be reliably extracted from a VIN and which remain opaque without access to proprietary OEM databases.

The International Organization for Standardization (ISO) established Standard 3779 to harmonize identification systems.1 This standard mandates a 17-character alphanumeric code divided into three specific sections: the World Manufacturer Identifier (WMI), the Vehicle Descriptor Section (VDS), and the Vehicle Identifier Section (VIS).2 However, the application of this standard varies significantly between the European Single Market and the North American Free Trade Agreement (NAFTA/USMCA) zone.



#### 1.1.1 The Check Digit Anomaly



In North America, under 49 CFR Part 565, the National Highway Traffic Safety Administration (NHTSA) strictly mandates that Position 9 of the VIN serves as a "Check Digit." This character is the result of a mathematical modulus 11 algorithm involving weighted values assigned to the other 16 characters.3 Its primary function is fraud detection and transcription error prevention.

In the European Union, however, adherence to the Check Digit requirement is voluntary. While some manufacturers, such as Volvo Trucks, voluntarily align their global production with the check digit system to streamline global database management 5, many others—most notably German manufacturers like Volkswagen and Daimler—frequently utilize Position 9 for other purposes in their domestic markets. They may employ it as a "fill-in" character (often 'Z') or to encode specific internal attributes unrelated to validation checksums.6 Consequently, standard VIN validation software designed for the US market will frequently flag valid European commercial vehicle VINs as "invalid" or "corrupt," creating significant friction in global fleet management systems.7



#### 1.1.2 The Model Year Ambiguity



A second critical divergence lies in the designation of the Model Year. FMVSS 565 mandates that Position 10 of the VIN identifies the Model Year according to a federally standardized 30-year alphanumeric cycle (e.g., 'K' for 2019, 'L' for 2020).8

In Europe, there is no such legal requirement for the VIN to encode the model year explicitly. European registration documents (such as the V5C in the UK or the Zulassungsbescheinigung in Germany) rely on the "Date of First Registration" rather than a manufacturing model year code. As a result, European OEMs often use Position 10 for internal logistical codes, such as a plant code overflow or a model generation identifier, rather than a calendar year.6 This creates a phenomenon known as "year lag," where a chassis manufactured in late 2019 (according to internal production records) might not be registered until mid-2021 after extensive bodywork application. Relying on US-centric decoders that interpret Position 10 as a year code can lead to gross misidentification of the vehicle's age and subsequent valuation errors.8



### 1.2 The Anatomy of the VIN Structure



To navigate this complexity, one must understand the tripartite structure of the VIN as applied to commercial assets.



|   |   |   |   |
|---|---|---|---|
|Section|Positions|Function|EU Commercial Vehicle Context|
|WMI|1-3|Manufacturer & Origin|Defines the OEM entity. Critical for distinguishing between passenger (e.g., WDD) and commercial (e.g., WDF) divisions within the same conglomerate.10|
|VDS|4-9|Vehicle Attributes|The "Genetic Code." For Mercedes, this holds the Baumuster. For Volvo/Renault, it indicates chassis height, axle config, and engine family.11|
|VIS|10-17|Unique Identifier|Contains the plant code and sequential serial number. For Scania and DAF, the final 7-8 digits here form the "Chassis Number," the primary index for all technical support.12|



## 2. The Daimler Paradigm: Decoding Mercedes-Benz Commercial Vehicles



Daimler Truck AG and Mercedes-Benz Vans utilize a highly systematic, engineering-driven approach to identification known as the Baumuster (Model Type) system. Unlike American manufacturers who might encode trim levels or marketing names directly into the VIN, Daimler uses the VIN as a reference key to an internal engineering taxonomy. Mastering the Baumuster code is the single most important skill for identifying Mercedes-Benz commercial vehicles.13



### 2.1 World Manufacturer Identifier (WMI) Evolution



The WMI codes for Daimler have evolved in lockstep with the company's complex corporate history—from Daimler-Benz AG to DaimlerChrysler, to Daimler AG, and finally the split into Mercedes-Benz Group AG (cars/vans) and Daimler Truck AG. Identifying the WMI is the first step in categorizing the asset.

- WDB: Historically the standard code for all Daimler-Benz vehicles (Germany). It remains common on older trucks and vans. Insight: A VIN starting with WDB indicates a vehicle produced by the unified Daimler entity, often pre-dating the strict corporate separation of the truck division.7

- WDF: Specific to Mercedes-Benz Commercial vehicles (Germany). This code is increasingly used to segregate vans and trucks from the passenger car lineage.10

- WDD: Assigned primarily to Daimler AG passenger vehicles and SUVs. If a "commercial" vehicle like a Vito carries a WDD code, it is likely a V-Class passenger variant rather than a panel van.10

- W1W, W1X, W1Y, W1Z: These codes denote Mercedes-Benz vehicles manufactured in the United States, specifically at the Charleston, South Carolina assembly plant.


- W1W: Multi-Purpose Vehicle (MPV).

- W1X: Incomplete Vehicle (Chassis Cab awaiting bodywork).

- W1Y: Truck (Cargo Van).

- W1Z: Bus.14


- WD3: Used for Sprinter and Metris trucks, indicating a specialized commercial designation often found on gliders or stripped chassis.7

- WD4: Often associated with Sprinter vans.14

- VSA: Indicates production at the Vitoria-Gasteiz plant in Spain, the global hub for the Vito and V-Class (W447/W639).16

- 9BM: Indicates production by Mercedes-Benz do Brasil, a major hub for heavy-duty trucks (Accelo, Atego) and bus chassis for the Latin American market.7




### 2.2 The Baumuster System: The DNA of the Vehicle



The Baumuster is encoded in VIN positions 4, 5, and 6. This three-digit numeric code (sometimes alphanumeric in recent US models) determines the exact technical generation and model series of the vehicle. It allows a technician to distinguish between visually similar vehicles that have completely different underpinnings (e.g., a 2005 Sprinter vs. a 2007 Sprinter).



#### 2.2.1 The Sprinter Lineage (Transporter Class)



The Sprinter, Daimler's flagship van, has passed through three distinct generations, each identified by a specific range of Baumuster codes.

- Generation 1 (T1N): 1995 – 2006


- Codes: 901, 902, 903, 904, 905.

- Technical Context: These codes roughly correspond to the Gross Vehicle Weight Rating (GVWR) class. A 903 indicates a 3-ton class vehicle, while a 904 indicates a 4-ton class with dual rear wheels. This generation is characterized by the OM601/OM602/OM611 engines.7

- Decoding Example: WDB903... identifies a T1N Sprinter 3-series (e.g., 311 CDI, 313 CDI).


- Generation 2 (NCV3 / New Concept Van 3): 2006 – 2018


- Code: 906.

- Technical Context: The introduction of the 906 code marked a complete platform redesign. This chassis introduced the OM642 V6 diesel and the widespread use of the ADAPTIVE ESP system. Identifying a 906 is crucial because parts from the T1N (901-905) are almost entirely incompatible.14

- Decoding Example: WDB906633... decodes to a Sprinter NCV3 (906), panel van (6), standard wheelbase (33).


- Generation 3 (VS30): 2018 – Present


- Codes: 907 and 910.

- The FWD Revolution: The split into two Baumuster codes for the VS30 generation is a critical development.


- 907: Designated for Rear-Wheel Drive (RWD) and All-Wheel Drive (AWD) configurations. This carries over the heavy-duty heritage of the Sprinter.17

- 910: Designated for Front-Wheel Drive (FWD) configurations. This was a strategic move to lower the loading floor and increase payload for the light logistics and camper van markets.


- Implication: An upfitter looking to install a rear underfloor liftgate must check the VIN immediately. If it reads W1V910..., the rear axle structure is fundamentally different from the 907, lacking the differential housing, which alters mounting points significantly.17




#### 2.2.2 The Mid-Size Sector: Vito, Viano, V-Class



The mid-size van segment follows a similar evolutionary coding structure.

- Generation 1 (W638): 1996 – 2003


- Code: 638.

- Context: The first front-wheel-drive van from Mercedes, produced in Vitoria, Spain.


- Generation 2 (W639): 2003 – 2014


- Code: 639.

- Context: A return to Rear-Wheel Drive (and 4x4). Sold as the "Vito" (Commercial) and "Viano" (Passenger).

- Decoding Note: Position 4-6 will show 639. This is the definitive identifier for the NCV2 platform.19


- Generation 3 (W447): 2014 – Present


- Code: 447.

- Context: Marketed as the Vito (Commercial), V-Class (Luxury Passenger), and Metris (North America). Even though the US market badges the vehicle as "Metris," the VIN retains the global 447 Baumuster code, ensuring global parts commonality.20




#### 2.2.3 Compact and Pickup Segments



- Citan: Codes 415 (Mark 1) and 420 (Mark 2). Based on the Renault Kangoo platform.

- X-Class: Code 470. A pickup truck based on the Nissan Navara platform. The 470 code is essential for distinguishing Mercedes-specific parts (body panels, interior) from Nissan shared components.7




### 2.3 The Heavy Truck Nomenclature: Actros, Arocs, Atego



Daimler Truck AG reorganized its heavy truck identification system significantly with the introduction of the Euro VI emission standards in 2011. This shift, often referred to as the "New Truck Generation" (NTG), introduced new Baumuster codes that separate vehicles by application rather than just tonnage.



#### 2.3.1 The New Generation (2011–Present)



- Baumuster 963: Covers both the Actros (Long-haul) and the Antos (Heavy-duty distribution).


- Differentiation: Since both share the 963 code, decoding the specific model requires looking at the cab variant codes or the sales designation. The Actros typically features wider, sleeper cabs, while the Antos features narrower, day cabs, but they share the same chassis architecture.22

- Insight: The 963 platform introduced the "common architecture" shared components, streamlining bodybuilder integration via the programmable special module (PSM).


- Baumuster 964: Designated for the Arocs.


- Application: The Arocs is the construction and off-road specialist. While it shares engines (OM470/471/473) with the Actros, the 964 chassis is engineered for higher torsional rigidity, greater ground clearance, and uses steel bumpers and distinct axle configurations (e.g., planetary hub reduction axles). Differentiating a 964 from a 963 is vital for ordering suspension and frame parts.7


- Baumuster 967: Designated for the Atego (Euro VI).


- Application: The light-to-medium duty distribution truck (6.5t – 16t). It uses the OM934/OM936 engines.22


- Baumuster 956: Designated for the Econic.


- Application: Low-entry cab for municipal waste collection and airport services. The 956 code identifies the unique low-frame architecture essential for high-visibility urban safety standards.25




#### 2.3.2 Legacy Truck Models (Pre-2011)



For fleets operating older assets, identifying the legacy codes is still relevant.

- Actros MP1 (1996-2002): Codes 950, 952, 953, 954.22

- Actros MP2/MP3 (2002-2011): Codes 930, 932, 933, 934.22

- Axor: Codes 940, 944, 950, 952. The Axor was a hybrid, using the Actros cab (narrowed) on a lighter chassis.7

- Legacy Atego: Codes 970, 972, 974, 975, 976.23




### 2.4 The Unimog: The Universal Implement Carrier



The Unimog (Universal-Motor-Gerät) is a unique vehicle class with its own dedicated Baumuster sequences (400-series). These codes are legendary in the off-road community and define the vehicle's generation and capability.

- Baumuster 404: The classic Unimog S (1955-1980). Primarily a gasoline-powered military vehicle. Decoding logic: 404.114.26

- Baumuster 406: The iconic mid-range Unimog (1963-1989). This series (U65-U900) introduced the diesel OM352 engine. It is the most common classic Unimog found in agricultural service.26

- Baumuster 437: The heavy series (U1700, U2150, U2450). These are high-capacity, high-mobility trucks often used as expedition vehicles or heavy utility carriers.26

- Baumuster 405: The modern Implement Carrier (U200 - U500 series). Designed for municipal work with rapid interchangeability of tools.

- Baumuster 437.4: The modern High Mobility series (U4000/U5000).


Decoding Insight: A VIN beginning with WDB406120... breaks down as: Manufacturer (Mercedes), Model (Unimog 406), Variant (120 = Cabriolet/Open Cab, Short Wheelbase). This level of detail allows restorers to source the correct axles and transmissions without needing the data card.26



### 2.5 Decoding the Engine and Axle from the VIN



While the Baumuster (Positions 4-6) gives the model series, the subsequent digits in the VDS (Positions 7-9) often provide specific attributes regarding the wheelbase, engine type, and steering configuration, though this varies by era.

- Steering (Position 10): Unlike the US "Year Code," European Mercedes commercial VINs often use Position 10 to denote the steering layout.


- 1 = Left Hand Drive (LHD).

- 2 = Right Hand Drive (RHD).


- Plant Code (Position 11): Critical for identifying the manufacturing origin.


- P: Düsseldorf (Sprinter Panel Vans).

- N: Ludwigsfelde (Sprinter Chassis Cabs, Vario).

- V: Vitoria (Vito/V-Class).

- L: Wörth am Rhein (Actros, Arocs, Atego, Unimog). The largest truck plant in the world.

- K: Graz (G-Class, some Unimogs).10




## 3. The Swedish Titans: Volvo Trucks and Scania



The identification philosophy of the Swedish manufacturers, Volvo and Scania, differs fundamentally from the German approach. Here, the concept of the "Chassis Number" serves as the primary key for vehicle identity, often superseding the VIN in workshop environments.



### 3.1 Volvo Trucks: The YV2 Identifier



Volvo Trucks maintains a distinct identity from Volvo Cars, a separation codified in the WMI system.

- WMI Codes:


- YV2: Volvo Trucks (Manufactured in Sweden/Europe).5

- YV3: Volvo Buses.27

- 4V4: Volvo Trucks North America (manufactured in the US).9

- YV5: Incomplete Vehicles (Glider kits/Chassis). This code is crucial for bodybuilders, as it flags the vehicle as requiring final stage certification.27




#### 3.1.1 Decoding the Volvo Structure (EU Market)



Volvo utilizes the Vehicle Descriptor Section (VDS, Positions 4-9) to encode technical specifications, though the key depends on the era.

- Positions 1-3: YV2 (Manufacturer).

- Position 4: Often denotes the Cab Type or Chassis Height (e.g., 'H' for High, 'L' for Low, 'T' for Tractor).

- Position 5: Axle Configuration and Brake Type (e.g., 'A' = 4x2 Rigid, 'B' = 6x2 Rigid).

- Positions 6-7: Engine Code. This correlates to the engine family (e.g., D11, D13, D16) and Euro standard level.

- Positions 10-17: The Chassis Number.


- Operational Insight: In the Volvo "Impact" parts and service system, the technician rarely inputs the full VIN. Instead, the last 6 or 7 digits (the Chassis Number) are entered. This number is unique and pulls the complete "Build Identity" of the truck, listing every single component installed at the factory.5




### 3.2 Scania: The Supremacy of the Chassis Number



Scania represents the most modular of all truck manufacturers. Their production philosophy allows for millions of permutations using a limited number of shared components. Consequently, the VIN VDS (Positions 4-9) is often less descriptive than the specific "Type Designation" found on the cab plate.

- WMI Codes:


- YS2: Scania Trucks (Södertälje, Sweden/Europe).28

- YS4: Scania Buses.29

- XLE: Scania (Assembly in the Netherlands).




#### 3.2.1 The Scania Identification Logic



- The VIN VDS (Positions 4-9): In the EU, these digits often define the cab series (P, G, R, S) and engine output, but due to the immense modularity, they are not the primary reference for parts.

- The Chassis Number: Located in the VIS (Positions 11-17), specifically the last 7 digits.


- Workshop Primacy: Scania technicians rely almost exclusively on the 7-digit chassis number. Entering this into the "Scania Multi" software retrieves the exact Bill of Materials (BOM).

- Physical Location: Scania is notable for stamping the chassis number prominently on the right-hand frame rail, often near the front wheel arch. This physical verification is a mandatory step in annual technical inspections (MOT/TÜV).30




#### 3.2.2 The Data Access Legal Precedent



Scania has been central to the debate over VIN data access in Europe. In the landmark case C-319/22, the European Court of Justice ruled against Scania, mandating that vehicle manufacturers must provide repair and maintenance information (RMI) in a machine-readable format to independent operators, searchable via VIN. Scania's previous method, which required manual searches or did not allow bulk VIN processing, was deemed insufficient under Regulation (EU) 2018/858. This ruling solidifies the VIN as a regulated "digital key" to vehicle data, ensuring that independent workshops have the same access as franchised dealers.31



## 4. The Continental Competitors: MAN, Iveco, DAF, and Renault




### 4.1 MAN Truck & Bus (WMA)



MAN (Maschinenfabrik Augsburg-Nürnberg) is a key part of the TRATON Group (alongside Scania).

- WMI: WMA (Germany).12

- Structure:


- VDS (Positions 4-9): Describes the vehicle type, such as TGX, TGS, TGL, or TGM.

- Decoding Nuance: MAN utilizes a "Basic Vehicle Number" (Grundtypmuster) which is linked to the VIN. For bodybuilders, accessing the MANTED (MAN Technical Data) portal using the VIN is essential to retrieve the specific PTO (Power Take-Off) parameterization and frame hole patterns. The VIN alone does not carry the geometric detail required for mounting a crane or mixer.33




### 4.2 DAF Trucks (XLR)



DAF (Van Doorne's Automobielfabriek), a PACCAR company, uses a straightforward identification system.

- WMI:


- XLR: DAF Trucks (Netherlands).34

- SFA: Leyland DAF (United Kingdom - manufacturing LF/CF series).


- Decoding Logic:


- Position 4: Line (e.g., 'T' for Tractor, 'F' for Rigid).

- Positions 5-8: Cab and Axle configuration.

- The 8-Digit Chassis Number: DAF emphasizes the last 8 digits of the VIN (Positions 10-17) as the "DAF Chassis Number" (e.g., 0E653688). This is the standard input for the DAF RMI (Repair and Maintenance Information) system and DAF Check for parts.35




### 4.3 Iveco (ZCF)



Iveco (Industrial Vehicles Corporation) presents a complex decoding challenge due to its myriad range of light (Daily) to heavy (S-Way) vehicles.

- WMI: ZCF (Italy) is the most common, though WJM (Germany - legacy Magirus) exists.

- The "Commercial Trade Name": Unlike Mercedes, where the VIN/Baumuster is king, Iveco relies heavily on the "Commercial Trade Name" or "Model Variant" found on the B-pillar data plate for identification.


- Example: A VIN might belong to a generic chassis family, but the plate reads 35S13.


- 35: 3.5 Ton GVW.

- S: Single Rear Wheels (vs 'C' for Dual/Commercial).

- 13: 130 Horsepower engine.


- Insight: The Iveco Bodybuilder Manual explicitly directs upfitters to use this variant code alongside the VIN to determine body mounting allowances.37




### 4.4 Renault Trucks (VF6)



Renault Trucks (part of the Volvo Group) shares technology with Volvo but retains a distinct identity.

- WMI: VF6 (France).39

- Family Codes (Positions 5-6):


- 17: Magnum (Legacy).

- 24/25: Premium (Legacy).

- 34: Kerax (Construction - Legacy).

- Current Generation: The T, K, C, and D ranges use updated codes within the VDS.


- The "CAM Plate": Renault trucks feature a "CAM Plate" or "Manufacturing Plate" in addition to the VIN. This plate contains technical codes essential for the "Dialogys" or "Impact" parts systems. Decoding the VIN without the data from the CAM plate (which lists paint codes, equipment levels, and specific axle ratios) can lead to ordering errors.39




## 5. The Trans-Atlantic Divide: Global Platforms with Dual Identities



A significant challenge in VIN decoding arises when a single vehicle platform is sold globally but carries different VIN structures depending on the market. This is most evident with Ford and Volkswagen.



### 5.1 The Ford Transit: A Tale of Two VINs



The Ford Transit is a "One Ford" global product, yet its identification is region-specific.

- European Transit:


- WMI: WF0 (Ford Germany).

- Structure: Follows EU Ford logic. Position 9 is not a check digit.

- Date Coding: Ford Europe uses a proprietary rotation table for Positions 11 (Year) and 12 (Month). For example, a character 'L' might represent 2020, and 'J' might represent September. This is not the standard FMVSS position 10 year code.


- North American Transit:


- WMI: 1FT (Ford Truck USA).

- Structure: Strictly follows FMVSS 565. Position 9 is a mandatory check digit. Position 10 is the standard Year Code (e.g., 'M' = 2021).


- Implication: A VIN decoder built for the US market will reject a European Transit VIN as "Invalid" because the checksum calculation will fail, and the date decoding logic will be incorrect. Analysts must use region-specific decoders for the Transit platform.40




### 5.2 Volkswagen Commercial Vehicles



- WMI: WV1 (Commercial Vehicles), WV2 (Bus/Van), WV3 (Trucks).

- The "Crafter" Split:


- Generation 1 (Crafter 2E): This vehicle was a rebadged Mercedes-Benz Sprinter (NCV3) built by Daimler.


- VIN: Starts with WV1... but the chassis code in the VDS is 2E.

- Tech Insight: Despite the VW badge and VIN, the underlying engineering is Mercedes. Parts lookup often requires cross-referencing the Mercedes equivalent part numbers.


- Generation 2 (Crafter SY/SZ): This is an all-new, in-house VW design built in Poland.


- VIN: Chassis codes SY or SZ.

- Tech Insight: This platform shares no DNA with the Sprinter. It is mechanically identical to the MAN TGE, which is a rebadged VW Crafter. Thus, a MAN TGE VIN (WMA...) and a VW Crafter VIN (WV1...) of this generation describe the exact same machine.6




## 6. The Bodybuilder Interface: Where VIN Meets Reality



In the commercial vehicle sector, the VIN designates an "Incomplete Vehicle." The final operational asset—a refrigerated box truck, a tipper, or an ambulance—is the result of multi-stage manufacturing.



### 6.1 The Certificate of Conformity (COC) Chain



European regulations require Whole Vehicle Type Approval (WVTA).

1. Stage 1: The OEM (e.g., Mercedes-Benz) produces the chassis and issues a COC linked to the VIN.

2. Stage 2: The Bodybuilder adds the superstructure. They must issue a second stage COC that references the original VIN but adds their own identification for the bodywork.

3. Registration: The final registration document incorporates data from both stages.




### 6.2 Electronic Integration and the VIN



Modern bodywork requires deep integration with the truck's CAN bus/electronic architecture. Bodybuilders use the VIN to determine if the chassis was ordered with the necessary pre-installations.

- Mercedes-Benz: The VIN option codes are checked for PSM (Parameterizable Special Module - Code ED5). This module acts as a gateway, allowing the bodybuilder to read signals (parking brake on, engine running) and send commands (throttle up for PTO) without splicing into the CAN bus.

- Scania: The BCI (Bodywork Communication Interface) serves a similar function.

- Volvo: The BBM (Body Builder Module).

- Insight: If a VIN decode reveals the absence of these codes, the upfitting process becomes significantly more expensive and technically hazardous, often requiring aftermarket CAN-bus readers that can void warranties.43




## 7. Reference Tables




### 7.1 Mercedes-Benz Commercial "Baumuster" Reference Guide (Positions 4-6)




|   |   |   |   |
|---|---|---|---|
|Code|Model Name|Configuration / Generation|Key Tech Note|
|638|Vito / V-Class|Gen 1 (FWD)|Early FWD platform 19|
|639|Vito / Viano|Gen 2 (RWD/AWD)|Introduction of RWD architecture 19|
|447|Vito / V-Class|Gen 3 (RWD/FWD/AWD)|Also "Metris" in USA 21|
|901-905|Sprinter (T1N)|Gen 1|3-ton (903) vs 4-ton (904) 17|
|906|Sprinter (NCV3)|Gen 2|New platform, OM642 V6 intro 17|
|907|Sprinter (VS30)|Gen 3 (RWD/AWD)|Current RWD platform 18|
|910|Sprinter (VS30)|Gen 3 (FWD)|FWD Low-floor variant 18|
|415|Citan|Mark 1|Renault Kangoo derivative|
|470|X-Class|Pickup|Nissan Navara derivative 7|
|963|Actros / Antos|New Truck Gen|Long Haul / Heavy Dist. 22|
|964|Arocs|New Truck Gen|Construction / Off-road 24|
|967|Atego|New Truck Gen|Medium Distribution 23|
|956|Econic|Low Entry|Municipal / Waste 25|
|406|Unimog|Classic Medium|U65-U900 Series 26|
|437|Unimog|Heavy Series|U1700 / U5000 Series 26|



### 7.2 Global Commercial Vehicle WMI Reference



|   |   |   |   |
|---|---|---|---|
|Manufacturer|WMI (Europe)|WMI (N. America)|Primary Identification Key|
|Mercedes-Benz|WDB, WDF, WDD|WD3, W1Y, W1W|Baumuster (Positions 4-6)|
|Volvo Trucks|YV2|4V4|Chassis Number (Last 7 digits)|
|Scania|YS2, YS4|-|Chassis Number (Last 7 digits)|
|MAN|WMA|-|MAN Number / VDS|
|Iveco|ZCF, WJM|-|Model Variant (e.g., 35S13)|
|DAF|XLR, SFA|-|Chassis Number (Last 8 digits)|
|Renault Trucks|VF6|-|Family Code (Pos 5-6) + CAM Plate|
|Ford|WF0|1FT|Date Code (Pos 11-12 EU / Pos 10 NA)|
|Volkswagen|WV1, WV2|-|Model Code (Pos 7-8, e.g., 2E, SY)|



## 8. Conclusion and Future Outlook



The decoding of commercial vehicle VINs is a discipline that requires shifting one's mindset from the standardized, year-centric model of North American passenger cars to a manufacturer-centric, chassis-number-driven approach in Europe.

For Mercedes-Benz, the Baumuster is the "Rosetta Stone." It is the link between the VIN and the engineering drawings. A user who cannot distinguish a 906 from a 907 is flying blind. For the Swedish giants, Volvo and Scania, the Chassis Number reigns supreme, acting as the index key for highly modular build sheets.

Professionals in this sector must adopt a hierarchical decoding strategy:

1. Identify the WMI to establish the OEM and the region of origin.

2. Extract the VDS to identify the Model Series (Mercedes Baumuster) or Vehicle Family (Renault/MAN).

3. Isolate the Chassis Number (VIS) for precise build sheet retrieval in OEM portals (e.g., Mercedes B2B Connect, Scania Multi).

4. Cross-reference with Bodybuilder Plates and Engine Data Plates to account for the multi-stage manufacturing process.


As the industry moves toward electrification, new codes (e.g., eSprinter engine designations, electric axle codes in Volvo VINs) are emerging. Simultaneously, the legal landscape, driven by EU rulings on data access, is transforming the VIN from a proprietary manufacturing tracking number into a regulated public utility key, democratizing access to the "digital twin" of every truck on the road.

---

Report compiled by the Senior Automotive Data Architect Desk.

Sources include official bodybuilder guidelines, NHTSA filings, OEM decoding manuals, and EU regulatory texts. 6

#### Works cited

1. Vehicle identification number - Wikipedia, accessed November 20, 2025, [https://en.wikipedia.org/wiki/Vehicle_identification_number](https://en.wikipedia.org/wiki/Vehicle_identification_number)

2. Get DAF VIN History Report | DAF Vindecoder, accessed November 20, 2025, [https://vindecoder.eu/daf](https://vindecoder.eu/daf)

3. 2015 Commercial Vehicle Indentification Manual - National Insurance Crime Bureau, accessed November 20, 2025, [https://www.nicb.org/media/2366/download](https://www.nicb.org/media/2366/download)

4. VIN Decoder | AutoZone, accessed November 20, 2025, [https://www.autozone.com/vin-decoder](https://www.autozone.com/vin-decoder)

5. Volvo Truck Vehicle Identification Number VIN | PDF - Scribd, accessed November 20, 2025, [https://www.scribd.com/document/489716986/Volvo-Truck-Vehicle-Identification-Number-VIN](https://www.scribd.com/document/489716986/Volvo-Truck-Vehicle-Identification-Number-VIN)

6. VW VIN Codes - Club VeeDub, accessed November 20, 2025, [https://www.clubvw.org.au/vwreference/vwvin/](https://www.clubvw.org.au/vwreference/vwvin/)

7. Vehicle Identification Numbers (VIN codes)/Mercedes-Benz/VIN Codes - Wikibooks, open books for an open world, accessed November 20, 2025, [https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/Mercedes-Benz/VIN_Codes](https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_\(VIN_codes\)/Mercedes-Benz/VIN_Codes)

8. VIN-to-Year Chart - ALLDATA, accessed November 20, 2025, [https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart](https://www.alldata.com/us/en/support/repair-collision/article/vin-to-year-chart)

9. The Vehicle Identification Number (VIN) - NISR - National Institute of Safety Research, accessed November 20, 2025, [https://cmvid.nisrinc.com/cmv_id/VIN_Visor.pdf](https://cmvid.nisrinc.com/cmv_id/VIN_Visor.pdf)

10. Mercedes-Benz VIN Decoder Phoenix, accessed November 20, 2025, [https://www.mbofnorthscottsdale.com/service/service-tips-and-tricks/mercedes-benz-vin-decoder/](https://www.mbofnorthscottsdale.com/service/service-tips-and-tricks/mercedes-benz-vin-decoder/)

11. Mercedes VIN Decoder - carVertical, accessed November 20, 2025, [https://www.carvertical.com/mercedes-benz-vin-decoder](https://www.carvertical.com/mercedes-benz-vin-decoder)

12. Man VIN Decoder, accessed November 20, 2025, [https://www.freevindecoder.eu/make/man](https://www.freevindecoder.eu/make/man)

13. Find Out What Your VIN Number Says About Your Car in This Mercedes-Benz VIN Breakdown, accessed November 20, 2025, [https://www.mbscottsdale.com/blog/mercedes-benz-vin-breakdown/amp/](https://www.mbscottsdale.com/blog/mercedes-benz-vin-breakdown/amp/)

14. Mercedes Sprinter VIN breakdown - sprntr.co, accessed November 20, 2025, [https://www.sprntr.co/blog/SprinterVan_VIN_number](https://www.sprntr.co/blog/SprinterVan_VIN_number)

15. Vehicle Identification Number (VIN) Coding Summary - StarTek Info, accessed November 20, 2025, [https://www.startekinfo.com/service/download-document/outside/226845/](https://www.startekinfo.com/service/download-document/outside/226845/)

16. VIN Codes: World Manufacturer's Identification - Angelfire, accessed November 20, 2025, [https://www.angelfire.com/ca/TORONTO/VIN/WMI.html](https://www.angelfire.com/ca/TORONTO/VIN/WMI.html)

17. Deciphering Sprinter Models - sprntr.co, accessed November 20, 2025, [https://www.sprntr.co/blog/deciphering-sprinter-models](https://www.sprntr.co/blog/deciphering-sprinter-models)

18. Sprinter Van Master Reference Guide, accessed November 20, 2025, [https://sprintermanual.com/sprinter-van-master-reference-guide/](https://sprintermanual.com/sprinter-van-master-reference-guide/)

19. Engine light on Mercedes V CLASS 447 (2014 - 2019) – clear the fault - klavkarr, accessed November 20, 2025, [https://www.klavkarr.com/compatible/mercedes-vclass-447-37000186](https://www.klavkarr.com/compatible/mercedes-vclass-447-37000186)

20. 2014-2024 Mercedes-Benz Vito Metris VIN Plate Location W447 - YouTube, accessed November 20, 2025, [https://www.youtube.com/watch?v=hqlWzK6SeSU](https://www.youtube.com/watch?v=hqlWzK6SeSU)

21. MODEL INDICATOR INDEX ® (1981-2020) - MB Wholesale Parts, accessed November 20, 2025, [https://www.mbwholesaleparts.com/content/dam/microsites/marketing-portal/parts/MODELDESIGNATIONHISTORICALDOCUMENT.pdf](https://www.mbwholesaleparts.com/content/dam/microsites/marketing-portal/parts/MODELDESIGNATIONHISTORICALDOCUMENT.pdf)

22. mercedes-benz trucks genuine accessories., accessed November 20, 2025, [https://asproddb.mercedes-benz-trucks.com/asproddb-static/catalog/catalogs/0000CB0C912.pdf?_noCache=1743832541623](https://asproddb.mercedes-benz-trucks.com/asproddb-static/catalog/catalogs/0000CB0C912.pdf?_noCache=1743832541623)

23. List of Mercedes-Benz trucks - Wikipedia, accessed November 20, 2025, [https://en.wikipedia.org/wiki/List_of_Mercedes-Benz_trucks](https://en.wikipedia.org/wiki/List_of_Mercedes-Benz_trucks)

24. Introduction of the New Truck Generation The New Arocs (Model 964) - Truckspares365, accessed November 20, 2025, [https://www.truckspares365.co.uk/hub/wp-content/uploads/2022/08/Mercedes-Benz-Arocs-964-Service-Manual.pdf](https://www.truckspares365.co.uk/hub/wp-content/uploads/2022/08/Mercedes-Benz-Arocs-964-Service-Manual.pdf)

25. Truck Retrofitting Mercedes-Benz - Autocommerce, accessed November 20, 2025, [https://www.mercedes-benz-trucks-autocommerce.si/content/dam/retail/si/sl/mercedes-benz-trucks-autocommerce_si/parts/Mercedes-Benz-Truck-Retrofitting-en.pdf.coredownload.pdf](https://www.mercedes-benz-trucks-autocommerce.si/content/dam/retail/si/sl/mercedes-benz-trucks-autocommerce_si/parts/Mercedes-Benz-Truck-Retrofitting-en.pdf.coredownload.pdf)

26. Unimog Models Sorted by Baumuster, accessed November 20, 2025, [https://www.cs.brandeis.edu/~zippy/unimog-model-baumuster.html](https://www.cs.brandeis.edu/~zippy/unimog-model-baumuster.html)

27. Vehicle Identification Numbers (VIN codes)/Volvo/VIN Codes - Wikibooks, open books for an open world, accessed November 20, 2025, [https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/Volvo/VIN_Codes](https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_\(VIN_codes\)/Volvo/VIN_Codes)

28. Vehicle Identification Numbers (VIN codes)/World Manufacturer Identifier (WMI) - Wikibooks, accessed November 20, 2025, [https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/World_Manufacturer_Identifier_(WMI)](https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_\(VIN_codes\)/World_Manufacturer_Identifier_\(WMI\))

29. Vehicle Identification Number - Wikipedia | PDF - Scribd, accessed November 20, 2025, [https://www.scribd.com/document/775111907/Vehicle-Identification-Number-Wikipedia](https://www.scribd.com/document/775111907/Vehicle-Identification-Number-Wikipedia)

30. Information package | Scania Group, accessed November 20, 2025, [https://www.scania.com/group/en/home/products-and-services/services/rmi/information-package.html](https://www.scania.com/group/en/home/products-and-services/services/rmi/information-package.html)

31. Case C‑319/22 - CURIA - Documents, accessed November 20, 2025, [https://curia.europa.eu/juris/document/document.jsf?docid=279492&pageIndex=0&doclang=EN&mode=req&occ=first%E2%88%82%3D1&cid=312775](https://curia.europa.eu/juris/document/document.jsf?docid=279492&pageIndex=0&doclang=EN&mode=req&occ=first%E2%88%82%3D1&cid=312775)

32. EU Court of Justice issues landmark judgement in Scania vehicle data case, accessed November 20, 2025, [https://www.taylorwessing.com/en/insights-and-events/insights/2023/11/eu-gerichtshof-faellt-wegweisendes-urteil-im-scania-fall](https://www.taylorwessing.com/en/insights-and-events/insights/2023/11/eu-gerichtshof-faellt-wegweisendes-urteil-im-scania-fall)

33. Builder process: Body production - MAN, accessed November 20, 2025, [https://www.man.eu/bodybuilder/en/body-manufacturing-process/step-2.html](https://www.man.eu/bodybuilder/en/body-manufacturing-process/step-2.html)

34. VEHICLE AND MANUFACTURER IDENTIFICATION BY VIN CODE As long ago as 1976 the ISO (International Organization for Standardization), accessed November 20, 2025, [https://www.angevaare.eu/pdf/VINENGLV.pdf](https://www.angevaare.eu/pdf/VINENGLV.pdf)

35. 3.1 Vehicle identification through use of the VIN - DAF, accessed November 20, 2025, [https://rmi.daf.com/en/standard-navigation/vehicle-identification/vin](https://rmi.daf.com/en/standard-navigation/vehicle-identification/vin)

36. DAF Vehicle Upgrades - DAF Trucks N.V., accessed November 20, 2025, [https://www.daf.global/en-us/daf-services/workshop-services/daf-vehicle-upgrades](https://www.daf.global/en-us/daf-services/workshop-services/daf-vehicle-upgrades)

37. DAILY RANGE BODYBUILDERS AND VEHICLE FITTING INSTRUCTIONS, accessed November 20, 2025, [http://www.giordanobenicchi.it/camper/IVECO-FIAT/Iveco_daily_bodybuilder_2005.pdf](http://www.giordanobenicchi.it/camper/IVECO-FIAT/Iveco_daily_bodybuilder_2005.pdf)

38. Daily E4 Bodybuilder Instructions | PDF - Scribd, accessed November 20, 2025, [https://www.scribd.com/document/768432406/daily-e4-bodybuilder-instructions](https://www.scribd.com/document/768432406/daily-e4-bodybuilder-instructions)

39. Vehicle Identification | PDF | Truck - Scribd, accessed November 20, 2025, [https://www.scribd.com/document/465961941/0-Vehicle-identification](https://www.scribd.com/document/465961941/0-Vehicle-identification)

40. 2020 VIN Guide.pdf - Ford Pro, accessed November 20, 2025, [https://content.fordpro.com/content/dam/fordpro/us/en-us/pdf/fleet-vehicles/vin-lookup-and-guides/2020-vin-guide.pdf](https://content.fordpro.com/content/dam/fordpro/us/en-us/pdf/fleet-vehicles/vin-lookup-and-guides/2020-vin-guide.pdf)

41. Wf0 D XX TT F D 8 A 52617: Decoding Ford Transit Chassis Number - Year 2002 To 2009, accessed November 20, 2025, [https://www.scribd.com/document/522475983/ford-transit-vincodes](https://www.scribd.com/document/522475983/ford-transit-vincodes)

42. Vehicle Identification Numbers (VIN codes)/Ford/VIN Codes - Wikibooks, open books for an open world, accessed November 20, 2025, [https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/Ford/VIN_Codes](https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_\(VIN_codes\)/Ford/VIN_Codes)

43. Body and Equipment Guideline - Mercedes-Benz Vans, accessed November 20, 2025, [https://www.mbvans.com/content/dam/mb-vans/us/upfitter/09112024_BEG_Sprinter_907.pdf](https://www.mbvans.com/content/dam/mb-vans/us/upfitter/09112024_BEG_Sprinter_907.pdf)

44. Daimler Trucks Body Builders Guide | PDF - Scribd, accessed November 20, 2025, [https://www.scribd.com/document/866796823/Daimler-Trucks-Body-Builders-Guide](https://www.scribd.com/document/866796823/Daimler-Trucks-Body-Builders-Guide)

45. Body/Equipment Mounting Directives Download Guide, accessed November 20, 2025, [https://www.daimlertruck.com.au/siteassets/downloading-of-bodybuilder-directives.pdf](https://www.daimlertruck.com.au/siteassets/downloading-of-bodybuilder-directives.pdf)


**
