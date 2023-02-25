package main

import (
	utils "example/project/mypackage"
)

type TeaSpoon float64
type TableSpoon float64
type Mililitres float64

func TeaSpoonToML(tsp TeaSpoon) Mililitres {
	return Mililitres(tsp * 4.92)
}

func TableSpoonToMililtres(tbs TableSpoon) Mililitres {
	return Mililitres(tbs * 14.79)
}

func (tsp TeaSpoon) ToMililitres() Mililitres {
	return Mililitres(tsp * 4.92)
}

func (tbs TableSpoon) ToMililitres() Mililitres {
	return Mililitres(tbs * 14.79)
}

func main() {
	ml1 := Mililitres(TeaSpoon(3) * 4.92)
	utils.F("3 TeaSpoons = %.2fml\n", ml1)

	ml2 := Mililitres(TableSpoon(3) * 14.79)
	utils.F("3 TableSpoons = %.2fml\n", ml2)

	utils.Pl("2 tsp + 4 tsp =", TeaSpoon(2)+TeaSpoon(4))
	utils.Pl("2 tsp > 4 tsp =", TeaSpoon(2) > TeaSpoon(4))

	utils.F("3 TeaSpoon = %.2fml\n", TeaSpoonToML(3))
	utils.F("3 TableSpoon = %.2fml\n", TeaSpoonToML(3))

	tsp1 := TeaSpoon(3)
	utils.F("%.2f tsp = %.2fml\n", tsp1, tsp1.ToMililitres())
}
