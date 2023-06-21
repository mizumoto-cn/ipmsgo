package encoding

import (
	"context"
)

var (
	decoderFactoryMap = DecoderFactoryMap{}
)

func init() {
	decoderFactoryMap.Register("utf8", DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
		return &UTF8Decoder{}, nil
	}))
	decoderFactoryMap.Register("gbk", DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
		return &GBKDecoder{}, nil
	}))
	decoderFactoryMap.Register("shiftjis", DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
		return &ShiftJISDecoder{}, nil
	}))
}

func (m DecoderFactoryMap) Register(name string, factory DecoderFactory) {
	m[name] = factory
}

func Decode(name string, src []byte, ctx context.Context) ([]byte, error) {
	decoder, err := CreateDecoder(name, ctx)
	if err != nil {
		return nil, err
	}
	return decoder.Decode(src, ctx)
}
