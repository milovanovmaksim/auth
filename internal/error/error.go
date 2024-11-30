package error


// InternalError внутренняя ошибка приложения.
// Удовлетворяет интерфейсу error.
type InternalError struct {
	String string
}


func (i InternalError) Error() string {
	return i.String
}