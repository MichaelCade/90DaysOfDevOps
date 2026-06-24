# 퍼징(Fuzzing) 고급

어제 우리는 퍼징이 무엇인지, 그리고 퍼즈 테스트(퍼지 입력을 사용하는 단위 테스트)를 작성하는 방법을 배웠습니다.
하지만 퍼즈 테스트는 단순한 단위 테스트 그 이상입니다.
이 방법론을 사용하여 서버로 전송되는 요청을 퍼징하여 웹 애플리케이션을 테스트할 수 있습니다.

오늘은 웹 서버 퍼지 테스트에 대한 실용적인 접근 방식을 살펴보겠습니다.

이를 위해 다양한 도구가 도움이 될 수 있습니다.

[Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder), [SmartBear](https://smartbear.com/) 등이 있습니다.
하지만 유료 라이선스가 필요한 독점 도구(proprietary tools)도 있습니다.


그렇기 때문에 오늘 데모에서는 `Burp Intruder`에서 영감을 받아 유사한 기능을 제공하는 Go로 작성된 간단한 오픈 소스 CLI를 사용하려고 합니다.
이름은 [httpfuzz](https://github.com/JonCooperWorks/httpfuzz)입니다.


## Getting started

이 도구는 아주 간단합니다.
요청에 대한 템플릿(퍼지 데이터에 대한 placeholder를 정의), 단어 목록(퍼지 데이터)을 제공하면 `httpfuzz`가 요청을 렌더링하여 서버로 전송합니다.

먼저 요청에 대한 템플릿을 정의해야 합니다.
다음 내용으로 `request.txt`라는 파일을 생성합니다.

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

이것은 JSON 본문이 있는 `/` 경로에 대한 유효한 HTTP `POST` 요청입니다.
본문의 "\`" 기호는 우리가 제공하는 데이터로 대체될 플레이스홀더를 정의합니다.

`httpfuzz`는 헤더, 경로, URL 매개변수를 퍼즈 처리할 수도 있습니다.

다음으로 요청에 넣을 입력의 단어 목록을 제공해야 합니다.
다음 내용으로 `data.txt`라는 파일을 생성합니다:

```text
SOME_NAME
Mozilla/5.0 (Linux; Android 7.0; SM-G930VC Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/58.0.3029.83 Mobile Safari/537.36
```

이 파일에서는 본문 내부에 대체될 두 개의 입력을 정의했습니다.
실제 시나리오에서는 적절한 퍼즈 테스트를 위해 훨씬 더 많은 데이터를 여기에 넣어야 합니다.

이제 템플릿과 입력이 준비되었으니 도구를 실행해 보겠습니다.
안타깝게도 이 도구는 바이너리로 배포되지 않으므로 소스에서 빌드해야 합니다.
리포지토리를 복제하고 실행합니다.

```shell
go build -o httpfuzz cmd/httpfuzz.go
```
(컴퓨터에 최신 버전의 Go가 설치되어 있어야 합니다.)

이제 바이너리를 얻었으니 실행해 보겠습니다.

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

- `httpfuzz`는 우리가 호출하는 바이너리입니다.
- `--wordlist data.txt`는 우리가 제공한 입력이 포함된 파일입니다.
- `--seed-request requests.txt`는 요청 템플릿입니다.
- `--target-header User-Agent`는 `httpfuzz`가 `User-Agent` 헤더 대신에 제공된 입력을 사용하도록 지시합니다.
- `--target-param fuzz`는 제공된 입력을 `fuzz` URL 매개변수의 값으로 사용하도록 `httpfuzz`에 지시합니다.
- `--delay-ms 50`은 요청 사이에 50ms를 기다리도록 `httpfuzz`에 지시합니다.
- `--skip-cert-verify`는 `httpfuzz`에게 TLS 확인을 수행하지 않도록 지시합니다.
- `--proxy-url http://localhost:8080`는 `httpfuzz`에게 HTTP 서버의 위치를 알려줍니다.

2개의 입력과 3개의 위치(본문, `User-Agent` 헤더, `fuzz` 매개변수)에 배치할 수 있습니다.
즉, `httpfuzz`는 6개의 요청을 생성하여 서버로 전송합니다.

실행해서 어떤 일이 일어나는지 봅시다.
서버로 들어오는 요청을 확인할 수 있도록 모든 요청을 기록하는 간단한 웹 서버를 작성했습니다.

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

그리고 서버 로그입니다.

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

6개의 HTTP 요청이 수신된 것을 확인할 수 있습니다.

이 중 2개는 `User-Agent` 헤더에 대한 값 파일의 값을 가지고 있고, 4개는 템플릿의 기본 헤더를 가지고 있습니다.

그 중 2개는 `fuzz` 쿼리 매개변수에 대한 값 파일의 값을 가지고 있고, 4개는 템플릿의 기본 헤더를 가지고 있습니다.

그 중 2개는 `Name` 본문 속성에 대한 값 파일에서 값을 가져오고, 4개는 템플릿의 기본 헤더를 사용합니다.

이 도구를 약간 개선하면 이러한 요청의 순열을 다르게 만들 수 있습니다(예: 값 파일의 값으로 `?fuzz=`와 `User-Agent`가 모두 있는 요청).

`httpfuzz`가 요청의 결과에 대한 정보를 제공하지 않는다는 점에 주목하세요.
이를 파악하려면 서버에 대한 일종의 모니터링을 설정하거나 우리에게 의미 있는 방식으로 결과를 처리하는 `httpfuzz` 플러그인을 작성해야 합니다.
그렇게 해봅시다.

사용자 정의 플러그인을 작성하려면 [`Listener`](https://github.com/JonCooperWorks/httpfuzz/blob/master/plugin.go#L13) 인터페이스를 구현해야합니다.

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

이제 플러그인을 빌드해야 합니다.

```shell
go build -buildmode=plugin -o log exampleplugins/log/log.go
```

그리고 `--post-request` 플래그를 통해 `httpfuzz`에 연결할 수 있습니다:

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

짜잔!
이제 최소한 서버의 응답 코드가 무엇인지 확인할 수 있습니다.

물론 훨씬 더 많은 데이터를 출력하는 훨씬 더 정교한 플러그인을 작성할 수도 있지만, 이 연습의 목적상 이 정도면 충분합니다.


## 요약

퍼징은 단위 테스트를 훨씬 뛰어넘는 매우 강력한 테스트 기법입니다.

퍼징은 유효한 HTTP 요청의 일부를 서버의 취약점이나 결함을 노출할 수 있는 데이터로 대체하여 HTTP 서버를 테스트하는 데 매우 유용할 수 있습니다.

웹 애플리케이션 퍼지 테스트에 도움이 되는 무료 및 유료 도구가 많이 있습니다.

## 리소스

[OWASP: Fuzzing](https://owasp.org/www-community/Fuzzing)

[OWASP: Fuzz Vectors](https://owasp.org/www-project-web-security-testing-guide/v41/6-Appendix/C-Fuzz_Vectors)

[Hacking HTTP with HTTPfuzz](https://medium.com/swlh/hacking-http-with-httpfuzz-67cfd061b616)

[Fuzzing the Stack for Fun and Profit at DefCamp 2019](https://www.youtube.com/watch?v=qCMfrbpuCBk&list=PLnwq8gv9MEKiUOgrM7wble1YRsrqRzHKq&index=33)

[HTTP Fuzzing Scan with SmartBear](https://support.smartbear.com/readyapi/docs/security/scans/types/fuzzing-http.html)

[Fuzzing Session: Finding Bugs and Vulnerabilities Automatically](https://youtu.be/DSJePjhBN5E)

[Fuzzing the CNCF Landscape](https://youtu.be/zIyIZxAZLzo)

[18일차](day18.md)에 뵙겠습니다.
