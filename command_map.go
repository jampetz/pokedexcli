package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func getIDFromURL(url string) (int, error) {
	segments := strings.Split(url, "/")
	idStr := segments[len(segments)-2]

	targetID, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("error converting ID to integer: %v", err)
	}

	return targetID, nil
}

func mapForward(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Available locations:")
	for _, location := range resp.Results {
		id, err := getIDFromURL(location.URL)
		if err != nil {
			return err
		}

		fmt.Printf("%v) %v\n", id, location.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaULR = resp.Previous
	return nil
}

func mapBackward(cfg *config, args ...string) error {
	if cfg.prevLocationAreaULR == nil {
		return errors.New("you're on first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaULR)
	if err != nil {
		return err
	}

	fmt.Println("Available locations:")
	for _, location := range resp.Results {
		id, err := getIDFromURL(location.URL)
		if err != nil {
			return err
		}

		fmt.Printf("%v) %v\n", id, location.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaULR = resp.Previous
	return nil
}
