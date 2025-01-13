package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// GetNestedValueFromMap gets a nested value from a map
func GetNestedValueFromMap(data map[string]any, key string) any {
	keys := strings.Split(key, ".")
	return parseMap(data, keys)
}

//nolint:forcetypeassert
func parseMap(currentMap map[string]any, keys []string) any {
	if len(keys) == 0 {
		return currentMap
	}
	k := keys[0]
	remainKeys := keys[1:]
	if strings.Contains(k, "[") {
		idx := strings.Index(k, "[")
		preKey := k[:idx]
		idx2 := strings.Index(k, "]")
		postKey := k[idx+1 : idx2]
		if postKey == "" {
			return currentMap
		}
		index, err := strconv.Atoi(postKey)
		if err != nil {
			return currentMap
		}
		if slice, ok := currentMap[preKey]; ok {
			if _, ok := slice.([]any); ok {
				if len(slice.([]any)) > index {
					if v, ok := slice.([]any)[index].(map[string]any); ok {
						return parseMap(v, remainKeys)
					}
					return currentMap[preKey].([]any)[index]
				}
			}
		}
	} else {
		if v, ok := currentMap[k]; ok {
			if mapV, ok := v.(map[string]any); ok {
				return parseMap(mapV, remainKeys)
			}
			if _, ok := v.([]any); ok {
				return currentMap
			}
			return currentMap[k]
		}
		return currentMap
	}

	return currentMap
}

// SetNestedValue sets a nested value in a viper instance
func SetNestedValue(v *viper.Viper, key string, value any) error {
	keys := strings.Split(key, ".")
	data := v.AllSettings()
	nestedMap := data
	newMap := parse(nestedMap, keys, value)
	if err := v.MergeConfigMap(newMap); err != nil {
		return fmt.Errorf("failed to set nested value: %w", err)
	}
	return nil
}

//nolint:forcetypeassert
func parse(currentMap map[string]any, keys []string, value any) map[string]any {
	if len(keys) == 0 {
		return currentMap
	}
	k := keys[0]
	remainKeys := keys[1:]
	if strings.Contains(k, "[") {
		idx := strings.Index(k, "[")
		preKey := k[:idx]
		idx2 := strings.Index(k, "]")
		postKey := k[idx+1 : idx2]
		if postKey == "" {
			return currentMap
		}
		index, err := strconv.Atoi(postKey)
		if err != nil {
			return currentMap
		}
		if slice, ok := currentMap[preKey]; ok {
			if _, ok := slice.([]any); ok {
				if len(slice.([]any)) > index {
					if v, ok := slice.([]any)[index].(map[string]any); ok {
						slice.([]any)[index] = parse(v, remainKeys, value)
						currentMap[preKey] = slice
						return currentMap
					}
					currentMap[preKey].([]any)[index] = value
					return currentMap
				}
			}
		}
	} else {
		if v, ok := currentMap[k]; ok {
			if mapV, ok := v.(map[string]any); ok {
				currentMap[k] = parse(mapV, remainKeys, value)
				return currentMap
			}
			_, ok := v.([]any)
			if ok {
				return currentMap
			}
			currentMap[k] = value
			return currentMap
		}
		return currentMap
	}

	return currentMap
}
