# goroutine-csv

Requirements:

- Each CSV line should be processed concurrently
- If a single CSV line is invalid the function should be canceled
- Goroutines should stop if a CSV line is invalid
