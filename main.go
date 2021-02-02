package main

import (
	"github.com/gabrielpsz/dm-challenge/internal"
	"github.com/gabrielpsz/dm-challenge/repository"
)

func main() {
	repository.StartDatabase();
	go router.StartRouter()
	internal.StartQueue()
}

