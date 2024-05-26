package measureconv

import "fmt"

type Centimeter float64
type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

func (cm Centimeter) String() string { return fmt.Sprintf("%g cm", cm) }
func (ft Feet) String() string       { return fmt.Sprintf("%g ft", ft) }
func (m Meter) String() string       { return fmt.Sprintf("%g m", m) }
func (p Pound) String() string       { return fmt.Sprintf("%g £", p) }
func (kg Kilogram) String() string   { return fmt.Sprintf("%g kg", kg) }
