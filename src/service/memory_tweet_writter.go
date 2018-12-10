package service

import "github.com/tweet/src/domain"

type MemoryTweetWriter struct {
	tweets []domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter{
 	memoryTweetWriter:=	MemoryTweetWriter{make([]domain.Tweet,0)}
 	return &memoryTweetWriter
}

func (memoryTweetWriter *MemoryTweetWriter)Write(tweet domain.Tweet){
	memoryTweetWriter.tweets = append(memoryTweetWriter.tweets,tweet)
}

func (memoryTweetWriter *MemoryTweetWriter)GetLastSavedTweet()domain.Tweet  {
	size := len(memoryTweetWriter.tweets)
	if size == 0 {
		return nil
	}
	return memoryTweetWriter.tweets[ size - 1 ]
}