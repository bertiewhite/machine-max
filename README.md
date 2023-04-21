## Machine Max Take Home Test

This can be run easy enough using `go run main.go` or you can build the executable and execute that. 

Firstly this was good fun! The challenge of stopping the in flight stuff only after it was complete was something I've not 
had to think about before. 

As always I feel like there's more that could be done. Testing around the job pool's running is a bit low. But testing
asynchronous code is hard and is very easily covered by manual tests, so I've got confidence in it. Allowing the number of
workers and the number of jobs needed to be run would be easy enough to extract into envvars or flags which could be used
to override the default variables. I was unable to get the server to return 422's so that's not been tested outside of 
unit tests, but the logic is simple enough that I have confidence in it!