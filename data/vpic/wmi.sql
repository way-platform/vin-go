SELECT
    Wmi.Id,
    Wmi.Wmi,
    Wmi.ManufacturerId,
    Manufacturer.Name AS ManufacturerName,
    Wmi.MakeId,
    Make.Name AS MakeName,
    Wmi.CountryId,
    Country.Name AS CountryName
FROM Wmi
    LEFT JOIN Manufacturer ON Wmi.ManufacturerId = Manufacturer.Id
    LEFT JOIN Make ON Wmi.MakeId = Make.Id
    LEFT JOIN Country ON Wmi.CountryId = Country.Id
