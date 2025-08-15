package main

import (
	"fmt"
	"time"
)

// Если длина переданных каналов равна 0 (т.е ничего не передано), то
// передаем канал и закрываем его.
// Если длина равна 1, возвращаем этот же канал
// Если длина равна 2, приходит значение и селект сам выберет, завершится
// и закроет канал
// Иначе рекурсивно запускаем эти функции, передавая первую половину и вторую
func or(channels ...<-chan interface{}) <-chan interface{} {
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
		return or(
			or(channels[:mid]...),
			or(channels[mid:]...),
		)

	}

}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or()
	fmt.Printf("done after %v\n", time.Since(start))
}
