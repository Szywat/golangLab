package main

import "time"

type PlikDoOdczytu struct {
	name       string
	content    []byte
	createdAt  time.Time
	modifiedAt time.Time
}

func (pd *PlikDoOdczytu) Name() string {return pd.name}
func (pd *PlikDoOdczytu) Path() string {return "/" + pd.name}
func (pd *PlikDoOdczytu) Size() int64 {return int64(len(pd.content))}
func (pd *PlikDoOdczytu) CreatedAt() time.Time {return pd.createdAt}
func (pd *PlikDoOdczytu) ModifiedAt() time.Time {return pd.modifiedAt}

func (pd *PlikDoOdczytu) Read(b []byte) (int, error) {
	n := copy(b, pd.content)
	return n, nil
}