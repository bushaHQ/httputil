package ip

import (
	"log"
	"os"
	"testing"
)

func TestForeignIP(t *testing.T) {
	Init(os.Getenv("IP_KEY"))

	v, err := ForeignIP("41.184.177.81")
	if err != nil {
		t.Error(err)
		return
	}

	log.Printf("%v", v)
}
