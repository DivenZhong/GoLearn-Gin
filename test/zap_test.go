/**
 * @Author Mr.ZhongQiWen
 * @Description TODO
 * @Date 2022/11/5 6:29 下午
 **/
package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	"testing"
)

func TestZap(t *testing.T) {
	fmt.Printf("DebugLevel: %v\n", zapcore.DebugLevel)
	fmt.Printf("InfoLevel: %v\n", zapcore.InfoLevel)
	fmt.Printf("WarnLevel: %v\n", zapcore.WarnLevel)
	fmt.Printf("-1: %v\n", zapcore.Level(-1))
	fmt.Printf("0: %v\n", zapcore.Level(0))
	fmt.Printf("1: %v\n", zapcore.Level(1))
	fmt.Printf("2: %v\n", zapcore.Level(2))
	fmt.Printf("3: %v\n", zapcore.Level(3))
}

func TestIsEqual(t *testing.T) {
	ok1 := IsEqual(1, 1)
	ok2 := IsEqual(1, 2)

	ok3, err3 := IsEqualWithErr(1, 1)
	ok4, err4 := IsEqualWithErr(1, 2)

	assert.True(t, ok1, "a:1 b:1 true")
	assert.False(t, ok2, "a:1 b:2 false")
	//断言相等
	assert.Equal(t, true, ok3, "a:1 b:1 true")
	assert.Nil(t, err3, "a:1 b:1 nil")

	assert.Equal(t, false, ok4, "a:1 b:2 false")
	assert.NotNil(t, err4, "a:1 b:2 not nil")
}
