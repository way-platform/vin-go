# Error Code

When a VIN is decoded, one or more errors might be provided in the output. Cleanly decoded VINs are those that have only one or more of the following error codes when they are decoded.

*   0 - VIN decoded clean. Check Digit (9th position) is correct.
*   1 - Check Digit (9th position) does not calculate properly.
*   10 - Off-road Vehicle Warning - This is not a vehicle identification number (VIN) for a motor vehicle. This indicates that the manufacturer did not certify this product as complying with the Federal motor vehicle safety standards which are applicable to motor vehicles as defined at 49 U.S.C. 30102.
*   400 - Invalid Characters Present

When a VIN is decoded, there can be errors applicable to the decode. In such a situation, the error codes are concatenated and separated by a comma. For example, error code combinations for cleanly decoded VINs can include (0), (0,10), (1,10), (1, 400), (1, 10, 400).

In the VIN decode files, even though **Error Code** data element is lookup data type, only concatenated error code IDs, not the concatenated error code names, are presented due to the size concern of the concatenated name string.