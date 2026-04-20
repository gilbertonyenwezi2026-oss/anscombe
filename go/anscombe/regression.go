package anscombe

import (
	"errors"
	"math"

	stats "github.com/montanaflynn/stats"
)

// Result stores simple linear regression outputs for one dataset.
type Result struct {
	Name      string  `json:"name"`
	Intercept float64 `json:"intercept"`
	Slope     float64 `json:"slope"`
	R         float64 `json:"r"`
	RSquared  float64 `json:"r_squared"`
}

// OLS computes intercept, slope, Pearson correlation, and R-squared.
func OLS(x, y []float64) (intercept, slope, r, rSquared float64, err error) {
	if len(x) != len(y) {
		return 0, 0, 0, 0, errors.New("x and y must have same length")
	}
	if len(x) < 2 {
		return 0, 0, 0, 0, errors.New("need at least two observations")
	}

	var sumX, sumY float64
	for i := range x {
		sumX += x[i]
		sumY += y[i]
	}
	meanX := sumX / float64(len(x))
	meanY := sumY / float64(len(y))

	var ssXX, ssXY, ssYY float64
	for i := range x {
		dx := x[i] - meanX
		dy := y[i] - meanY
		ssXX += dx * dx
		ssXY += dx * dy
		ssYY += dy * dy
	}
	if ssXX == 0 {
		return 0, 0, 0, 0, errors.New("x variance is zero")
	}

	slope = ssXY / ssXX
	intercept = meanY - slope*meanX
	if ssYY > 0 {
		r = ssXY / math.Sqrt(ssXX*ssYY)
		rSquared = r * r
	}
	return intercept, slope, r, rSquared, nil
}

// StatsPkgRegression computes coefficients by fitting the external stats package's
// regression line and reconstructing slope/intercept from predicted endpoints.
func StatsPkgRegression(x, y []float64) (intercept, slope float64, err error) {
	series := make(stats.Series, len(x))
	for i := range x {
		series[i] = stats.Coordinate{X: x[i], Y: y[i]}
	}

	reg, err := stats.LinearRegression(series)
	if err != nil {
		return 0, 0, err
	}
	if len(reg) < 2 {
		return 0, 0, errors.New("regression output too short")
	}

	minPt := reg[0]
	maxPt := reg[0]

	for _, pt := range reg[1:] {
		if pt.X < minPt.X {
			minPt = pt
		}
		if pt.X > maxPt.X {
			maxPt = pt
		}
	}

	if maxPt.X == minPt.X {
		return 0, 0, errors.New("cannot derive slope from regression output with only one unique x value")
	}

	slope = (maxPt.Y - minPt.Y) / (maxPt.X - minPt.X)
	intercept = minPt.Y - slope*minPt.X
	return intercept, slope, nil
}

// RunQuartet computes OLS results for all four datasets.
func RunQuartet() ([]Result, error) {
	results := make([]Result, 0, 4)
	for _, ds := range Quartet() {
		intercept, slope, r, rSquared, err := OLS(ds.X, ds.Y)
		if err != nil {
			return nil, err
		}
		results = append(results, Result{
			Name:      ds.Name,
			Intercept: intercept,
			Slope:     slope,
			R:         r,
			RSquared:  rSquared,
		})
	}
	return results, nil
}
