# Product Information Catalog and Vehicle Listing (vPIC) Analytical User's Manual, 2023

## Introduction

This document provides guidance to analytical users of Vehicle Identification Number (VIN) decoding information generated for the National Highway Traffic Safety Administration's crash studies using NHTSA's Product Information Catalog and Vehicle Listing (vPIC). The vPIC VIN decoder is built based on the vehicle manufacturer submissions to NHTSA as mandated by the Federal Motor Vehicle Safety Standard (FMVSS) 49 Code of Federal Regulation (CFR) Part 565. The vPIC VIN decoder is intended for use on VINs with model years (MY) 1981 and later. Vehicles prior to 1981 did not have standardized VINs and are not included in the decoding capability for the system. The 49 CFR Part 565 requires vehicle manufacturers to submit VIN deciphering information to NHTSA for the following seven vehicle types: passenger cars, multipurpose passenger vehicles, trucks, buses, trailers/trailer kits, motorcycles, incomplete vehicles, and low-speed vehicles. Therefore, vPIC is capable of decoding VINs for these seven vehicle types. In addition to these seven vehicle types, NHTSA receives voluntary submissions from some manufacturers for their off-road vehicle VINs that can also be decoded through vPIC VIN decoder.

Primarily, the VIN decoding information in vPIC is based on the manufacturer reported data from 49 CFR Part 565 submissions. Some of this information is mandatory per the regulation, but manufacturers may also submit a variety of information that is not mandatory. The information submitted varies by manufacturer, vehicle type and MY. The mandated information for each vehicle type can be found in 49 CFR 565, Table I.

In certain areas of key interest, such as driver assistance technologies, NHTSA does research and adds additional supplemental data. NHTSA searches OEM websites for vehicle manuals, press releases, etc., to identify new technologies that may be optional or standard equipment on specific vehicles by make, model, MY, and trim level. The search is performed for light-duty and medium-/heavy-duty vehicles manufactured by the major light-duty vehicle manufacturers for vehicle MYs 2017 and newer.

The following is a list of vehicle makes for which supplemental data is gathered from the internet.

*   Acura
*   Alfa Romeo
*   Aston Martin
*   Audi
*   Bentley
*   BMW
*   Buick
*   Cadillac
*   Chevrolet
*   Chrysler
*   Dodge
*   Ferrari
*   Fiat
*   Ford
*   Genesis
*   GMC
*   Honda
*   Hyundai
*   Infiniti
*   Jaguar
*   Jeep
*   Kia
*   Lamborghini
*   Land Rover
*   Lexus
*   Lincoln
*   Lotus
*   Lucid
*   Maserati
*   Mazda
*   McLaren
*   Mercedes-Benz
*   MINI
*   Mitsubishi
*   Nissan
*   Polestar
*   Porsche
*   Ram
*   Rivian
*   Rolls-Royce
*   Smart
*   Subaru
*   Tesla
*   Toyota
*   Volkswagen
*   Volvo

The following is a list of driver assistance technologies that are researched on the internet and coded into the vPIC supplemental module.

*   Adaptive Cruise Control (ACC)
*   Adaptive Driving Beam (ADB)
*   Anti-lock Braking System (ABS)
*   Automatic Crash Notification (ACN) / Advanced Automatic Crash Notification (AACN)
*   Automatic Pedestrian Alerting Sound (for Hybrid and EV only)
*   Auto-Reverse System for Windows and Sunroofs
*   Backup Camera
*   Blind Spot Intervention (BSI)
*   Blind Spot Warning (BSW)
*   Crash Imminent Braking (CIB)
*   Daytime Running Light (DRL)
*   Dynamic Brake Support (DBS)
*   Electronic Stability Control (ESC)
*   Event Data Recorder (EDR)
*   Forward Collision Warning (FCW)
*   Headlamp Light Source
*   Keyless Ignition
*   Lane Centering Assistance
*   Lane Departure Warning (LDW)
*   Lane Keeping Assistance (LKA)
*   Parking Assist
*   Pedestrian Automatic Emergency Braking (PAEB)
*   Rear Automatic Emergency Braking
*   Rear Cross Traffic Alert
*   Semiautomatic Headlamp Beam Switching
*   Tire Pressure Monitoring system (TPMS)
*   Traction Control

If manufacturers provide information about driver assistance technologies in their submissions, the manufacturer submitted information is used for VIN decoding instead of the supplemental information.

For motor vehicles and trailers involved in crashes, VINs are collected in NHTSA's crash data collection systems, the Fatality Analysis Reporting System (FARS), the Crash Report Sampling System (CRSS), and the Crash Investigation Sampling System (CISS), and are decoded using the vPIC VIN decoder. The vehicle information decoded by vPIC from a VIN is provided as auxiliary files for the systems. Two vPIC VIN decode files are created for each data collection system, one for motor vehicles (called *vPICDecode*) and one for trailing units (called *vPICTrailerDecode*).

In the vPIC VIN decode files, information will be included only for the VINs that can be cleanly decoded, i.e., with no major errors. See Section **Error Code** for the definition of cleanly decoded VINs.

## Data Elements

The *vPICDecode* file has 128 VIN decoded data elements, excluding key data elements. The number of key data elements vary for different studies that have *vPICDecode* files. A list of the key data elements can be found in each of NHTSA's crash data collection program's Analytical User's Manual (AUM), such as FARS AUM, CRSS AUM, and CISS AUM, etc.

The *vPICTrailerDecode* file has 24 VIN decoded data elements excluding key data elements. This file only includes the data elements that are applicable to trailing units.

## Error Code

When a VIN is decoded, one or more errors might be provided in the output. Cleanly decoded VINs are those that have only one or more of the following error codes when they are decoded.

*   0 - VIN decoded clean. Check Digit (9th position) is correct.
*   1 - Check Digit (9th position) does not calculate properly.
*   10 - Off-road Vehicle Warning - This is not a vehicle identification number (VIN) for a motor vehicle. This indicates that the manufacturer did not certify this product as complying with the Federal motor vehicle safety standards which are applicable to motor vehicles as defined at 49 U.S.C. 30102.
*   400 - Invalid Characters Present

When a VIN is decoded, there can be errors applicable to the decode. In such a situation, the error codes are concatenated and separated by a comma. For example, error code combinations for cleanly decoded VINs can include (0), (0,10), (1,10), (1, 400), (1, 10, 400).

In the VIN decode files, even though **Error Code** data element is lookup data type, only concatenated error code IDs, not the concatenated error code names, are presented due to the size concern of the concatenated name string.

## Vehicle Descriptor

To avoid releasing the full VIN that might be used to link to personally identifiable information, vPIC decode file only contains vehicle descriptors. A vehicle descriptor is 17 characters long. It has the same characters as the VIN except that the check digit number and sequential numbers are replaced by asterisks (*). The check digit is the 9th position. The positions for sequential numbers depend on the manufacturer production capabilities.

*   If the 3rd position of the VIN is 9, i.e., the vehicle manufacturer is a low-volume manufacturer, positions 9 and 15 to 17 are replaced with asterisks (*).
*   If the 3rd position of the VIN is not 9, i.e., the vehicle manufacturer is a high-volume manufacturer, positions 9 and 12 to 17 are replaced with asterisks (*).

## Model Year

There are two sources for vehicle MY:

*   The 10th digit of a VIN specifies vehicle MY, and Vehicle MY can also be reported on a police crash report (PCR).

If there is a MY reported on the PCR, PCR MY is used by vPIC to decode the VIN. If a PCR MY is not reported, MY decoded from the 10th position of the VIN is used by vPIC to decode the VIN.

## Studies With vPIC Decode Files

Please refer to the Analytical User Manuals for the following studies for study specific information.

*   FARS: https://crashstats.nhtsa.dot.gov/Api/Public/ViewPublication/813706
*   CRSS: https://crashstats.nhtsa.dot.gov/Api/Public/ViewPublication/813707
*   CISS: https://crashstats.nhtsa.dot.gov/Api/Public/ViewPublication/813664

## Document Convention for Data Elements, Definitions, and Codes

This document includes several conventions to describe and group the data elements and definitions provided. There are generally five sections for each data element.

1.  The heading contains the group/subgroup of the data element, data element name, the file in which the data element is included, and the data element type as follows.

    **Group/Sub-Group: Data Element Name \[File as T, V] (Data Element Type)**
    *Example: Exterior/Body: Body Class \[VT] (lookup)*

2.  The second line contains the column name for data element as it appears in the files. For lookup data elements, there are two columns, one for *ElementNameID* and one for *ElementName*. For non-lookup data elements, there will be only one column for *ElementName*.

    | Type                                  | Column Names                               |
    | :------------------------------------ | :----------------------------------------- |
    | Data Element Column Names in files    | ElementNameId, ElementName                 |
    | Example for a lookup data element     | Column Names: BodyClassId, BodyClass       |
    | Example for a non-lookup data element | Column Names: ModelYear                    |

3.  The third line contains the datatypes for the columns. For lookup data elements, there are two datatypes, one for *ElementNameID* and one for *ElementName*. For non-lookup elements, there will be only one datatype for *ElementName*.

    | Type                                  | Datatypes                                                |
    | :------------------------------------ | :------------------------------------------------------- |
    | Column Datatypes in files             | DataType for ElementNameId, Data Type for ElementName    |
    | Example for lookup data element       | Column Datatypes: tinyint, varchar (80)                  |
    | Example for non-lookup data element   | Column Datatypes: int                                    |

4.  The fourth section contains a general description of the data element.
5.  The fifth section is a table with the list of attribute names and attribute IDs for the data element for lookup data element types.

### Group and Sub-Group

The decoded output data elements are grouped by vehicle systems/sub-systems. These groupings are as follows.

*   General
*   Active Safety System
*   Active Safety System/Backing Up and Parking
*   Active Safety System/Forward Collision Prevention
*   Active Safety System/Maintaining Safe Distance
*   Active Safety System/Lane and Side Assist
*   Active Safety System/Lighting Technologies
*   Active Safety System/911 Notification
*   Engine
*   Exterior/Body
*   Exterior/Bus
*   Exterior/Dimension
*   Exterior/Motorcycle
*   Exterior/Trailer
*   Exterior/Truck
*   Exterior/Wheel tire
*   Interior
*   Interior/Seat
*   Mechanical/Battery
*   Mechanical/Brake
*   Mechanical/Drivetrain
*   Mechanical/Transmission
*   Passive Safety System
*   Passive Safety System/Air Bag

### Files

There are two files described in this manual: *vPICDecode* file and *vPICTrailerDecode* file. The *vPICDecode* file includes VIN decoding information for a motor vehicle, while *vPICTrailerDecode* file contains decoding information for a trailing unit. For each data element, the following notation is used to indicate in which file this data element appears.

*   [VT]: this data element is included in both *vPICDecode* file and *vPICTrailerDecode* file.
*   [T]: this data element is only included in *vPICTrailerDecode* file.
*   The default, i.e., no bracket provided: this data element is only included in *vPICDecode* file.

### Data Element Type

The data element type describes whether a given data element's values are a fixed list of possible values or an open field for numbers or text. The following is a list of data element types.

*   **Lookup**: the data element values are a finite list of possible values.
*   **String**: this data element values are strings with free-form text.
*   **Int**: the data element values are integers.
*   **Decimal**: the data element values are decimal numbers.

### Column Names

For lookup data elements, there are two columns in the files, one for *ElementNameID* and one for *ElementName*. For non-lookup elements, there is only one column for *ElementName*. For an example, for the data element Axle Count, there is only the name of data element name, *AxleCount*, because this is an integer type data element. For another example, for the data element Plant Country, which is a lookup data type, this line will contain both ID and name, i.e., *PlantCountryId* and *PlantCountry*.

### Column Datatypes

The following is a list of the data types for the columns in the files.

*   **Decimal (x, y)**: the data element is a decimal with a total of x digits and number of decimal digits is y.
*   **Int**: the data element is an integer, i.e., 4 bytes.
*   **Smallint**: the data element is a small integer, i.e., 2 bytes.
*   **Tinyint**: the data element is a tiny integer, i.e., 1 byte.
*   **Varchar(x)**: the column contains data as character string with the maximum length as x.
*   **Datetime**: the data element is a date/time type.

## Data Elements, Definitions, and Codes

### Error Code [VT] (lookup)

Column Names: VINDecodeError
Column Datatypes: varchar(50)

Error Code is a numerical code that determines the nature of the error from VIN decode, and why it occurred.

| Attribute Name                                                                                                                                                                                                                                                                                                                                                     | Attribute Id |
| :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------- |
| 0 - VIN decoded clean. Check Digit (9th position) is correct                                                                                                                                                                                                                                                                                                       | 0            |
| 1 - Check Digit (9th position) does not calculate properly                                                                                                                                                                                                                                                                                                         | 1            |
| 10 - Off-road Vehicle Warning – This is not a vehicle identification number (VIN) for a motor vehicle. This indicates that the manufacturer did not certify this product as complying with the Federal motor vehicle safety standards which are applicable to motor vehicles as defined at 49 U.S.C. 30102.                                                          | 10           |
| 400 - Invalid Characters Present                                                                                                                                                                                                                                                                                                                                   | 400          |

### Vehicle Descriptor [VT] (string)

Column Names: VehicleDescriptor
Column Datatypes: `varchar(17)`

Vehicle Descriptor is a string that can be decoded like a VIN but does not include information to identify an individual vehicle. The sequential numbers and the check digit are replaced by asterisks to avoid any personally identifiable information that would be present in the VIN.

### General

#### Vehicle Type [VT] (lookup)

Column Names: VehicleTypeId, VehicleType
Column Datatypes: tinyint, varchar (40)

This field defines the type of the vehicle based on the World Manufacturer Identifier (WMI).

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| Bus                                 | 5            |
| Incomplete Vehicle                  | 10           |
| Low Speed Vehicle (LSV)             | 9            |
| Motorcycle                          | 1            |
| Multipurpose Passenger Vehicle (MPV) | 7            |
| Off Road Vehicle                    | 13           |
| Passenger Car                       | 2            |
| Trailer                             | 6            |
| Truck                               | 3            |

#### Manufacturer Name [VT] (lookup)

Column Names: ManufacturerFullNameId, ManufacturerFullName
Column Datatypes: int, varchar(135)

Name of the vehicle manufacturer.

The API provides XML lists in pages of 100 manufacturers per page:
`https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=XML&page=1`. For more information, please visit the vPIC API page at `https://vpic.nhtsa.dot.gov/api/`.

#### Make [VT] (lookup)

*   **Column Names:** MakeId, Make
*   **Column Datatypes:** int, varchar(80)

Per 49 CFR 565, Make is a name that a manufacturer applies to a group of vehicles or engines. Please visit the below URL for a complete listing of makes in vPIC.

https://vpic.nhtsa.dot.gov/api/vehicles/getallmakes?format=csv

#### Model [VT] (lookup)

Column Names: ModelId, Model
Column Datatypes: int, varchar(140)

Per 49 CFR 565, Model means a name that a manufacturer applies to a family of vehicles of the same type, make, line, series, and body type.

To lookup the models for a particular make, please pass a valid vPIC Make ID or vPIC Make Text in the URLs below.

*   Replace `*` in the URL with vPIC Make ID:
    https://vpic.nhtsa.dot.gov/api/vehicles/GetModelsForMakeId/*?format=csv
*   Replace `*` in the URL with vPIC Make Text:
    https://vpic.nhtsa.dot.gov/api/vehicles/getmodelsformake/*?format=csv

Example 1: Use the following URL to see all the models for Buick:

*   Use Buick Make ID 468 as parameter:
    https://vpic.nhtsa.dot.gov/api/vehicles/GetModelsForMakeId/468?format=csv
*   Use the Make Name "Buick" as parameter:
    https://vpic.nhtsa.dot.gov/api/vehicles/getmodelsformake/Buick?format=csv

Example 2: Use the following URL to see all the models for Toyota

*   Use Toyota Make ID 448 as parameter:
    https://vpic.nhtsa.dot.gov/api/vehicles/GetModelsForMakeId/448?format=csv
*   Use the Make Name "Toyota" as parameter:
    https://vpic.nhtsa.dot.gov/api/vehicles/getmodelsformake/Toyota?format=csv

#### Model Year [VT] (int)

Column Names: ModelYear
Column Datatypes: int

If the model year (MY) is supplied when the VIN is decoded, such as from a crash report or a vehicle registration record, the MY value will be the supplied MY, even if the MY decoded from the VIN differs from the supplied MY. If the MY is not supplied when the VIN is decoded, the MY value will be decoded from the 10th character in the VIN.

#### Series [VT] (string)

Column Names: Series
Column Datatypes: varchar(165)

Per 49 CFR 565, Series means a name that a manufacturer applies to a subdivision of a "line" denoting price, size, or weight identification and that is used by the manufacturer for marketing purposes.

#### Series2 (string)

Column Names: Series2
Column Datatypes: varchar(65)

This data element captures additional information about series of the vehicle.

#### Trim [VT] (string)

Column Names: Trim
Column Datatypes: varchar(160)

Trim levels further identify a vehicle by a particular set of special features. Higher trim levels either will add to the features of the base (entry-level model) or replace them with something else.

#### Trim2 (string)

**Column Names**: Trim2
**Column Datatypes**: varchar(160)

This data element captures additional information about trim of the vehicle.

#### Plant Country [VT] (lookup)

Column Names: PlantCountryId, PlantCountry
Column Datatypes: tinyint, varchar(40)

This data element captures the country of the manufacturing plant where the manufacturer affixes the VIN.

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| AFGHANISTAN                         | 59           |
| ALBANIA                             | 60           |
| ALGERIA                             | 61           |
| ANDORRA                             | 62           |
| ANGOLA                              | 224          |
| ANTIGUA AND BARBUDA                 | 63           |
| ARGENTINA                           | 17           |
| ARMENIA                             | 64           |
| AUSTRALIA                           | 14           |
| AUSTRIA                             | 7            |
| AZERBAIJAN                          | 65           |
| BAHAMAS                             | 66           |
| BAHRAIN                             | 67           |
| BANGLADESH                          | 68           |
| BARBADOS                            | 69           |
| BELARUS                             | 70           |
| BELGIUM                             | 22           |
| BELIZE                              | 71           |
| BENIN                               | 72           |
| BHUTAN                              | 73           |
| BOLIVIA                             | 74           |
| BOSNIA AND HERZEGOVINA              | 75           |
| BOTSWANA                            | 76           |
| BRAZIL                              | 18           |
| BRUNEI                              | 77           |
| BULGARIA                            | 78           |
| BURKINA FASO                        | 79           |
| BURUNDI                             | 81           |
| CABO VERDE                          | 82           |
| CAMBODIA                            | 83           |
| CAMEROON                            | 84           |
| CANADA                              | 1            |
| CENTRAL AFRICAN REPUBLIC            | 85           |
| CHAD                                | 86           |
| CHILE                               | 87           |
| CHINA                               | 8            |
| COLOMBIA                            | 34           |
| COMOROS                             | 223          |
| CONGO (BRAZZAVILLE)                 | 88           |
| CONGO (KINSHASA)                    | 89           |
| COSTA RICA                          | 49           |
| CROATIA                             | 91           |
| CUBA                                | 92           |
| CYPRUS                              | 93           |
| CZECH REPUBLIC                      | 48           |
| CÔTE D'IVOIRE                       | 90           |
| DENMARK                             | 54           |
| DJIBOUTI                            | 94           |
| DOMINICA                            | 95           |
| DOMINICAN REPUBLIC                  | 96           |
| ECUADOR                             | 97           |
| EGYPT                               | 19           |
| EL SALVADOR                         | 98           |
| ENGLAND                             | 28           |
| EQUATORIAL GUINEA                   | 99           |
| ERITREA                             | 100          |
| ESTONIA                             | 55           |
| ESWATINI                            | 225          |
| ETHIOPIA                            | 101          |
| FEDERATED STATES OF MICRONESIA      | 145          |
| FIJI                                | 102          |
| FINLAND                             | 25           |
| FRANCE                              | 11           |
| GABON                               | 103          |
| GAMBIA                              | 104          |
| GEORGIA                             | 105          |
| GERMANY                             | 2            |
| GHANA                               | 106          |
| GREECE                              | 107          |
| GRENADA                             | 108          |
| GUATEMALA                           | 109          |
| GUINEA-BISSAU                       | 110          |
| GUYANA                              | 111          |
| HAITI                               | 112          |
| HONDURAS                            | 114          |
| HONG KONG                           | 57           |
| HUNGARY                             | 32           |
| ICELAND                             | 115          |
| INDIA                               | 36           |
| INDONESIA                           | 42           |
| IRAN                                | 116          |
| IRAQ                                | 117          |
| IRELAND                             | 118          |
| ISRAEL                              | 21           |
| ITALY                               | 9            |
| JAMAICA                             | 119          |
| JAPAN                               | 3            |
| JORDAN                              | 120          |
| KAZAKHSTAN                          | 121          |
| KENYA                               | 122          |
| KIRIBATI                            | 123          |
| KOSOVO                              | 124          |
| KUWAIT                              | 125          |
| KYRGYZSTAN                          | 126          |
| LAOS                                | 52           |
| LATVIA                              | 128          |
| LEBANON                             | 129          |
| LESOTHO                             | 130          |
| LIBERIA                             | 131          |
| LIBYA                               | 132          |
| LIECHTENSTEIN                       | 133          |
| LITHUANIA                           | 134          |
| LUXEMBOURG                          | 135          |
| MADAGASCAR                          | 137          |
| MALAWI                              | 138          |
| MALAYSIA                            | 43           |
| MALDIVES                            | 139          |
| MALI                                | 140          |
| MALTA                               | 141          |
| MARSHALL ISLANDS                    | 142          |
| MAURITANIA                          | 143          |
| MAURITIUS                           | 144          |
| MEXICO                              | 12           |
| MOLDOVA                             | 146          |
| MONACO                              | 147          |
| MONGOLIA                            | 148          |
| MONTENEGRO                          | 149          |
| MOROCCO                             | 58           |
| MOZAMBIQUE                          | 150          |
| MYANMAR (BURMA)                     | 53           |
| NAMIBIA                             | 151          |
| NAURU                               | 152          |
| NEPAL                               | 153          |
| NETHERLANDS                         | 31           |
| NEW ZEALAND                         | 56           |
| NICARAGUA                           | 156          |
| NIGER                               | 157          |
| NIGERIA                             | 158          |
| NORTH KOREA                         | 29           |
| NORTH MACEDONIA                     | 35           |
| NORWAY                              | 47           |
| OMAN                                | 160          |
| PAKISTAN                            | 161          |
| PALAU                               | 162          |
| PALESTINE                           | 227          |
| PANAMA                              | 163          |
| PAPUA NEW GUINEA                    | 164          |
| PARAGUAY                            | 165          |
| PERU                                | 40           |
| PHILIPPINES                         | 167          |
| POLAND                              | 39           |
| PORTUGAL                            | 41           |
| PUERTO RICO                         | 45           |
| QATAR                               | 170          |
| ROMANIA                             | 44           |
| RUSSIA                              | 37           |
| RWANDA                              | 173          |
| SAINT KITTS AND NEVIS               | 174          |
| SAINT LUCIA                         | 175          |
| SAINT VINCENT AND THE GRENADINES    | 176          |
| SAMOA                               | 177          |
| SAN MARINO                          | 178          |
| SAO TOME AND PRINCIPE               | 180          |
| SAUDI ARABIA                        | 181          |
| SENEGAL                             | 182          |
| SERBIA                              | 24           |
| SEYCHELLES                          | 184          |
| SIERRA LEONE                        | 185          |
| SINGAPORE                           | 46           |
| SLOVAKIA                            | 33           |
| SLOVENIA                            | 27           |
| SOLOMON ISLANDS                     | 189          |
| SOMALIA                             | 190          |
| SOUTH AFRICA                        | 10           |
| SOUTH KOREA                         | 15           |
| SOUTH SUDAN                         | 192          |
| SPAIN                               | 26           |
| SRI LANKA                           | 194          |
| SUDAN                               | 195          |
| SURINAME                            | 196          |
| SWEDEN                              | 23           |
| SWITZERLAND                         | 50           |
| SYRIA                               | 200          |
| TAIWAN                              | 4            |
| TAJIKISTAN                          | 201          |
| TANZANIA                            | 202          |
| THAILAND                            | 30           |
| TIMOR-LESTE                         | 204          |
| TOGO                                | 205          |
| TONGA                               | 226          |
| TRINIDAD AND TOBAGO                 | 206          |
| TUNISIA                             | 207          |
| TURKEY                              | 13           |
| TURKMENISTAN                        | 209          |
| TUVALU                              | 210          |
| UGANDA                              | 211          |
| UKRAINE                             | 212          |
| UNITED ARAB EMIRATES                | 213          |
| UNITED KINGDOM (UK)                 | 5            |
| UNITED STATES (USA)                 | 6            |
| URUGUAY                             | 51           |
| UZBEKISTAN                          | 216          |
| VANUATU                             | 217          |
| VATICAN CITY (HOLY SEE)             | 113          |
| VENEZUELA                           | 16           |
| VIETNAM                             | 38           |
| YEMEN                               | 220          |
| ZAMBIA                              | 221          |
| ZIMBABWE                            | 222          |

#### Plant State [VT] (string)

Column Names: PlantState
Column Datatypes: varchar(30)

This data element captures the State or Province name within the Plant Country of the manufacturing plant where the manufacturer affixes the VIN.

#### Plant City [VT] (string)

Column Names: PlantCity
Column Datatypes: varchar(45)

This data element captures the city of the manufacturing plant where the manufacturer affixes the VIN.

#### Plant Company Name [VT] (string)

Column Names: PlantCompanyName
Column Datatypes: varchar(55)

This data element captures the name of the company that owns the manufacturing plant where the manufacturer affixes the VIN.

#### Destination Market (lookup)

Column Names: DestinationMarketId, DestinationMarket
Column Datatypes: tinyint, varchar(50)

Destination Market is the market where the vehicle is intended to be sold.

| Attribute Name                                  | Attribute Id |
| :---------------------------------------------- | :----------- |
| 49 States (Except California)                   | 1            |
| 50 States                                       | 2            |
| Alaska                                          | 23           |
| Arabian Countries                               | 33           |
| Australia                                       | 25           |
| Brazil                                          | 27           |
| California                                      | 3            |
| Canada                                          | 4            |
| Canada and Other Export Market (BUX)            | 31           |
| Canada, Mexico, Other Export Market (BUX)       | 28           |
| Continental US (excluding Hawaii and Alaska)    | 24           |
| England                                         | 26           |
| Europe                                          | 5            |
| Gulf States                                     | 8            |
| Hawaii                                          | 22           |
| Holland                                         | 30           |
| Japan                                           | 7            |
| Mexico                                          | 11           |
| Mexico and Other Export Market (BUX)            | 34           |
| Non-U.S., Canada, or Mexico                     | 12           |
| Non-U.S., Non-Canada                            | 10           |
| Other Export Market (BUX)                       | 13           |
| Rest of the World (ROW)                         | 32           |
| South Korea                                     | 9            |
| U.S.                                            | 17           |
| U.S., Canada                                    | 14           |
| U.S., Canada, Mexico                            | 18           |
| U.S., Canada, Mexico, Other Export Market (BUX) | 16           |
| U.S., Canada, Other Export Market (BUX)         | 15           |
| U.S., Mexico                                    | 19           |
| U.S., Mexico, Other Export Market (BUX)         | 21           |
| U.S., Other Export Market (BUX)                 | 20           |
| U.S., Puerto Rico, Canada                       | 35           |
| United Kingdom (UK)                             | 29           |
| Worldwide                                       | 6            |

#### Base Price ($) (decimal)

| Label            | Value         |
| :--------------- | :------------ |
| Column Names     | BasePrice     |
| Column Datatypes | decimal(9, 2) |

Base price of the vehicle is the cost of a new vehicle with only the standard equipment and factory warranty. It is the cost without any optional packages.

#### Non-Land Use [VT] (lookup)

**Column Names:** NonLandUseId, NonLandUse
**Column Datatypes:** tinyint, varchar(20)

Non-Land Use data element identifies the non-land use of the vehicle when a vehicle is designed to be used off land (e.g., an amphibious vehicle).

*   Air: identifies vehicles that can be driven on land and in the air
*   Air and Water: identifies vehicles that can be driven on land, in the air and on or under water
*   Water: identifies vehicles that can be driven on land and on or under water

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Air            | 1            |
| Air and Water  | 3            |
| Water          | 2            |

#### Note [VT] (string)

Column Names: Note
Column Datatypes: varchar(500)

Note is used to store any additional information that does not correspond to any of the specified fields on the interface. This is a catch-all element for systems other than for engine, restraint system, brake, and battery. Engine, restraint system, brake, and battery have their own note elements.

### Active Safety System

#### Anti-lock Braking System (ABS) (lookup)

**Column Names:** AntilockBrakeSystemId, AntilockBrakeSystem
**Column Datatypes:** tinyint, varchar(20)

Anti-lock Braking System (ABS) means a portion of a service brake system that automatically controls the degree of rotational wheel slip during braking by:

1.  Sensing the rate of angular rotation of the wheels;
2.  Transmitting signals regarding the rate of wheel angular rotation to one or more controlling devices that interpret those signals and generate responsive controlling output signals; and
3.  Transmitting those controlling signals to one or more modulator devices that adjust brake actuating forces in response to those signals.

| Attribute Name  | Attribute Id |
| :-------------- | :----------- |
| Not Available   | 2            |
| Optional        | 3            |
| Standard        | 1            |

#### Automatic Pedestrian Alerting Sound (for Hybrid and EV only) (lookup)

Column Names: AutoPedestrianAlertingSoundId, AutoPedestrianAlertingSound
Column Datatypes: tinyint, varchar(20)

Electric vehicle warning sounds are a series of sounds designed to alert pedestrians to the presence of electric drive vehicles such as hybrid electric vehicles (HEVs), plug-in hybrid electric vehicles (PHEVs), and all-electric vehicles (EVs) travelling at low speeds. Vehicles operating in all-electric mode produce less noise than traditional combustion engine vehicles and can make it more difficult for pedestrians, the blind, cyclists, and others to be aware of their presence.

| Attribute Name  | Attribute Id |
| :-------------- | :----------- |
| Not Available   | 2            |
| Optional        | 3            |
| Standard        | 1            |

#### Auto-Reverse System for Windows and Sunroofs (lookup)

Column Names: AutoReverseSystemId, AutoReverseSystem
Column Datatypes: tinyint, varchar(20)

An auto-reverse system enables power windows and sunroofs on motor vehicles to automatically reverse direction when such power windows and panels detect an obstruction. This feature can prevent children and others from being trapped, injured, or killed by the power windows and sunroofs.

| Attribute Name  | Attribute Id |
| :-------------- | :----------- |
| Not Available   | 2            |
| Optional        | 3            |
| Standard        | 1            |

#### Electronic Stability Control (ESC) (lookup)

**Column Names**: ElectronicStabilityControlId, ElectronicStabilityControl
**Column Datatypes**: tinyint, varchar(20)

ESC is a computerized technology that improves a vehicle's stability by detecting and reducing loss of traction (skidding). When ESC detects loss of steering control, it automatically applies the brakes to help steer the vehicle in the driver's intended direction. Braking is automatically applied to wheels individually, such as the outer front wheel to counter oversteer, or the inner rear wheel to counter understeer. Some ESC systems also reduce engine power until control is regained.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Event Data Recorder (EDR) (lookup)

Column Names: EventDataRecorderId, EventDataRecorder
Column Datatypes: tinyint, varchar(20)

An EDR is a device installed in motor vehicles to record technical vehicle and occupant information for a brief period before, during, and after a triggering event, typically a crash or near-crash event. Sometimes referred to as "black-box" data, these data or event recorders can be valuable when analyzing and reconstructing crashes.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Keyless Ignition (lookup)

**Column Names:** KeylessIgnitionId, KeylessIgnition
**Column Datatypes:** tinyint, varchar(20)

A keyless ignition system permits starting a car without a physical key being inserted into an ignition. Instead, a small device known as a "key fob" transmits a code to a computer in the vehicle when the fob is within a certain close range. When the coded signal matches the code embedded in the vehicle's computer, a number of systems within the car are activated, including the starter system. This allows the car to be started by simply pressing a button on the dashboard while the key fob is left in a pocket or a purse. The vehicle is usually shut down by pushing the same button.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Tire Pressure Monitoring System (TPMS) Type (lookup)

Column Names: TPMSId, TPMS
Column Datatypes: tinyint, varchar(15)

A TPMS is an electronic system designed to monitor the air pressure inside the pneumatic tires on various types of vehicles. TPMS systems can be divided into two different types - direct and indirect. Direct TPMS employs pressure sensors on each wheel, either internal or external. The sensors physically measure the tire pressure in each tire and report it to the vehicle's instrument cluster or a corresponding monitor. Indirect TPMS does not use physical pressure sensors but measure air pressures by monitoring individual wheel rotational speeds and other signals available outside of the tire itself.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Direct         | 1            |
| Indirect       | 2            |

#### Traction Control (lookup)

**Column Names:** TractionControlId, TractionControl
**Column Datatypes:** tinyint, varchar(20)

When the traction control computer detects a driven wheel or wheels spinning significantly faster than another, it invokes an electronic control unit to apply brake friction to wheels spinning due to loss of traction. This braking action on slipping wheels will cause power transfer to the wheels with traction due to the mechanical action within the differential.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### SAE Automation Level From (int)

Column Names: SAEAutomationLevel_from
Column Datatypes: tinyint

SAE stands for the Society of Automotive Engineers, which changed its name to SAE International in 2006. This field indicates the lower bound of intelligence level and automation capabilities of vehicles, ranking from 0 to 5, per SAE Standard J3016 2018.

#### SAE Automation Level To (int)

Column Names: SAEAutomationLevel_to
Column Datatypes: tinyint

SAE stands for the Society of Automotive Engineers, which changed its name to SAE International in 2006. This field indicates the higher bound of intelligence level and automation capabilities of vehicles, ranking from 0 to 5, per SAE Standard J3016 2018.

#### Active Safety System Note (string)

**Column Names:** ActiveSafetySysNote
**Column Datatypes:** varchar(500)

This field stores additional information about active safety systems in a vehicle.

### Backing Up and Parking

#### Backup Camera (lookup)

Column Names: BackupCameraId, BackupCamera
Column Datatypes: tinyint, varchar(20)

A backup camera, also known as a rearview video system, helps prevent back-over crashes and protects our most vulnerable people - children and senior citizens - by providing an image of the area behind the vehicle. A backup camera helps the driver see behind the vehicle while in reverse.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Parking Assist (lookup)

Column Names: ParkAssistId, ParkAssist
Column Datatypes: tinyint, varchar(20)

A parking assist system uses computer processors, back up cameras, surround-view cameras, and sensors to assist with steering and other functions during parking. Drivers may be required to accelerate, brake, or select gear position. Some systems are capable of parallel and perpendicular parking. Drivers must constantly supervise this support feature and maintain responsibility for parking.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Rear Cross Traffic Alert (lookup)

Column Names: RearCrossTrafficAlertId, RearCrossTrafficAlert
Column Datatypes: tinyint, varchar(20)

A rear cross traffic alert system warns the driver of a potential collision while in reverse, which may be outside the view of the backup camera.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 3            |
| Optional       | 2            |
| Standard       | 1            |

#### Rear Automatic Emergency Braking (lookup)

**Column Names:** RearAutomaticEmergencyBrakingId, RearAutomaticEmergencyBraking
**Column Datatypes:** tinyint, varchar(20)

A rear automatic braking system uses sensors, like parking sensors and the backup camera, to detect objects behind the vehicle. If the system detects a potential collision while in reverse, it automatically applies the brakes if a crash is imminent.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

### Forward Collision Prevention

#### Crash Imminent Braking (CIB) (lookup)

**Column Names:** CrashImminentBrakingId, CrashImminentBraking
**Column Datatypes:** tinyint, varchar(20)

A CIB system is an automatic emergency braking system designed to detect an impending forward crash with another vehicle. CIB systems automatically apply the brakes in a crash imminent situation to slow or stop the vehicle, avoiding the crash or reducing its severity, if the driver does not brake in response to a forward collision alert.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Dynamic Brake Support (DBS) (lookup)

Column Names: DynamicBrakeSupportId, DynamicBrakeSupport
Column Datatypes: tinyint, varchar(20)

A DBS system is an automatic emergency braking system designed to detect an impending forward crash with another vehicle. DBS systems automatically supplement the driver's braking in an effort to avoid a crash if the driver does not brake hard enough to avoid it.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Forward Collision Warning (FCW) (lookup)

**Column Names:** ForwardCollisionWarningId, ForwardCollisionWarning
**Column Datatypes:** tinyint, varchar(20)

An FCW system monitors a vehicle's speed, the speed of the vehicle in front of it, and the distance between the vehicles. If the vehicles get too close due to the speed of either vehicle, the FCW system will warn the driver of the rear vehicle of an impending crash so that the driver can apply the brakes or take evasive action, such as steering, to prevent a potential crash. FCW systems provide an audible, visual, or haptic warning, or any combination thereof, to alert the driver of an FCW-equipped vehicle of a potential collision.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Pedestrian Automatic Emergency Braking (PAEB) (lookup)

**Column Names:** PedestrianAutoEmergencyBrakingId, PedestrianAutoEmergencyBraking
**Column Datatypes:** tinyint, varchar(20)

A PAEB system provides automatic braking for vehicles when pedestrians are in front of the vehicle and the driver has not acted to avoid a crash.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

### Maintaining Safe Distance

#### Adaptive Cruise Control (ACC) (lookup)

**Column Names:** AdaptiveCruiseControlId, AdaptiveCruiseControl
**Column Datatypes:** tinyint, varchar(20)

ACC automatically adjusts the vehicle's speed to keep a pre-set distance from the vehicle in front of it.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

### Lane and Side Assist

#### Blind Spot Intervention (BSI) (lookup)

Column Names: BlindSpotInterventionId, BlindSpotIntervention
Column Datatypes: tinyint, varchar(20)

BSI helps prevent a collision with a vehicle in the driver's blind spot. If the driver ignores the blind spot warning and starts to change to a lane where there's a vehicle, the system activates and automatically applies light braking pressure, or provides steering input, to guide the vehicle back into the original lane. The system monitors for vehicles in the driver's blind spot using rear-facing cameras or proximity sensors.

| Attribute Name  | Attribute Id |
| :-------------- | :----------- |
| Not Available   | 2            |
| Optional        | 3            |
| Standard        | 1            |

#### Blind Spot Warning (BSW) (lookup)

Column Names: BlindSpotWarningId, BlindSpotWarning
Column Datatypes: tinyint, varchar(20)

BSW alerts drivers with an audio or visual warning if there are vehicles in adjacent lanes that the driver may not see when making a lane change.

| Attribute Name  | Attribute Id |
| :-------------- | :----------- |
| Not Available   | 2            |
| Optional        | 3            |
| Standard        | 1            |

#### Lane Centering Assistance (lookup)

**Column Names**: LaneCenteringAssistanceId, LaneCenteringAssistance
**Column Datatypes**: tinyint, varchar(20)

A lane centering assistance system utilizes a camera-based vision system designed to monitor the vehicle's lane position and automatically and continuously apply steering inputs needed to keep the vehicle centered within its lane.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Lane Departure Warning (LDW) (lookup)

**Column Names:** LaneDepartureWarningId, LaneDepartureWarning
**Column Datatypes:** tinyint, varchar(20)

An LDW system monitors lane markings and alerts the driver if their vehicle drifts out of their lane without a turn signal or any control input indicating the lane departure is intentional. An audio, visual or other alert warns the driver of the unintentional lane shift so the driver can steer the vehicle back into its lane.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Lane Keeping Assistance (LKA) (lookup)

Column Names: LaneKeepingAssistanceId, LaneKeepingAssistance
Column Datatypes: tinyint, varchar(20)

An LKA system prevents a driver from unintentionally drifting out of the intended travel lane. LKA systems use information provided by Lane Departure Warning (LDW) system sensors to determine whether a vehicle is about to unintentionally move out of its lane of travel. If so, LKA activates and corrects the steering, brakes or accelerates one or more wheels, or does both, resulting in the vehicle returning to its intended lane of travel.

| Attribute Name  | Attribute Id |
| :-------------- | :----------- |
| Not Available   | 4            |
| Optional        | 5            |
| Standard        | 1            |

### Lighting Technologies

#### Adaptive Driving Beam (ADB) (lookup)

**Column Names:** AdaptiveDrivingBeamId, AdaptiveDrivingBeam
**Column Datatypes:** tinyint, varchar(20)

ADB is a type of front-lighting system that lets upper beam headlamps adapt their beam patterns to create shaded areas around oncoming and preceding vehicles to improve long-range visibility for the driver without causing discomfort, distraction, or glare to other road users.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Daytime Running Light (DRL) (lookup)

**Column Names:** DaytimeRunningLightId, DaytimeRunningLight
**Column Datatypes:** tinyint, varchar(20)

DRL is an automotive lighting system on the front of a vehicle or bicycle, that automatically switches on when the vehicle is in drive, and emits white, yellow, or amber light to increase the conspicuity of the vehicle during daylight conditions.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

#### Headlamp Light Source (lookup)

**Column Names:** HeadlampLightSourceId, HeadlampLightSource
**Column Datatypes:** tinyint, varchar(30)

A headlamp light source provides a distribution of light designed to provide adequate forward and lateral illumination with limits on light directed towards the eyes of other road users, to control glare. This beam is intended for use whenever other vehicles are present ahead. Halogen, high-intensity discharge (HID), light-emitting diode (LED), and laser are the most common headlights on the market.

| Attribute Name            | Attribute Id |
| :------------------------ | :----------- |
| Halogen                   | 1            |
| Halogen, HID              | 5            |
| Halogen, HID, LED         | 11           |
| Halogen, HID, LED, Other  | 15           |
| Halogen, HID, Other       | 12           |
| Halogen, LED              | 6            |
| Halogen, LED, Other       | 13           |
| Halogen, Other            | 7            |
| HID                       | 2            |
| HID, LED                  | 8            |
| HID, LED, Other           | 14           |
| HID, Other                | 9            |
| Laser                     | 16           |
| LED                       | 3            |
| LED, Other                | 10           |
| Other                     | 4            |

#### Semiautomatic Headlamp Beam Switching (lookup)

Column Names: SemiAutoHeadlampBeamSwitchingId, SemiAutoHeadlampBeamSwitching
Column Datatypes: tinyint, varchar(20)

A semi-automatic headlamp beam switching device provides automatic or manual control of beam switching at the option of the driver. When the control is automatic, the headlamps switch from the upper beam to the lower beam when illuminated by the headlamps on an approaching car and switch back to the upper beam when the road ahead is dark. When the control is manual, the driver may obtain either beam manually regardless of the condition of lights ahead of the vehicle.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

### 911 Notification

#### Automatic Crash Notification (ACN) / Advanced Automatic Crash Notification (AACN) (lookup)

**Column Names:** AutomaticCrashNotificationId, AutomaticCrashNotification
**Column Datatypes:** tinyint, varchar(20)

An ACN system notifies emergency responders that a crash has occurred and provides its location.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Not Available  | 2            |
| Optional       | 3            |
| Standard       | 1            |

### Engine

#### Engine Manufacturer (string)

Column Names: EngineManufacturer
Column Datatypes: varchar(40)

This data element captures the name of the engine manufacturer.

#### Engine Model (string)

Column Names: EngineModel
Column Datatypes: varchar(115)

Engine model is a name that a manufacturer applies to a family of engine.

#### Engine Configuration (lookup)

Column Names: EngineConfigurationId, EngineConfiguration
Column Datatypes: tinyint, varchar(35)

Engine configuration defines how engine cylinders are arranged. Common values are V6 for V-shaped arrangement, I4 or L4 for in-line arrangement.

| Attribute Name               | Attribute Id |
| :--------------------------- | :----------- |
| Horizontally opposed (boxer) | 5            |
| In-Line                      | 1            |
| Rotary                       | 3            |
| V-Shaped                     | 2            |
| W Shaped                     | 4            |

#### Engine Power (kW) (decimal)

**Column Names:** EnginePowerKW
**Column Datatypes:** decimal(7, 3)

This field stores engine power in kilowatts (kW).

#### Engine Stroke Cycles (int)

Column Names: EngineStrokeCycles
Column Datatypes: tinyint

Engine stroke cycle is a numerical field for the number of strokes used by an internal combustion engine to complete a power cycle.

#### Engine Number of Cylinders (int)

**Column Names:** EngineCylindersCount
**Column Datatypes:** tinyint

This is a numerical field to store the number of cylinders in an engine. Common values for passenger cars are 4 or 6.

#### Engine Brake (hp) From (decimal)

*   **Column Names:** EngineBrakeHP\_from
*   **Column Datatypes:** decimal(8, 4)

Engine brake is the horsepower (hp) at the engine output shaft. Engine Brake (hp) From is the lower value of the range.

#### Engine Brake (hp) To (decimal)

Column Names: EngineBrakeHP_to
Column Datatypes: decimal(8, 4)

Engine brake is the horsepower (hp) at the engine output shaft. Engine Brake (hp) To is the upper value of the range.

#### Cooling Type (lookup)

Column Names: EngineCoolingTypeId, EngineCoolingType
Column Datatypes: tinyint, varchar(10)

Cooling type defines the cooling system used to control the engine temperature. It can be either air-cooled or water-cooled.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Air            | 1            |
| Water          | 2            |

#### Displacement (CI) (decimal)

Column Names: DisplacementCI
Column Datatypes: decimal(7, 2)

Engine displacement (in cubic inches) is the volume swept by all the pistons inside the cylinders of a reciprocating engine in a single movement from top dead center to bottom dead center.

#### Displacement (CC) (decimal)

| Category         | Value          |
| :--------------- | :------------- |
| Column Names     | DisplacementCC |
| Column Datatypes | decimal(9, 3)  |

Engine displacement (in cubic centimeters) is the volume swept by all the pistons inside the cylinders of a reciprocating engine in a single movement from top dead center to bottom dead center.

#### Displacement (L) (decimal)

Column Names: DisplacementL
Column Datatypes: decimal(9, 6)

Engine displacement (in liters) is the volume swept by all the pistons inside the cylinders of a reciprocating engine in a single movement from top dead center to bottom dead center.

#### Electrification Level (lookup)

Column Names: EngineElectrificationLevelId, EngineElectrificationLevel
Column Datatypes: tinyint, varchar(45)

Electrification level defines to what level the vehicle is powered by electric system. The common electric system configurations are mild hybrid, strong hybrid, plug-in hybrid, battery electric, and fuel cell vehicles.

1.  Mild hybrid is the system such as 12-volt start-stop or 48-volt belt integrator starter generator (BISG) system that uses an electric motor to add assisting power to the internal combustion engine. The system has features such as stop-start, power assist, and mild level of generative braking features.
2.  Strong hybrid systems, in vehicles such as the Toyota Prius, mainly consist of motors, conventional gasoline engine, and battery, but the source of electrical charge for the battery power is provided by the conventional engine and/or regenerative braking.
3.  Plug-in hybrid systems, in vehicles such as the Toyota Rav4 Prime, mainly consist of motors, conventional gasoline engine and battery. Plug-in hybrid vehicles are like strong hybrids, but they have a larger battery pack and can be charged with an external source of electricity by electric vehicle supply equipment (EVSE).
4.  Battery electric vehicles (BEV), such as the Tesla Model S or Nissan Leaf, have only a battery and electrical motor components and use electricity as the only power source.
5.  Fuel cell electric vehicles (FCEV) use full electric drive platforms but consume electricity generated by onboard fuel cells and hydrogen fuel.

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| BEV (Battery Electric Vehicle)      | 4            |
| FCEV (Fuel Cell Electric Vehicle)   | 5            |
| HEV (Hybrid Electric Vehicle) - Level Unknown | 9            |
| Mild HEV (Hybrid Electric Vehicle)  | 1            |
| PHEV (Plug-in Hybrid Electric Vehicle) | 3            |
| Strong HEV (Hybrid Electric Vehicle) | 2            |

#### Fuel Delivery / Fuel Injection Type (lookup)

Column Names: FuelDeliveryInjectionTypeId, FuelDeliveryInjectionType
Column Datatypes: tinyint, varchar(50)

Fuel Delivery / Fuel Injection Type defines the mechanism through which fuel is delivered to the engine.

| Attribute Name                                  | Attribute Id |
| :---------------------------------------------- | :----------- |
| Common Rail Direct Injection Diesel (CRDI)      | 6            |
| Compression Ignition                            | 9            |
| Lean-Burn Gasoline Direct Injection (LBGDI)     | 2            |
| Multipoint Fuel Injection (MPFI)                | 3            |
| Sequential Fuel Injection (SFI)                 | 4            |
| Stoichiometric Gasoline Direct Injection (SGDI) | 1            |
| Throttle Body Fuel Injection (TBI)              | 5            |
| Transistor Controlled Ignition (TCI)            | 10           |
| Unit Injector Direct Injection Diesel (UDI)     | 7            |

#### Fuel Type - Primary (lookup)

Column Names: FuelTypePrimaryId, FuelTypePrimary
Column Datatypes: tinyint, varchar(45)

Fuel type defines the fuel used to power the vehicle. For vehicles that have two power sources, such as plug-in hybrid vehicle, both primary fuel type and secondary fuel type will be provided.

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| Compressed Hydrogen/Hydrogen        | 8            |
| Compressed Natural Gas (CNG)        | 6            |
| Diesel                              | 1            |
| Electric                            | 2            |
| Ethanol (E85)                       | 10           |
| Flexible Fuel Vehicle (FFV)         | 15           |
| Fuel Cell                           | 18           |
| Gasoline                            | 4            |
| Liquefied Natural Gas (LNG)         | 7            |
| Liquefied Petroleum Gas (propane or LPG) | 9            |
| Methanol (M85)                      | 13           |
| Natural Gas                         | 17           |
| Neat Ethanol (E100)                 | 11           |
| Neat Methanol (M100)                | 14           |

#### Fuel Type - Secondary (lookup)

Column Names: FuelTypeSecondaryId, FuelTypeSecondary
Column Datatypes: tinyint, varchar(45)

Fuel type defines the fuel used to power the vehicle. For vehicles that have two power sources, such as plug-in hybrid vehicle, both primary fuel type and secondary fuel type will be provided.

| Attribute Name                           | Attribute Id |
| :--------------------------------------- | :----------- |
| Compressed Hydrogen/Hydrogen             | 8            |
| Compressed Natural Gas (CNG)             | 6            |
| Diesel                                   | 1            |
| Electric                                 | 2            |
| Ethanol (E85)                            | 10           |
| Flexible Fuel Vehicle (FFV)              | 15           |
| Fuel Cell                                | 18           |
| Gasoline                                 | 4            |
| Liquefied Natural Gas (LNG)              | 7            |
| Liquefied Petroleum Gas (propane or LPG) | 9            |
| Methanol (M85)                           | 13           |
| Natural Gas                              | 17           |
| Neat Ethanol (E100)                      | 11           |
| Neat Methanol (M100)                     | 14           |

#### Top Speed (MPH) (int)

Column Names: TopSpeedMPH
Column Datatypes: tinyint

This data element captures the top speed of the vehicle in miles per hour (mph).

#### Turbo (lookup)

Column Names: EngineTurboId, EngineTurbo
Column Datatypes: tinyint, varchar(10)

Turbo is a yes/no field to identify whether the engine is turbo-charged or not.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| No             | 2            |
| Yes            | 1            |

#### Valve Train Design (lookup)

**Column Names:** EngineValveTrainDesignId, EngineValveTrainDesign
**Column Datatypes:** tinyint, varchar(35)

Valve train design defines engine camshaft design and control. Common values are single overhead cam (SOHC), dual overhead cam (DOHC), overhead valve (OHV), etc.

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| Camless Valve Actuation (CVA)       | 1            |
| Dual Overhead Cam (DOHC)            | 2            |
| Overhead Valve (OHV)                | 3            |
| Single Overhead Cam (SOHC)          | 4            |

#### Other Engine Info (string)

Column Names: OtherEngineInfo
Column Datatypes: varchar(500)

This is a catch-all field for storing additional engine information that does not fit in any of the other engine fields.

### Exterior

### Body

#### Body Class [VT] (lookup)

Column Names: BodyClassId, BodyClass
Column Datatypes: tinyint, varchar(80)

Body Class presents the body type based on 49 CFR 565.12(b): "Body type means the general configuration or shape of a vehicle distinguished by such characteristics as the number of doors or windows, cargo-carrying features and the roofline (e.g., sedan, fastback, hatchback)." Definitions are not provided for individual body types in the regulation.

| Attribute Name                                                      | Attribute Id |
| :------------------------------------------------------------------ | :----------- |
| Ambulance                                                           | 128          |
| Bus                                                                 | 16           |
| Bus - School Bus                                                    | 73           |
| Cargo Van                                                           | 95           |
| Convertible/Cabriolet                                               | 1            |
| Coupe                                                               | 3            |
| Crossover Utility Vehicle (CUV)                                     | 8            |
| Fire Apparatus                                                      | 130          |
| Hatchback/Liftback/Notchback                                        | 5            |
| Incomplete                                                          | 65           |
| Incomplete - Bus Chassis                                            | 107          |
| Incomplete - Chassis Cab (Double Cab)                               | 70           |
| Incomplete - Chassis Cab (Number of Cab Unknown)                    | 74           |
| Incomplete - Chassis Cab (Single Cab)                               | 63           |
| Incomplete - Commercial Bus Chassis                                 | 72           |
| Incomplete - Commercial Chassis                                     | 112          |
| Incomplete - Cutaway                                                | 62           |
| Incomplete - Glider                                                 | 64           |
| Incomplete - Motor Coach Chassis                                    | 76           |
| Incomplete - Motor Home Chassis                                     | 78           |
| Incomplete - School Bus Chassis                                     | 71           |
| Incomplete - Shuttle Bus Chassis                                    | 77           |
| Incomplete - Stripped Chassis                                       | 67           |
| Incomplete - Trailer Chassis                                        | 116          |
| Incomplete - Transit Bus Chassis                                    | 75           |
| Limousine                                                           | 117          |
| Low Speed Vehicle (LSV) / Neighborhood Electric Vehicle (NEV)       | 4            |
| Minivan                                                             | 2            |
| Motorcycle - Competition                                            | 114          |
| Motorcycle - Cross Country                                          | 109          |
| Motorcycle - Cruiser                                                | 82           |
| Motorcycle - Custom                                                 | 94           |
| Motorcycle - Dual Sport / Adventure / Supermoto / On/Off-road       | 85           |
| Motorcycle - Enclosed Three Wheeled or Enclosed Autocycle [1Rear Wheel] | 100          |
| Motorcycle - Moped                                                  | 104          |
| Motorcycle - Scooter                                                | 12           |
| Motorcycle - Side Car                                               | 90           |
| Motorcycle - Small / Minibike                                       | 87           |
| Motorcycle - Sport                                                  | 80           |
| Motorcycle - Standard                                               | 6            |
| Motorcycle - Street                                                 | 98           |
| Motorcycle - Three Wheeled, Unknown Enclosure or Autocycle, Unknown Enclosure | 131          |
| Motorcycle - Three-Wheeled Motorcycle (2 Rear Wheels)               | 83           |
| Motorcycle - Touring / Sport Touring                                | 81           |
| Motorcycle - Underbone                                              | 110          |
| Motorcycle - Unenclosed Three Wheeled or Open Autocycle [1 Rear Wheel] | 103          |
| Motorcycle - Unknown Body Class                                     | 125          |
| Motorhome                                                           | 108          |
| Off-road Vehicle - All Terrain Vehicle (ATV) (Motorcycle-style)     | 69           |
| Off-road Vehicle - Construction Equipment                           | 127          |
| Off-road Vehicle - Dirt Bike / Off-Road                             | 84           |
| Off-road Vehicle - Enduro (Off-road long distance racing)           | 86           |
| Off-road Vehicle - Farm Equipment                                   | 126          |
| Off-road Vehicle - Go Kart                                          | 88           |
| Off-road Vehicle - Golf Cart                                        | 124          |
| Off-road Vehicle - Motocross (Off-road short distance, closed track racing) | 113          |
| Off-road Vehicle - Multipurpose Off-Highway Utility Vehicle [MOHUV] or Recreational Off-Highway Vehicle [ROV] | 105          |
| Off-road Vehicle - Snowmobile                                       | 97           |
| Pickup                                                              | 60           |
| Roadster                                                            | 10           |
| Sedan/Saloon                                                        | 13           |
| Sport Utility Truck (SUT)                                           | 119          |
| Sport Utility Vehicle (SUV)/Multi-Purpose Vehicle (MPV)             | 7            |
| Step Van / Walk-in Van                                              | 111          |
| Street Sweeper                                                      | 129          |
| Streetcar / Trolley                                                 | 68           |
| Trailer                                                             | 61           |
| Truck                                                               | 11           |
| Truck-Tractor                                                       | 66           |
| Van                                                                 | 9            |
| Wagon                                                               | 15           |

#### Doors (int)

Column Names: DoorsCount
Column Datatypes: tinyint

This is a numerical field to store the number of doors on a vehicle.

#### Track Width (inches) (decimal)

Column Names: TrackWidthIN
Column Datatypes: decimal(6, 3)

A vehicle's track, or track width, is the distance in inches between the center line of each of the two wheels on the same axle on any given vehicle.

#### Wheel Base Type (lookup)

**Column Names:** WheelBaseTypeId, WheelBaseType
**Column Datatypes:** tinyint, varchar(15)

This field describes the wheel base variations for some trucks and passenger cars, relative to other variants of the vehicle model. This field may have values such as short, standard, long, extra long, or super long.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Extra Long     | 4            |
| Long           | 1            |
| Medium         | 6            |
| Short          | 2            |
| Standard       | 5            |
| Super Long     | 3            |

#### Windows (int)

Column Names: Windows
Column Datatypes: tinyint

This field defines the number of windows on a vehicle.

### Bus

#### Bus Type (lookup)

Column Names: BusTypeId, BusType
Column Datatypes: tinyint, varchar(20)

This field defines the type of bus. Common types include commuter coach, double deck coach, tourist coach, urban bus, transit bus, entertainer coach and motorhome.

| Attribute Name    | Attribute Id |
| :---------------- | :----------- |
| Commuter Coach    | 1            |
| Double Deck Coach | 2            |
| Entertainer Coach | 6            |
| Motorhome         | 7            |
| Tourist Coach     | 3            |
| Transit Bus       | 5            |
| Urban Bus         | 4            |

#### Bus Floor Configuration Type (lookup)

Column Names: BusFloorConfigurationTypeId, BusFloorConfigurationType
Column Datatypes: tinyint, varchar(20)

This field defines the relative height of the bus floor from the ground. Common values include normal, lift/raised, low floor, and sleeper coach. Low floor refers to a bus deck that is accessible from the sidewalk with only a single step with a small height difference, caused solely by the difference between the bus deck and sidewalk. This is distinct from high/raised floor, a bus deck design that requires climbing one or more steps to access the interior floor that is placed at a higher height. A sleeper coach is a type of specially adapted coach designed for the passengers to sleep in.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Lift/Raised    | 2            |
| Low Floor      | 4            |
| Normal         | 1            |
| Sleeper Coach  | 3            |

#### Bus Length (feet) (int)

Column Names: BusLengthFT
Column Datatypes: smallint

This field captures the length of bus in feet.

#### Other Bus Info (string)

Column Names: OtherBusInfo
Column Datatypes: varchar(500)

Additional information about bus is captured in this field.

### Dimension

#### Bed Length (inches) (int)

*   **Column Names:** TruckBedLengthIN
*   **Column Datatypes:** smallint

Defines the length of the pickup truck bed in inches.

#### Curb Weight (pounds) (int)

Column Names: CurbWeightLB
Column Datatypes: smallint

Curb weight (in pounds) is the total weight of a vehicle with standard equipment, all necessary operating consumables such as motor oil, transmission oil, coolant, air conditioning refrigerant, and a full tank of fuel, while not loaded with either passengers or cargo.

#### Gross Vehicle Weight Rating From [VT] (lookup)

Column Names: GrossVehicleWeightRatingFromId, GrossVehicleWeightRatingFrom
Column Datatypes: tinyint, varchar(55)

Gross vehicle weight rating (GVWR) is the maximum operating weight of a vehicle including the vehicle's chassis, body, engine, engine fluids, fuel, accessories, driver, passengers, and cargo, but excluding that of the trailers. Per 49 CFR 565.15, Class 1 is further broken down to Class A-D; Class 2 is further broken down to Class E-H. This field captures the lower bound of GVWR range for the vehicle.

| Attribute Name                                    | Attribute Id |
| :------------------------------------------------ | :----------- |
| Class 1: 6,000 lb or less (2,722 kg or less)      | 1            |
| Class 1A: 3,000 lb or less (1,360 kg or less)     | 10           |
| Class 1B: 3,001 - 4,000 lb (1,360 - 1,814 kg)     | 11           |
| Class 1C: 4,001 - 5,000 lb (1,814 - 2,268 kg)     | 12           |
| Class 1D: 5,001 - 6,000 lb (2,268 - 2,722 kg)     | 13           |
| Class 2: 6,001 - 10,000 lb (2,722 - 4,536 kg)     | 2            |
| Class 2E: 6,001 - 7,000 lb (2,722 - 3,175 kg)     | 14           |
| Class 2F: 7,001 - 8,000 lb (3,175 - 3,629 kg)     | 15           |
| Class 2G: 8,001 - 9,000 lb (3,629 - 4,082 kg)     | 16           |
| Class 2H: 9,001 - 10,000 lb (4,082 - 4,536 kg)    | 17           |
| Class 3: 10,001 - 14,000 lb (4,536 - 6,350 kg)    | 4            |
| Class 4: 14,001 - 16,000 lb (6,350 - 7,258 kg)    | 5            |
| Class 5: 16,001 - 19,500 lb (7,258 - 8,845 kg)    | 6            |
| Class 6: 19,501 - 26,000 lb (8,845 - 11,794 kg)   | 7            |
| Class 7: 26,001 - 33,000 lb (11,794 - 14,969 kg)  | 8            |
| Class 8: 33,001 lb and above (14,969 kg and above) | 9            |

#### Gross Vehicle Weight Rating To [VT] (lookup)

Column Names: GrossVehicleWeightRatingToId, GrossVehicleWeightRatingTo
Column Datatypes: tinyint, varchar(55)

Gross vehicle weight rating (GVWR) is the maximum operating weight of a vehicle including the vehicle's chassis, body, engine, engine fluids, fuel, accessories, driver, passengers, and cargo, but excluding that of the trailers. Per 49 CFR 565.15, Class 1 is further broken down to Class A-D; Class 2 is further broken down to Class E-H. This field captures the higher bound of GVWR range for the vehicle.

| Attribute Name                                     | Attribute Id |
| :------------------------------------------------- | :----------- |
| Class 1: 6,000 lb or less (2,722 kg or less)       | 1            |
| Class 1A: 3,000 lb or less (1,360 kg or less)      | 10           |
| Class 1B: 3,001 - 4,000 lb (1,360 - 1,814 kg)      | 11           |
| Class 1C: 4,001 - 5,000 lb (1,814 - 2,268 kg)      | 12           |
| Class 1D: 5,001 - 6,000 lb (2,268 - 2,722 kg)      | 13           |
| Class 2: 6,001 - 10,000 lb (2,722 - 4,536 kg)      | 2            |
| Class 2E: 6,001 - 7,000 lb (2,722 - 3,175 kg)      | 14           |
| Class 2F: 7,001 - 8,000 lb (3,175 - 3,629 kg)      | 15           |
| Class 2G: 8,001 - 9,000 lb (3,629 - 4,082 kg)      | 16           |
| Class 2H: 9,001 - 10,000 lb (4,082 - 4,536 kg)     | 17           |
| Class 3: 10,001 - 14,000 lb (4,536 - 6,350 kg)     | 4            |
| Class 4: 14,001 - 16,000 lb (6,350 - 7,258 kg)     | 5            |
| Class 5: 16,001 - 19,500 lb (7,258 - 8,845 kg)     | 6            |
| Class 6: 19,501 - 26,000 lb (8,845 - 11,794 kg)    | 7            |
| Class 7: 26,001 - 33,000 lb (11,794 - 14,969 kg)   | 8            |
| Class 8: 33,001 lb and above (14,969 kg and above) | 9            |

#### Gross Combination Weight Rating From (lookup)

Column Names: GrossCombWeightRatingFromId, GrossCombWeightRatingFrom
Column Datatypes: tinyint, varchar(55)

Gross combination weight rating (GCWR) is the maximum allowable combined mass of a road vehicle, the passengers and cargo in the tow vehicle, plus the mass of the trailer and cargo in the trailer. This rating is set by the vehicle manufacturer. This field captures the lower bound of GCWR range for the vehicle.

| Attribute Name                                    | Attribute Id |
| :------------------------------------------------ | :----------- |
| Class 1: 6,000 lb or less (2,722 kg or less)      | 1            |
| Class 1A: 3,000 lb or less (1,360 kg or less)     | 10           |
| Class 1B: 3,001 - 4,000 lb (1,360 - 1,814 kg)     | 11           |
| Class 1C: 4,001 - 5,000 lb (1,814 - 2,268 kg)     | 12           |
| Class 1D: 5,001 - 6,000 lb (2,268 - 2,722 kg)     | 13           |
| Class 2: 6,001 - 10,000 lb (2,722 - 4,536 kg)     | 2            |
| Class 2E: 6,001 - 7,000 lb (2,722 - 3,175 kg)     | 14           |
| Class 2F: 7,001 - 8,000 lb (3,175 - 3,629 kg)     | 15           |
| Class 2G: 8,001 - 9,000 lb (3,629 - 4,082 kg)     | 16           |
| Class 2H: 9,001 - 10,000 lb (4,082 - 4,536 kg)    | 17           |
| Class 3: 10,001 - 14,000 lb (4,536 - 6,350 kg)    | 4            |
| Class 4: 14,001 - 16,000 lb (6,350 - 7,258 kg)    | 5            |
| Class 5: 16,001 - 19,500 lb (7,258 - 8,845 kg)    | 6            |
| Class 6: 19,501 - 26,000 lb (8,845 - 11,794 kg)   | 7            |
| Class 7: 26,001 - 33,000 lb (11,794 - 14,969 kg)  | 8            |
| Class 8: 33,001 lb and above (14,969 kg and above)| 9            |

#### Gross Combination Weight Rating To (lookup)

Column Names: GrossCombWeightRatingToId, GrossCombWeightRatingTo
Column Datatypes: tinyint, varchar(55)

Gross combination weight rating (GCWR) is the maximum allowable combined mass of a road vehicle, the passengers and cargo in the tow vehicle, plus the mass of the trailer and cargo in the trailer. This rating is set by the vehicle manufacturer. This field captures the higher bound of GCWR range for the vehicle.

| Attribute Name                                    | Attribute Id |
| :------------------------------------------------ | :----------- |
| Class 1: 6,000 lb or less (2,722 kg or less)      | 1            |
| Class 1A: 3,000 lb or less (1,360 kg or less)     | 10           |
| Class 1B: 3,001 - 4,000 lb (1,360 - 1,814 kg)     | 11           |
| Class 1C: 4,001 - 5,000 lb (1,814 - 2,268 kg)     | 12           |
| Class 1D: 5,001 - 6,000 lb (2,268 - 2,722 kg)     | 13           |
| Class 2: 6,001 - 10,000 lb (2,722 - 4,536 kg)     | 2            |
| Class 2E: 6,001 - 7,000 lb (2,722 - 3,175 kg)     | 14           |
| Class 2F: 7,001 - 8,000 lb (3,175 - 3,629 kg)     | 15           |
| Class 2G: 8,001 - 9,000 lb (3,629 - 4,082 kg)     | 16           |
| Class 2H: 9,001 - 10,000 lb (4,082 - 4,536 kg)    | 17           |
| Class 3: 10,001 - 14,000 lb (4,536 - 6,350 kg)    | 4            |
| Class 4: 14,001 - 16,000 lb (6,350 - 7,258 kg)    | 5            |
| Class 5: 16,001 - 19,500 lb (7,258 - 8,845 kg)    | 6            |
| Class 6: 19,501 - 26,000 lb (8,845 - 11,794 kg)   | 7            |
| Class 7: 26,001 - 33,000 lb (11,794 - 14,969 kg)  | 8            |
| Class 8: 33,001 lb and above (14,969 kg and above)| 9            |

#### Wheel Base (inches) From (decimal)

**Column Names:** WheelBaseIN_from
**Column Datatypes:** decimal(7, 4)

Wheel base is the distance between the centers of the front and rear wheels measured in inches. This field captures the lower bound of the wheel base range.

#### Wheel Base (inches) To (decimal)

**Column Names:** `WheelBaseIN_to`
**Column Datatypes:** `decimal(7, 4)`

Wheel base is the distance between the centers of the front and rear wheels measured in inches. This field captures the upper bound of the wheel base range.

### Motorcycle

#### Custom Motorcycle Type (lookup)

Column Names: CustomMotorcycleTypeId, CustomMotorcycleType
Column Datatypes: tinyint, varchar(25)

Defines the type of customized motorcycle, including values such as bagger, chopper, cruise, touring. Custom motorcycle type is provided by the manufacturer but not defined.

| Attribute Name      | Attribute Id |
| :------------------ | :----------- |
| Bagger              | 13           |
| Base                | 15           |
| Bobber              | 7            |
| Cafe Racer          | 20           |
| Chopper/Chopped     | 2            |
| Cruising            | 18           |
| Drag                | 17           |
| Dresser             | 6            |
| Drop Seat           | 10           |
| Hot Rod             | 12           |
| Reverse Trike       | 3            |
| Rigid Frame         | 4            |
| Roller Kit          | 8            |
| Scooter             | 19           |
| Show Bike           | 9            |
| Side-by-Side Seating| 21           |
| Sportster           | 5            |
| Standard            | 16           |
| Street/Prostreet    | 11           |
| Tour/Touring        | 14           |
| Utility             | 22           |

#### Motorcycle Chassis Type (lookup)

**Column Names:** MotorcycleChassisTypeId, MotorcycleChassisType
**Column Datatypes:** tinyint, varchar(50)

Defines the type of motorcycle chassis, including values such as trike, reversed trike, autocycle, three-wheeler, etc.

| Attribute Name              | Attribute Id |
| :-------------------------- | :----------- |
| Autocycle (Open or Closed)  | 3            |
| Reverse Trike               | 4            |
| Three-Wheeler - Other       | 2            |
| Trike                       | 1            |

#### Motorcycle Suspension Type (lookup)

Column Names: MotorcycleSuspensionTypeId, MotorcycleSuspensionType
Column Datatypes: tinyint, varchar(40)

Defines the type of suspension in motorcycles, including values such as hardtail, softail, wing fork, etc. Motorcycle Suspension Type is provided by the manufacturer but not defined.

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| Hardtail/Rigid Frame                | 1            |
| Rubber Mount                        | 4            |
| Softail                             | 2            |
| Swingarm/Wing Fork/Pivoted Fork     | 3            |

#### Other Motorcycle Info (string)

Column Names: OtherMotorcycleInfo
Column Datatypes: varchar(500)

Additional information about motorcycles is captured in this field.

### Trailer

#### Trailer Body Type [T] (lookup)

Column Names: TrailerBodyTypeId, TrailerBodyType
Column Datatypes: tinyint, varchar(55)

Trailer Body Type describes the purpose of the trailer or how it is designed to be used.

| Attribute Name                        | Attribute Id |
| :------------------------------------ | :----------- |
| A Frame Trailer                       | 64           |
| ATV Trailer                           | 108          |
| Adventure Trailer                     | 165          |
| Aerial (Tiller)                       | 67           |
| Air Bag Trailer                       | 107          |
| Angle Iron                            | 114          |
| Asphalt Trailer                       | 145          |
| Auto Hauler                           | 95           |
| Auto Transporter                      | 58           |
| Boat Trailer                          | 2            |
| Box Trailer                           | 51           |
| Box or Van Enclosed Trailer           | 5            |
| Breast Board Trailer                  | 110          |
| Bulk                                  | 103          |
| Bunk Trailer                          | 161          |
| CVT                                   | 127          |
| Cable/Reel Trailer                    | 85           |
| Cage Trailer                          | 158          |
| Camping or Travel Trailer             | 10           |
| Car Hauler Trailer                    | 3            |
| Car Trailer                           | 66           |
| Cargo Trailer                         | 160          |
| Cemetery Trailer                      | 73           |
| Chemical Trailer                      | 172          |
| Chipper Trailer                       | 34           |
| Clam Shell Trailer                    | 56           |
| Combined Travel/Utility               | 87           |
| Commercial Trailer                    | 184          |
| Communication Trailer                 | 99           |
| Concession Trailer                    | 162          |
| Contractor Trailer                    | 94           |
| Cryogenic Transport Trailer           | 138          |
| Curtain Side Trailer                  | 7            |
| Custom Trailer                        | 129          |
| Deckover Trailer                      | 55           |
| Deer Stand Trailer                    | 149          |
| Dining Trailer                        | 79           |
| Dolly                                 | 45           |
| Dovetail                              | 125          |
| Drop Deck Trailer                     | 61           |
| Dry Bulk Trailer                      | 174          |
| Dry Freight                           | 186          |
| Dump Trailer                          | 4            |
| Dynamometer                           | 128          |
| Elliptical Floor                      | 178          |
| Emergency Traffic Control             | 47           |
| Equipment Trailer                     | 23           |
| Fifth Wheel                           | 26           |
| Fire Training Trailer                 | 98           |
| Flagger Signal Trailer                | 48           |
| Flatbed or Platform Trailer           | 6            |
| Flatbed - A Train                     | 92           |
| Flatbed - B Train                     | 93           |
| Flatbed - Double Drop                 | 9            |
| Flatbed - Extendable                  | 89           |
| Flatbed - Flip Axle                   | 96           |
| Flatbed - Single Drop                 | 88           |
| Flatbed - Tilt                        | 134          |
| Flat Deck Trailer                     | 57           |
| Flat Floor                            | 177          |
| Flat Top Trailer                      | 159          |
| Float Trailer                         | 139          |
| Frame Type                            | 168          |
| Generator                             | 167          |
| Golf Cart Trailer                     | 133          |
| Goose Trailer                         | 52           |
| Grain Trailer                         | 33           |
| Grain, Chip, or Gravel Trailer        | 180          |
| Grill/Smoker Trailer                  | 35           |
| Grinder                               | 37           |
| Agricultural Trailer                  | 136          |
| Hoist Trailer                         | 175          |
| Horse Trailer                         | 14           |
| Hospitality Trailer                   | 153          |
| House Trailer                         | 82           |
| Hybrid Trailer                        | 120          |
| Hydratail                             | 24           |
| Hydraulic Trailer                     | 8            |
| Incomplete Trailer Chassis            | 46           |
| Intermodal Container Chassis or Trailer | 109          |
| Isothermal                            | 173          |
| Jeep                                  | 90           |
| Kitchen Trailer                       | 77           |
| Landscape Trailer                     | 17           |
| Large Hauler                          | 156          |
| Laundry Trailer                       | 84           |
| Leaf Collector                        | 144          |
| Live Animal Trailer                   | 13           |
| Living Quarters                       | 155          |
| Livestock - Cattle Trailer            | 53           |
| Livestock - Horse Trailer             | 60           |
| Log Splitter                          | 146          |
| Logging Trailer                       | 19           |
| Lowboy/Low bed Trailer                | 11           |
| Material Handling Trailer             | 86           |
| Mixer Trailer                         | 187          |
| Mobile Equipment                      | 40           |
| Mobile Facility Trailer               | 189          |
| Motor Home                            | 30           |
| Motorcycle Trailer                    | 12           |
| Multi-Passenger Vehicle Trailer       | 163          |
| Non-Walk-In Body                      | 70           |
| Off-Road Use                          | 118          |
| Office Trailer                        | 154          |
| Open Bed Trailer                      | 32           |
| Open Top Trailer                      | 132          |
| Other Trailer                         | 101          |
| Palift                                | 116          |
| Park Trailer                          | 27           |
| Paver                                 | 126          |
| Personal Watercraft Trailer/PWC Trailer | 135          |
| Pipe Trailer                          | 18           |
| Pole Trailer                          | 91           |
| Pony Pup                              | 41           |
| Porta-Toilet Trailer                  | 81           |
| Portable Traffic Signal               | 49           |
| Pressure Washer Trailer               | 147          |
| Pro-Comp                              | 157          |
| Pup Trailer                           | 83           |
| Ramp                                  | 39           |
| Rapiscan (X-Ray)                      | 68           |
| Rec. Trailer                          | 80           |
| Recycling Trailer                     | 63           |
| Reefer Trailer                        | 143          |
| Reel                                  | 166          |
| Refrigerated Trailer                  | 65           |
| Rest Room Trailer                     | 123          |
| Roll Off Trailer                      | 43           |
| Round Floor                           | 179          |
| Rumble Strip Machine                  | 42           |
| Sand and Gravel Trailer               | 183          |
| Satellite Trailer                     | 140          |
| Scow                                  | 44           |
| Self-Contained                        | 20           |
| Sewer Cleaning Equipment              | 185          |
| Shower Trailer                        | 78           |
| Side Dump Trailer                     | 22           |
| Skid Unit                             | 113          |
| Slide-In Truck Camper                 | 29           |
| Slot Saw                              | 38           |
| Snowmobile Trailer                    | 1            |
| Special Design Trailer                | 111          |
| Special Purpose Trailer               | 188          |
| Sport Utility                         | 121          |
| Sprayer Trailer                       | 148          |
| Stacker Trailer/Stack Trailer         | 100          |
| Standard                              | 102          |
| Steel Body/Frame                      | 119          |
| Step Bed                              | 181          |
| Step Deck                             | 104          |
| Stinger                               | 97           |
| Stock Trailer                         | 182          |
| Storage Shed Trailer                  | 122          |
| Straw Trailer                         | 169          |
| Tank Trailer                          | 15           |
| Tear Drop Camper                      | 62           |
| Telescoping Trailer                   | 54           |
| Tender Trailer                        | 137          |
| Tent Camper                           | 28           |
| Tile Cart                             | 71           |
| Tilt/Angled Trailer                   | 76           |
| Tow Dolly                             | 105          |
| Towing Trailer                        | 117          |
| Toy Hauler                            | 170          |
| Tractor                               | 112          |
| Trailer Busses                        | 164          |
| Trailer on Flat Car (TOFC)            | 171          |
| Transfer Trailer                      | 141          |
| Travel Trailer                        | 21           |
| Truck Campers                         | 131          |
| Truck Caps                            | 130          |
| Truck Loader                          | 151          |
| Tube                                  | 115          |
| Utility Trailer                       | 16           |
| Vacuum Tank                           | 25           |
| Van Trailer                           | 36           |
| Van/Covered Cargo/Enclosed            | 59           |
| Vault Trailer                         | 72           |
| Vending Trailer                       | 152          |
| Walk-In Body                          | 69           |
| Water Trailer                         | 31           |
| Wedge Trailer                         | 142          |
| Work Over Rig Carrier                 | 176          |
| Workforce Trailer                     | 124          |
| Yard Trailer                          | 106          |

#### Trailer Length (feet) [T] (decimal)

Column Names: TrailerLengthFT
Column Datatypes: decimal(6, 3)

Trailer length is the length of the trailer in feet from the front of the connector to the end of the trailer.

#### Other Trailer Info [T] (string)

Column Names: OtherTrailerInfo
Column Datatypes: varchar(500)

Additional information about trailers is captured in this field.

#### Trailer Type Connection [T] (lookup)

Column Names: TrailerTypeConnectionId, TrailerTypeConnection
Column Datatypes: tinyint, varchar(35)

Trailer Type Connection describes the connector or tongue used for the trailers.

| Attribute Name              | Attribute Id |
| :-------------------------- | :----------- |
| Ball Hitch                  | 31           |
| Ball Type Pull              | 1            |
| Bumper Pull                 | 14           |
| Fifth Wheel                 | 5            |
| Gooseneck                   | 3            |
| Kingpin                     | 6            |
| Other                       | 30           |
| Pintle Hitch/Hook           | 2            |
| Straight Semi/Semi Trailer  | 4            |

### Truck

#### Bed Type (lookup)

Column Names: TruckBedTypeId, TruckBedType
Column Datatypes: tinyint, varchar(15)

Bed type is the type of bed (the open back) used for pickup trucks. The common values are standard, short, long, extended.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Extended       | 3            |
| Long           | 1            |
| Short          | 2            |
| Standard       | 4            |

#### Cab Type (lookup)

Column Names: `TruckBodyCabTypeId`, `TruckBodyCabType`
Column Datatypes: `tinyint`, `varchar(45)`

Cab type applies to both pickup truck and other medium- and heavy-duty trucks. The cab or cabin of a truck is the inside space in a truck where the driver is seated. For pickup trucks, the cab type is categorized by the combination number of doors and number of rows for seating. For medium- and heavy-duty trucks (MDHD), the cab type is categorized by the relative location of engine and cab.

For pickup trucks, there are four cab types.

*   Regular: 2 door, 1 row of seats
*   Extra/Super/Quad/Double/King/Extended: 2 doors, 2 rows of seats
*   Crew/Super Crew/Crew Max: 4 doors, 2 rows of seats
*   Mega: 4 doors, 2 rows of seats (with a bigger cabin than crew cab type)

For medium- and heavy-duty (MDHD) trucks, there are several categories as listed below.

*   Cab Beside Engine
*   CAE: Cab Above Engine
*   CBE: Cab Behind Engine
*   COE: Cab Over Engine or Flat Nose: Driver sits on top of the front axle and engine
*   LCF: Low Cab Forward
*   Conventional: Driver sits behind the engine
*   Non-Tilt
*   Tilt

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| Crew/ Super Crew/ Crew Max          | 4            |
| Extra/Super/Quad/Double/King/Extended | 2            |
| MDHD: CAE (Cab Above Engine)        | 8            |
| MDHD: CBE (Cab Behind Engine)       | 7            |
| MDHD: COE (Cab Over Engine)         | 6            |
| MDHD: Cab Beside Engine             | 12           |
| MDHD: Conventional                  | 5            |
| MDHD: LCF (Low Cab Forward)         | 9            |
| MDHD: Non-Tilt                      | 11           |
| MDHD: Tilt                          | 10           |
| Mega                                | 13           |
| Regular                             | 3            |

### Wheel tire

#### Number of Wheels (int)

Column Names:
: WheelsCount

Column Datatypes:
: tinyint

This field captures the number of wheels on a vehicle.

#### Wheel Size Front (inches) (int)

Column Names: WheelSizeFrontIN
Column Datatypes: tinyint

This field captures the diameter of the front wheel in inches.

#### Wheel Size Rear (inches) (int)

Column Names: WheelSizeRearIN
Column Datatypes: tinyint

This field captures the diameter of the rear wheel in inches.

### Interior

#### Steering Location (lookup)

Column Names: SteeringLocationId, SteeringLocation
Column Datatypes: tinyint, varchar(25)

This data element captures the location of steering column, either on left- (LHD) or right-hand side (RHD).

| Attribute Name        | Attribute Id |
| :-------------------- | :----------- |
| Left-Hand Drive (LHD) | 1            |
| Right-Hand Drive (RHD)| 2            |

#### Entertainment System (lookup)

Column Names: EntertainmentSystemId, EntertainmentSystem
Column Datatypes: tinyint, varchar(30)

This field defines the type of different entertainment systems in vehicles.

| Attribute Name          | Attribute Id |
| :---------------------- | :----------- |
| CD + Stereo             | 2            |
| Rear Entertainment System | 1            |

### Seat

#### Number of Seats (int)

Column Names: SeatsCount
Column Datatypes: tinyint

This data element is a numeric field to store the number of seats in a vehicle.

#### Number of Seat Rows (int)

|                   |               |
| :---------------- | :------------ |
| Column Names:     | SeatRowsCount |
| Column Datatypes: | tinyint       |

This data element is a numeric field to capture the number of rows of seats in a vehicle.

### Mechanical

### Battery

#### Battery Current (Amps) From (int)

Column Names: BatteryA_from
Column Datatypes: int
Battery Current (Amps) From field stores the lower value for battery current range in the unit of amperes (amps).

#### Battery Current (Amps) To (int)

**Column Names**: BatteryA_to
**Column Datatypes**: int

Battery Current (Amps) To field stores the higher value for battery current range in the unit of amperes (amps).

#### Battery Energy (kWh) From (decimal)

Column Names: BatteryKWh_from
Column Datatypes: decimal(7, 2)

Battery Energy (kWh) From field stores the lower bound of battery energy range in the unit of kilowatt-hour (kWh).

#### Battery Energy (kWh) To (decimal)

Column Names: BatteryKWh_to
Column Datatypes: decimal(7, 2)

Battery Energy (kWh) To field stores the upper bound of battery energy range in the unit of kilowatt-hour (kWh).

#### Battery Type (lookup)

Column Names: BatteryTypeId, BatteryType
Column Datatypes: tinyint, varchar(30)

Battery type field stores the battery chemistry type for anode, cathode.

| Attribute Name                      | Attribute Id |
| :---------------------------------- | :----------- |
| Cobalt Dioxide/Cobalt               | 4            |
| Iron Phosphate/FePo                 | 8            |
| Lead Acid/Lead                      | 1            |
| Lithium-Ion/Li-Ion                  | 3            |
| Manganese Oxide Spinel/MnO          | 7            |
| Nickel-Metal-Hydride/NiMH           | 2            |
| Nickle-Cobalt-Aluminum/NCA          | 6            |
| Nickle-Cobalt-Manganese/NCM         | 5            |
| Silicon                             | 9            |

#### Battery Voltage (Volts) From (int)

Column Names: BatteryV_from
Column Datatypes: smallint

Battery Voltage (Volts) From field stores the lower bound for battery voltage range in the unit of volts.

#### Battery Voltage (Volts) To (int)

Column Names: BatteryV_to
Column Datatypes: smallint

Battery Voltage (Volts) To field stores the upper bound for battery voltage range in the unit of volts.

#### EV Drive Unit (lookup)

Column Names: EVDriveUnitId, EVDriveUnit
Column Datatypes: tinyint, varchar(15)

EV Drive Unit field stores electric vehicle (EV) motor configuration for single or dual motor.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| Dual Motor     | 1            |
| Single Motor   | 2            |

#### Number of Battery Packs per Vehicle (int)

Column Names: BatteryPacksPerVehicle
Column Datatypes: int

This field stores the information about how battery packs are arranged for a vehicle.

#### Number of Battery Modules per Pack (int)

Column Names: BatteryModulesPerPack
Column Datatypes: tinyint

This field stores the information about how battery modules are arranged for a vehicle.

#### Number of Battery Cells per Module (int)

Column Names: BatteryCellsPerModule
Column Datatypes: int

This field stores the information about how battery cells are arranged for a vehicle.

#### Other Battery Info (string)

| Column Names:   | OtherBatteryInfo |
| :-------------- | :--------------- |
| Column Datatypes: | varchar(500)     |

This field stores any other battery information that does not belong to any of the other battery related fields.

### Charger

#### Charger Level (lookup)

Column Names: ChargerLevelId, ChargerLevel
Column Datatypes: tinyint, varchar(135)

There are three levels of battery chargers currently. Level 1 and 2 are AC chargers; Level 3 is a DC charger. Chargers at each level charge batteries with different voltage and current levels. Level 3 is the fastest charging.

*   **Level 1**
    *   AC charger.
    *   In North America this typically means 16 amps at 120 volts delivering 1.9 kW of power.
    *   In Europe it may be 13 or 16 amps at 240 volts delivering 3 kW of power.
    *   The EV may incorporate a standard domestic power cord to connect the vehicle to a domestic socket outlet or a Level 1 charging station.
*   **Level 2**
    *   AC charger.
    *   It delivers up to 20 kW of power from either single- or three-phase alternating current (AC) sources of 208-240 volts at up to 80 amps.
*   **Level 3**
    *   DC charging, or "fast charging."
    *   To achieve very short charging times, Level 3 chargers supply very high currents of up to 400 amps at voltages up to 600 volts DC delivering a maximum power of 240 kW.

| Attribute Name                                                                                                                                                                         | Attribute Id |
| :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------- |
| Level 1 AC Charger (typically 16A 120V 1.9kW or 13-16A 240V 3kW) may incorporate standard domestic power cord.                                                                         | 1            |
| Level 2 AC Charger (up to 80A, 208-240V AC, up to 20kW from single- or three-phase AC) cables permanently fixed to charging station.                                                     | 2            |
| Level 3 DC Charger or fast charger (up to 400A, up to 600V DC, up to 240kW)                                                                                                            | 3            |

#### Charger Power (kW) (int)

Column Names: ChargerPowerKW
Column Datatypes: smallint
Charger power stores the power of the charging circuit on board in kilowatts (kW).

### Brake

#### Brake System Type (lookup)

Column Names: BrakeSystemTypeId, BrakeSystemType
Column Datatypes: tinyint, varchar(20)

Brake system type is the type of brake systems used to stop and hold the vehicle or combination of motor vehicles.

| Attribute Name    | Attribute Id |
| :---------------- | :----------- |
| Air               | 1            |
| Air and Hydraulic | 3            |
| Electric          | 4            |
| Hydraulic         | 2            |
| Mechanical        | 5            |

#### Brake System Description (string)

**Column Names:** BrakeSystemDesc
**Column Datatypes:** varchar(500)

This field provides additional details about the brake system.

### Drivetrain

#### Axles [VT] (int)

Column Names: AxlesCount
Column Datatypes: tinyint

This numerical field defines the number of axles on a vehicle.

#### Axle Configuration [VT] (lookup)

Column Names: AxleConfigurationId, AxleConfiguration
Column Datatypes: tinyint, varchar(25)

Axle configuration describes the placement of the axles in a vehicle. This field is mainly used for trucks.

| Attribute Name        | Attribute Id |
| :-------------------- | :----------- |
| Independent Axle      | 4            |
| SBA - Set-Back Axle   | 1            |
| SFA - Set-Forward Axle| 2            |
| Single                | 5            |
| Tandem                | 3            |

#### Drive Type (lookup)

Column Names: DriveTypeId, DriveType
Column Datatypes: tinyint, varchar(25)

Drive type stores information about vehicle drivetrain configuration. The most common drive types for passenger cars, crossover vehicles, and pickup trucks are front-wheel drive (FWD), rear-wheel drive (RWD), all-wheel drive (AWD), and 4-wheel drive (4WD).

| Attribute Name        | Attribute Id |
| :-------------------- | :----------- |
| 10x10                 | 21           |
| 10x4                  | 15           |
| 10x6                  | 16           |
| 10x8                  | 23           |
| 12x4                  | 17           |
| 12x6                  | 18           |
| 14x4                  | 19           |
| 14x6                  | 22           |
| 2WD/4WD               | 24           |
| 4WD/4-Wheel Drive/4x4 | 2            |
| 4x2                   | 7            |
| 6x2                   | 9            |
| 6x4                   | 5            |
| 6x6                   | 13           |
| 8x2                   | 20           |
| 8x4                   | 12           |
| 8x6                   | 14           |
| 8x8                   | 10           |
| AWD/All-Wheel Drive   | 3            |
| FWD/Front-Wheel Drive | 1            |
| Glider                | 8            |
| Other                 | 6            |
| RWD/Rear-Wheel Drive  | 4            |

### Transmission

#### Transmission Speeds (int)

**Column Names**: TransmissionSpeeds
**Column Datatypes**: tinyint

Transmission speed is a numerical field to capture the number of speeds for a transmission, such as 6 for a 6-speed automatic or manual transmission.

#### Transmission Style (lookup)

Column Names: TransmissionStyleId, TransmissionStyle
Column Datatypes: tinyint, varchar(45)

Transmission style provides information about the type of transmissions. The major types of transmissions are manual transmission, automatic transmission, continuously variable transmission (CVT), and dual-clutch transmission (DCT).

| Attribute Name                          | Attribute Id |
| :-------------------------------------- | :----------- |
| Automated Manual Transmission (AMT)     | 8            |
| Automatic                               | 2            |
| Continuously Variable Transmission (CVT) | 7            |
| Dual-Clutch Transmission (DCT)          | 14           |
| Electronic Continuously Variable (e-CVT) | 4            |
| Manual/Standard                         | 3            |
| Motorcycle - Chain Drive                | 10           |
| Motorcycle - Chain Drive Off-Road       | 13           |
| Motorcycle - CVT Belt Drive             | 12           |
| Motorcycle - Shaft Drive                | 9            |
| Motorcycle - Shaft Drive Off-Road       | 11           |

### Passive Safety System

#### Pretensioner (lookup)

Column Names: PretensionerId, Pretensioner
Column Datatypes: tinyint, varchar(10)

This yes/no field captures whether or not the vehicle has a pretensioner, a device designed to make seat belts even more effective by removing the slack from a seat belt as soon as a crash is detected or if the system senses excessive seat belt tension on the driver's or passenger's seat belt.

| Attribute Name | Attribute Id |
| :------------- | :----------- |
| No             | 2            |
| Yes            | 1            |

#### Seat Belt Type (lookup)

Column Names: SeatBeltTypeId, SeatBeltType
Column Datatypes: tinyint, varchar(25)

This field describes the type of seat belt, such as manual or automatic. Automatic seat belts automatically close over riders in a vehicle. Automatic seat belts were mainly used in some older model GM, Nissan, and Honda vehicles and are rarely seen now.

| Attribute Name        | Attribute Id |
| :-------------------- | :----------- |
| Automatic             | 2            |
| Manual                | 1            |
| Manual and Automatic  | 3            |

#### Other Restraint System Info (string)

**Column Names**: OtherRestraintSystemInfo
**Column Datatypes**: varchar(500)

Other Restraint Info field is used to code additional information about restraint system that cannot be coded in any other restraint fields.

### Air Bag Location

#### Curtain Air Bag Locations (lookup)

**Column Names:** AirBagLocCurtainId, AirBagLocCurtain
**Column Datatypes:** tinyint, varchar(35)

This field captures the location of curtain air bags. Curtain air bags are side air bags that protect the head.

| Attribute Name              | Attribute Id |
| :-------------------------- | :----------- |
| 1st and 2nd and 3rd Rows    | 5            |
| 1st and 2nd Rows            | 4            |
| 1st Row (Driver and Passenger) | 3            |
| All Rows                    | 6            |
| Driver Seat Only            | 1            |
| Passenger Seat Only         | 2            |

#### Front Air Bag Locations (lookup)

Column Names: AirBagLocFrontId, AirBagLocFront
Column Datatypes: tinyint, varchar(35)

This field captures the location of frontal air bags. Frontal air bags are generally designed to deploy in "moderate to severe" frontal or near-frontal crashes.

| Attribute Name              | Attribute Id |
| :-------------------------- | :----------- |
| 1st Row (Driver and Passenger) | 3            |
| Driver Seat Only            | 1            |

#### Knee Air Bag Locations (lookup)

Column Names: AirBagLocKneeId, AirBagLocKnee
Column Datatypes: tinyint, varchar(35)

This field captures the location of knee air bags, which deploy from a car's lower dashboard, are meant to distribute impact forces on an occupant's legs in the case of a crash, thereby reducing leg injuries.

| Attribute Name              | Attribute Id |
| :-------------------------- | :----------- |
| 1st Row (Driver and Passenger) | 3            |
| Driver Seat Only            | 1            |
| Passenger Seat Only         | 2            |

#### Side Air Bag Locations (lookup)

Column Names: AirBagLocSideId, AirBagLocSide
Column Datatypes: tinyint, varchar(35)

This field captures the location of side air bags, typically designed for three areas of added protection: chest/torso, head, or both.

| Attribute Name              | Attribute Id |
| :-------------------------- | :----------- |
| 1st and 2nd and 3rd Rows    | 5            |
| 1st and 2nd Rows            | 4            |
| 1st Row (Driver and Passenger) | 3            |
| All Rows                    | 6            |
| Driver Seat Only            | 1            |
| Passenger Seat Only         | 2            |

#### Seat Cushion Air Bag Locations (lookup)

Column Names: AirBagLocSeatCushionId, AirBagLocSeatCushion
Column Datatypes: tinyint, varchar(35)

This field captures the location of seat cushion air bags, used as a supplement to the seat belts to help prevent the front passenger from sliding forward in the event of a front impact collision.

| Attribute Name              | Attribute Id |
| :-------------------------- | :----------- |
| 1st and 2nd and 3rd Rows    | 5            |
| 1st and 2nd Rows            | 4            |
| 1st Row (Driver and Passenger) | 3            |
| All Rows                    | 6            |
| Driver Seat Only            | 1            |
| Passenger Seat Only         | 2            |