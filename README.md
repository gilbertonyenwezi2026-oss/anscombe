# Go vs Python vs R: Statistical Validation & Performance Benchmark  
## Anscombe Quartet Analysis


 ## Overview
This project evaluates whether the Go programming language can serve as a viable alternative to Python and R for data science tasks. Specifically, we test the statistical accuracy and performance of Go using the **Anscombe Quartet**, a classic dataset designed to highlight the importance of data visualization.

- Validate Go’s statistical correctness using the **Anscombe Quartet**
- Compare regression outputs across:
  - Go
  - Python
  - R
- Benchmark:
  - Execution time
  - Memory usage
- Provide a **recommendation to management**

---
## Dataset: Anscombe Quartet
The Anscombe Quartet consists of four datasets that:
- Have **identical summary statistics**
- Produce the **same linear regression equation**
- Exhibit **very different distributions when visualized**

This makes it ideal for testing both:
- Statistical correctness  
- Analytical interpretation  

This demonstrates:
> Statistical equivalence does not imply structural similarity

---

## Technologies Used

| Language | Libraries / Packages |
|----------|---------------------|
| Go       | `github.com/montanaflynn/stats`, `testing` |
| Python   | `numpy`, `statsmodels` |
| R        | `lm` (base R) |

---

## Methodology

1. Perform **linear regression (y ~ x)** on all four datasets
2. Validate:
   - Intercept
   - Slope
   - Correlation (r)
   - R²
3. Compare outputs across Go, Python, and R
4. Benchmark:
   - Execution time
   - Memory usage (where available)
5. Validate correctness using **tolerance-based testing**

---

## Results

### Statistical Accuracy

| Dataset | Intercept | Slope | r     | R²     |
|--------|----------|-------|---------|--------|
| Set I  | ~3.000   | ~0.500 | ~0.816 | ~0.666 |
| Set II | ~3.001   | ~0.500 | ~0.816 | ~0.666 |
| Set III| ~3.002   | ~0.500 | ~0.816 | ~0.666 |
| Set IV | ~3.002   | ~0.500 | ~0.816 | ~0.666 |

Results are **consistent across Go, Python, and R**

---

### Go Benchmark Results (Execution Time)

Benchmarks were executed in an unrestricted Linux environment.

| Benchmark                    | ns/op | B/op | allocs/op |
|------------------------------|------:|-----:|----------:|
| BenchmarkRunQuartet          | 152.5 |  0   | 0         |
| BenchmarkStatsPackageQuartet | 1527  | 2496 | 16        |

**Interpretation:**  
The custom Go OLS implementation significantly outperforms the external stats-package wrapper. It achieves extremely low latency and zero memory allocation, making it highly suitable for high-throughput production workloads.


### Cross-Language Performance Summary

| Language / Implementation     | Execution Time | Memory / Allocations    | Notes                                  |
|-------------------------------|---------------:|-------------------------|----------------------------------------|
| Go (custom OLS benchmark)     | 152.5 ns/op    | 0 B/op, 0 allocs/op     | Fastest and most efficient             |
| Go (stats package benchmark)  | 1527 ns/op     | 2496 B/op, 16 allocs/op | Correct but less efficient             |
| Python                        | 0.010236 sec   | 12,164 bytes peak       | Very fast, strong usability            |
| R                             | 0.492317 sec   | Not captured locally    | Correct but slowest in this comparison |

---

### Developer Productivity & Ecosystem (Qualitative)

| Category             | Go        | Python                       | R                     |
| -------------------- | --------- | ---------------------------- | --------------------- |
| Ease of Use          | Moderate  | Easy                         | Easy                  |
| Libraries (ML/Stats) | Limited   | Extensive                    | Extensive             |
| Visualization        | Limited   | Strong (Matplotlib, Seaborn) | Very Strong (ggplot2) |
| Performance          | High      | Moderate                     | Moderate              |
| Production Ready     | Excellent | Good                         | Limited               |

---

## Key Insights

### 1. Statistical Correctness (Accuracy is NOT the issue)
Go produces results that are **fully consistent** with Python and R.

### 2. Performance Advantage
Go significantly outperforms both Python and R in:
- Execution speed  
- Memory efficiency  
Ideal for **production systems**

### 3. Visualization Matters
Despite identical regression outputs, the Anscombe Quartet demonstrates:
- The datasets behave differently
- Regression alone can be misleading 
> Statistical results alone can be misleading without visualization.

### 4. Ecosystem Differences
- Python & R → Strong ecosystems for analytics, ML, visualization  
- Go → Strong for backend systems and performance-critical workloads  

---

## Limitations

- Go lacks mature libraries for:
  - Advanced statistical modeling  
  - Machine learning  
  - Visualization  
- Less support for interactive analysis (e.g., Jupyter, RStudio)

---

## Testing & Validation

- Unit tests confirm:
  - Regression coefficients match expected values
- Tolerance-based comparisons handle floating-point precision
- Benchmarks implemented using Go’s `testing` package

## Environment Notes

Local execution on a managed Windows system was partially restricted by **Application Control policies**, which blocked execution of temporary Go test binaries and R scripts.

Final benchmarking and validation were completed in an **unrestricted Linux environment**, ensuring accurate performance measurements.

---

## Final Conclusion

Benchmarking confirms that Go is the strongest language in this comparison for raw execution efficiency. The custom OLS implementation was an order of magnitude faster than the external Go stats-package wrapper and dramatically faster than the Python and R script timings collected for the same problem. However, Python and R remain more convenient for exploratory analysis and modeling. These findings support a hybrid recommendation: use Go for production-grade analytics and backend services, while retaining Python and R for data science workflows that require rapid experimentation, visualization, and richer statistical ecosystems.

---

## Recommendation to Management

While Go matches Python and R in statistical accuracy, it significantly outperforms both in execution speed and memory efficiency. However, Python and R remain superior for exploratory analysis, visualization, and advanced statistical modeling due to their mature ecosystems. These results support a hybrid approach, where Go is used for performance-critical production systems, and Python/R are used for analytical workflows.
Do not fully replace Python/R with Go. Instead, adopt a hybrid strategy. Go can reproduce the same basic linear regression coefficients as Python and R for this problem, and Go is likely to show strong performance in execution speed and memory efficiency. However, data scientists may still prefer Python or R for exploratory analysis, visualization, and broader ecosystem support. A practical recommendation is to use Go where performance and deployment matter most, while allowing Python or R where analytical flexibility matters most.

### Use Go for:
- Backend services  
- Data pipelines  
- High-performance applications  
- Scalable cloud systems  

### Use Python/R for:
- Exploratory data analysis (EDA)  
- Statistical modeling  
- Machine learning  
- Visualization  

---

## Repository Structure

```text
anscombe/
│── go/
│   ├── anscombe/
│   │   ├── data.go
│   │   ├── regression.go
│   │   ├── regression_test.go
│
│── python/
│   ├── anscombe.py
│
│── r/
│   ├── anscombe.R
│
│── data/
│   ├── anscombe.csv
│
│── results/
│   ├── go-benchmark.txt
│   ├── python-output.txt
│   ├── r-output.txt
│
│── README.md
```
---

## File-by-File Guide

🟦 /go/ — Go Implementation
go/main.go
 - Entry point for the Go application
 - Loads the Anscombe datasets
 - Executes regression using:
 - Custom OLS implementation
 - External stats package
 - Prints regression results for comparison

👉 Purpose: Demonstrates Go-based statistical computation

go/anscombe/data.go
 - Defines the Dataset struct
 - Stores all four Anscombe datasets
 - Provides the Quartet() function to access data

👉 Purpose: Centralized, reusable data source for all Go computations

go/anscombe/regression.go
Implements:
 - Custom OLS regression function
 - Wrapper for external stats package (montanaflynn/stats)
Computes:
 - Intercept
 - Slope
 - Correlation (r)
 - R²

👉 Purpose: Core statistical logic of the project

go/anscombe/regression_test.go
Contains:
 - Unit tests for regression correctness
 - Cross-validation between OLS and stats package
 - Benchmark functions
Key functions:
 - TestOLSMatchesKnownCoefficients
 - TestStatsPackageMatchesOLS
 - BenchmarkRunQuartet
 - BenchmarkStatsPackageQuartet

👉 Purpose: Ensures correctness and measures performance

go/go.mod
 - Defines Go module name and dependencies
Includes external package:
 - github.com/montanaflynn/stats

👉 Purpose: Dependency management and reproducibility

🟨 /python/ — Python Implementation
python/anscombe.py
 - Loads Anscombe datasets
Performs regression using:
 - numpy
 - statsmodels
Outputs:
 - Intercept
 - Slope
 - Correlation
 - R²
Measures:
 - Execution time
 - Memory usage

👉 Purpose: Baseline comparison for Go results

🟩 /r/ — R Implementation
r/anscombe.R
 - Performs regression using R’s lm() function
Outputs:
 - Intercept
 - Slope
 - Correlation
 - R²
 - Measures execution time

👉 Purpose: Statistical reference implementation

🟪 /data/ — Shared Dataset
data/anscombe.csv
 - Contains all four datasets in structured format
 - Used for consistency across languages

👉 Purpose: Single source of truth for data

🟥 /results/ — Output & Benchmark Results
results/go-benchmark.txt

Contains Go benchmark output from:

go test -bench=. -benchmem

👉 Includes:
 - Execution time (ns/op)
 - Memory usage (B/op)
 - Allocation counts
results/python-output.txt
 - Contains Python regression output
Includes:
 - Statistical results
 - Execution time
 - Memory usage
results/r-output.txt
 - Contains R regression output
Includes:
 - Statistical results
 - Execution time

👉 Purpose: Enables reproducibility and comparison

🟫 /scripts/ — Execution Helpers
scripts/run_all.sh
Bash script to run:
 - Go tests and benchmarks
 - Python script
 - R script

👉 Purpose: One-command execution for Unix-based systems

⚠️ Note:
 - Designed for macOS/Linux environments
 - Windows users should run commands manually or via Git Bash

---

## What the Go Project Includes

- `data.go`: the four Anscombe Quartet datasets
- `regression.go`: a direct ordinary least squares implementation plus a comparison wrapper for `github.com/montanaflynn/stats`
- `regression_test.go`: unit tests and Go benchmarks using the standard `testing` package
- `main.go`: prints regression results in both human-readable and JSON form

## How to Run

### Go

```bash
cd go
go mod tidy
go test ./anscombe -count=1
go test ./anscombe -run=^$ -bench=. -benchmem -count=1
go run ./...
```

### Python

```bash
pip install numpy statsmodels
python python/anscombe.py
```

### R

```bash
Rscript r/anscombe.R
```

### Run Everything

```bash
./scripts/run_all.sh
```

## Expected Statistical Result

For all four datasets, the fitted simple linear regression should be approximately:

- Intercept: `3.00`
- Slope: `0.50`
- R-squared: `0.67`

That is the key property of the Anscombe Quartet and the core reason it is useful for this assignment.
> Important analytic note: identical regression summaries do **not** mean the datasets are equivalent. The four datasets have very different shapes, which is why visualization remains essential in real analysis.

## References

- Anscombe, F. J. 1973. *Graphs in Statistical Analysis.* The American Statistician 27: 17-21.
- Miller, T. W. 2015. *Modeling Techniques in Predictive Analytics with Python and R.* Pearson.
- Go documentation: https://go.dev/
- `montanaflynn/stats` package: https://pkg.go.dev/github.com/montanaflynn/stats

---

## Author

**Gilbert Onyenwezi**  
 - MSDS 431 - GO & AI Programming  
 - Northwestern University
 - School of Professional Studies
 - Masters in Data Science - Data Engineering Specialization
 - 2026