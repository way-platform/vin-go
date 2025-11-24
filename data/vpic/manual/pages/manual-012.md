# Vehicle Descriptor

To avoid releasing the full VIN that might be used to link to personally identifiable information, vPIC decode file only contains vehicle descriptors. A vehicle descriptor is 17 characters long. It has the same characters as the VIN except that the check digit number and sequential numbers are replaced by asterisks (*). The check digit is the 9th position. The positions for sequential numbers depend on the manufacturer production capabilities.

*   If the 3rd position of the VIN is 9, i.e., the vehicle manufacturer is a low-volume manufacturer, positions 9 and 15 to 17 are replaced with asterisks (*).
*   If the 3rd position of the VIN is not 9, i.e., the vehicle manufacturer is a high-volume manufacturer, positions 9 and 12 to 17 are replaced with asterisks (*).