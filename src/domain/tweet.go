package domain

import "time"

type Tweet interface {
	PrintableTweet() string
	String() string
	GetUser() string
	GetText() string
	GetId() int
	GetDate() *time.Time
	SetId(int)
	UserIsEmpty() bool
	TextIsEmpty() bool
	TextHasMoreCharactersThanMaxCharactersPerTweet() bool
}
