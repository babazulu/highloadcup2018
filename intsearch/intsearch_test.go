package intsearch

import (
	"math/rand"
	"sort"
	"testing"
)

const maxLimit = 1e6

var Ints []uint32

func fillInts() {
	rand.Seed(0)

	Ints = make([]uint32, maxLimit)

	for i := 0; i < maxLimit; i++ {
		Ints[i] = uint32(rand.Int() % maxLimit)
	}
}

func benchmarkSearch(b *testing.B, limit int, search func([]uint32, uint32) uint32) {

	if Ints == nil {
		fillInts()
	}

	ints := make([]uint32, limit)
	for i := range ints {
		ints[i] = Ints[rand.Int()%maxLimit]
	}

	sort.Slice(ints, func(i, j int) bool { return ints[i] < ints[j] })

	delta := uint32(maxLimit / limit)
	elt := uint32(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(ints, elt)
		elt += delta
		if elt >= maxLimit {
			elt = 0
		}
	}
}

func BenchmarkInterp100(b *testing.B)   { benchmarkSearch(b, 100, SearchInts) }
func BenchmarkInterp1000(b *testing.B)  { benchmarkSearch(b, 1000, SearchInts) }
func BenchmarkInterp10000(b *testing.B) { benchmarkSearch(b, 10000, SearchInts) }
func BenchmarkInterp1e5(b *testing.B)   { benchmarkSearch(b, 1e5, SearchInts) }
func BenchmarkInterp1e6(b *testing.B)   { benchmarkSearch(b, 1e6, SearchInts) }

func BenchmarkStd100(b *testing.B)   { benchmarkSearch(b, 100, StdSearchInts) }
func BenchmarkStd1000(b *testing.B)  { benchmarkSearch(b, 1000, StdSearchInts) }
func BenchmarkStd10000(b *testing.B) { benchmarkSearch(b, 10000, StdSearchInts) }
func BenchmarkStd1e5(b *testing.B)   { benchmarkSearch(b, 1e5, StdSearchInts) }
func BenchmarkStd1e6(b *testing.B)   { benchmarkSearch(b, 1e6, StdSearchInts) }

func BenchmarkAsm100(b *testing.B)   { benchmarkSearch(b, 100, AsmSearchInts) }
func BenchmarkAsm1000(b *testing.B)  { benchmarkSearch(b, 1000, AsmSearchInts) }
func BenchmarkAsm10000(b *testing.B) { benchmarkSearch(b, 10000, AsmSearchInts) }
func BenchmarkAsm1e5(b *testing.B)   { benchmarkSearch(b, 1e5, AsmSearchInts) }
func BenchmarkAsm1e6(b *testing.B)   { benchmarkSearch(b, 1e6, AsmSearchInts) }

func BenchmarkBin100(b *testing.B)   { benchmarkSearch(b, 100, BinSearchInts) }
func BenchmarkBin200(b *testing.B)   { benchmarkSearch(b, 200, BinSearchInts) }
func BenchmarkBin500(b *testing.B)   { benchmarkSearch(b, 500, BinSearchInts) }
func BenchmarkBin1000(b *testing.B)  { benchmarkSearch(b, 1000, BinSearchInts) }
func BenchmarkBin10000(b *testing.B) { benchmarkSearch(b, 10000, BinSearchInts) }
func BenchmarkBin1e5(b *testing.B)   { benchmarkSearch(b, 1e5, BinSearchInts) }
func BenchmarkBin1e6(b *testing.B)   { benchmarkSearch(b, 1e6, BinSearchInts) }

func BenchmarkBinApprox100(b *testing.B)   { benchmarkSearch(b, 100, BinApproxSearchInts) }
func BenchmarkBinApprox200(b *testing.B)   { benchmarkSearch(b, 200, BinApproxSearchInts) }
func BenchmarkBinApprox500(b *testing.B)   { benchmarkSearch(b, 500, BinApproxSearchInts) }
func BenchmarkBinApprox1000(b *testing.B)  { benchmarkSearch(b, 1000, BinApproxSearchInts) }
func BenchmarkBinApprox10000(b *testing.B) { benchmarkSearch(b, 10000, BinApproxSearchInts) }
func BenchmarkBinApprox1e5(b *testing.B)   { benchmarkSearch(b, 1e5, BinApproxSearchInts) }
func BenchmarkBinApprox1e6(b *testing.B)   { benchmarkSearch(b, 1e6, BinApproxSearchInts) }

func TestSearchSmall(t *testing.T) {

	rand.Seed(0)

	const limit = 100

	ints := make([]uint32, limit)

	for i := range ints {
		ints[i] = uint32(i)
	}

	for want, q := range ints {
		if idx := AsmSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := BinSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := BinApproxSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := StdSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := SearchInts(ints, q); idx != uint32(want) {
			t.Errorf("SearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
	}

	for i := 0; i < 100; i++ {

		q := uint32(rand.Int() % (limit + 2))

		want := sort.Search(len(ints), func(ii int) bool { return ints[ii] >= q })

		if idx := AsmSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := BinSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := BinApproxSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := StdSearchInts(ints, q); idx != uint32(want) {
			t.Errorf("StdSearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
		if idx := SearchInts(ints, q); idx != uint32(want) {
			t.Errorf("SearchInts(ints, %v)=%v, want %v", q, idx, want)
		}
	}
}
