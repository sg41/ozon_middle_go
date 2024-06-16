package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JSONValue interface{}

type JSONObject map[string]JSONValue
type JSONArray []JSONValue

func prettify(sourceVal JSONValue) JSONValue {
	switch val := sourceVal.(type) {
	case map[string]interface{}:
		newObj := make(JSONObject)
		for k, v := range val {
			prettified := prettify(v)
			if prettified != nil {
				newObj[k] = prettified
			}
		}
		if len(newObj) > 0 {
			return newObj
		}
	case []interface{}:
		newArr := make(JSONArray, 0, len(val))
		for _, v := range val {
			prettified := prettify(v)
			if prettified != nil {
				newArr = append(newArr, prettified)
			}
		}
		if len(newArr) > 0 {
			return newArr
		}
	case string:
		return val
	}
	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadString('\n')

		lines := make([]string, n)
		for j := 0; j < n; j++ {
			line, _ := in.ReadString('\n')
			lines[j] = strings.TrimSpace(line)
		}
		jsonString := strings.Join(lines, "")
		var jsonValue JSONValue
		json.Unmarshal([]byte(jsonString), &jsonValue)

		prettified := prettify(jsonValue)

		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.Encode(prettified)
		if i < t-1 {
			fmt.Fprintf(out, "%s,", buf.String()[0:len(buf.String())-1])
		} else {

			fmt.Fprintf(out, "%s", buf.String())
		}

	}
}
