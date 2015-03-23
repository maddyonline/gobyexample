## My Forked Version of Go by Example
(To be continued...)

### Table of Contents (and my status of completion)

	0. Hello World (Done!)
	1. Values (Done!)
	2. Variables (Done!)
	3. Constants (Done!)
	4. For  (Done!)
	5. If/Else (Done!)
	6. Switch (Done!)
	7. Arrays (Done!)
	8. Slices (Done!)
	9. Maps
	10. Range
	11. Functions
	12. Multiple Return Values
	13. Variadic Functions
	14. Closures
	15. Recursion
	16. Pointers
	17. Structs
	18. Methods
	19. Interfaces
	20. Errors
	21. Goroutines
	22. Channels
	23. Channel Buffering
	24. Channel Synchronization
	25. Channel Directions
	26. Select
	27. Timeouts
	28. Non-Blocking Channel Operations
	29. Closing Channels
	30. Range over Channels
	31. Timers
	32. Tickers
	33. Worker Pools
	34. Rate Limiting
	35. Atomic Counters
	36. Mutexes
	37. Stateful Goroutines
	38. Sorting
	39. Sorting by Functions
	40. Panic
	41. Defer
	42. Collection Functions
	43. String Functions
	44. String Formatting
	45. Regular Expressions
	46. JSON
	47. Time
	48. Epoch
	49. Time Formatting / Parsing
	50. Random Numbers
	51. Number Parsing
	52. URL Parsing
	53. SHA1 Hashes
	54. Base64 Encoding
	55. Reading Files
	56. Writing Files
	57. Line Filters
	58. Command-Line Arguments
	59. Command-Line Flags
	60. Environment Variables
	61. Spawning Processes
	62. Exec'ing Processes
	63. Signals
	64. Exit


*The original README.md follows.*

## Go by Example

Content, toolchain, and web server for [Go by Example](https://gobyexample.com).


### Overview

The Go by Example site is built by extracting code &
comments from source files in `examples` and rendering
that data via the site `templates`. The programs
implementing this build process are in `tools`.

The build process produces a directory of static files -
`public` - suitable for serving by any modern HTTP server.
We include a lightweight Go server in `server.go`.


### Building

To build the site:

```console
$ go get github.com/russross/blackfriday
$ tools/build
$ open public/index.html
```

To build continuously in a loop:

```console
$ tools/build-loop
```


### Local Deploy

```bash
$ mkdir -p $GOPATH/src/github.com/mmcgrana
$ cd $GOPATH/src/github.com/mmcgrana
$ git clone git@github.com:mmcgrana/gobyexample.git
$ cd gobyexample
$ go get
$ foreman start
$ foreman open
```


### Platform Deploy

Basic setup:

```bash
$ export DEPLOY=$USER
$ export APP=gobyexample-$USER
$ heroku create $APP -r $DEPLOY
$ heroku config:add -a $APP
    BUILDPACK_URL=https://github.com/mmcgrana/buildpack-go.git
    CANONICAL_HOST=$APP.herokuapp.com \
    FORCE_HTTPS=1 \
    AUTH=go:byexample
$ heroku labs:enable dot-profile-d -a $APP
$ git push $DEPLOY master
$ heroku open -a $APP
```

Add a domain + SSL:

```bash
$ heroku domains:add $DOMAIN
$ heroku addons:add ssl -r $DEPLOY
# order ssl cert for domain
$ cat > /tmp/server.key
$ cat > /tmp/server.crt.orig
$ curl https://knowledge.rapidssl.com/library/VERISIGN/ALL_OTHER/RapidSSL%20Intermediate/RapidSSL_CA_bundle.pem > /tmp/rapidssl_bundle.pem
$ cat /tmp/server.crt.orig /tmp/rapidssl_bundle.pem > /tmp/server.crt
$ heroku certs:add /tmp/server.crt /tmp/server.key -r $DEPLOY
# add ALIAS record from domain to ssl endpoint dns
$ heroku config:add CANONICAL_HOST=$DOMAIN -r $DEPLOY
$ heroku open -r $DEPLOY
```


### License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](http://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Translations

Contributor translations of the Go by Example site are available in:

* [Chinese](http://everyx.github.io/gobyexample/) by [everyx](https://github.com/everyx)


### Thanks

Thanks to [Jeremy Ashkenas](https://github.com/jashkenas)
for [Docco](http://jashkenas.github.com/docco/), which
inspired this project.
