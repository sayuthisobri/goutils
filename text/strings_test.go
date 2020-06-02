package text

import (
	Assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings(t *testing.T) {
	assert := Assert.New(t)
	testString := "test"
	builder := NewString(testString)
	assert.Equal(builder.toString(), testString)

	testString = "test123"
	assert.Equal(builder.Append("123").toString(), testString)

	testString = "123"
	assert.Equal(builder.Set(testString).toString(), testString)

	testString = "123.456|789"
	builder.Set(testString)
	split := builder.Split("[.|]")
	assert.Len(split,3)
	assert.Equal(split[0],"123")
	assert.Equal(split[1],"456")
	assert.Equal(split[2],"789")

}
