# Mechanical/Battery/Charger: Charger Level (lookup)

Column Names: ChargerLevelId, ChargerLevel
Column Datatypes: tinyint, varchar(135)

There are three levels of battery chargers currently. Level 1 and 2 are AC chargers; Level 3 is a DC charger. Chargers at each level charge batteries with different voltage and current levels. Level 3 is the fastest charging.

*   **Level 1**
    *   AC charger.
    *   In North America this typically means 16 amps at 120 volts delivering 1.9 kW of power.
    *   In Europe it may be 13 or 16 amps at 240 volts delivering 3 kW of power.
    *   The EV may incorporate a standard domestic power cord to connect the vehicle to a domestic socket outlet or a Level 1 charging station.
*   **Level 2**
    *   AC charger.
    *   It delivers up to 20 kW of power from either single- or three-phase alternating current (AC) sources of 208-240 volts at up to 80 amps.
*   **Level 3**
    *   DC charging, or "fast charging."
    *   To achieve very short charging times, Level 3 chargers supply very high currents of up to 400 amps at voltages up to 600 volts DC delivering a maximum power of 240 kW.

| Attribute Name                                                                                                                                                                         | Attribute Id |
| :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------- |
| Level 1 AC Charger (typically 16A 120V 1.9kW or 13-16A 240V 3kW) may incorporate standard domestic power cord.                                                                         | 1            |
| Level 2 AC Charger (up to 80A, 208-240V AC, up to 20kW from single- or three-phase AC) cables permanently fixed to charging station.                                                     | 2            |
| Level 3 DC Charger or fast charger (up to 400A, up to 600V DC, up to 240kW)                                                                                                            | 3            |