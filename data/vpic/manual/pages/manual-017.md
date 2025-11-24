# Data Elements, Definitions, and Codes

## Error Code [VT] (lookup)

Column Names: VINDecodeError
Column Datatypes: varchar(50)

Error Code is a numerical code that determines the nature of the error from VIN decode, and why it occurred.

| Attribute Name                                                                                                                                                                                                                                                                                                                                                     | Attribute Id |
| :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------- |
| 0 - VIN decoded clean. Check Digit (9th position) is correct                                                                                                                                                                                                                                                                                                       | 0            |
| 1 - Check Digit (9th position) does not calculate properly                                                                                                                                                                                                                                                                                                         | 1            |
| 10 - Off-road Vehicle Warning – This is not a vehicle identification number (VIN) for a motor vehicle. This indicates that the manufacturer did not certify this product as complying with the Federal motor vehicle safety standards which are applicable to motor vehicles as defined at 49 U.S.C. 30102.                                                          | 10           |
| 400 - Invalid Characters Present                                                                                                                                                                                                                                                                                                                                   | 400          |