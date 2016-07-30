package handlers

/*
0 — ОК,
1 — запрашиваемый объект не найден,
2 — невалидный запрос (например, не парсится json),
3 — некорректный запрос (семантически),
4 — неизвестная ошибка.
5 — такой юзер уже существует
 */
const (
	StatusOK = iota
	StatusNotFound
	StatusNotValid
	StatusNotCorrect
	StatusUnknown
	StatusUserCreated
)
