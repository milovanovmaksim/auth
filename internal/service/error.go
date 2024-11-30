package service

// ValidationError представляет ошибку валидации входных данных.
// Удовлетворяет интерфейсу erors.Error.
type ValidationError struct {
	String string
}

// Error возвращает текст ошибки.
func (v ValidationError) Error() string {
	return v.String
}
