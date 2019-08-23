package backend

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/baudtime/baudtime/msg"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AddRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Series")
		return
	}
	if cap(z.Series) >= int(zb0002) {
		z.Series = (z.Series)[:zb0002]
	} else {
		z.Series = make([]*msg.Series, zb0002)
	}
	for za0001 := range z.Series {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "Series", za0001)
				return
			}
			z.Series[za0001] = nil
		} else {
			if z.Series[za0001] == nil {
				z.Series[za0001] = new(msg.Series)
			}
			err = z.Series[za0001].DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Series", za0001)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AddRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Series)))
	if err != nil {
		err = msgp.WrapError(err, "Series")
		return
	}
	for za0001 := range z.Series {
		if z.Series[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Series[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Series", za0001)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AddRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Series)))
	for za0001 := range z.Series {
		if z.Series[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Series[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Series", za0001)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AddRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Series")
		return
	}
	if cap(z.Series) >= int(zb0002) {
		z.Series = (z.Series)[:zb0002]
	} else {
		z.Series = make([]*msg.Series, zb0002)
	}
	for za0001 := range z.Series {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			z.Series[za0001] = nil
		} else {
			if z.Series[za0001] == nil {
				z.Series[za0001] = new(msg.Series)
			}
			bts, err = z.Series[za0001].UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Series", za0001)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AddRequest) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize
	for za0001 := range z.Series {
		if z.Series[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Series[za0001].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LabelValuesRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "matchers":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Matchers")
				return
			}
			if cap(z.Matchers) >= int(zb0002) {
				z.Matchers = (z.Matchers)[:zb0002]
			} else {
				z.Matchers = make([]*Matcher, zb0002)
			}
			for za0001 := range z.Matchers {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Matchers", za0001)
						return
					}
					z.Matchers[za0001] = nil
				} else {
					if z.Matchers[za0001] == nil {
						z.Matchers[za0001] = new(Matcher)
					}
					err = z.Matchers[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Matchers", za0001)
						return
					}
				}
			}
		case "spanCtx":
			z.SpanCtx, err = dc.ReadBytes(z.SpanCtx)
			if err != nil {
				err = msgp.WrapError(err, "SpanCtx")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *LabelValuesRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "name"
	err = en.Append(0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "matchers"
	err = en.Append(0xa8, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Matchers)))
	if err != nil {
		err = msgp.WrapError(err, "Matchers")
		return
	}
	for za0001 := range z.Matchers {
		if z.Matchers[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Matchers[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Matchers", za0001)
				return
			}
		}
	}
	// write "spanCtx"
	err = en.Append(0xa7, 0x73, 0x70, 0x61, 0x6e, 0x43, 0x74, 0x78)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.SpanCtx)
	if err != nil {
		err = msgp.WrapError(err, "SpanCtx")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LabelValuesRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "name"
	o = append(o, 0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "matchers"
	o = append(o, 0xa8, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Matchers)))
	for za0001 := range z.Matchers {
		if z.Matchers[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Matchers[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Matchers", za0001)
				return
			}
		}
	}
	// string "spanCtx"
	o = append(o, 0xa7, 0x73, 0x70, 0x61, 0x6e, 0x43, 0x74, 0x78)
	o = msgp.AppendBytes(o, z.SpanCtx)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LabelValuesRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "matchers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Matchers")
				return
			}
			if cap(z.Matchers) >= int(zb0002) {
				z.Matchers = (z.Matchers)[:zb0002]
			} else {
				z.Matchers = make([]*Matcher, zb0002)
			}
			for za0001 := range z.Matchers {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Matchers[za0001] = nil
				} else {
					if z.Matchers[za0001] == nil {
						z.Matchers[za0001] = new(Matcher)
					}
					bts, err = z.Matchers[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Matchers", za0001)
						return
					}
				}
			}
		case "spanCtx":
			z.SpanCtx, bts, err = msgp.ReadBytesBytes(bts, z.SpanCtx)
			if err != nil {
				err = msgp.WrapError(err, "SpanCtx")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *LabelValuesRequest) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.ArrayHeaderSize
	for za0001 := range z.Matchers {
		if z.Matchers[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Matchers[za0001].Msgsize()
		}
	}
	s += 8 + msgp.BytesPrefixSize + len(z.SpanCtx)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MatchType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 byte
		zb0001, err = dc.ReadByte()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = MatchType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z MatchType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteByte(byte(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MatchType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendByte(o, byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MatchType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 byte
		zb0001, bts, err = msgp.ReadByteBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = MatchType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MatchType) Msgsize() (s int) {
	s = msgp.ByteSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Matcher) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Type":
			{
				var zb0002 byte
				zb0002, err = dc.ReadByte()
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = MatchType(zb0002)
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Value":
			z.Value, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Matcher) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Type"
	err = en.Append(0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteByte(byte(z.Type))
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Matcher) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Type"
	o = append(o, 0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendByte(o, byte(z.Type))
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Matcher) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Type":
			{
				var zb0002 byte
				zb0002, bts, err = msgp.ReadByteBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = MatchType(zb0002)
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Matcher) Msgsize() (s int) {
	s = 1 + 5 + msgp.ByteSize + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.StringPrefixSize + len(z.Value)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SelectRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "mint":
			z.Mint, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Mint")
				return
			}
		case "maxt":
			z.Maxt, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Maxt")
				return
			}
		case "interval":
			z.Interval, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Interval")
				return
			}
		case "matchers":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Matchers")
				return
			}
			if cap(z.Matchers) >= int(zb0002) {
				z.Matchers = (z.Matchers)[:zb0002]
			} else {
				z.Matchers = make([]*Matcher, zb0002)
			}
			for za0001 := range z.Matchers {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Matchers", za0001)
						return
					}
					z.Matchers[za0001] = nil
				} else {
					if z.Matchers[za0001] == nil {
						z.Matchers[za0001] = new(Matcher)
					}
					err = z.Matchers[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Matchers", za0001)
						return
					}
				}
			}
		case "spanCtx":
			z.SpanCtx, err = dc.ReadBytes(z.SpanCtx)
			if err != nil {
				err = msgp.WrapError(err, "SpanCtx")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SelectRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "mint"
	err = en.Append(0x85, 0xa4, 0x6d, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Mint)
	if err != nil {
		err = msgp.WrapError(err, "Mint")
		return
	}
	// write "maxt"
	err = en.Append(0xa4, 0x6d, 0x61, 0x78, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Maxt)
	if err != nil {
		err = msgp.WrapError(err, "Maxt")
		return
	}
	// write "interval"
	err = en.Append(0xa8, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Interval)
	if err != nil {
		err = msgp.WrapError(err, "Interval")
		return
	}
	// write "matchers"
	err = en.Append(0xa8, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Matchers)))
	if err != nil {
		err = msgp.WrapError(err, "Matchers")
		return
	}
	for za0001 := range z.Matchers {
		if z.Matchers[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Matchers[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Matchers", za0001)
				return
			}
		}
	}
	// write "spanCtx"
	err = en.Append(0xa7, 0x73, 0x70, 0x61, 0x6e, 0x43, 0x74, 0x78)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.SpanCtx)
	if err != nil {
		err = msgp.WrapError(err, "SpanCtx")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SelectRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "mint"
	o = append(o, 0x85, 0xa4, 0x6d, 0x69, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.Mint)
	// string "maxt"
	o = append(o, 0xa4, 0x6d, 0x61, 0x78, 0x74)
	o = msgp.AppendInt64(o, z.Maxt)
	// string "interval"
	o = append(o, 0xa8, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.Interval)
	// string "matchers"
	o = append(o, 0xa8, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Matchers)))
	for za0001 := range z.Matchers {
		if z.Matchers[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Matchers[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Matchers", za0001)
				return
			}
		}
	}
	// string "spanCtx"
	o = append(o, 0xa7, 0x73, 0x70, 0x61, 0x6e, 0x43, 0x74, 0x78)
	o = msgp.AppendBytes(o, z.SpanCtx)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SelectRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "mint":
			z.Mint, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Mint")
				return
			}
		case "maxt":
			z.Maxt, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Maxt")
				return
			}
		case "interval":
			z.Interval, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Interval")
				return
			}
		case "matchers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Matchers")
				return
			}
			if cap(z.Matchers) >= int(zb0002) {
				z.Matchers = (z.Matchers)[:zb0002]
			} else {
				z.Matchers = make([]*Matcher, zb0002)
			}
			for za0001 := range z.Matchers {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Matchers[za0001] = nil
				} else {
					if z.Matchers[za0001] == nil {
						z.Matchers[za0001] = new(Matcher)
					}
					bts, err = z.Matchers[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Matchers", za0001)
						return
					}
				}
			}
		case "spanCtx":
			z.SpanCtx, bts, err = msgp.ReadBytesBytes(bts, z.SpanCtx)
			if err != nil {
				err = msgp.WrapError(err, "SpanCtx")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SelectRequest) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int64Size + 5 + msgp.Int64Size + 9 + msgp.Int64Size + 9 + msgp.ArrayHeaderSize
	for za0001 := range z.Matchers {
		if z.Matchers[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Matchers[za0001].Msgsize()
		}
	}
	s += 8 + msgp.BytesPrefixSize + len(z.SpanCtx)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SelectResponse) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "status":
			err = z.Status.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Status")
				return
			}
		case "series":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Series")
				return
			}
			if cap(z.Series) >= int(zb0002) {
				z.Series = (z.Series)[:zb0002]
			} else {
				z.Series = make([]*msg.Series, zb0002)
			}
			for za0001 := range z.Series {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Series", za0001)
						return
					}
					z.Series[za0001] = nil
				} else {
					if z.Series[za0001] == nil {
						z.Series[za0001] = new(msg.Series)
					}
					err = z.Series[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Series", za0001)
						return
					}
				}
			}
		case "errorMsg":
			z.ErrorMsg, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ErrorMsg")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SelectResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "status"
	err = en.Append(0x83, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return
	}
	err = z.Status.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	// write "series"
	err = en.Append(0xa6, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Series)))
	if err != nil {
		err = msgp.WrapError(err, "Series")
		return
	}
	for za0001 := range z.Series {
		if z.Series[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Series[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Series", za0001)
				return
			}
		}
	}
	// write "errorMsg"
	err = en.Append(0xa8, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.ErrorMsg)
	if err != nil {
		err = msgp.WrapError(err, "ErrorMsg")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SelectResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "status"
	o = append(o, 0x83, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	o, err = z.Status.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	// string "series"
	o = append(o, 0xa6, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Series)))
	for za0001 := range z.Series {
		if z.Series[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Series[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Series", za0001)
				return
			}
		}
	}
	// string "errorMsg"
	o = append(o, 0xa8, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67)
	o = msgp.AppendString(o, z.ErrorMsg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SelectResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "status":
			bts, err = z.Status.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Status")
				return
			}
		case "series":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Series")
				return
			}
			if cap(z.Series) >= int(zb0002) {
				z.Series = (z.Series)[:zb0002]
			} else {
				z.Series = make([]*msg.Series, zb0002)
			}
			for za0001 := range z.Series {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Series[za0001] = nil
				} else {
					if z.Series[za0001] == nil {
						z.Series[za0001] = new(msg.Series)
					}
					bts, err = z.Series[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Series", za0001)
						return
					}
				}
			}
		case "errorMsg":
			z.ErrorMsg, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ErrorMsg")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SelectResponse) Msgsize() (s int) {
	s = 1 + 7 + z.Status.Msgsize() + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Series {
		if z.Series[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Series[za0001].Msgsize()
		}
	}
	s += 9 + msgp.StringPrefixSize + len(z.ErrorMsg)
	return
}