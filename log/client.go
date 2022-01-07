package log

import (
	"bytes"
	"fmt"
	"go-l-plurasight/registry"
	stlog "log"
	"net/http"
)

type clientLogger struct{
	url string
}

func(c clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer(data)
	res, err := http.Post(c.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message")
	}

	return len(data), nil
}

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{})
}