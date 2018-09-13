package util

import (
	"bytes"
	"testing"

	"github.com/ontio/neo-go-compiler/vm/api/runtime"
	"github.com/stretchr/testify/assert"
)

func TestWriteVarUint1(t *testing.T) {
	var (
		val = uint64(1)
		buf = new(bytes.Buffer)
	)
	if err := WriteVarUint(buf, val); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, buf.Len())
}

func TestWriteVarUint1000(t *testing.T) {
	var (
		val = uint64(1000)
		buf = new(bytes.Buffer)
	)

	if err := WriteVarUint(buf, val); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, buf.Len())
	assert.Equal(t, byte(0xfd), buf.Bytes()[0])
	res := ReadVarUint(buf)
	assert.Equal(t, val, res)
}

func TestWriteVarUint100000(t *testing.T) {
	var (
		val = uint64(100000)
		buf = new(bytes.Buffer)
	)

	if err := WriteVarUint(buf, val); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 5, buf.Len())
	assert.Equal(t, byte(0xfe), buf.Bytes()[0])
	res := ReadVarUint(buf)
	assert.Equal(t, val, res)
}

func TestWriteVarUint100000000000(t *testing.T) {
	var (
		val = uint64(1000000000000)
		buf = new(bytes.Buffer)
	)

	if err := WriteVarUint(buf, val); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 9, buf.Len())
	assert.Equal(t, byte(0xff), buf.Bytes()[0])
	res := ReadVarUint(buf)
	assert.Equal(t, val, res)
}

func TestEq(t *testing.T) {

	a := []byte("abcde")
	b := []byte{'a', 'b', 'c', 'd', 'e'}
	assert.True(t, bytesEquals(a, b))
}

func bytesEquals(a []byte, b []byte) bool {
	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		runtime.Notify(i)
		if a[i] != b[i] {
			return false
		}
	}

	return true

}
