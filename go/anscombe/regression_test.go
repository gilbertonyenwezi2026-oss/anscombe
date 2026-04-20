package anscombe

import (
	"math"
	"testing"
)

const tol = 1e-3

const (
	knownTol   = 1e-3
	compareTol = 1e-6
)

func TestOLSMatchesKnownCoefficients(t *testing.T) {
	for _, ds := range Quartet() {
		intercept, slope, _, _, err := OLS(ds.X, ds.Y)
		if err != nil {
			t.Fatalf("OLS(%s) error = %v", ds.Name, err)
		}
		if math.Abs(intercept-3.0) > knownTol {
			t.Fatalf("OLS(%s) intercept = %v, want %v", ds.Name, intercept, 3.0)
		}
		if math.Abs(slope-0.5) > knownTol {
			t.Fatalf("OLS(%s) slope = %v, want %v", ds.Name, slope, 0.5)
		}
	}
}

func TestStatsPackageMatchesOLS(t *testing.T) {
	for _, ds := range Quartet() {
		wantIntercept, wantSlope, _, _, err := OLS(ds.X, ds.Y)
		if err != nil {
			t.Fatalf("OLS(%s) error = %v", ds.Name, err)
		}

		gotIntercept, gotSlope, err := StatsPkgRegression(ds.X, ds.Y)
		if err != nil {
			t.Fatalf("StatsPkgRegression(%s) error = %v", ds.Name, err)
		}

		if math.Abs(gotIntercept-wantIntercept) > compareTol {
			t.Fatalf("StatsPkgRegression(%s) intercept = %v, want %v", ds.Name, gotIntercept, wantIntercept)
		}
		if math.Abs(gotSlope-wantSlope) > compareTol {
			t.Fatalf("StatsPkgRegression(%s) slope = %v, want %v", ds.Name, gotSlope, wantSlope)
		}
	}
}

func BenchmarkRunQuartet(b *testing.B) {
	data := Quartet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ds := range data {
			_, _, _, _, _ = OLS(ds.X, ds.Y)
		}
	}
}

func BenchmarkStatsPackageQuartet(b *testing.B) {
	data := Quartet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ds := range data {
			_, _, _ = StatsPkgRegression(ds.X, ds.Y)
		}
	}
}
