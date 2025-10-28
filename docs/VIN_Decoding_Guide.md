# Vehicle Identification Number (VIN) Decoding Guide

This document provides a comprehensive guide to understanding and decoding a Vehicle Identification Number (VIN) based on the ISO 3779 standard.

## Introduction

The Vehicle Identification Number (VIN) is a unique 17-character alphanumeric code assigned to every motor vehicle. The VIN system is designed to provide a universal method for identifying vehicles, and it contains specific information about the vehicle's manufacturer, model, and features. The standard for VINs is defined by ISO 3779.

The VIN is composed of three main sections:

1.  **World Manufacturer Identifier (WMI):** Characters 1-3
2.  **Vehicle Descriptor Section (VDS):** Characters 4-9
3.  **Vehicle Identifier Section (VIS):** Characters 10-17

---

## 1. World Manufacturer Identifier (WMI)

The first three characters of the VIN uniquely identify the manufacturer of the vehicle.

- **Character 1:** Represents the geographic area or country where the vehicle was manufactured.
- **Character 2:** In combination with the first character, this identifies the country of manufacture.
- **Character 3:** In combination with the first two characters, this identifies the specific manufacturer. For manufacturers that produce fewer than 1000 vehicles per year, this character is often a '9'.

---

## 2. Vehicle Descriptor Section (VDS)

Characters 4 through 9 of the VIN provide a detailed description of the vehicle.

- **Characters 4-8:** These characters are used by the manufacturer to identify vehicle attributes such as the model, body style, engine type, and transmission. The meaning of these characters is specific to each manufacturer.
- **Character 9: Check Digit:** This is a security code that is used to verify the authenticity of the VIN. It is calculated based on the other characters in the VIN using a weighted mathematical formula.

### Check Digit Calculation Algorithm

The check digit is calculated as follows:

1.  **Assign Numeric Values:** Each letter in the VIN is assigned a numeric value.
2.  **Apply Weights:** Each position in the VIN is assigned a specific weight.
3.  **Multiply and Sum:** Multiply the numeric value of each character by its weight and sum the results.
4.  **Calculate Remainder:** Divide the total sum by 11. The remainder is the check digit. If the remainder is 10, the check digit is represented by the letter 'X'.

---

## 3. Vehicle Identifier Section (VIS)

The final section of the VIN is used to uniquely identify the specific vehicle.

- **Character 10: Model Year:** This character indicates the vehicle's model year. The coding is standardized, using a combination of letters and numbers that cycle every 30 years.
- **Character 11: Plant Code:** This character identifies the specific manufacturing plant where the vehicle was assembled. The codes are assigned by the manufacturer.
- **Characters 12-17: Serial Number:** These characters represent the unique serial number assigned to the vehicle during production.
