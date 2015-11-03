// Copyright (c) 2015 Arista Networks, Inc.  All rights reserved.
// Arista Networks, Inc. Confidential and Proprietary.

package key

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Stringify transforms an arbitrary interface into its string
// representation.  We need to do this because some entities use the string
// representation of their keys as their names.
func Stringify(key interface{}) (string, error) {
	if key == nil {
		return "", errors.New("Unable to stringify nil")
	}
	var str string
	switch key := key.(type) {
	case bool:
		str = strconv.FormatBool(key)
	case uint8:
		str = strconv.FormatUint(uint64(key), 10)
	case uint16:
		str = strconv.FormatUint(uint64(key), 10)
	case uint32:
		str = strconv.FormatUint(uint64(key), 10)
	case uint64:
		str = strconv.FormatUint(key, 10)
	case int8:
		str = strconv.FormatInt(int64(key), 10)
	case int16:
		str = strconv.FormatInt(int64(key), 10)
	case int32:
		str = strconv.FormatInt(int64(key), 10)
	case int64:
		str = strconv.FormatInt(key, 10)
	case float32:
		str = "f" + strconv.FormatInt(int64(math.Float32bits(key)), 10)
	case float64:
		str = "f" + strconv.FormatInt(int64(math.Float64bits(key)), 10)
	case string:
		str = key
	case map[string]interface{}:
		keys := SortedKeys(key)
		for _, k := range keys {
			v := key[k]
			if len(str) > 0 {
				str += "_"
			}
			s, err := Stringify(v)
			if err != nil {
				return str, err
			}
			str += s
		}
	case *map[string]interface{}:
		return Stringify(*key)
	case Keyable:
		return key.KeyString(), nil

	default:
		return "", fmt.Errorf("Unable to stringify type %T", key)
	}

	return str, nil
}