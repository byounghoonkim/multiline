package multiline

import (
	"fmt"
	"sync"
)

type MultiLine struct {
	lines []*Line
	out   chan string
}

func New() *MultiLine {
	return &MultiLine{}
}

func (m *MultiLine) GetLine(prefix string) *Line {
	if m.lines == nil {
		m.lines = []*Line{}
		m.out = make(chan string)
	}

	l := &Line{
		prefix:  prefix,
		lineNum: len(m.lines),
		in:      make(chan string),
		out:     m.out,
	}

	m.lines = append(m.lines, l)
	return l
}

func (m *MultiLine) Print() error {
	for i := 0; i < len(m.lines); i++ {
		// make blank (lines count) * line
		fmt.Println("")
	}
	fmt.Printf("\x1b[%dA\r", len(m.lines)+1)

	wg := sync.WaitGroup{}
	for _, l := range m.lines {
		wg.Add(1)
		go func(myl *Line) {
			defer wg.Done()
			myl.Read()
		}(l)
	}

	go func() {
		wg.Wait()
		close(m.out)
	}()

	for s := range m.out {
		fmt.Printf(s)
	}

	fmt.Printf("\x1b[%dB", len(m.lines)+1)
	return nil
}

type Line struct {
	prefix  string
	lineNum int
	in      chan string
	out     chan string
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

func (l *Line) Read() error {
	for s := range l.in {
		clearline := "\x1b[K"
		up := fmt.Sprintf("\x1b[%dA\r", l.lineNum+1)
		down := fmt.Sprintf("\x1b[%dB", l.lineNum+1)
		l.out <- fmt.Sprint(down, clearline, l.prefix, s, up)
	}
	return nil
}
