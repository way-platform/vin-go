## General: Model [VT] (lookup)

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