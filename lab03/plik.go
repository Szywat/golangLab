package main

import (
	"time"
)

type Plik struct {
	name       string
	content    []byte
	createdAt  time.Time
	modifiedAt time.Time
}

func (p *Plik) Name() string          { return p.name }
func (p *Plik) Path() string          { return "/" + p.name }
func (p *Plik) Size() int64           { return int64(len(p.content)) }
func (p *Plik) CreatedAt() time.Time  { return p.createdAt }
func (p *Plik) ModifiedAt() time.Time { return p.modifiedAt }

func (p *Plik) Read(b []byte) (int, error) {
	n := copy(b, p.content)
	return n, nil
}

func (p *Plik) Write(b []byte) (int, error) {
	p.content = b
	p.modifiedAt = time.Now()
	return len(b), nil
}