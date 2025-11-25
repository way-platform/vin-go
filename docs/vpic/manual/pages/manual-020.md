# Data Elements, Definitions, and Codes

## General: Manufacturer Name [VT] (lookup)

Column Names: ManufacturerFullNameId, ManufacturerFullName
Column Datatypes: int, varchar(135)

Name of the vehicle manufacturer.

The API provides XML lists in pages of 100 manufacturers per page:
`https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=XML&page=1`. For more information, please visit the vPIC API page at `https://vpic.nhtsa.dot.gov/api/`.