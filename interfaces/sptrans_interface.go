package interfaces

type BusLine struct {
	Cl int32
	Lc bool
	Lt string
	Sl int32
	Tl int32
	Tp string
	Ts string
}

type SptransApiInterface interface {
	Authentication() ([]byte, error)
	SearchLine(line string) ([]BusLine, error)
}
