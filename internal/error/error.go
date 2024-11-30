package error


// InternalError внутренняя ошибка приложения.
// Удовлетворяет интерфейсу error.
type InternalError struct {
	String string
}


func (i InternalError) Error() string {
	return i.String
}


// ValidationError представляет ошибку валидации входных данных.
// Удовлетворяет интерфейсу error.
type ValidationError struct {
	String string
}

// Error возвращает текст ошибки.
func (v ValidationError) Error() string {
	return v.String
}