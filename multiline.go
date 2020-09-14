package multiline

import (
	"fmt"
	"sync"
)

type MultiLine struct {
	lines []*Line
}

func New() *MultiLine {
	return &MultiLine{}
}

func (m *MultiLine) GetLine(prefix string) *Line {
	if m.lines == nil {
		m.lines = []*Line{}
	}

	l := &Line{
		prefix:  prefix,
		lineNum: len(m.lines),
		in:      make(chan string),
	}

	m.lines = append(m.lines, l)
	return l
}

func (m *MultiLine) Print() error {
	for i := 0; i < len(m.lines); i++ {
		// make initial (lines count) * line
		fmt.Println(m.lines[i].prefix)
	}

	out := make(chan string)

	wg := sync.WaitGroup{}
	for _, line := range m.lines {
		wg.Add(1)
		go func(line *Line) {
			defer wg.Done()
			line.Read(out, len(m.lines))
		}(line)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for s := range out {
		fmt.Printf(s)
	}

	return nil
}

type Line struct {
	prefix  string
	lineNum int
	in      chan string
}

func (l *Line) Write(b []byte) (int, error) {
	l.in <- string(b)
	return len(b), nil
}

func (l *Line) WriteS(s string) error {
	_, err := l.Write([]byte(s))
	return err
}

func (l *Line) Close() {
	close(l.in)
}

func (l *Line) Read(out chan<- string, lineTotal int) error {
	for s := range l.in {
		move := lineTotal - l.lineNum
		clearline := "\x1b[K"
		up := fmt.Sprintf("\x1b[%dA\r", move)
		down := fmt.Sprintf("\x1b[%dB\r", move)
		out <- fmt.Sprint(up, clearline, l.prefix, s, down)
	}
	return nil
}
