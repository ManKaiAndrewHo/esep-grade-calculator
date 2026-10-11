# ESEP Grade Calculator

This library is a simple utility to calculate your grades. This is a GO library, there is not 'main' program, so you use it by importing it or by running its unit test. 

A final grade is computed by averaging each kind of grade: assignments, exams, essays, multiply each average by its weight, adding the results, and converting the number to a letter grade. 

# Requirement

- Golang 1.24+ (Check with 'go version')
    - macOS: 'brew install go' 
- Optional: the "Go" extension on VSCode

# Project layout

```
ESEP-GRADE-CALCULATER/
├── README.ED
└── grade-calculator/  <- the Go module, rull all the go cmds from here
    ├── go.mod
    ├── grade_calculator_test.go
    └── grade_calculator.go
```

# Run unit tests:

```
cd grade_calculator/
go test ./...
```

# Run unit tests with code coverage:

```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

# Grading rules

| Category    | Weight |
|-------------|--------|
| Assignments | 50%    |
| Exams       | 35%    |
| Essays      | 15%    |

| Final score | Letter |
|-------------|--------|
| 90 and up   | A      |
| 80 - 89     | B      |
| 70 - 79     | C      |
| 60 - 69     | D      |
| below 60    | F      |

Passing requires a C or better.