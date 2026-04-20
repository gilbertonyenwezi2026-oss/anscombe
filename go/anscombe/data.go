package anscombe

// Dataset holds a single x/y pair series from the Anscombe Quartet.
type Dataset struct {
	Name string
	X    []float64
	Y    []float64
}

// Quartet returns the four classic Anscombe datasets.
func Quartet() []Dataset {
	return []Dataset{
		{
			Name: "Set I",
			X:    []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y:    []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		{
			Name: "Set II",
			X:    []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y:    []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
		},
		{
			Name: "Set III",
			X:    []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y:    []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		{
			Name: "Set IV",
			X:    []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
			Y:    []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89},
		},
	}
}
