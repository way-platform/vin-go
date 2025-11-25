# Document Convention for Data Elements, Definitions, and Codes

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

## Group and Sub-Group

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