package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *StateBlock) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "type":
			err = z.Type.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "token":
			err = dc.ReadExtension(&z.Token)
			if err != nil {
				return
			}
		case "address":
			err = dc.ReadExtension(&z.Address)
			if err != nil {
				return
			}
		case "balance":
			err = dc.ReadExtension(&z.Balance)
			if err != nil {
				return
			}
		case "vote":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Vote = nil
			} else {
				if z.Vote == nil {
					z.Vote = new(Balance)
				}
				err = dc.ReadExtension(z.Vote)
				if err != nil {
					return
				}
			}
		case "network":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Network = nil
			} else {
				if z.Network == nil {
					z.Network = new(Balance)
				}
				err = dc.ReadExtension(z.Network)
				if err != nil {
					return
				}
			}
		case "storage":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Storage = nil
			} else {
				if z.Storage == nil {
					z.Storage = new(Balance)
				}
				err = dc.ReadExtension(z.Storage)
				if err != nil {
					return
				}
			}
		case "oracle":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Oracle = nil
			} else {
				if z.Oracle == nil {
					z.Oracle = new(Balance)
				}
				err = dc.ReadExtension(z.Oracle)
				if err != nil {
					return
				}
			}
		case "previous":
			err = dc.ReadExtension(&z.Previous)
			if err != nil {
				return
			}
		case "link":
			err = dc.ReadExtension(&z.Link)
			if err != nil {
				return
			}
		case "sender":
			z.Sender, err = dc.ReadBytes(z.Sender)
			if err != nil {
				return
			}
		case "receiver":
			z.Receiver, err = dc.ReadBytes(z.Receiver)
			if err != nil {
				return
			}
		case "message":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Message = nil
			} else {
				if z.Message == nil {
					z.Message = new(Hash)
				}
				err = dc.ReadExtension(z.Message)
				if err != nil {
					return
				}
			}
		case "data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				return
			}
		case "povHeight":
			z.PoVHeight, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "timestamp":
			z.Timestamp, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "extra":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Extra = nil
			} else {
				if z.Extra == nil {
					z.Extra = new(Hash)
				}
				err = dc.ReadExtension(z.Extra)
				if err != nil {
					return
				}
			}
		case "representative":
			err = dc.ReadExtension(&z.Representative)
			if err != nil {
				return
			}
		case "priFrom":
			z.PrivateFrom, err = dc.ReadString()
			if err != nil {
				return
			}
		case "priFor":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.PrivateFor) >= int(zb0002) {
				z.PrivateFor = (z.PrivateFor)[:zb0002]
			} else {
				z.PrivateFor = make([]string, zb0002)
			}
			for za0001 := range z.PrivateFor {
				z.PrivateFor[za0001], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "priGid":
			z.PrivateGroupID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "work":
			err = dc.ReadExtension(&z.Work)
			if err != nil {
				return
			}
		case "signature":
			err = dc.ReadExtension(&z.Signature)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StateBlock) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 23
	// write "type"
	err = en.Append(0xde, 0x0, 0x17, 0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = z.Type.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "token"
	err = en.Append(0xa5, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Token)
	if err != nil {
		return
	}
	// write "address"
	err = en.Append(0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Address)
	if err != nil {
		return
	}
	// write "balance"
	err = en.Append(0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Balance)
	if err != nil {
		return
	}
	// write "vote"
	err = en.Append(0xa4, 0x76, 0x6f, 0x74, 0x65)
	if err != nil {
		return
	}
	if z.Vote == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteExtension(z.Vote)
		if err != nil {
			return
		}
	}
	// write "network"
	err = en.Append(0xa7, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b)
	if err != nil {
		return
	}
	if z.Network == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteExtension(z.Network)
		if err != nil {
			return
		}
	}
	// write "storage"
	err = en.Append(0xa7, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	if z.Storage == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteExtension(z.Storage)
		if err != nil {
			return
		}
	}
	// write "oracle"
	err = en.Append(0xa6, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65)
	if err != nil {
		return
	}
	if z.Oracle == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteExtension(z.Oracle)
		if err != nil {
			return
		}
	}
	// write "previous"
	err = en.Append(0xa8, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Previous)
	if err != nil {
		return
	}
	// write "link"
	err = en.Append(0xa4, 0x6c, 0x69, 0x6e, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Link)
	if err != nil {
		return
	}
	// write "sender"
	err = en.Append(0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Sender)
	if err != nil {
		return
	}
	// write "receiver"
	err = en.Append(0xa8, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Receiver)
	if err != nil {
		return
	}
	// write "message"
	err = en.Append(0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	if z.Message == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteExtension(z.Message)
		if err != nil {
			return
		}
	}
	// write "data"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		return
	}
	// write "povHeight"
	err = en.Append(0xa9, 0x70, 0x6f, 0x76, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.PoVHeight)
	if err != nil {
		return
	}
	// write "timestamp"
	err = en.Append(0xa9, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Timestamp)
	if err != nil {
		return
	}
	// write "extra"
	err = en.Append(0xa5, 0x65, 0x78, 0x74, 0x72, 0x61)
	if err != nil {
		return
	}
	if z.Extra == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteExtension(z.Extra)
		if err != nil {
			return
		}
	}
	// write "representative"
	err = en.Append(0xae, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Representative)
	if err != nil {
		return
	}
	// write "priFrom"
	err = en.Append(0xa7, 0x70, 0x72, 0x69, 0x46, 0x72, 0x6f, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteString(z.PrivateFrom)
	if err != nil {
		return
	}
	// write "priFor"
	err = en.Append(0xa6, 0x70, 0x72, 0x69, 0x46, 0x6f, 0x72)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.PrivateFor)))
	if err != nil {
		return
	}
	for za0001 := range z.PrivateFor {
		err = en.WriteString(z.PrivateFor[za0001])
		if err != nil {
			return
		}
	}
	// write "priGid"
	err = en.Append(0xa6, 0x70, 0x72, 0x69, 0x47, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.PrivateGroupID)
	if err != nil {
		return
	}
	// write "work"
	err = en.Append(0xa4, 0x77, 0x6f, 0x72, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Work)
	if err != nil {
		return
	}
	// write "signature"
	err = en.Append(0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Signature)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StateBlock) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 23
	// string "type"
	o = append(o, 0xde, 0x0, 0x17, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o, err = z.Type.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "token"
	o = append(o, 0xa5, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	o, err = msgp.AppendExtension(o, &z.Token)
	if err != nil {
		return
	}
	// string "address"
	o = append(o, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o, err = msgp.AppendExtension(o, &z.Address)
	if err != nil {
		return
	}
	// string "balance"
	o = append(o, 0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	o, err = msgp.AppendExtension(o, &z.Balance)
	if err != nil {
		return
	}
	// string "vote"
	o = append(o, 0xa4, 0x76, 0x6f, 0x74, 0x65)
	if z.Vote == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = msgp.AppendExtension(o, z.Vote)
		if err != nil {
			return
		}
	}
	// string "network"
	o = append(o, 0xa7, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b)
	if z.Network == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = msgp.AppendExtension(o, z.Network)
		if err != nil {
			return
		}
	}
	// string "storage"
	o = append(o, 0xa7, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65)
	if z.Storage == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = msgp.AppendExtension(o, z.Storage)
		if err != nil {
			return
		}
	}
	// string "oracle"
	o = append(o, 0xa6, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65)
	if z.Oracle == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = msgp.AppendExtension(o, z.Oracle)
		if err != nil {
			return
		}
	}
	// string "previous"
	o = append(o, 0xa8, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73)
	o, err = msgp.AppendExtension(o, &z.Previous)
	if err != nil {
		return
	}
	// string "link"
	o = append(o, 0xa4, 0x6c, 0x69, 0x6e, 0x6b)
	o, err = msgp.AppendExtension(o, &z.Link)
	if err != nil {
		return
	}
	// string "sender"
	o = append(o, 0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Sender)
	// string "receiver"
	o = append(o, 0xa8, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72)
	o = msgp.AppendBytes(o, z.Receiver)
	// string "message"
	o = append(o, 0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if z.Message == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = msgp.AppendExtension(o, z.Message)
		if err != nil {
			return
		}
	}
	// string "data"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.Data)
	// string "povHeight"
	o = append(o, 0xa9, 0x70, 0x6f, 0x76, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendUint64(o, z.PoVHeight)
	// string "timestamp"
	o = append(o, 0xa9, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.Timestamp)
	// string "extra"
	o = append(o, 0xa5, 0x65, 0x78, 0x74, 0x72, 0x61)
	if z.Extra == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = msgp.AppendExtension(o, z.Extra)
		if err != nil {
			return
		}
	}
	// string "representative"
	o = append(o, 0xae, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65)
	o, err = msgp.AppendExtension(o, &z.Representative)
	if err != nil {
		return
	}
	// string "priFrom"
	o = append(o, 0xa7, 0x70, 0x72, 0x69, 0x46, 0x72, 0x6f, 0x6d)
	o = msgp.AppendString(o, z.PrivateFrom)
	// string "priFor"
	o = append(o, 0xa6, 0x70, 0x72, 0x69, 0x46, 0x6f, 0x72)
	o = msgp.AppendArrayHeader(o, uint32(len(z.PrivateFor)))
	for za0001 := range z.PrivateFor {
		o = msgp.AppendString(o, z.PrivateFor[za0001])
	}
	// string "priGid"
	o = append(o, 0xa6, 0x70, 0x72, 0x69, 0x47, 0x69, 0x64)
	o = msgp.AppendString(o, z.PrivateGroupID)
	// string "work"
	o = append(o, 0xa4, 0x77, 0x6f, 0x72, 0x6b)
	o, err = msgp.AppendExtension(o, &z.Work)
	if err != nil {
		return
	}
	// string "signature"
	o = append(o, 0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	o, err = msgp.AppendExtension(o, &z.Signature)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StateBlock) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "type":
			bts, err = z.Type.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "token":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Token)
			if err != nil {
				return
			}
		case "address":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Address)
			if err != nil {
				return
			}
		case "balance":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Balance)
			if err != nil {
				return
			}
		case "vote":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Vote = nil
			} else {
				if z.Vote == nil {
					z.Vote = new(Balance)
				}
				bts, err = msgp.ReadExtensionBytes(bts, z.Vote)
				if err != nil {
					return
				}
			}
		case "network":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Network = nil
			} else {
				if z.Network == nil {
					z.Network = new(Balance)
				}
				bts, err = msgp.ReadExtensionBytes(bts, z.Network)
				if err != nil {
					return
				}
			}
		case "storage":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Storage = nil
			} else {
				if z.Storage == nil {
					z.Storage = new(Balance)
				}
				bts, err = msgp.ReadExtensionBytes(bts, z.Storage)
				if err != nil {
					return
				}
			}
		case "oracle":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Oracle = nil
			} else {
				if z.Oracle == nil {
					z.Oracle = new(Balance)
				}
				bts, err = msgp.ReadExtensionBytes(bts, z.Oracle)
				if err != nil {
					return
				}
			}
		case "previous":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Previous)
			if err != nil {
				return
			}
		case "link":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Link)
			if err != nil {
				return
			}
		case "sender":
			z.Sender, bts, err = msgp.ReadBytesBytes(bts, z.Sender)
			if err != nil {
				return
			}
		case "receiver":
			z.Receiver, bts, err = msgp.ReadBytesBytes(bts, z.Receiver)
			if err != nil {
				return
			}
		case "message":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Message = nil
			} else {
				if z.Message == nil {
					z.Message = new(Hash)
				}
				bts, err = msgp.ReadExtensionBytes(bts, z.Message)
				if err != nil {
					return
				}
			}
		case "data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				return
			}
		case "povHeight":
			z.PoVHeight, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "timestamp":
			z.Timestamp, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "extra":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Extra = nil
			} else {
				if z.Extra == nil {
					z.Extra = new(Hash)
				}
				bts, err = msgp.ReadExtensionBytes(bts, z.Extra)
				if err != nil {
					return
				}
			}
		case "representative":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Representative)
			if err != nil {
				return
			}
		case "priFrom":
			z.PrivateFrom, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "priFor":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.PrivateFor) >= int(zb0002) {
				z.PrivateFor = (z.PrivateFor)[:zb0002]
			} else {
				z.PrivateFor = make([]string, zb0002)
			}
			for za0001 := range z.PrivateFor {
				z.PrivateFor[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "priGid":
			z.PrivateGroupID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "work":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Work)
			if err != nil {
				return
			}
		case "signature":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Signature)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StateBlock) Msgsize() (s int) {
	s = 3 + 5 + z.Type.Msgsize() + 6 + msgp.ExtensionPrefixSize + z.Token.Len() + 8 + msgp.ExtensionPrefixSize + z.Address.Len() + 8 + msgp.ExtensionPrefixSize + z.Balance.Len() + 5
	if z.Vote == nil {
		s += msgp.NilSize
	} else {
		s += msgp.ExtensionPrefixSize + z.Vote.Len()
	}
	s += 8
	if z.Network == nil {
		s += msgp.NilSize
	} else {
		s += msgp.ExtensionPrefixSize + z.Network.Len()
	}
	s += 8
	if z.Storage == nil {
		s += msgp.NilSize
	} else {
		s += msgp.ExtensionPrefixSize + z.Storage.Len()
	}
	s += 7
	if z.Oracle == nil {
		s += msgp.NilSize
	} else {
		s += msgp.ExtensionPrefixSize + z.Oracle.Len()
	}
	s += 9 + msgp.ExtensionPrefixSize + z.Previous.Len() + 5 + msgp.ExtensionPrefixSize + z.Link.Len() + 7 + msgp.BytesPrefixSize + len(z.Sender) + 9 + msgp.BytesPrefixSize + len(z.Receiver) + 8
	if z.Message == nil {
		s += msgp.NilSize
	} else {
		s += msgp.ExtensionPrefixSize + z.Message.Len()
	}
	s += 5 + msgp.BytesPrefixSize + len(z.Data) + 10 + msgp.Uint64Size + 10 + msgp.Int64Size + 6
	if z.Extra == nil {
		s += msgp.NilSize
	} else {
		s += msgp.ExtensionPrefixSize + z.Extra.Len()
	}
	s += 15 + msgp.ExtensionPrefixSize + z.Representative.Len() + 8 + msgp.StringPrefixSize + len(z.PrivateFrom) + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.PrivateFor {
		s += msgp.StringPrefixSize + len(z.PrivateFor[za0001])
	}
	s += 7 + msgp.StringPrefixSize + len(z.PrivateGroupID) + 5 + msgp.ExtensionPrefixSize + z.Work.Len() + 10 + msgp.ExtensionPrefixSize + z.Signature.Len()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StateBlockList) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(StateBlockList, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(StateBlock)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z StateBlockList) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z StateBlockList) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StateBlockList) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(StateBlockList, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(StateBlock)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z StateBlockList) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
