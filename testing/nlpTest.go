package main

import (
	"log"
  "os"
	//"github.com/go-ego/gse"
)
var (
  //seg gse.Segmenter
  targetText = "喺時候瞓覺"
  targetTextByte = []byte(targetText)
)

func nlp_test(){
  f, err := os.OpenFile("test_log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
      log.Fatalf("error opening file: %v", err)
  }
  defer f.Close()

  log.SetOutput(f)
  log.Println("-------------------------This is a test log entry-------------------------------")

  seg.LoadDict()

  a := seg.Segment(targetTextByte)
  log.Print(seg.String(targetTextByte))
  //for {
  for i := 0; i < len(a); i++ {
		log.Print(a[i].Token().Pos())
	}
  
  //}
  //
}
