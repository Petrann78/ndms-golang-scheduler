package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ReplacePlaceholders takes a JSON template string and replaces
// macros like {{NOW}}, {{RANDOM_FLOAT min max}}, {{RANDOM_INT min max}}, {{CHOICE ...}}
func ReplacePlaceholders(tmpl string) string {
	now := time.Now().UTC().Format(time.RFC3339)

	// {{NOW}}
	tmpl = strings.ReplaceAll(tmpl, "{{NOW}}", now)

	// {{RANDOM_FLOAT min max}}
	reFloat := regexp.MustCompile(`\{\{RANDOM_FLOAT ([\d.]+) ([\d.]+)\}\}`)
	tmpl = reFloat.ReplaceAllStringFunc(tmpl, func(s string) string {
		matches := reFloat.FindStringSubmatch(s)
		min, _ := strconv.ParseFloat(matches[1], 64)
		max, _ := strconv.ParseFloat(matches[2], 64)
		val := min + rand.Float64()*(max-min)
		return fmt.Sprintf("%.2f", val)
	})

	// {{RANDOM_INT min max}}
	reInt := regexp.MustCompile(`\{\{RANDOM_INT (\d+) (\d+)\}\}`)
	tmpl = reInt.ReplaceAllStringFunc(tmpl, func(s string) string {
		matches := reInt.FindStringSubmatch(s)
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		val := rand.Intn(max-min+1) + min
		return strconv.Itoa(val)
	})

	// {{CHOICE A,B,C}}
	reChoice := regexp.MustCompile(`\{\{CHOICE ([^}]+)\}\}`)
	tmpl = reChoice.ReplaceAllStringFunc(tmpl, func(s string) string {
		matches := reChoice.FindStringSubmatch(s)
		options := strings.Split(matches[1], ",")
		return strings.TrimSpace(options[rand.Intn(len(options))])
	})

	return tmpl
}
