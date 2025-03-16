package interfaces

type SearchLine struct {
	Cl int32
	Lc bool
	Lt string
	Sl int32
	Tl int32
	Tp string
	Ts string
}

type SearchStop struct {
	Cp int32
	Np string
	Ed string
	Py float32
	Px float32
}

type SptransApiInterface interface {
	Authentication() ([]byte, error)
	SearchLine(line string) ([]SearchLine, error)
	SearchLineDirection(line string, direction string) ([]SearchLine, error)
	SearchStops(stop string) ([]SearchStop, error)
	SearchStopsByLine(line string) ([]SearchStop, error)
}
