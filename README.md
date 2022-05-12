# solapi-go

[Site](https://www.solapi.com/) |
[Docs](https://docs.solapi.com/) |
[Examples](https://github.com/solapi/solapi-go/tree/master/_examples) |

문자 메시지 발송 및 조회 관련 기능들을 쉽게 사용하실 수 있도록 만들어진 SDK 입니다.

## Example

```go
require (
        github.com/solapi/solapi-go
)

func main() {
	client := solapi.NewClient()

	// Message Data
	// 관련 파라미터들은 https://docs.solapi.com에서 확인 가능합니다.
	message := make(map[string]interface{})
	message["to"] = "01000000000"
	message["from"] = "029302266"
	message["text"] = "Test Message"
	message["type"] = "SMS"

	params := make(map[string]interface{})
	params["message"] = message

	// Call API Resource
	result, err := client.Message.SendSimpleMessage(params)
	if err != nil {
		fmt.Println(err)
	}
}
```

[examples folder](https://github.com/solapi/solapi-go/tree/master/_examples)에서 자세한 예제파일들을 확인하세요.

## Installation

```
go get github.com/solapi/solapi-go
```

## Configs

```
> OS의 환경설정 기반으로 값을 가져옵니다.
SOLAPI_API_KEY=
SOLAPI_API_SECRET=
SOLAPI_PROTOCOL=https
SOLAPI_DOMAIN=api.solapi.com
SOLAPI_PREFIX=
SOLAPI_APP_ID=
```
