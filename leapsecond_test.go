package leapsecond

import (
	"fmt"
	"testing"
	"time"
)

var haveTests = []struct {
	t    time.Time
	have time.Duration
}{
	{time.Date(1945, 1, 1, 0, 0, 0, 0, time.UTC), 0},
	{time.Date(1985, 11, 11, 0, 0, 0, 0, time.UTC), 13 * time.Second},
	{time.Date(1999, 9, 21, 0, 0, 0, 0, time.UTC), 22 * time.Second},
	{time.Date(2010, 12, 1, 0, 0, 0, 0, time.UTC), 24 * time.Second},
	{time.Date(2011, 3, 11, 0, 0, 0, 0, time.UTC), 24 * time.Second},
	{time.Date(2014, 3, 18, 0, 0, 0, 0, time.UTC), 25 * time.Second},
	{time.Date(2014, 9, 26, 0, 0, 0, 0, time.UTC), 25 * time.Second},
	{time.Date(2017, 3, 28, 0, 0, 0, 0, time.UTC), 27 * time.Second},
	{time.Date(2017, 8, 15, 0, 0, 0, 0, time.UTC), 27 * time.Second},
	{time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC), 27 * time.Second},
	{time.Date(2018, 4, 13, 0, 0, 0, 0, time.UTC), 27 * time.Second},
}

var diffTests = []struct {
	t1   time.Time
	t2   time.Time
	diff time.Duration
}{
	{time.Date(1985, 11, 11, 0, 0, 0, 0, time.UTC), time.Date(1999, 9, 21, 0, 0, 0, 0, time.UTC), 9 * time.Second},
	{time.Date(2010, 12, 1, 0, 0, 0, 0, time.UTC), time.Date(1999, 9, 21, 0, 0, 0, 0, time.UTC), 2 * time.Second},
	{time.Date(2011, 3, 11, 0, 0, 0, 0, time.UTC), time.Date(2014, 3, 18, 0, 0, 0, 0, time.UTC), 1 * time.Second},
	{time.Date(2017, 3, 28, 0, 0, 0, 0, time.UTC), time.Date(2017, 8, 15, 0, 0, 0, 0, time.UTC), 0},
	{time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC), time.Date(2018, 4, 13, 0, 0, 0, 0, time.UTC), 0},
}

func TestHave(t *testing.T) {
	for _, tt := range haveTests {
		t.Run(tt.t.String(), func(t *testing.T) {
			got := Have(tt.t)
			if got != tt.have {
				t.Errorf("got %q, want %q", got, tt.have)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	for _, tt := range diffTests {
		t.Run(tt.t1.String(), func(t *testing.T) {
			got := Diff(tt.t1, tt.t2)
			if got != tt.diff {
				t.Errorf("got %q, want %q", got, tt.diff)
			}
		})
	}
}

func BenchmarkHave(b *testing.B) {
	t := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		Have(t)
	}
}

func BenchmarkDiff(b *testing.B) {
	t1 := time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2010, 12, 1, 0, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		Diff(t1, t2)
	}
}

func ExampleHave() {
	t1 := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v", Have(t1))
	// Output: 27s
}

func ExampleDiff() {
	t1 := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2010, 12, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v", Diff(t1, t2))
	// Output: 3s
}
