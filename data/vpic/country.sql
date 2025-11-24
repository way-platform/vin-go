SELECT
    Country.Id,
    Country.Name,
    COUNT(Wmi.CountryId) as WmiCount
FROM Country
LEFT OUTER JOIN Wmi ON Wmi.CountryId = Country.Id
GROUP BY Country.Id, Country.Name
ORDER BY WmiCount DESC
