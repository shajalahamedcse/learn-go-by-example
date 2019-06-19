## Introduction To Testing

`Testing` is hugely important in all software development. It is very important to able to ensure the correctness of code. By taking the time to test our go programs we can allow ourself to develop faster and with a greater sense of confidence that what we are writing will continue to work properly when we release it to production.

### Goals

By the end of this tutorial, we will have a good understanding of testing basic functions. We will use the standard "testing" package of `Go`.


### Introduction

In this tutorial we are going to run tests for our go code using the `go test`command.

### Go Test Files

The feature file and the unit test file will be in the same folder always . 

    myproject/
    - add.go
    - add_test.go
    - main.go
    - main_test.go



### A Simple Test File

Imagine we had a very simple `go program` that is just an `add()` function. This `add()' function simply addes two number and retruns it. It is a nice and simple example to play around.


    package add

    func Addition(a, b int) int {
        return a + b
    }

If we wished to test this we could create a `addition_test.go` file within the same directory and write the following test:

    package add

    import "testing"

    func TestAddition(t *testing.T) {
        if Addition(4, 5) != 9 {
            t.Error("Expected 4 + 5 to equal 9")
        }
        if Addition(0, -1) != -1 {
            t.Error("Expected 0 + (-1) to equal -1")
        }
    }


### Running Our Tests

Now that we have created our first go test, it’s time to run this and see if our code behaves the way we expect it to. We can execute our tests by running:

    $ go test

This should then output something similar to the following:

    PASS
    ok      projects/go-learn-by-example/testing-in-go/introduction-to-testing      0.001s



    === RUN   TestAddition
    --- PASS: TestAddition (0.00s)
    PASS
    ok      projects/go-learn-by-example/testing-in-go/introduction-to-testing      0.001s



### Table Driven Testing

Now, we can add more test cases in our code to improve confidence. If we want to gradually build up a series of test cases that are always tested, we can take leverage an array of tests :

func TestTableCalculate(t *testing.T) {
    var tests = []struct {
        input    int
        expected int
    }{
        {2,2, 4},
        {-1,1, 0},
        {0,4, 4},
        {-5, -3, -8},
        {99999,2, 100001},
    }

    for _, test := range tests {
        if output := Calculate(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}


Here we declared a struct that contains both input numbers and the expected value. Then we iterate through the list of tests with our `for _, test := range tests` call and check to see that our function will always return the expected `results`, regardless of `input`.

When we run our test suite now, we should see the same output as before:

    PASS
    ok      projects/go-learn-by-example/testing-in-go/introduction-to-testing      0.001s



#### Verbose Test Output

Sometimes you may wish to see exactly what tests are running and how long they took. This is available if you use the `-v` flag when running your tests . so:

    $ go test -v                                                                                                                   18:44:56
    === RUN   TestAddition
    --- PASS: TestAddition (0.00s)
    === RUN   TestTableCalculate
    --- PASS: TestTableCalculate (0.00s)
    PASS


### Checking Test Coverage

As a developer we have to write test cases around the critical business logic  within our systems and ensuring more edge cases are covered. We can use test coverage matrix to measure our coverage.

#### Using the -cover flag

let’s look at how you can check the test coverage of your system using the go test command:

Within the same directory as our addition.go and our addition_test.go files, run the following:

    $ go test -cover

    PASS
    coverage: 100.0% of statements
    ok      projects/go-learn-by-example/testing-in-go/introduction-to-testing      0.001s

we will see that we have 100.0% of our total Go code covered by test cases.

### Visualizing Coverage

We can use the `go test` tool to generate a coverprofile which can then be converted to a HTML visualization using the `go tool cover` command:

    $ go test -coverprofile=coverage.out


    PASS
    coverage: 100.0% of statements
    ok      projects/go-learn-by-example/testing-in-go/introduction-to-testing      0.001s

You can then take this generated coverage.out file and use it to generate a HTML page which shows exactly what lines have been covered like so:

    $ go tool cover -html=coverage.out

