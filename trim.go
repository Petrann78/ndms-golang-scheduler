package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)


// call after ReplacePlaceholders
var reQuotedFloat = regexp.MustCompile(`"\s*\{\{\s*RANDOM_FLOAT\s+([+-]?\d+(?:\.\d+)?)\s+([+-]?\d+(?:\.\d+)?)\s*\}\}\s*"`)
var reQuotedInt   = regexp.MustCompile(`"\s*\{\{\s*RANDOM_INT\s+(-?\d+)\s+(-?\d+)\s*\}\}\s*"`)

func UnquoteNumericPlaceholders(s string) string {
    // FLOAT: replace the whole quoted placeholder with numeric literal
    s = reQuotedFloat.ReplaceAllStringFunc(s, func(m string) string {
        parts := reQuotedFloat.FindStringSubmatch(m)
        min, _ := strconv.ParseFloat(parts[1], 64)
        max, _ := strconv.ParseFloat(parts[2], 64)
        if max < min { min, max = max, min }
        val := min + rand.Float64()*(max-min)
        return fmt.Sprintf("%.6f", val)
    })

    // INT:
    s = reQuotedInt.ReplaceAllStringFunc(s, func(m string) string {
        parts := reQuotedInt.FindStringSubmatch(m)
        min, _ := strconv.Atoi(parts[1])
        max, _ := strconv.Atoi(parts[2])
        if max < min { min, max = max, min }
        val := rand.Intn(max-min+1) + min
        return strconv.Itoa(val)
    })

    return s
}
