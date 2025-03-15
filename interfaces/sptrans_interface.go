package interfaces

type SptransApiInterface interface {
	Authentication() ([]byte, error)
	SearchLine(lineNumber int) ([]byte, error)
}
