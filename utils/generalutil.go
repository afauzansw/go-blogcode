package utils

import (
	"log"
	"regexp"
	"strings"
	"time"
)

func GenerateSlug(input string) string {
	slug := strings.ToLower(input)

	slug = strings.ReplaceAll(slug, " ", "-")

	reg := regexp.MustCompile("[^a-z0-9-]+")
	slug = reg.ReplaceAllString(slug, "")

	slug = strings.Trim(slug, "-")

	return slug
}

func DateFormat(input time.Time) string {
	zone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}

	localTime := input.In(zone)

	formatedTime := localTime.Format("01 January 2024 15:04:05")

	return formatedTime
}
