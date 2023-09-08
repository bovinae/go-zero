package netx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalIp(t *testing.T) {
	fmt.Println(InternalIp())
	assert.True(t, len(InternalIp()) > 0)
}
