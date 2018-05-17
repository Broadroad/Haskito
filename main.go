package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

)

type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	PrevHash string
}


var Blockchain []Block
