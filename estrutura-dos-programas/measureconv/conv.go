package measureconv

func CMToFT(cm Centimeter) Feet { return Feet(cm * 0.0328084) }
func CMToM(cm Centimeter) Meter { return Meter(cm * 0.01) }
func PToKG(p Pound) Kilogram    { return Kilogram(p * 0.453592) }
