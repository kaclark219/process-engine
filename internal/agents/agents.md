# Types of Device Agents + Uses

## Pumps
Controls:
- Start/Stop (`Pump.Control.Start` + `Pump.Control.Stop`)
- Speed Control (`Pump.Control.Speed`)

Variables:
- Wear (`Pump.Wear`)
- Temperature (`Pump.Temperature`)
- Speed (`Pump.Speed`)
- Flow Rate (`Pump.Flow`)
- Discharge Presure (`Pump.DischargePressure`)
- Suction Pressure (`Pump.SuctionPressure`)

Alarms:
- Overcurrent (`Pump.Alarm.Overcurrent`)
- Loss of Flow (`Pump.Alarm.LowFlow`)
- Running Dry (`Pump.Alarm.Dry`)
- Exceeding Pressure (`Pump.Alarm.HighPressure`)

## Tanks
Variables:
- Level (`Tank.Level`)
- Temperature (`Tank.Temperature`)
- Pressure (`Tank.Pressure`)
- Inflow/Outflow Rates (`Tank.Inflow` + `Outflow.Inflow`)

Alarms:
- High/Low Level (`Tank.Alarm.HighLevel` + `Tank.Alarm.LowLevel`)
- High/Low Pressure (`Tank.Alarm.HighPressure` + `Tank.Alarm.LowPressure`)
- High/Low Temperature (`Tank.Alarm.HighTemperature` + `Tank.Alarm.LowTemperature`)

## Valves
Controls:
- Open/Close (`Valve.Control.Open` + `Valve.Control.Close`)
- Set Position (`Valve.Control.Position`)

Variables:
- Position (`Valve.Position`)
- Flow Rate (`Valve.Flow`)
- Pressure Drop (`Valve.Pressure`)

Alarms:
- Failure to Open/Close (`Valve.Alarm.Open` + `Valve.Alarm.Close`)
- Position Deviation (`Valve.Alarm.Position`)
- High Differential Pressure (`Valve.Alarm.Pressure`)

## Heat Exchangers
Controls:
- Setpoint for Outlet Temperature (`HX.Control.OutletTemperature`)
- Flow Control (`HX.Control.Flow`)

Variables:
- Fouling (`HX.Fouling`)
- Efficiency (`HX.Efficiency`)
- Inlet/Outlet Temperature (`HX.InletTemperature` + `HX.OutletTemperature`)
- Flow Rate (`HX.Flow`)
- Inlet/Outlet Pressure (`HX.InletPressure` + `HX.OutletPressure`)

Alarms:
- High/Low Temperature (`HX.Alarm.HighTemperature` + `HX.Alarm.LowTemperature`)
- Fouling Fouling (`HX.Alarm.Fouling`)
- Pressure Drop (`HX.Alarm.Pressure`)
- Flow Imbalance (`HX.Alarm.Flow`)

## Compressors
Controls:
- Start/Stop (`Compressor.Control.Start` + `Compressor.Control.Stop`)
- Speed Control (`Compressor.Control.Speed`)
- Discharge Pressure Setpoint (`Compressor.Control.DischargePressure`)

Variables:
- Suction Pressure (`Compressor.SuctionPressure`)
- Discharge Pressure (`Compressor.DischargePressure`)
- Inlet/Outlet Temperature (`Compressor.InletTemperature` + `Compressor.OutletTemperature`)
- Vibration Level (`Compressor.Vibration`)
- Motor Current (`Compressor.Current`)

Alarms:
- High Vibration (`Compressor.Alarm.Vibration`)
- High Discharge Temperature (`Compressor.Alarm.OutletTemperature`)
- Surge (`Compressor.Alarm.Surge`)
- Low Suction Pressure (`Compressor.Alarm.LowSuctionPressure`)