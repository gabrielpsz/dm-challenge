package main

import (
	"github.com/gabrielpsz/dm-challenge/repository"
)

func main() {
	repository.StartDatabase();
	internal.StartQueue()
}

