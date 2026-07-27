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

## Tanks
Variables:
- Level (`Tank.Level`)
- Temperature (`Tank.Temperature`)
- Pressure (`Tank.Pressure`)
- Inflow/Outflow Rates (`Tank.Inflow` + `Tank.Outflow`)

## Valves
Controls:
- Open/Close (`Valve.Control.Open` + `Valve.Control.Close`)
- Set Position (`Valve.Control.Position`)

Variables:
- Position (`Valve.Position`)
- Flow Rate (`Valve.Flow`)
- Pressure Drop (`Valve.Pressure`)

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