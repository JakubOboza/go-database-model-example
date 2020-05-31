# Description

This is how I build database models in GO (in this example with Postgres)

Key things:

* Easy to use in "normal code" just do `users.Store.FindByEmail(...`
* Easy to mock in tests where you don't need to have database setup/teardown 
* Doesn't add a lot fo complexity to the code and doesn't require a lot of interface injection or other techniques

# How to make it run

To run example binary:

* Start up postgres via docker eg. `docker-compose up`
* build it `make build` 
* run it root of project `./bin/example` this will make it pick up migrations

To run tests:

* just run `make test`

# Maybe you know something better

Tell me, how to improve it. How we can make it easier to use, more testable. At this point writing mocks is kinda not super easy. 
Maybe in future i will make a code generator for this part.

Mocking is maybe not the best way to cover database layer in app but often avoiding mocking makes testing very hard and slow.
So all of the options should be considered.

# Author

Jakub Oboza
