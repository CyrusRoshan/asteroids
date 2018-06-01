package physics

import "math"

var (
	G                          = 6.67 * math.Pow(10, -11)
	AssumedFrictionCoefficient = 0.6
)

func velocity(dt float64, ds float64) float64 {
	return ds / dt
}

func acceleration(dt float64, dv float64) float64 {
	return dv / dt
}

func gravity(m1 float64, m2 float64, r float64) float64 {
	return G * m1 * m2 / math.Pow(r, 2)
}

// welll.. I never said this was accurate, or a physics sim..
func frictionalForce(m float64, g float64) float64 {
	n := m * g
	return AssumedFrictionCoefficient * n
}

// not physics but w/e
func pythagoreanTheorem(dx float64, dy float64) float64 {
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
