package or

// Если длина переданных каналов равна 0 (т.е ничего не передано), то
// передаем канал и закрываем его.
// Если длина равна 1, возвращаем этот же канал
// Если длина равна 2, приходит значение и селект сам выберет, завершится
// и закроет канал
// Иначе рекурсивно запускаем эти функции, передавая первую половину и вторую
func Or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		out := make(chan interface{})
		close(out)
		return out
	case 1:
		return channels[0]
	case 2:
		out := make(chan interface{})
		go func() {
			defer close(out)

			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		}()
		return out
	default:
		mid := len(channels) / 2
		return Or(
			Or(channels[:mid]...),
			Or(channels[mid:]...),
		)

	}

}
