package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var results []interface{}

	for i := 0; i < n; i++ {
		var count int
		fmt.Fscan(in, &count)

		var jsonStrBuffer bytes.Buffer

		for j := 0; j <= count; j++ {
			line, err := in.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка при чтении строки:", err)
				return
			}
			jsonStrBuffer.WriteString(strings.TrimSpace(line))
		}

		jsonStr := jsonStrBuffer.String()

		var result interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			fmt.Println("Ошибка при декодировании JSON:", err)
			return
		}

		prettifiedResult := prettify(result)

		results = append(results, prettifiedResult)
	}

	finalResult, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Ошибка при кодировании JSON:", err)
		return
	}

	fmt.Fprintln(out, string(finalResult))
}

func prettify(v interface{}) interface{} {
	switch v := v.(type) {
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, value := range v {
			processedValue := prettify(value)
			if processedValue != nil {
				newMap[key] = processedValue
			}
		}
		if len(newMap) == 0 {
			return nil
		}
		return newMap
	case []interface{}:
		newSlice := make([]interface{}, 0, len(v))
		for _, item := range v {
			processedItem := prettify(item)
			if processedItem != nil {
				newSlice = append(newSlice, processedItem)
			}
		}
		if len(newSlice) == 0 {
			return nil
		}
		return newSlice
	default:
		return v
	}
}
