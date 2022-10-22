package file

import (
	"fmt"
	"testing"
)

func TestConfigFromJsonFile(t *testing.T) {

	ddnsConfig := ConfigFromJsonFile()
	fmt.Println(ddnsConfig)
}
