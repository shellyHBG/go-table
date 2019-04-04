package table

import (
	"errors"
	"strings"
)

func Parse(name, data string) (result Table, err error) {
	//fmt.Println("[table] " + name + " starts to parse.")

	// initialize
	result = Table{name, make(map[string]map[string]string)}
	err = nil

	// string split
	// https://golang.org/pkg/strings/
	// func Split(s, sep string) []string
	// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
	// If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
	// If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.
	lines := strings.Split(data, "\n")
	if len(lines) == 0 {
		err = errors.New("no lines")
		return result, err
	}
	//fmt.Printf("%q\n", lines)

	// types
	types := strings.Split(lines[0], "\t")
	types = types[:len(types)-1]
	//fmt.Printf("type: %q\n", types)

	// key
	keys := strings.Split(lines[1], "\t")
	keys = keys[:len(keys)-1]
	//fmt.Printf("key: %q\n", keys)

	for _, line := range lines[2 : len(lines)-1] {
		items := strings.Split(line, "\t")
		items = items[:len(items)-1]
		if len(items) == 0 {
			err = errors.New("no items")
			return result, err
		}
		//fmt.Printf("%q\n", items)

		key := items[0]
		result.content[key] = make(map[string]string)
		for index, subKey := range keys[1:] {
			result.content[key][subKey] = items[index+1]
		}

		//fmt.Printf("%v > %q\n", key, result.content[key])
	}

	return result, err
}
