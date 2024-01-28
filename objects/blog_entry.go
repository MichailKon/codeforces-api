package objects

import (
	"github.com/rkennedy/optional"
)

type BlogEntry struct {
	Id                      int                    `json:"id"`
	OriginalLocale          string                 `json:"originalLocale"`
	CreationTimeSeconds     int                    `json:"creationTimeSeconds"`
	AuthorHandle            string                 `json:"authorHandle"`
	Title                   string                 `json:"title"`
	Content                 optional.Value[string] `json:"content"`
	Locale                  string                 `json:"locale"`
	ModificationTimeSeconds int                    `json:"modificationTimeSeconds"`
	AllowViewHistory        bool                   `json:"allowViewHistory"`
	Tags                    []string               `json:"tags"`
	Rating                  int                    `json:"rating"`
}
