# Fuzzing

Fuzzing, also known as "fuzz testing," is a software testing technique that involves providing invalid, unexpected, or random data as input to a computer program.
The goal of fuzzing is to identify security vulnerabilities and other bugs in the program by causing it to crash or exhibit unintended behaviour.

Fuzzing can be performed manually or by using a testing library/framework to craft the inputs for us.

To better understand fuzzing let's go with an example, imagine this code:

```go
func DontPanic(s string) {
    if len(s) == 4 {
        if s[0] == 'f' {
            if s[1] == 'u' {
                if s[2] == 'z' {
                    if s[3] == 'z' {
                        panic("error: wrong input")
                    }
                }
            }
        }
    }
}
```

This is a Go function that accepts a `string` as the one and only argument.

Looking at the functions it looks like the function will `panic` (e.g. crash) in only one condition - if the provided input is the word `fuzz`.

Obviously, this function is really simple and by just looking at it we can see its behaviour.
However, in more complex systems such fail points may not be obvious, and may be missed by the person testing it/writing the unit test cases.

This is where fuzzing comes in handy.

The Go Fuzzing library (part of the standard language library since Go 1.18) generates many inputs for a test case, and then based on the coverage and the results determine which inputs are "interesting".

If we write a fuzz test for this function what will happen is:

1. The fuzzing library will start providing random strings starting from smaller strings and increasing their size.
2. Once the library provides a string of length 4 it will notice a change in the test-coverage (`if (len(s) == 4)` is now `true`) and will continue to generate inputs with this length.
3. Once the library provides a string of length 4 that starts with `f` it will notice another change in the test-coverage (`if s[0] == "f"` is now `true`) and will continue to generate inputs that start with `f`.
4. The same thing will repeat for `u` and the double `z`.
5. Once it provides `fuzz` as input the function will panic and the test will fail.
6. We have _fuzzed_ successfully!

Another good practice for fuzzing is to save the inputs that made the code crash and run them every time to make sure that the original error we caught via fuzzing will never be introduced into our code again.

This could again be a feature of the fuzzing framework.

Most fuzzing libraries allow us to add specific values we want to test with.
This also helps the libraries because it shows them values we already know are "interesting" so they can model the generated values on them.

## When fuzzing is not enough

Fuzzing is a useful technique, but there are situations in which it might not be helpful.

For example, if the input that fails our code is too specific and there are no clues to help, the fuzzing library might not be able to guess it.

If we change the example code from the previous paragraph to something like this:

```go
func DontPanic(s input) {
    if (len(s) == 4) && s[0] == 'f' && s[1] == 'u' && s[2] == 'z' && s[3] == 'z' {
        panic("error")
    }
}
```

or just:

```go
func DontPanic(s input) {
    if s == "fuzz" {
        panic("error")
    }
}
```

then fuzzing won't help us, because there's a small chance it will generate the exact string `fuzz` without having any clues,
and none of the inputs that triggered a code-coverage change in the previous case (string of size 4, string of size 4 starting with `z`, etc.) will trigger a code-coverage now (because we only have one `if` check, compared to 5 in the previous example).

So it's important to understand that while fuzzing is a good way to detect anomalies and corner-cases in your code it is not a silver-bullet for 100% correct code.

## Practical example

If you want to get your hands dirty with fuzzing in Go check out [my repo with examples on the topic](https://github.com/asankov/go-fuzzing-101/tree/v1).

It contains the example I used in this article + a fuzz test that triggers a failure, and instructions on how to run the test yourself.

## Resources

- <https://en.wikipedia.org/wiki/Fuzzing>
- [Fuzzing in Go by Valentin Deleplace, Devoxx Belgium 2022](https://www.youtube.com/watch?v=Zlf3s4EjnFU)
- [Write applications faster and securely with Go by Cody Oss, Go Day 2022](https://www.youtube.com/watch?v=aw7lFSFGKZs)

See you on [Day 17](day17.md).
