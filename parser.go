package imgi

import (
	"fmt"
	"strconv"
	"strings"
)

func parseOperations(raw string) ([]Operation, error) {
	if raw == "" {
		return nil, nil
	}
	raws := strings.Split(raw, "|")
	ops := make([]Operation, 0, len(raws))
	for _, s := range raws {
		rawOp := strings.Split(s, ":")
		if l := len(rawOp); l < 1 || l > 2 {
			return nil, fmt.Errorf("error parsing operation with [%s]", s)
		}

		opName := strings.TrimSpace(rawOp[0])
		if opName == "" {
			continue
		}

		action, exist := actions[opName]
		if !exist {
			return nil, fmt.Errorf("not supported operation: [%s]", opName)
		}
		var opts = Options{}
		if len(rawOp) == 2 {
			rawOpts := normalizeOptions(rawOp[1])
			for k, v := range rawOpts {
				parse, ok := parsers[k]
				if !ok {
					continue
				}
				if err := parse(&opts, v); err != nil {
					msg := `error while parsing "%s" with value "%s" for [%s], error: %s`
					return nil, fmt.Errorf(msg, k, v, opName, err.Error())
				}
			}

		}
		op := Operation{
			Name:    opName,
			Options: opts,
			Action:  action,
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func normalizeOptions(raw string) map[string]string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	raws := strings.Split(raw, ",")
	opts := make(map[string]string, len(raws))
	for _, s := range raws {
		s = strings.TrimSpace(s)
		kv := strings.Split(s, "_")
		len := len(kv)
		if len == 0 {
			continue
		}
		k := kv[0]
		v := ""
		if len == 2 {
			v = kv[1]
		}
		opts[k] = v
	}
	return opts
}

const (
	maxWidth  = 10000
	maxHeight = 10000
)

type parser func(*Options, string) error

var parsers = map[string]parser{
	"m": parseMode,
	"w": parseWidth,
	"h": parseHeight,
}

var modes = map[string]Mode{
	"fit":  Fit,
	"fill": Fill,
	"flex": Flex,
}

func parseMode(opts *Options, val string) (err error) {
	if mode, ok := modes[val]; ok {
		opts.Mode = mode
		return nil
	}
	return fmt.Errorf(`fail parse mode with value "%s"`, val)
}

func parseIntInRange(val string, min, max int) (int, error) {
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	if v < min || v > max {
		return 0, fmt.Errorf("value must in [%d, %d]", min, max)
	}
	return v, nil
}

func parseWidth(opts *Options, val string) (err error) {
	opts.Width, err = parseIntInRange(val, 1, maxWidth)
	return err
}

func parseHeight(opts *Options, val string) (err error) {
	opts.Height, err = parseIntInRange(val, 1, maxHeight)
	return err
}
