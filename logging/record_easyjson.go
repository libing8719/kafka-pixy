// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package logging

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

func easyjson500470b6DecodeGithubComMailgunLogrusHooksCommon(in *jlexer.Lexer, out *LogRecord) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "context":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Context = make(map[string]interface{})
				} else {
					out.Context = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					if m, ok := v1.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v1.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v1 = in.Interface()
					}
					(out.Context)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "category":
			out.Category = string(in.String())
		case "appname":
			out.AppName = string(in.String())
		case "hostname":
			out.HostName = string(in.String())
		case "logLevel":
			out.LogLevel = string(in.String())
		case "filename":
			out.FileName = string(in.String())
		case "funcName":
			out.FuncName = string(in.String())
		case "lineno":
			out.LineNo = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "timestamp":
			out.Timestamp = number(in.Float64())
		case "cid":
			out.CID = string(in.String())
		case "pid":
			out.PID = int(in.Int())
		case "tid":
			out.TID = string(in.String())
		case "excType":
			out.ExcType = string(in.String())
		case "excText":
			out.ExcText = string(in.String())
		case "excValue":
			out.ExcValue = string(in.String())
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
func easyjson500470b6EncodeGithubComMailgunLogrusHooksCommon(out *jwriter.Writer, in LogRecord) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Context) != 0 {
		const prefix string = ",\"context\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Context {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if m, ok := v2Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v2Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v2Value))
				}
			}
			out.RawByte('}')
		}
	}
	if in.Category != "" {
		const prefix string = ",\"category\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"appname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AppName))
	}
	{
		const prefix string = ",\"hostname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HostName))
	}
	{
		const prefix string = ",\"logLevel\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LogLevel))
	}
	{
		const prefix string = ",\"filename\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FileName))
	}
	{
		const prefix string = ",\"funcName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FuncName))
	}
	{
		const prefix string = ",\"lineno\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.LineNo))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Timestamp).MarshalJSON())
	}
	if in.CID != "" {
		const prefix string = ",\"cid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CID))
	}
	if in.PID != 0 {
		const prefix string = ",\"pid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.PID))
	}
	if in.TID != "" {
		const prefix string = ",\"tid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TID))
	}
	if in.ExcType != "" {
		const prefix string = ",\"excType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExcType))
	}
	if in.ExcText != "" {
		const prefix string = ",\"excText\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExcText))
	}
	if in.ExcValue != "" {
		const prefix string = ",\"excValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExcValue))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LogRecord) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson500470b6EncodeGithubComMailgunLogrusHooksCommon(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LogRecord) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson500470b6EncodeGithubComMailgunLogrusHooksCommon(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LogRecord) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson500470b6DecodeGithubComMailgunLogrusHooksCommon(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LogRecord) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson500470b6DecodeGithubComMailgunLogrusHooksCommon(l, v)
}
