package controller

type ErrorConverter interface {
	Convert(err error) int
}
type ErrorConvertImpl struct {
}

func (e ErrorConvertImpl) Convert(err error) int {
	//TODO implement me
	panic("implement me")
}

var _ ErrorConverter = &ErrorConvertImpl{}
