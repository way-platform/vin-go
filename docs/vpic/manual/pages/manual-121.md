# Data Elements, Definitions, and Codes

## Exterior/Truck: Cab Type (lookup)

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