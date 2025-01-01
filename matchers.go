//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/it
//

package it

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//
// String
//

type String string

func (x String) HavePrefix(y string) error {
	assert := fmt.Errorf("string %s have prefix %s", x, y)

	if !strings.HasPrefix(string(x), y) {
		return assert
	}

	return passed(assert)
}

func (x String) HaveSuffix(y string) error {
	assert := fmt.Errorf("string %s have suffix %s", x, y)

	if !strings.HasSuffix(string(x), y) {
		return assert
	}

	return passed(assert)
}

func (x String) Contain(y string) error {
	assert := fmt.Errorf("string %s contain %s", x, y)

	if !strings.Contains(string(x), y) {
		return assert
	}

	return passed(assert)
}

//
// Sequence of elements
//

type SeqOf[A any] []A

func Seq[A any](seq []A) SeqOf[A] {
	return SeqOf[A](seq)
}

func (xs SeqOf[A]) BeEmpty() error {
	assert := fmt.Errorf("seq %v be empty", xs)

	if len(xs) != 0 {
		return assert
	}

	return passed(assert)
}

func (xs SeqOf[A]) Equal(ys ...A) error {
	if len(xs) != len(ys) {
		return fmt.Errorf("seq %v length be equal to %v", xs, ys)
	}

	for i, x := range xs {
		if !equal(x, ys[i]) {
			return fmt.Errorf("seq %dth element of %v be equal to %v", i, x, ys[i])
		}
	}

	return passed(fmt.Errorf("seq %v is equal to %v", xs, ys))
}

func (xs SeqOf[A]) Contain(ys ...A) SeqContainIt[A] {
	assert := fmt.Errorf("seq %v contain %v", xs, ys)

	for _, y := range ys {
		has := false
		for _, x := range xs {
			if equal(x, y) {
				has = true
				break
			}
		}
		if !has {
			return SeqContainIt[A]{assert, xs}
		}
	}

	return SeqContainIt[A]{passed(assert), xs}
}

// FailIt extend Fail assert
type SeqContainIt[A any] struct {
	assert error
	xs     []A
}

func (x SeqContainIt[A]) Error() string      { return x.assert.Error() }
func (x SeqContainIt[A]) As(target any) bool { return errors.As(x.assert, target) }

func (x SeqContainIt[A]) AllOf(ys ...A) error {
	assert := fmt.Errorf("seq %v contain all of %v", x.xs, ys)

	for _, y := range ys {
		has := false
		for _, x := range x.xs {
			if equal(x, y) {
				has = true
				break
			}
		}
		if !has {
			return assert
		}
	}

	return passed(assert)
}

func (x SeqContainIt[A]) OneOf(ys ...A) error {
	assert := fmt.Errorf("seq %v contain one of %v", x.xs, ys)

	for _, y := range ys {
		for _, x := range x.xs {
			if equal(x, y) {
				return passed(assert)
			}
		}
	}

	return assert
}

//
// Map of elements
//

type MapOf[K comparable, V any] map[K]V

func Map[K comparable, V any](val map[K]V) MapOf[K, V] {
	return MapOf[K, V](val)
}

func (xs MapOf[K, V]) Have(key K, y V) error {
	x, exists := xs[key]
	if !exists {
		return fmt.Errorf("map %v have key %v", xs, key)
	}

	assert := fmt.Errorf("key %v value %v of %T be equal to %v", key, x, (map[K]V)(xs), y)

	if !equal(x, y) {
		return assert
	}

	return passed(assert)
}

//
// Json
//

type JsonOf[A any] struct{ obj A }

func Json[A any](obj A) JsonOf[A] {
	return JsonOf[A]{obj: obj}
}

func (obj JsonOf[A]) Equiv(shapes ...string) error {
	var val any
	raw, err := json.Marshal(obj.obj)
	if err != nil {
		return fmt.Errorf("input be valid JSON")
	}
	if err := json.Unmarshal([]byte(raw), &val); err != nil {
		return fmt.Errorf("input be valid JSON")
	}

	for _, shape := range shapes {
		var pat any
		if err := json.Unmarshal([]byte(shape), &pat); err != nil {
			return fmt.Errorf("pattern be valid JSON")
		}

		dv := diffVal(pat, val)
		if dv != nil {
			var sb strings.Builder
			p := newPrinter(&sb)
			p.print("", dv)

			return fmt.Errorf("be matching\n%s", sb.String())
		}
	}

	return passed(fmt.Errorf("be matching"))
}

type diff struct {
	expect any // -
	actual any // +
}

func diffVal(pat, val any) any {
	if pp, ok := pat.(string); ok && pp == "_" {
		return nil
	}

	switch vv := val.(type) {
	case string:
		pp, ok := pat.(string)
		if !ok {
			return diff{expect: pat, actual: val}
		}
		if strings.HasPrefix(pp, "m/") && pp[len(pp)-1] == '/' {
			re := regexp.MustCompile(pp[2 : len(pp)-1])
			if !re.MatchString(vv) {
				return diff{expect: pat, actual: val}
			}
		} else if vv != pp {
			return diff{expect: pat, actual: val}
		}

		return nil
	case float64:
		pp, ok := pat.(float64)
		if !ok || vv != pp {
			return diff{expect: pat, actual: val}
		}
		return nil
	case bool:
		pp, ok := pat.(bool)
		if !ok || vv != pp {
			return diff{expect: pat, actual: val}
		}
		return nil
	case []any:
		pp, ok := pat.([]any)
		if !ok {
			return diff{expect: pat, actual: val}
		}

		add := make([]any, 0)
		sub := make([]any, 0)
		for i, vvx := range vv {
			if len(pp)-1 < i {
				sub = append(add, vv[i:])
				break
			}
			if pp[i] == "..." {
				break
			}
			if dv := diffVal(pp[i], vvx); dv != nil {
				kv := dv.(diff)
				if kv.actual != nil {
					add = append(add, kv.actual)
				}
				if kv.expect != nil {
					sub = append(sub, kv.expect)
				}
			}
		}

		if len(add) != 0 || len(sub) != 0 {
			return diff{expect: sub, actual: add}
		}

		return nil
	case map[string]any:
		pp, ok := pat.(map[string]any)
		if !ok {
			return diff{expect: pat, actual: val}
		}
		return diffMap(pp, vv)
	}

	return nil
}

func diffMap(pat, val map[string]any) any {
	d := make(map[string]any)

	for k, p := range pat {
		v, has := val[k]
		if !has {
			d[k] = diff{expect: p}
		}

		if dv := diffVal(p, v); dv != nil {
			d[k] = dv
		}
	}

	if len(d) != 0 {
		return d
	}

	return nil
}

//------------------------------------------------------------------------------

type printer struct {
	sb *strings.Builder
}

func newPrinter(sb *strings.Builder) printer {
	sb.WriteString("\x1b[0m")
	return printer{sb: sb}
}

func (p printer) atos(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func (p printer) value(sign, indent string, key string, v any) {
	p.sb.WriteString(fmt.Sprintf("%s %s", sign, indent))
	if key != "" {
		p.sb.WriteString(fmt.Sprintf("\"%s\": ", key))
	}
	p.sb.WriteString(p.atos(v))
}

func (p printer) diff(indent string, key string, v diff) {
	if v.expect != nil {
		if v.actual == nil {
			p.sb.WriteString("\x1b[1;43m")
		} else {
			p.sb.WriteString("\x1b[1;33m")
		}
		p.value("-", indent, key, v.expect)
		p.sb.WriteString("\x1b[0m")
		p.sb.WriteString("\n")
	}

	if v.actual != nil {
		p.sb.WriteString("\x1b[1;41m")
		p.value("+", indent, key, v.actual)
		p.sb.WriteString("\x1b[0m")
		p.sb.WriteString("\n")
	}
}

func (p printer) print(indent string, val any) {
	switch obj := val.(type) {
	case diff:
		p.diff(indent, "", obj)
	case map[string]any:
		p.sb.WriteString("  {\n")
		for k, v := range obj {
			switch vv := v.(type) {
			case diff:
				p.sb.WriteString("\n")
				p.diff(indent+"  ", k, vv)
				p.sb.WriteString("\n")
			default:
				p.sb.WriteString(fmt.Sprintf("  %s\"%s\": ", indent+"  ", k))
				p.print(indent+"  ", v)
			}
			p.sb.WriteString(fmt.Sprintf("  %s}\n", indent))
		}
	default:
		p.sb.WriteString(p.atos(obj))
	}
}
