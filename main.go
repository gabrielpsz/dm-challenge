package main

import (
	"github.com/gabrielpsz/dm-challenge/internal"
	"github.com/gabrielpsz/dm-challenge/repository"
	"github.com/gabrielpsz/dm-challenge/router"
)

func main() {
	repository.StartDatabase();
	go router.StartRouter()
	internal.StartQueue()
}

