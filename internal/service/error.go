package service

// ValidationError представляет ошибку валидации входных данных.
// Удовлетворяет интерфейсу error.
type ValidationError struct {
	String string
}

// Error возвращает текст ошибки.
func (v ValidationError) Error() string {
	return v.String
}
