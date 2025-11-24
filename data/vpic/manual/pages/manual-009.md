# Introduction

The following is a list of driver assistance technologies that are researched on the internet and coded into the vPIC supplemental module.

*   Adaptive Cruise Control (ACC)
*   Adaptive Driving Beam (ADB)
*   Anti-lock Braking System (ABS)
*   Automatic Crash Notification (ACN) / Advanced Automatic Crash Notification (AACN)
*   Automatic Pedestrian Alerting Sound (for Hybrid and EV only)
*   Auto-Reverse System for Windows and Sunroofs
*   Backup Camera
*   Blind Spot Intervention (BSI)
*   Blind Spot Warning (BSW)
*   Crash Imminent Braking (CIB)
*   Daytime Running Light (DRL)
*   Dynamic Brake Support (DBS)
*   Electronic Stability Control (ESC)
*   Event Data Recorder (EDR)
*   Forward Collision Warning (FCW)
*   Headlamp Light Source
*   Keyless Ignition
*   Lane Centering Assistance
*   Lane Departure Warning (LDW)
*   Lane Keeping Assistance (LKA)
*   Parking Assist
*   Pedestrian Automatic Emergency Braking (PAEB)
*   Rear Automatic Emergency Braking
*   Rear Cross Traffic Alert
*   Semiautomatic Headlamp Beam Switching
*   Tire Pressure Monitoring system (TPMS)
*   Traction Control

If manufacturers provide information about driver assistance technologies in their submissions, the manufacturer submitted information is used for VIN decoding instead of the supplemental information.

For motor vehicles and trailers involved in crashes, VINs are collected in NHTSA's crash data collection systems, the Fatality Analysis Reporting System (FARS), the Crash Report Sampling System (CRSS), and the Crash Investigation Sampling System (CISS), and are decoded using the vPIC VIN decoder. The vehicle information decoded by vPIC from a VIN is provided as auxiliary files for the systems. Two vPIC VIN decode files are created for each data collection system, one for motor vehicles (called *vPICDecode*) and one for trailing units (called *vPICTrailerDecode*).

In the vPIC VIN decode files, information will be included only for the VINs that can be cleanly decoded, i.e., with no major errors. See Section **Error Code** for the definition of cleanly decoded VINs.