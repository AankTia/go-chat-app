## Tracing code to get a look under the hood

- Tracing is a debuging technique.
- Tracing is a practice by which we log or print key steps in the flow of a program to make what is going on under the covers visible.

### Writing a package using TDD

- Packages in Go are organized into folders, with one package per folder
- Go has no concept of subpackages, which means nested packages (in nested folders) exist only for aesthetic or informational reasons but do not inherit any functionality or visibility from super packages
- Go has no concept of subpackages, which means nested packages (in nested folders) exist only for aesthetic or informational reasons but do not inherit any functionality or visibility from super packages

## Interfaces

- Interfaces in Go are an extremely powerful language feature that allows us to define an API without being strict or specific on the implementation details.

## Unit tests

- Testing was built into the Go tool chain from the very beginning, making writing automatable tests a first-class citizen.
- The test code lives alongside the production code in files suffixed with `_test.go`.
- The Go tools will treat any function that starts with Test (taking a single `*testing.T ` argument) as a unit test, and it will be executed when we run our tests.
- To run them for this package, navigate to the folder in a terminal and do the following:

  ```bash
  go test
  ```

- The ever-thoughtful core team has even solved this problem for us by providing code coverage statistics. The following command provides code statistics:

  ```bash
  go test -cover
  ```

## Red-green testing

- Red-green testing proposes that we first write a unit test, see it fail (or produce an error), write the minimum amount of code possible to make that test pass, and rinse and repeat it again. The key point here being that we want to make sure the code we add is actually doing something as well as ensuring that the test code we write is testing something meaningful.
