package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("emojimaker: missing file operand")
		fmt.Println("Try 'emojimaker --help' for more information.")
		return
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("usage: emojimaker FILE")
		fmt.Println("Searches a file for emoji matches and replace with the equivalent emoji.")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("emojimaker: cannot", err)
		return
	}
	strdata := string(data)

	m := make(map[string]string)
	m["airplane"] = "✈️"
	m["bicycle"] = "🚲"
	m["cat"] = "🐈"
	m["check"] = "✅"
	m["crossmark"] = "❌"
	m["locked"] = "🔐"
	m["mindblown"] = "🤯"
	m["thumbsdown"] = "👎"
	m["thumbsup"] = "👍"
	m["unlocked"] = "🔓"

	r := regexp.MustCompile(`:\w+:`)
	strdata = r.ReplaceAllStringFunc(strdata, func(match string) string {
		key := match[1 : len(match)-1]
		if emoji, ok := m[key]; ok {
			return emoji
		}
		return match
	})

	err = os.WriteFile(os.Args[1], []byte(strdata), 0644)
	if err != nil {
		fmt.Println("emojimaker: cannot", err)
	}
}
