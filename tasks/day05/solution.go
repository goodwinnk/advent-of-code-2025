package day05

import (
	"AdventOfCode2025/internal/util"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = 5

func input() string {
	return util.MustReadInput(day, "task.txt")
}

func Part1() int {
	return Part1Text(input())
}

func Part2() int64 {
	return Part2Text(input())
}

type EventType int

const (
	IntervalStart EventType = iota
	ID
	IntervalEnd
)

type Event struct {
	Type  EventType
	Id    int64
	Range Range
}

func Part1Text(input string) int {
	ranges, ids, err := parse(input)
	if err != nil {
		panic(err)
	}

	events := make([]Event, 0, len(ranges)*2+len(ids))
	for _, r := range ranges {
		events = append(events, Event{Type: IntervalStart, Id: r.Start, Range: r})
		events = append(events, Event{Type: IntervalEnd, Id: r.End, Range: r})
	}
	for _, id := range ids {
		events = append(events, Event{Type: ID, Id: id})
	}

	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Id < events[j].Id ||
			(events[i].Id == events[j].Id && events[i].Type < events[j].Type)
	})

	intervals := 0
	fresh := 0
	for _, e := range events {
		switch e.Type {
		case IntervalStart:
			{
				intervals++
			}
		case IntervalEnd:
			{
				intervals--
				if intervals < 0 {
					panic("negative intervals")
				}
			}
		case ID:
			{
				if intervals != 0 {
					fresh++
				}
			}
		}
	}

	return fresh
}

func Part2Text(input string) int64 {
	ranges, _, err := parse(input)
	if err != nil {
		panic(err)
	}

	events := make([]Event, 0, len(ranges)*2)
	for _, r := range ranges {
		events = append(events, Event{Type: IntervalStart, Id: r.Start, Range: r})
		events = append(events, Event{Type: IntervalEnd, Id: r.End, Range: r})
	}

	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Id < events[j].Id ||
			(events[i].Id == events[j].Id && events[i].Type < events[j].Type)
	})

	all := int64(0)
	lastStart := int64(-1)
	intervals := 0
	for _, e := range events {
		if e.Id < 0 {
			panic("negative ID")
		}
		switch e.Type {
		case IntervalStart:
			{
				intervals++
				if intervals == 1 && lastStart < 0 {
					lastStart = e.Id
				}
			}
		case IntervalEnd:
			{
				intervals--
				if intervals == 0 {
					if lastStart < 0 {
						panic("no start interval")
					}
					all += 1 + e.Id - lastStart
					lastStart = -1
				}

				if intervals < 0 {
					panic("negative intervals")
				}
			}
		case ID:
			panic("ID event")
		}
	}

	return all
}

type Range struct {
	Start int64 // Exported fields (capitalized) are more idiomatic
	End   int64
}

func parse(input string) ([]Range, []int64, error) {
	var ranges []Range
	var ids []int64

	scanner := bufio.NewScanner(strings.NewReader(input))
	parsingIds := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if len(ranges) > 0 && !parsingIds {
				parsingIds = true
			}
			continue
		}

		if !parsingIds {
			r, err := parseRange(line)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid range %q: %w", line, err)
			}
			ranges = append(ranges, r)
		} else {
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid ID %q: %w", line, err)
			}
			ids = append(ids, id)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scanning input: %w", err)
	}

	return ranges, ids, nil
}

func parseRange(line string) (Range, error) {
	parts := strings.SplitN(line, "-", 2)
	if len(parts) != 2 {
		return Range{}, fmt.Errorf("expected format 'start-end'")
	}

	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return Range{}, fmt.Errorf("invalid start: %w", err)
	}

	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return Range{}, fmt.Errorf("invalid end: %w", err)
	}

	return Range{Start: start, End: end}, nil
}
