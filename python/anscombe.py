import json
import time
import tracemalloc

import numpy as np
import statsmodels.api as sm

QUARTET = {
    "Set I": {
        "x": [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        "y": [8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68],
    },
    "Set II": {
        "x": [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        "y": [9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74],
    },
    "Set III": {
        "x": [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        "y": [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73],
    },
    "Set IV": {
        "x": [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
        "y": [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89],
    },
}


def run_regression(x, y):
    x_arr = np.asarray(x, dtype=float)
    y_arr = np.asarray(y, dtype=float)
    design = sm.add_constant(x_arr)
    model = sm.OLS(y_arr, design).fit()
    return {
        "intercept": float(model.params[0]),
        "slope": float(model.params[1]),
        "r": float(np.corrcoef(x_arr, y_arr)[0, 1]),
        "r_squared": float(model.rsquared),
    }


def run_all():
    return {name: run_regression(v["x"], v["y"]) for name, v in QUARTET.items()}


if __name__ == "__main__":
    tracemalloc.start()
    start = time.perf_counter()
    results = run_all()
    elapsed = time.perf_counter() - start
    current, peak = tracemalloc.get_traced_memory()
    tracemalloc.stop()

    print(json.dumps(results, indent=2))
    print(f"elapsed_seconds={elapsed:.6f}")
    print(f"peak_memory_bytes={peak}")
