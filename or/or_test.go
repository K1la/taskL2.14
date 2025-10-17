package or

import (
	"testing"
	"time"
)

// helper: создает канал, который закроется через d
func sigAfter(d time.Duration) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		time.Sleep(d)
	}()
	return ch
}

func TestOr_NoChannels(t *testing.T) {
	out := Or()
	select {
	case <-out:
		// ожидаем мгновенное закрытие
	default:
		t.Fatal("expected closed channel immediately for zero input")
	}
}

func TestOr_SingleChannel(t *testing.T) {
	ch := sigAfter(10 * time.Millisecond)
	start := time.Now()
	out := Or(ch)
	<-out
	if time.Since(start) < 10*time.Millisecond {
		t.Fatal("Or should return same channel and close after delay")
	}
}

func TestOr_TwoChannels_FirstCloses(t *testing.T) {
	ch1 := sigAfter(5 * time.Millisecond)
	ch2 := sigAfter(50 * time.Millisecond)

	start := time.Now()
	out := Or(ch1, ch2)
	<-out

	elapsed := time.Since(start)
	if elapsed > 20*time.Millisecond {
		t.Fatalf("expected early close, got delay %v", elapsed)
	}
}

func TestOr_TwoChannels_SecondCloses(t *testing.T) {
	ch1 := sigAfter(50 * time.Millisecond)
	ch2 := sigAfter(5 * time.Millisecond)

	start := time.Now()
	out := Or(ch1, ch2)
	<-out

	elapsed := time.Since(start)
	if elapsed > 20*time.Millisecond {
		t.Fatalf("expected early close when second closes, got delay %v", elapsed)
	}
}

func TestOr_MultipleChannels(t *testing.T) {
	ch1 := sigAfter(50 * time.Millisecond)
	ch2 := sigAfter(100 * time.Millisecond)
	ch3 := sigAfter(5 * time.Millisecond)
	ch4 := sigAfter(150 * time.Millisecond)

	start := time.Now()
	out := Or(ch1, ch2, ch3, ch4)
	<-out

	elapsed := time.Since(start)
	if elapsed > 20*time.Millisecond {
		t.Fatalf("expected early close for recursive case, got delay %v", elapsed)
	}
}
