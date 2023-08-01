// Code generated by github.com/geneva/gqlgen, DO NOT EDIT.

package selection

import (
	"time"
)

type Event interface {
	IsEvent()
	GetSelection() []string
	GetCollected() []string
}

type Like struct {
	Reaction  string    `json:"reaction"`
	Sent      time.Time `json:"sent"`
	Selection []string  `json:"selection,omitempty"`
	Collected []string  `json:"collected,omitempty"`
}

func (Like) IsEvent() {}
func (this Like) GetSelection() []string {
	if this.Selection == nil {
		return nil
	}
	interfaceSlice := make([]string, 0, len(this.Selection))
	for _, concrete := range this.Selection {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
func (this Like) GetCollected() []string {
	if this.Collected == nil {
		return nil
	}
	interfaceSlice := make([]string, 0, len(this.Collected))
	for _, concrete := range this.Collected {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type Post struct {
	Message   string    `json:"message"`
	Sent      time.Time `json:"sent"`
	Selection []string  `json:"selection,omitempty"`
	Collected []string  `json:"collected,omitempty"`
}

func (Post) IsEvent() {}
func (this Post) GetSelection() []string {
	if this.Selection == nil {
		return nil
	}
	interfaceSlice := make([]string, 0, len(this.Selection))
	for _, concrete := range this.Selection {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
func (this Post) GetCollected() []string {
	if this.Collected == nil {
		return nil
	}
	interfaceSlice := make([]string, 0, len(this.Collected))
	for _, concrete := range this.Collected {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
