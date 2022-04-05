// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bangumi

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson54bdebb0DecodeBangumiEntityBangumi(in *jlexer.Lexer, out *BilibiliData) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "season_id":
			out.SeasonID = int(in.Int())
		case "media_id":
			out.MediaID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "is_vip":
			out.IsVip = bool(in.Bool())
		case "is_exclusive":
			out.IsExclusive = bool(in.Bool())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson54bdebb0EncodeBangumiEntityBangumi(out *jwriter.Writer, in BilibiliData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"season_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.SeasonID))
	}
	{
		const prefix string = ",\"media_id\":"
		out.RawString(prefix)
		out.Int(int(in.MediaID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"is_vip\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsVip))
	}
	{
		const prefix string = ",\"is_exclusive\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsExclusive))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BilibiliData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson54bdebb0EncodeBangumiEntityBangumi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BilibiliData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson54bdebb0EncodeBangumiEntityBangumi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BilibiliData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson54bdebb0DecodeBangumiEntityBangumi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BilibiliData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson54bdebb0DecodeBangumiEntityBangumi(l, v)
}
func easyjson54bdebb0DecodeBangumiEntityBangumi1(in *jlexer.Lexer, out *Bilibili) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]BilibiliData, 0, 1)
					} else {
						out.Data = []BilibiliData{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 BilibiliData
					(v1).UnmarshalEasyJSON(in)
					out.Data = append(out.Data, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson54bdebb0EncodeBangumiEntityBangumi1(out *jwriter.Writer, in Bilibili) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bilibili) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson54bdebb0EncodeBangumiEntityBangumi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bilibili) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson54bdebb0EncodeBangumiEntityBangumi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bilibili) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson54bdebb0DecodeBangumiEntityBangumi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bilibili) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson54bdebb0DecodeBangumiEntityBangumi1(l, v)
}
