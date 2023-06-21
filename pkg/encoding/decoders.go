package encoding

import (
	"context"
	"errors"

	jp "golang.org/x/text/encoding/japanese"
	zh_cn "golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

type Decoder interface {
	Decode(src []byte, ctx context.Context) ([]byte, error)
}

type DecoderFactory interface {
	CreateDecoder(ctx context.Context) (Decoder, error)
}

type DecoderFactoryFunc func(ctx context.Context) (Decoder, error)

func (f DecoderFactoryFunc) CreateDecoder(ctx context.Context) (Decoder, error) {
	return f(ctx)
}

// DecoderFactoryMap is a map of decoder factories.
type DecoderFactoryMap map[string]DecoderFactory

// Register adds a new decoder factory to the DecoderFactoryMap.
func (m DecoderFactoryMap) Register(name string, factory DecoderFactory) {
	m[name] = factory
}

// CreateDecoder creates a new decoder based on the given name and context.
func (m DecoderFactoryMap) CreateDecoder(name string, ctx context.Context) (Decoder, error) {
	factory, ok := m[name]
	if !ok {
		return nil, errors.New("no such decoder")
	}
	return factory.CreateDecoder(ctx)
}

type UTF8Decoder struct{}

func (d *UTF8Decoder) Decode(src []byte, ctx context.Context) ([]byte, error) {
	return unicode.UTF8.NewDecoder().Bytes(src)
}

type GBKDecoder struct{}

func (d *GBKDecoder) Decode(src []byte, ctx context.Context) ([]byte, error) {
	return zh_cn.GBK.NewDecoder().Bytes(src)
}

type ShiftJISDecoder struct{}

func (d *ShiftJISDecoder) Decode(src []byte, ctx context.Context) ([]byte, error) {
	return jp.ShiftJIS.NewDecoder().Bytes(src)
}
