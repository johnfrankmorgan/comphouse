package e2e

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/johnfrankmorgan/comphouse"
)

func TestFunctionality(t *testing.T) {
	client := comphouse.NewClient("", nil)

	client.Hooks.BeforeRequest = append(client.Hooks.BeforeRequest, func(_ *http.Request) {
		fmt.Println("executed before sending a request!")
	})

	client.Hooks.AfterRequest = append(client.Hooks.AfterRequest, func(_ *http.Response) {
		fmt.Println("executed after sending a request!")
	})
}
