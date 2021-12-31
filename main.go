package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

//input : size of piece in KB, number of pieces, begin-end
//1000 * 256KB = 25600 KB = 256 MB
var sizeOfPiece = int64(256 * 1024)
var numberOfPieces = int64(1000)
var sizeOfFile = numberOfPieces * sizeOfPiece

func makeFile(fileName string, numPieces int64, offset int64) {

	from := strconv.FormatInt(offset, 10)
	end := strconv.FormatInt(offset+numPieces-1, 10)
	offset -= 1
	path := "peer_" + from + "_to_" + end

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Create(path + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Truncate(sizeOfFile); err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteAt(bytes.Repeat([]byte("a"), int(numPieces*sizeOfPiece)), offset*sizeOfPiece)

	_, err = f.WriteAt(bytes.Repeat([]byte("b"), int(offset*sizeOfPiece)), 0)
	_, err = f.WriteAt(bytes.Repeat([]byte("b"), int((numberOfPieces-(offset+numPieces))*sizeOfPiece)), (offset+numPieces)*sizeOfPiece)
	if err != nil {
		log.Fatalf("failed to write")
	}
}

func main() {
	fileName := os.Args[1]
	from, _ := strconv.Atoi(os.Args[2])
	to, _ := strconv.Atoi(os.Args[3])
	to = int(math.Min(float64(numberOfPieces), float64(to)))
	from = int(math.Max(1, float64(from)))
	numberOfPiecesToWrite := to - from + 1
	offset := from
	fmt.Println("fileName : " + fileName + " nofPieces : " + strconv.FormatInt(int64(numberOfPiecesToWrite), 10) + " offset: " + strconv.FormatInt(int64(offset), 10))
	makeFile(fileName, int64(numberOfPiecesToWrite), int64(offset))

}
