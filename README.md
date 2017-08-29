dp-cmd-api-test-spike
================

This spike was to determine a suitable approach for integration testing against API's.

### Getting started

`go test ./...` - run all the tests

## Considerations:

### Language

The consideration of languages came down to Java and Go. Java has a lot of mature testing libraries and frameworks and 
appears to be a popular choice for automation testers. Go is the primary language used for production code within 
ONS digital, so is considered to keep the languages consistent and for all the same reasons that Go was chosen in 
production code.

#### Recommendation:

Go. 
The developers of the production code will be creating and maintaining the automated tests, and minimising the 
context switch between the languages is important. Part of the spike was to investigate frameworks / libraries for
automated testing in Go. There are plenty of libraries that provide equivalent functionality to the Java libraries.
Go is significantly faster to compile and run compared to Java, which will save resources on the CI infrastructure.

### Location

We have to decide whether to maintain tests as a separate repository, or to maintain them in the same code repositories
as the production code. There seems to be no significant advantage to either, and is a matter of preference. Having 
a single repository for all the integration tests will keep them all together and have only one build job. 
The integration tests have a scope of the API interface, so keeping the integration test code alongside the API can
help keep focus on that scope.

#### Recommendation

Keep the integration test code alongside the API's they are testing.
For code maintenance and CI perspective its appealing to have the tests alongside the API's. However if we were to
consider tests that go outside the scope of a single API, then separating the test code would make more sense.

### Framework / libraries

There were two styles of frameworks / libraries considered - high level Cucumber style testing frameworks, and lower level frameworks 
where tests are written just in the plain programming language. 

Java: 
https://github.com/rest-assured/rest-assured
https://github.com/intuit/karate

Go:
https://github.com/gavv/httpexpect

#### Recommendation

It was decided that using a lower level testing framework was more appropriate. As it will be the developers maintaining
the tests there was no appeal for using the higher level cucumber style frameworks. Using a lower level framework allows
more flexibility and means less of a context switch as the tests are written using the plain programming language.
It is recommended to use the standard tooling's unit test framework along with a library that provides a DSL for testing API's. 

### Configuration / environment

The test suite needs a way of injecting configuration values allowing the tests to be run in different environments.
Other services follow the 12 factor app guidelines of having config injected via environment variables so that seems like
a sensible default. The go test package creates it own main method which has a built in call to flags.Parse(). If we 
were to use flags we would have to ensure that config values were supplied as flags instead of environment vars. This 
depends on how the tests are integrated into the CI pipeline. To use environment variables you can specify a main method:

`func TestMain(m *testing.M)`

This method is run at the start of the test suite and can be used to load environment configuration as we do elsewhere.
One thing to note about the TestMain method - it only appears to get run if there are tests in the same file. I tried to
have a seperate file to contains the TestMain method as it did not get run.

#### Recommendation

Consider both depending on how it fits into CI - but leaning towards environment variables with TestMain function.
Having all services consistent in terms of configuration is important. Using the synthetic main method that parses flags 
would be nice to save some code, but the additional complexity in getting it integrated with CI should be considered.

### Test suites

We need a way of specifying a subset of tests to run. If the integration tests live alongside the API's they are testing,
we need a way to separate the integration tests from the unit tests. We also need a way of running tests that are safe
to run on production, aka smoke tests.

There are a few possible solutions:
 - there is a -short flag built into the test package that can be used as a switch to run a test or not https://golang.org/pkg/testing/#Short
 - use the -run flag with a naming convention for the tests https://golang.org/cmd/go/#hdr-Description_of_testing_flags
 - use environment variables to set flags and only run the test if the flag is set
 - use build tags:  `// +build integration` at top of file and pass `-tags=integration` to go test to include that file
   - described here : http://peter.bourgon.org/go-in-production/#testing-and-validation
 
#### Recommendation

Build flags.
The build flags solution appears to be the cleanest solution. Any files with the tag will not be included by default, and
as many flags as required can be used.

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
