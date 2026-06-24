# Fuzzing Advanced

Yesterday we learned what fuzzing is and how to write fuzz tests (unit tests with fuzzy inputs).
However, fuzz testing goes beyond just unit testing.
We can use this methodology to test our web application by fuzzing the requests sent to our server.

Today, we will take a practical approach to fuzzy testing a web server.

Different tools can help us do this.

Such tools are [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder) and [SmartBear](https://smartbear.com/).
However, there are proprietary tools that require a paid license to use them.

That is why for our demonstration today we are going to use a simple open-source CLI written in Go that was inspired by Burp Intruder and provides similar functionality.
It is called [httpfuzz](https://github.com/JonCooperWorks/httpfuzz).


## Getting started

This tool is quite simple.
We provide it a template for our requests (in which we have defined placeholders for the fuzzy data), a wordlist (the fuzzy data) and `httpfuzz` will render the requests and send them to our server.

First, we need to define a template for our requests.
Create a file named `request.txt` with the following content:

```text
POST / HTTP/1.1
Content-Type: application/json
User-Agent: PostmanRuntime/7.26.3
Accept: */*
Cache-Control: no-cache
Host: localhost:8000
Accept-Encoding: gzip, deflate
Connection: close
Content-Length: 35

{
    "name": "`S9`",
}
```

This is a valid HTTP `POST` request to the `/` route with JSON body.
The "\`" symbol in the body defines a placeholder that will be substituted with the data we provide.

`httpfuzz` can also fuzz the headers, path, and URL params.

Next, we need to provide a wordlist of inputs that will be placed in the request.
Create a file named `data.txt` with the following content:

```text
SOME_NAME
Mozilla/5.0 (Linux; Android 7.0; SM-G930VC Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/58.0.3029.83 Mobile Safari/537.36
```

In this file, we defined two inputs that will be substituted inside the body.
In a real-world scenario, you should put much more data here for proper fuzz testing.

Now that we have our template and our inputs, let's run the tool.
Unfortunately, this tool is not distributed as a binary, so we will have to build it from source.
Clone the repo and run:

```shell
go build -o httpfuzz cmd/httpfuzz.go
```

(requires to have a recent version of Go installed on your machine).

Now that we have the binary let's run it:

```shell
./httpfuzz \
   --wordlist data.txt \
   --seed-request request.txt \
   --target-header User-Agent \
   --target-param fuzz \
   --delay-ms 50 \
   --skip-cert-verify \
   --proxy-url http://localhost:8080 \
```

- `httpfuzz` is the binary we are invoking.
- `--wordlist data.txt` is the file with inputs we provided.
- `--seed-request requests.txt` is the request template.
- `--target-header User-Agent` tells `httpfuzz` to use the provided inputs in the place of the `User-Agent` header.
- `--target-param fuzz` tells `httpfuzz` to use the provided inputs as values for the `fuzz` URL parameter.
- `--delay-ms 50` tells `httpfuzz` to wait 50 ms between the requests.
- `--skip-cert-verify` tells `httpfuzz` to not do any TLS verification.
- `--proxy-url http://localhost:8080` tells `httpfuzz` where our HTTP server is.

We have 2 inputs and 3 places to place them (in the body, the `User-Agent` header, and the `fuzz` parameter).
This means that `httpfuzz` will generate 6 requests and send them to our server.

Let's run it and see what happens.
I wrote a simple web server that logs all requests so that we can see what is coming into our server:

```shell
$ ./httpfuzz \
   --wordlist data.txt \
   --seed-request request.txt \
   --target-header User-Agent \
   --target-param fuzz \
   --delay-ms 50 \
   --skip-cert-verify \
   --proxy-url http://localhost:8080 \

httpfuzz: httpfuzz.go:164: Sending 6 requests
```

and the server logs:

```text
-----
Got request to http://localhost:8000/
User-Agent header = [SOME_NAME]
Name = S9
-----
Got request to http://localhost:8000/?fuzz=SOME_NAME
User-Agent header = [PostmanRuntime/7.26.3]
Name = S9
-----
Got request to http://localhost:8000/
User-Agent header = [PostmanRuntime/7.26.3]
Name = SOME_NAME
-----
Got request to http://localhost:8000/
User-Agent header = [Mozilla/5.0 (Linux; Android 7.0; SM-G930VC Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/58.0.3029.83 Mobile Safari/537.36]
Name = S9
-----
Got request to http://localhost:8000/?fuzz=Mozilla%2F5.0+%28Linux%3B+Android+7.0%3B+SM-G930VC+Build%2FNRD90M%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.083+Mobile+Safari%2F537.36
User-Agent header = [PostmanRuntime/7.26.3]
Name = S9
-----
Got request to http://localhost:8000/
User-Agent header = [PostmanRuntime/7.26.3]
Name = Mozilla/5.0 (Linux; Android 7.0; SM-G930VC Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/58.0.3029.83 Mobile Safari/537.36
```

We see that we have received 6 HTTP requests.

Two of them have a value from our values file for the `User-Agent` header, and 4 have the default header from the template.
Two of them have a value from our values file for the `fuzz` query parameter, and 4 have the default header from the template.
Two of them have a value from our values file for the `Name` body property, and 4 have the default header from the template.

A slight improvement of the tool could be to make different permutations of these requests (for example, a request that has both `?fuzz=` and `User-Agent` as values from the values file).

Notice how `httpfuzz` does not give us any information about the outcome of the requests.
To figure that out, we need to either set up some sort of monitoring for our server or write a `httpfuzz` plugin that will process the results in a meaningful for us way.
Let's do that.

To write a custom plugin, we need to implement the [`Listener`](https://github.com/JonCooperWorks/httpfuzz/blob/master/plugin.go#L13) interface:

```go
// Listener must be implemented by a plugin to users to hook the request - response transaction.
// The Listen method will be run in its own goroutine, so plugins cannot block the rest of the program, however panics can take down the entire process.
type Listener interface {
    Listen(results <-chan *Result)
}
```

```go
package main

import (
    "bytes"
    "io/ioutil"
    "log"

    "github.com/joncooperworks/httpfuzz"
)

type logResponseCodePlugin struct {
    logger *log.Logger
}

func (b *logResponseCodePlugin) Listen(results <-chan *httpfuzz.Result) {
    for result := range results {
        b.logger.Printf("Got %d response from the server\n", result.Response.StatusCode)
    }
}

// New returns a logResponseCodePlugin plugin that simple logs the response code of the response.
func New(logger *log.Logger) (httpfuzz.Listener, error) {
    return &logResponseCodePlugin{logger: logger}, nil
}
```

Now we need to build our plugin first:

```shell
go build -buildmode=plugin -o log exampleplugins/log/log.go
```

and then we can plug it into `httpfuzz` via the `--post-request` flag:

```shell
$ ./httpfuzz \
   --wordlist data.txt \
   --seed-request request.txt \
   --target-header User-Agent \
   --target-param fuzz \
   --delay-ms 50 \
   --skip-cert-verify \
   --proxy-url http://localhost:8080 \
   --post-request log

httpfuzz: httpfuzz.go:164: Sending 6 requests
httpfuzz: log.go:15: Got 200 response from the server
httpfuzz: log.go:15: Got 200 response from the server
httpfuzz: log.go:15: Got 200 response from the server
httpfuzz: log.go:15: Got 200 response from the server
httpfuzz: log.go:15: Got 200 response from the server
httpfuzz: log.go:15: Got 200 response from the server
```

Voila!
Now we can at least see what the response code from the server was.

Of course, we can write much more sophisticated plugins that output much more data, but for the purpose of this exercise, that is enough.

## Summary

Fuzzing is a really powerful testing technique that goes way beyond unit testing.

Fuzzing can be extremely useful for testing HTTP servers by substituting parts of valid HTTP requests with data that could potentially expose vulnerabilities or deficiencies in our server.

There are many tools that can help us in fuzzy testing our web applications, both free and paid ones.

## Resources

[OWASP: Fuzzing](https://owasp.org/www-community/Fuzzing)

[OWASP: Fuzz Vectors](https://owasp.org/www-project-web-security-testing-guide/v41/6-Appendix/C-Fuzz_Vectors)

[Hacking HTTP with HTTPfuzz](https://medium.com/swlh/hacking-http-with-httpfuzz-67cfd061b616)

[Fuzzing the Stack for Fun and Profit at DefCamp 2019](https://www.youtube.com/watch?v=qCMfrbpuCBk&list=PLnwq8gv9MEKiUOgrM7wble1YRsrqRzHKq&index=33)

[HTTP Fuzzing Scan with SmartBear](https://support.smartbear.com/readyapi/docs/security/scans/types/fuzzing-http.html)

[Fuzzing Session: Finding Bugs and Vulnerabilities Automatically](https://youtu.be/DSJePjhBN5E)

[Fuzzing the CNCF Landscape](https://youtu.be/zIyIZxAZLzo)

See you on [Day 18](day18.md).
