package main

import (
	"time"
)

type VirtualFileSystem struct {
	root *Katalog
}

func NewVirtualFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{
		root: &Katalog{
			name:       "root",
			items:      []FileSystemItem{},
			createdAt:  time.Now(),
			modifiedAt: time.Now(),
		},
	}
}

func (vfs *VirtualFileSystem) CreateFile(name string, content []byte) error {
	file := &Plik{
		name:       name,
		content:    content,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return vfs.root.AddItem(file)
}

func (vfs *VirtualFileSystem) CreateReadFile(name string, content []byte) error {
	file := &PlikDoOdczytu{
		name:       name,
		content:    content,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return vfs.root.AddItem(file)
}

func (vfs *VirtualFileSystem) CreateDirectory(name string) error {
	directory := &Katalog{
		name:       name,
		items:      []FileSystemItem{},
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return vfs.root.AddItem(directory)
}

func (vfs *VirtualFileSystem) FindItem(name string) (FileSystemItem, error) {
	for _, item := range vfs.root.Items() {
		if item.Name() == name {
			return item, nil
		}
	}
	return nil, ErrItemNotFound
}

func (vfs *VirtualFileSystem) ReadFile(name string) ([]byte, error) {
	item, err := vfs.FindItem(name)
	if err != nil {
		return nil, err
	}

	readable, ok := item.(Readable)
	if !ok {
		return nil, ErrPermissionDenied
	}

	buffer := make([]byte, item.Size())
	_, err = readable.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil

}

func (vfs *VirtualFileSystem) WriteFile(name string, content []byte) error {
	item, err := vfs.FindItem(name)
	if err != nil {
		return err
	}

	writable, ok := item.(Writable)
	if !ok {
		return ErrPermissionDenied
	}

	_, err = writable.Write(content)
	return err
}

func (vfs *VirtualFileSystem) DeleteItem(name string) error {
	return vfs.root.RemoveItem(name)
}
