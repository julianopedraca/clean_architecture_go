package interfaces

type SptransApiInterface interface {
	Authentication() ([]byte, error)
	SearchLine(line string) ([]byte, error)
}
