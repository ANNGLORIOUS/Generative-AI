package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}<>?/|"
)

type config struct {
	length         int
	count          int
	includeLower   bool
	includeUpper   bool
	includeNumbers bool
	includeSymbols bool
}

func main() {
	cfg := parseFlags()

	passwords, err := generatePasswords(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	for _, password := range passwords {
		fmt.Println(password)
	}
}

func parseFlags() config {
	cfg := config{}

	flag.IntVar(&cfg.length, "length", 16, "length of each generated password")
	flag.IntVar(&cfg.count, "count", 1, "number of passwords to generate")
	flag.BoolVar(&cfg.includeLower, "lower", true, "include lowercase letters")
	flag.BoolVar(&cfg.includeUpper, "upper", true, "include uppercase letters")
	flag.BoolVar(&cfg.includeNumbers, "numbers", true, "include digits")
	flag.BoolVar(&cfg.includeSymbols, "symbols", true, "include symbols")
	flag.Parse()

	return cfg
}

func generatePasswords(cfg config) ([]string, error) {
	if cfg.length < 4 {
		return nil, fmt.Errorf("length must be at least 4")
	}

	if cfg.count < 1 {
		return nil, fmt.Errorf("count must be at least 1")
	}

	selectedSets := buildSelectedSets(cfg)
	if len(selectedSets) == 0 {
		return nil, fmt.Errorf("enable at least one character set")
	}

	if cfg.length < len(selectedSets) {
		return nil, fmt.Errorf("length must be at least %d when using %d character sets", len(selectedSets), len(selectedSets))
	}

	passwords := make([]string, 0, cfg.count)
	for i := 0; i < cfg.count; i++ {
		password, err := generatePassword(cfg.length, selectedSets)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	return passwords, nil
}

func buildSelectedSets(cfg config) []string {
	sets := make([]string, 0, 4)

	if cfg.includeLower {
		sets = append(sets, lowercase)
	}
	if cfg.includeUpper {
		sets = append(sets, uppercase)
	}
	if cfg.includeNumbers {
		sets = append(sets, numbers)
	}
	if cfg.includeSymbols {
		sets = append(sets, symbols)
	}

	return sets
}

func generatePassword(length int, selectedSets []string) (string, error) {
	allChars := strings.Join(selectedSets, "")
	passwordRunes := make([]byte, 0, length)

	// Guarantee at least one character from each enabled group.
	for _, charSet := range selectedSets {
		char, err := randomChar(charSet)
		if err != nil {
			return "", err
		}
		passwordRunes = append(passwordRunes, char)
	}

	for len(passwordRunes) < length {
		char, err := randomChar(allChars)
		if err != nil {
			return "", err
		}
		passwordRunes = append(passwordRunes, char)
	}

	if err := shuffle(passwordRunes); err != nil {
		return "", err
	}

	return string(passwordRunes), nil
}

func randomChar(charSet string) (byte, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
	if err != nil {
		return 0, fmt.Errorf("failed to generate secure random index: %w", err)
	}

	return charSet[index.Int64()], nil
}

func shuffle(values []byte) error {
	for i := len(values) - 1; i > 0; i-- {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return fmt.Errorf("failed to shuffle password characters: %w", err)
		}

		j := int(index.Int64())
		values[i], values[j] = values[j], values[i]
	}

	return nil
}
