# Telegram GO

This project is just a study case for me. If you looking for a telegram framework write in golang
you should search something in google, but, if you looking at a simple code to understand how to
implement a simple [questions and answers](https://www.youtube.com/watch?v=MShfaRUzrL0&list=RDMShfaRUzrL0)
telegram bot, you're welcome.

## Usage
```go
import (
    "fmt"

    "github.com/drgarcia1986/telegram-go/server"
)

func greeting(msg string, mv map[string]string) (string, error) {
    return fmt.Sprintf("Hello %s nice to meet you", mv["name"]), nil
}

func startServer() {
    s := server.New("your telegram-bot token")
    s.HandleFunc("Hi I'm (?P<name>\\w+)", greeting)
    s.Run()
}
```
The `server.HandlerFunc` expects a regex pattern and a function with this signature:
```go
func (msg string, mv map[string]string) (string, error) {}
```
* **msg**: original message received.
* **mv**: matched values.

The return of this function is the answer will be sent to the incoming message in Telegram.
