#!/bin/bash
# This script generates a CSV file with the 20 most recent data samples for each table in the database.

set -e

OUTPUT_DIR="data/vpic/example-data"
mkdir -p "$OUTPUT_DIR"
echo "Output directory '$OUTPUT_DIR' created or already exists."

QUERY_SCRIPT="./tools/query-vpic.sh"
if [ ! -f "$QUERY_SCRIPT" ]; then
    echo "Error: Query script not found at '$QUERY_SCRIPT'"
    exit 1
fi

echo "Querying sample data for ABS..."
echo 'SELECT TOP 20 * FROM [ABS] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ABS.csv"

echo "Querying sample data for AdaptiveCruiseControl..."
echo 'SELECT TOP 20 * FROM [AdaptiveCruiseControl] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AdaptiveCruiseControl.csv"

echo "Querying sample data for AdaptiveDrivingBeam..."
echo 'SELECT TOP 20 * FROM [AdaptiveDrivingBeam] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AdaptiveDrivingBeam.csv"

echo "Querying sample data for AirBagLocations..."
echo 'SELECT TOP 20 * FROM [AirBagLocations] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AirBagLocations.csv"

echo "Querying sample data for AirBagLocFront..."
echo 'SELECT TOP 20 * FROM [AirBagLocFront] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AirBagLocFront.csv"

echo "Querying sample data for AirBagLocKnee..."
echo 'SELECT TOP 20 * FROM [AirBagLocKnee] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AirBagLocKnee.csv"

echo "Querying sample data for AutoBrake..."
echo 'SELECT TOP 20 * FROM [AutoBrake] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AutoBrake.csv"

echo "Querying sample data for AutomaticPedestrainAlertingSound..."
echo 'SELECT TOP 20 * FROM [AutomaticPedestrainAlertingSound] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AutomaticPedestrainAlertingSound.csv"

echo "Querying sample data for AutoReverseSystem..."
echo 'SELECT TOP 20 * FROM [AutoReverseSystem] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AutoReverseSystem.csv"

echo "Querying sample data for AxleConfiguration..."
echo 'SELECT TOP 20 * FROM [AxleConfiguration] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/AxleConfiguration.csv"

echo "Querying sample data for BatteryType..."
echo 'SELECT TOP 20 * FROM [BatteryType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BatteryType.csv"

echo "Querying sample data for BedType..."
echo 'SELECT TOP 20 * FROM [BedType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BedType.csv"

echo "Querying sample data for BlindSpotIntervention..."
echo 'SELECT TOP 20 * FROM [BlindSpotIntervention] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BlindSpotIntervention.csv"

echo "Querying sample data for BlindSpotMonitoring..."
echo 'SELECT TOP 20 * FROM [BlindSpotMonitoring] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BlindSpotMonitoring.csv"

echo "Querying sample data for BodyCab..."
echo 'SELECT TOP 20 * FROM [BodyCab] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BodyCab.csv"

echo "Querying sample data for BodyStyle..."
echo 'SELECT TOP 20 * FROM [BodyStyle] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BodyStyle.csv"

echo "Querying sample data for BrakeSystem..."
echo 'SELECT TOP 20 * FROM [BrakeSystem] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BrakeSystem.csv"

echo "Querying sample data for BusFloorConfigType..."
echo 'SELECT TOP 20 * FROM [BusFloorConfigType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BusFloorConfigType.csv"

echo "Querying sample data for BusType..."
echo 'SELECT TOP 20 * FROM [BusType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/BusType.csv"

echo "Querying sample data for CAN_AACN..."
echo 'SELECT TOP 20 * FROM [CAN_AACN] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/CAN_AACN.csv"

echo "Querying sample data for ChargerLevel..."
echo 'SELECT TOP 20 * FROM [ChargerLevel] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ChargerLevel.csv"

echo "Querying sample data for CombinedBrakingSystem..."
echo 'SELECT TOP 20 * FROM [CombinedBrakingSystem] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/CombinedBrakingSystem.csv"

echo "Querying sample data for Conversion..."
echo 'SELECT TOP 20 * FROM [Conversion] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Conversion.csv"

echo "Querying sample data for CoolingType..."
echo 'SELECT TOP 20 * FROM [CoolingType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/CoolingType.csv"

echo "Querying sample data for Country..."
echo 'SELECT TOP 20 * FROM [Country] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Country.csv"

echo "Querying sample data for CustomMotorcycleType..."
echo 'SELECT TOP 20 * FROM [CustomMotorcycleType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/CustomMotorcycleType.csv"

echo "Querying sample data for DaytimeRunningLight..."
echo 'SELECT TOP 20 * FROM [DaytimeRunningLight] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DaytimeRunningLight.csv"

echo "Querying sample data for DecodingOutput..."
echo 'SELECT TOP 20 * FROM [DecodingOutput] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DecodingOutput.csv"

echo "Querying sample data for DefaultValue..."
echo 'SELECT TOP 20 * FROM [DefaultValue] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DefaultValue.csv"

echo "Querying sample data for DEFS_Body..."
echo 'SELECT TOP 20 * FROM [DEFS_Body] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DEFS_Body.csv"

echo "Querying sample data for DEFS_Make..."
echo 'SELECT TOP 20 * FROM [DEFS_Make] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DEFS_Make.csv"

echo "Querying sample data for DEFS_Model..."
echo 'SELECT TOP 20 * FROM [DEFS_Model] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DEFS_Model.csv"

echo "Querying sample data for DestinationMarket..."
echo 'SELECT TOP 20 * FROM [DestinationMarket] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DestinationMarket.csv"

echo "Querying sample data for DriveType..."
echo 'SELECT TOP 20 * FROM [DriveType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DriveType.csv"

echo "Querying sample data for DynamicBrakeSupport..."
echo 'SELECT TOP 20 * FROM [DynamicBrakeSupport] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/DynamicBrakeSupport.csv"

echo "Querying sample data for ECS..."
echo 'SELECT TOP 20 * FROM [ECS] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ECS.csv"

echo "Querying sample data for EDR..."
echo 'SELECT TOP 20 * FROM [EDR] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/EDR.csv"

echo "Querying sample data for ElectrificationLevel..."
echo 'SELECT TOP 20 * FROM [ElectrificationLevel] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ElectrificationLevel.csv"

echo "Querying sample data for Element..."
echo 'SELECT TOP 20 * FROM [Element] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Element.csv"

echo "Querying sample data for EngineConfiguration..."
echo 'SELECT TOP 20 * FROM [EngineConfiguration] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/EngineConfiguration.csv"

echo "Querying sample data for EngineModel..."
echo 'SELECT TOP 20 * FROM [EngineModel] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/EngineModel.csv"

echo "Querying sample data for EngineModelPattern..."
echo 'SELECT TOP 20 * FROM [EngineModelPattern] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/EngineModelPattern.csv"

echo "Querying sample data for EntertainmentSystem..."
echo 'SELECT TOP 20 * FROM [EntertainmentSystem] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/EntertainmentSystem.csv"

echo "Querying sample data for ErrorCode..."
echo 'SELECT TOP 20 * FROM [ErrorCode] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ErrorCode.csv"

echo "Querying sample data for EVDriveUnit..."
echo 'SELECT TOP 20 * FROM [EVDriveUnit] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/EVDriveUnit.csv"

echo "Querying sample data for ForwardCollisionWarning..."
echo 'SELECT TOP 20 * FROM [ForwardCollisionWarning] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ForwardCollisionWarning.csv"

echo "Querying sample data for FuelDeliveryType..."
echo 'SELECT TOP 20 * FROM [FuelDeliveryType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/FuelDeliveryType.csv"

echo "Querying sample data for FuelTankMaterial..."
echo 'SELECT TOP 20 * FROM [FuelTankMaterial] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/FuelTankMaterial.csv"

echo "Querying sample data for FuelTankType..."
echo 'SELECT TOP 20 * FROM [FuelTankType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/FuelTankType.csv"

echo "Querying sample data for FuelType..."
echo 'SELECT TOP 20 * FROM [FuelType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/FuelType.csv"

echo "Querying sample data for GrossVehicleWeightRating..."
echo 'SELECT TOP 20 * FROM [GrossVehicleWeightRating] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/GrossVehicleWeightRating.csv"

echo "Querying sample data for KeylessIgnition..."
echo 'SELECT TOP 20 * FROM [KeylessIgnition] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/KeylessIgnition.csv"

echo "Querying sample data for LaneCenteringAssistance..."
echo 'SELECT TOP 20 * FROM [LaneCenteringAssistance] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/LaneCenteringAssistance.csv"

echo "Querying sample data for LaneDepartureWarning..."
echo 'SELECT TOP 20 * FROM [LaneDepartureWarning] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/LaneDepartureWarning.csv"

echo "Querying sample data for LaneKeepSystem..."
echo 'SELECT TOP 20 * FROM [LaneKeepSystem] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/LaneKeepSystem.csv"

echo "Querying sample data for LowerBeamHeadlampLightSource..."
echo 'SELECT TOP 20 * FROM [LowerBeamHeadlampLightSource] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/LowerBeamHeadlampLightSource.csv"

echo "Querying sample data for Make..."
echo 'SELECT TOP 20 * FROM [Make] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Make.csv"

echo "Querying sample data for Make_Model..."
echo 'SELECT TOP 20 * FROM [Make_Model] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Make_Model.csv"

echo "Querying sample data for Manufacturer..."
echo 'SELECT TOP 20 * FROM [Manufacturer] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Manufacturer.csv"

echo "Querying sample data for Manufacturer_Make..."
echo 'SELECT TOP 20 * FROM [Manufacturer_Make] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Manufacturer_Make.csv"

echo "Querying sample data for Model..."
echo 'SELECT TOP 20 * FROM [Model] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Model.csv"

echo "Querying sample data for MotorcycleChassisType..."
echo 'SELECT TOP 20 * FROM [MotorcycleChassisType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/MotorcycleChassisType.csv"

echo "Querying sample data for MotorcycleSuspensionType..."
echo 'SELECT TOP 20 * FROM [MotorcycleSuspensionType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/MotorcycleSuspensionType.csv"

echo "Querying sample data for NonLandUse..."
echo 'SELECT TOP 20 * FROM [NonLandUse] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/NonLandUse.csv"

echo "Querying sample data for ParkAssist..."
echo 'SELECT TOP 20 * FROM [ParkAssist] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ParkAssist.csv"

echo "Querying sample data for Pattern..."
echo 'SELECT TOP 20 * FROM [Pattern] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Pattern.csv"

echo "Querying sample data for PedestrianAutomaticEmergencyBraking..."
echo 'SELECT TOP 20 * FROM [PedestrianAutomaticEmergencyBraking] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/PedestrianAutomaticEmergencyBraking.csv"

echo "Querying sample data for Pretensioner..."
echo 'SELECT TOP 20 * FROM [Pretensioner] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Pretensioner.csv"

echo "Querying sample data for RearAutomaticEmergencyBraking..."
echo 'SELECT TOP 20 * FROM [RearAutomaticEmergencyBraking] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/RearAutomaticEmergencyBraking.csv"

echo "Querying sample data for RearCrossTrafficAlert..."
echo 'SELECT TOP 20 * FROM [RearCrossTrafficAlert] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/RearCrossTrafficAlert.csv"

echo "Querying sample data for RearVisibilityCamera..."
echo 'SELECT TOP 20 * FROM [RearVisibilityCamera] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/RearVisibilityCamera.csv"

echo "Querying sample data for SeatBeltsAll..."
echo 'SELECT TOP 20 * FROM [SeatBeltsAll] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/SeatBeltsAll.csv"

echo "Querying sample data for SemiautomaticHeadlampBeamSwitching..."
echo 'SELECT TOP 20 * FROM [SemiautomaticHeadlampBeamSwitching] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/SemiautomaticHeadlampBeamSwitching.csv"

echo "Querying sample data for Steering..."
echo 'SELECT TOP 20 * FROM [Steering] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Steering.csv"

echo "Querying sample data for TPMS..."
echo 'SELECT TOP 20 * FROM [TPMS] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/TPMS.csv"

echo "Querying sample data for TractionControl..."
echo 'SELECT TOP 20 * FROM [TractionControl] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/TractionControl.csv"

echo "Querying sample data for TrailerBodyType..."
echo 'SELECT TOP 20 * FROM [TrailerBodyType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/TrailerBodyType.csv"

echo "Querying sample data for TrailerType..."
echo 'SELECT TOP 20 * FROM [TrailerType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/TrailerType.csv"

echo "Querying sample data for Transmission..."
echo 'SELECT TOP 20 * FROM [Transmission] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Transmission.csv"

echo "Querying sample data for Turbo..."
echo 'SELECT TOP 20 * FROM [Turbo] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Turbo.csv"

echo "Querying sample data for ValvetrainDesign..."
echo 'SELECT TOP 20 * FROM [ValvetrainDesign] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/ValvetrainDesign.csv"

echo "Querying sample data for VehicleSpecPattern..."
echo 'SELECT TOP 20 * FROM [VehicleSpecPattern] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VehicleSpecPattern.csv"

echo "Querying sample data for VehicleSpecSchema..."
echo 'SELECT TOP 20 * FROM [VehicleSpecSchema] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VehicleSpecSchema.csv"

echo "Querying sample data for VehicleSpecSchema_Model..."
echo 'SELECT TOP 20 * FROM [VehicleSpecSchema_Model] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VehicleSpecSchema_Model.csv"

echo "Querying sample data for VehicleSpecSchema_Year..."
echo 'SELECT TOP 20 * FROM [VehicleSpecSchema_Year] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VehicleSpecSchema_Year.csv"

echo "Querying sample data for VehicleType..."
echo 'SELECT TOP 20 * FROM [VehicleType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VehicleType.csv"

echo "Querying sample data for VinDescriptor..."
echo 'SELECT TOP 20 * FROM [VinDescriptor] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VinDescriptor.csv"

echo "Querying sample data for VinException..."
echo 'SELECT TOP 20 * FROM [VinException] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VinException.csv"

echo "Querying sample data for VinSchema..."
echo 'SELECT TOP 20 * FROM [VinSchema] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VinSchema.csv"

echo "Querying sample data for vNCSABodyType..."
echo 'SELECT TOP 20 * FROM [vNCSABodyType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/vNCSABodyType.csv"

echo "Querying sample data for vNCSAMake..."
echo 'SELECT TOP 20 * FROM [vNCSAMake] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/vNCSAMake.csv"

echo "Querying sample data for vNCSAModel..."
echo 'SELECT TOP 20 * FROM [vNCSAModel] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/vNCSAModel.csv"

echo "Querying sample data for VSpecSchemaPattern..."
echo 'SELECT TOP 20 * FROM [VSpecSchemaPattern] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/VSpecSchemaPattern.csv"

echo "Querying sample data for WheelBaseType..."
echo 'SELECT TOP 20 * FROM [WheelBaseType] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/WheelBaseType.csv"

echo "Querying sample data for WheelieMitigation..."
echo 'SELECT TOP 20 * FROM [WheelieMitigation] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/WheelieMitigation.csv"

echo "Querying sample data for Wmi..."
echo 'SELECT TOP 20 * FROM [Wmi] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Wmi.csv"

echo "Querying sample data for Wmi_VinSchema..."
echo 'SELECT TOP 20 * FROM [Wmi_VinSchema] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Wmi_VinSchema.csv"

echo "Querying sample data for WMIYearValidChars..."
echo 'SELECT TOP 20 * FROM [WMIYearValidChars] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/WMIYearValidChars.csv"

echo "Querying sample data for WMIYearValidChars_CacheExceptions..."
echo 'SELECT TOP 20 * FROM [WMIYearValidChars_CacheExceptions] ORDER BY Id DESC' | "$QUERY_SCRIPT" "$OUTPUT_DIR/WMIYearValidChars_CacheExceptions.csv"

# --- Tables without an 'Id' column ---
echo "Querying sample data for Wmi_Make..."
echo 'SELECT TOP 20 * FROM [Wmi_Make]' | "$QUERY_SCRIPT" "$OUTPUT_DIR/Wmi_Make.csv"

echo "Sample data generation complete."

echo "Querying sample data for AdaptiveDrivingBeam..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AdaptiveDrivingBeam] ORDER BY Id DESC') "$OUTPUT_DIR/AdaptiveDrivingBeam.csv"

echo "Querying sample data for AirBagLocations..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AirBagLocations] ORDER BY Id DESC') "$OUTPUT_DIR/AirBagLocations.csv"

echo "Querying sample data for AirBagLocFront..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AirBagLocFront] ORDER BY Id DESC') "$OUTPUT_DIR/AirBagLocFront.csv"

echo "Querying sample data for AirBagLocKnee..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AirBagLocKnee] ORDER BY Id DESC') "$OUTPUT_DIR/AirBagLocKnee.csv"

echo "Querying sample data for AutoBrake..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AutoBrake] ORDER BY Id DESC') "$OUTPUT_DIR/AutoBrake.csv"

echo "Querying sample data for AutomaticPedestrainAlertingSound..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AutomaticPedestrainAlertingSound] ORDER BY Id DESC') "$OUTPUT_DIR/AutomaticPedestrainAlertingSound.csv"

echo "Querying sample data for AutoReverseSystem..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AutoReverseSystem] ORDER BY Id DESC') "$OUTPUT_DIR/AutoReverseSystem.csv"

echo "Querying sample data for AxleConfiguration..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [AxleConfiguration] ORDER BY Id DESC') "$OUTPUT_DIR/AxleConfiguration.csv"

echo "Querying sample data for BatteryType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BatteryType] ORDER BY Id DESC') "$OUTPUT_DIR/BatteryType.csv"

echo "Querying sample data for BedType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BedType] ORDER BY Id DESC') "$OUTPUT_DIR/BedType.csv"

echo "Querying sample data for BlindSpotIntervention..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BlindSpotIntervention] ORDER BY Id DESC') "$OUTPUT_DIR/BlindSpotIntervention.csv"

echo "Querying sample data for BlindSpotMonitoring..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BlindSpotMonitoring] ORDER BY Id DESC') "$OUTPUT_DIR/BlindSpotMonitoring.csv"

echo "Querying sample data for BodyCab..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BodyCab] ORDER BY Id DESC') "$OUTPUT_DIR/BodyCab.csv"

echo "Querying sample data for BodyStyle..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BodyStyle] ORDER BY Id DESC') "$OUTPUT_DIR/BodyStyle.csv"

echo "Querying sample data for BrakeSystem..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BrakeSystem] ORDER BY Id DESC') "$OUTPUT_DIR/BrakeSystem.csv"

echo "Querying sample data for BusFloorConfigType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BusFloorConfigType] ORDER BY Id DESC') "$OUTPUT_DIR/BusFloorConfigType.csv"

echo "Querying sample data for BusType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [BusType] ORDER BY Id DESC') "$OUTPUT_DIR/BusType.csv"

echo "Querying sample data for CAN_AACN..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [CAN_AACN] ORDER BY Id DESC') "$OUTPUT_DIR/CAN_AACN.csv"

echo "Querying sample data for ChargerLevel..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [ChargerLevel] ORDER BY Id DESC') "$OUTPUT_DIR/ChargerLevel.csv"

echo "Querying sample data for CombinedBrakingSystem..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [CombinedBrakingSystem] ORDER BY Id DESC') "$OUTPUT_DIR/CombinedBrakingSystem.csv"

echo "Querying sample data for Conversion..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Conversion] ORDER BY Id DESC') "$OUTPUT_DIR/Conversion.csv"

echo "Querying sample data for CoolingType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [CoolingType] ORDER BY Id DESC') "$OUTPUT_DIR/CoolingType.csv"

echo "Querying sample data for Country..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Country] ORDER BY Id DESC') "$OUTPUT_DIR/Country.csv"

echo "Querying sample data for CustomMotorcycleType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [CustomMotorcycleType] ORDER BY Id DESC') "$OUTPUT_DIR/CustomMotorcycleType.csv"

echo "Querying sample data for DaytimeRunningLight..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DaytimeRunningLight] ORDER BY Id DESC') "$OUTPUT_DIR/DaytimeRunningLight.csv"

echo "Querying sample data for DecodingOutput..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DecodingOutput] ORDER BY Id DESC') "$OUTPUT_DIR/DecodingOutput.csv"

echo "Querying sample data for DefaultValue..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DefaultValue] ORDER BY Id DESC') "$OUTPUT_DIR/DefaultValue.csv"

echo "Querying sample data for DEFS_Body..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DEFS_Body] ORDER BY Id DESC') "$OUTPUT_DIR/DEFS_Body.csv"

echo "Querying sample data for DEFS_Make..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DEFS_Make] ORDER BY Id DESC') "$OUTPUT_DIR/DEFS_Make.csv"

echo "Querying sample data for DEFS_Model..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DEFS_Model] ORDER BY Id DESC') "$OUTPUT_DIR/DEFS_Model.csv"

echo "Querying sample data for DestinationMarket..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DestinationMarket] ORDER BY Id DESC') "$OUTPUT_DIR/DestinationMarket.csv"

echo "Querying sample data for DriveType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DriveType] ORDER BY Id DESC') "$OUTPUT_DIR/DriveType.csv"

echo "Querying sample data for DynamicBrakeSupport..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [DynamicBrakeSupport] ORDER BY Id DESC') "$OUTPUT_DIR/DynamicBrakeSupport.csv"

echo "Querying sample data for ECS..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [ECS] ORDER BY Id DESC') "$OUTPUT_DIR/ECS.csv"

echo "Querying sample data for EDR..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [EDR] ORDER BY Id DESC') "$OUTPUT_DIR/EDR.csv"

echo "Querying sample data for ElectrificationLevel..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [ElectrificationLevel] ORDER BY Id DESC') "$OUTPUT_DIR/ElectrificationLevel.csv"

echo "Querying sample data for Element..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Element] ORDER BY Id DESC') "$OUTPUT_DIR/Element.csv"

echo "Querying sample data for EngineConfiguration..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [EngineConfiguration] ORDER BY Id DESC') "$OUTPUT_DIR/EngineConfiguration.csv"

echo "Querying sample data for EngineModel..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [EngineModel] ORDER BY Id DESC') "$OUTPUT_DIR/EngineModel.csv"

echo "Querying sample data for EngineModelPattern..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [EngineModelPattern] ORDER BY Id DESC') "$OUTPUT_DIR/EngineModelPattern.csv"

echo "Querying sample data for EntertainmentSystem..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [EntertainmentSystem] ORDER BY Id DESC') "$OUTPUT_DIR/EntertainmentSystem.csv"

echo "Querying sample data for ErrorCode..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [ErrorCode] ORDER BY Id DESC') "$OUTPUT_DIR/ErrorCode.csv"

echo "Querying sample data for EVDriveUnit..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [EVDriveUnit] ORDER BY Id DESC') "$OUTPUT_DIR/EVDriveUnit.csv"

echo "Querying sample data for ForwardCollisionWarning..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [ForwardCollisionWarning] ORDER BY Id DESC') "$OUTPUT_DIR/ForwardCollisionWarning.csv"

echo "Querying sample data for FuelDeliveryType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [FuelDeliveryType] ORDER BY Id DESC') "$OUTPUT_DIR/FuelDeliveryType.csv"

echo "Querying sample data for FuelTankMaterial..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [FuelTankMaterial] ORDER BY Id DESC') "$OUTPUT_DIR/FuelTankMaterial.csv"

echo "Querying sample data for FuelTankType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [FuelTankType] ORDER BY Id DESC') "$OUTPUT_DIR/FuelTankType.csv"

echo "Querying sample data for FuelType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [FuelType] ORDER BY Id DESC') "$OUTPUT_DIR/FuelType.csv"

echo "Querying sample data for GrossVehicleWeightRating..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [GrossVehicleWeightRating] ORDER BY Id DESC') "$OUTPUT_DIR/GrossVehicleWeightRating.csv"

echo "Querying sample data for KeylessIgnition..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [KeylessIgnition] ORDER BY Id DESC') "$OUTPUT_DIR/KeylessIgnition.csv"

echo "Querying sample data for LaneCenteringAssistance..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [LaneCenteringAssistance] ORDER BY Id DESC') "$OUTPUT_DIR/LaneCenteringAssistance.csv"

echo "Querying sample data for LaneDepartureWarning..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [LaneDepartureWarning] ORDER BY Id DESC') "$OUTPUT_DIR/LaneDepartureWarning.csv"

echo "Querying sample data for LaneKeepSystem..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [LaneKeepSystem] ORDER BY Id DESC') "$OUTPUT_DIR/LaneKeepSystem.csv"

echo "Querying sample data for LowerBeamHeadlampLightSource..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [LowerBeamHeadlampLightSource] ORDER BY Id DESC') "$OUTPUT_DIR/LowerBeamHeadlampLightSource.csv"

echo "Querying sample data for Make..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Make] ORDER BY Id DESC') "$OUTPUT_DIR/Make.csv"

echo "Querying sample data for Make_Model..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Make_Model] ORDER BY Id DESC') "$OUTPUT_DIR/Make_Model.csv"

echo "Querying sample data for Manufacturer..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Manufacturer] ORDER BY Id DESC') "$OUTPUT_DIR/Manufacturer.csv"

echo "Querying sample data for Manufacturer_Make..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Manufacturer_Make] ORDER BY Id DESC') "$OUTPUT_DIR/Manufacturer_Make.csv"

echo "Querying sample data for Model..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Model] ORDER BY Id DESC') "$OUTPUT_DIR/Model.csv"

echo "Querying sample data for MotorcycleChassisType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [MotorcycleChassisType] ORDER BY Id DESC') "$OUTPUT_DIR/MotorcycleChassisType.csv"

echo "Querying sample data for MotorcycleSuspensionType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [MotorcycleSuspensionType] ORDER BY Id DESC') "$OUTPUT_DIR/MotorcycleSuspensionType.csv"

echo "Querying sample data for NonLandUse..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [NonLandUse] ORDER BY Id DESC') "$OUTPUT_DIR/NonLandUse.csv"

echo "Querying sample data for ParkAssist..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [ParkAssist] ORDER BY Id DESC') "$OUTPUT_DIR/ParkAssist.csv"

echo "Querying sample data for Pattern..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Pattern] ORDER BY Id DESC') "$OUTPUT_DIR/Pattern.csv"

echo "Querying sample data for PedestrianAutomaticEmergencyBraking..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [PedestrianAutomaticEmergencyBraking] ORDER BY Id DESC') "$OUTPUT_DIR/PedestrianAutomaticEmergencyBraking.csv"

echo "Querying sample data for Pretensioner..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Pretensioner] ORDER BY Id DESC') "$OUTPUT_DIR/Pretensioner.csv"

echo "Querying sample data for RearAutomaticEmergencyBraking..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [RearAutomaticEmergencyBraking] ORDER BY Id DESC') "$OUTPUT_DIR/RearAutomaticEmergencyBraking.csv"

echo "Querying sample data for RearCrossTrafficAlert..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [RearCrossTrafficAlert] ORDER BY Id DESC') "$OUTPUT_DIR/RearCrossTrafficAlert.csv"

echo "Querying sample data for RearVisibilityCamera..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [RearVisibilityCamera] ORDER BY Id DESC') "$OUTPUT_DIR/RearVisibilityCamera.csv"

echo "Querying sample data for SeatBeltsAll..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [SeatBeltsAll] ORDER BY Id DESC') "$OUTPUT_DIR/SeatBeltsAll.csv"

echo "Querying sample data for SemiautomaticHeadlampBeamSwitching..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [SemiautomaticHeadlampBeamSwitching] ORDER BY Id DESC') "$OUTPUT_DIR/SemiautomaticHeadlampBeamSwitching.csv"

echo "Querying sample data for Steering..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Steering] ORDER BY Id DESC') "$OUTPUT_DIR/Steering.csv"

echo "Querying sample data for TPMS..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [TPMS] ORDER BY Id DESC') "$OUTPUT_DIR/TPMS.csv"

echo "Querying sample data for TractionControl..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [TractionControl] ORDER BY Id DESC') "$OUTPUT_DIR/TractionControl.csv"

echo "Querying sample data for TrailerBodyType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [TrailerBodyType] ORDER BY Id DESC') "$OUTPUT_DIR/TrailerBodyType.csv"

echo "Querying sample data for TrailerType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [TrailerType] ORDER BY Id DESC') "$OUTPUT_DIR/TrailerType.csv"

echo "Querying sample data for Transmission..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Transmission] ORDER BY Id DESC') "$OUTPUT_DIR/Transmission.csv"

echo "Querying sample data for Turbo..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Turbo] ORDER BY Id DESC') "$OUTPUT_DIR/Turbo.csv"

echo "Querying sample data for ValvetrainDesign..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [ValvetrainDesign] ORDER BY Id DESC') "$OUTPUT_DIR/ValvetrainDesign.csv"

echo "Querying sample data for VehicleSpecPattern..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VehicleSpecPattern] ORDER BY Id DESC') "$OUTPUT_DIR/VehicleSpecPattern.csv"

echo "Querying sample data for VehicleSpecSchema..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VehicleSpecSchema] ORDER BY Id DESC') "$OUTPUT_DIR/VehicleSpecSchema.csv"

echo "Querying sample data for VehicleSpecSchema_Model..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VehicleSpecSchema_Model] ORDER BY Id DESC') "$OUTPUT_DIR/VehicleSpecSchema_Model.csv"

echo "Querying sample data for VehicleSpecSchema_Year..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VehicleSpecSchema_Year] ORDER BY Id DESC') "$OUTPUT_DIR/VehicleSpecSchema_Year.csv"

echo "Querying sample data for VehicleType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VehicleType] ORDER BY Id DESC') "$OUTPUT_DIR/VehicleType.csv"

echo "Querying sample data for VinDescriptor..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VinDescriptor] ORDER BY Id DESC') "$OUTPUT_DIR/VinDescriptor.csv"

echo "Querying sample data for VinException..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VinException] ORDER BY Id DESC') "$OUTPUT_DIR/VinException.csv"

echo "Querying sample data for VinSchema..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VinSchema] ORDER BY Id DESC') "$OUTPUT_DIR/VinSchema.csv"

echo "Querying sample data for vNCSABodyType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [vNCSABodyType] ORDER BY Id DESC') "$OUTPUT_DIR/vNCSABodyType.csv"

echo "Querying sample data for vNCSAMake..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [vNCSAMake] ORDER BY Id DESC') "$OUTPUT_DIR/vNCSAMake.csv"

echo "Querying sample data for vNCSAModel..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [vNCSAModel] ORDER BY Id DESC') "$OUTPUT_DIR/vNCSAModel.csv"

echo "Querying sample data for VSpecSchemaPattern..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [VSpecSchemaPattern] ORDER BY Id DESC') "$OUTPUT_DIR/VSpecSchemaPattern.csv"

echo "Querying sample data for WheelBaseType..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [WheelBaseType] ORDER BY Id DESC') "$OUTPUT_DIR/WheelBaseType.csv"

echo "Querying sample data for WheelieMitigation..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [WheelieMitigation] ORDER BY Id DESC') "$OUTPUT_DIR/WheelieMitigation.csv"

echo "Querying sample data for Wmi..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Wmi] ORDER BY Id DESC') "$OUTPUT_DIR/Wmi.csv"

echo "Querying sample data for Wmi_VinSchema..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Wmi_VinSchema] ORDER BY Id DESC') "$OUTPUT_DIR/Wmi_VinSchema.csv"

echo "Querying sample data for WMIYearValidChars..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [WMIYearValidChars] ORDER BY Id DESC') "$OUTPUT_DIR/WMIYearValidChars.csv"

echo "Querying sample data for WMIYearValidChars_CacheExceptions..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [WMIYearValidChars_CacheExceptions] ORDER BY Id DESC') "$OUTPUT_DIR/WMIYearValidChars_CacheExceptions.csv"

# --- Tables without an 'Id' column ---
echo "Querying sample data for Wmi_Make..."
"$QUERY_SCRIPT" <(echo 'SELECT TOP 20 * FROM [Wmi_Make]') "$OUTPUT_DIR/Wmi_Make.csv"

echo "Sample data generation complete."
