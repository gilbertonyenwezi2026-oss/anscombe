quartet <- list(
  "Set I" = list(
    x = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
    y = c(8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68)
  ),
  "Set II" = list(
    x = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
    y = c(9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74)
  ),
  "Set III" = list(
    x = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
    y = c(7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73)
  ),
  "Set IV" = list(
    x = c(8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8),
    y = c(6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89)
  )
)

run_regression <- function(x, y) {
  fit <- lm(y ~ x)
  list(
    intercept = unname(coef(fit)[1]),
    slope = unname(coef(fit)[2]),
    r = cor(x, y),
    r_squared = summary(fit)$r.squared
  )
}

run_all <- function() {
  lapply(quartet, function(ds) run_regression(ds$x, ds$y))
}

start <- Sys.time()
results <- run_all()
elapsed <- Sys.time() - start

print(results)
cat(sprintf("elapsed_seconds=%.6f\n", as.numeric(elapsed, units = "secs")))
# Memory tracking is platform/tool dependent in base R; benchmark externally if needed.
