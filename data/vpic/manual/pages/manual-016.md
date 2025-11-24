## Document Convention for Data Elements, Definitions, and Codes

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