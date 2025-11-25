## General: Model Year [VT] (int)

Column Names: ModelYear
Column Datatypes: int

If the model year (MY) is supplied when the VIN is decoded, such as from a crash report or a vehicle registration record, the MY value will be the supplied MY, even if the MY decoded from the VIN differs from the supplied MY. If the MY is not supplied when the VIN is decoded, the MY value will be decoded from the 10th character in the VIN.