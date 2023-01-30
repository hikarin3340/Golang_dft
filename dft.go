package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"math"
	"strconv"
	"time"
)

func read_csv(read_file string ) []float64 {
	file, err := os.Open(read_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	
	readdata := []float64{}
	for _, v := range rows {
		s, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		readdata = append(readdata, s)
	}
	return readdata
	
}

func dft(y []float64) ([]string , []string ) {
	hz := []string{} //hertx string
	spt:= []string{} //amplitude string
	ycos, ysin := 0.0, 0.0
	dft_now := time.Now()
	
	for l:=0; l<len(y); l++ {
		ycos, ysin = 0.0, 0.0
		if l % 10000 == 0 {
			fmt.Printf("%d周目 %fs\n", l,  time.Since(dft_now).Seconds())
			dft_now = time.Now()
		}
		
		for k:=0; k<len(y); k++ {
			tf := 2.0 * math.Pi * float64(l) * float64(k) / float64(len(y)) // twiddle factor
			ycos += y[k] * math.Cos(tf) // real part
			ysin += y[k] * math.Sin(tf) // imag part
		}
		hz_f := float64(l) * 44100.0 / float64(len(y)) // hz float
		spt_f:= math.Sqrt(ycos*ycos + ysin*ysin) // spt float
		
		hz = append(hz, strconv.FormatFloat(hz_f, 'f', -4, 64)) // float to string hz
		spt = append(spt, strconv.FormatFloat(spt_f, 'f', -4, 64)) // float to string spt
	}
	return hz, spt
}

func out_csv (out_file string , out1 []string , out2 []string ) {
	file2 , err := os.Create(out_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	
	writer := csv.NewWriter(file2)
	defer writer.Flush()
	
	for i:=0; i< len(out1); i++ {
		writer.Write([]string{out1[i], out2[i]})
	}
}

func main() {
	now := time.Now()
	
	if len(os.Args) != 2 {
		fmt.Println("input > dft xxx.csv")
		os.Exit(0)
	}
	open_file := os.Args[1]
	output_file := "dft_" + open_file

	fmt.Println("計測開始")

	x := read_csv(open_file)
	
	fmt.Printf("csv読み込み完了 %fs\n", time.Since(now).Seconds())
	
	out_hz , out_spt := dft(x)
	
	fmt.Printf("dft完了 %fs\n", time.Since(now).Seconds())
	
	out_csv(output_file, out_hz, out_spt)
	
	fmt.Printf("csv出力完了 %fs\n", time.Since(now).Seconds())
}
