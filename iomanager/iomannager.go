package iomanager

type IOManager interface {
	Read() ([]string, error)
	WriteResults(data interface{}) error
}
