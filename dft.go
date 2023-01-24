package main

import (
  "encoding/csv"
  "log"
  "os"
  "fmt"
  "math"
  //"math/cmplx"
  "strconv"
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
	yr := []float64{} //real part
	yi := []float64{} //imag part
	spt:= []string{} //sinpuku spectol
	spt_f:= []float64{}
	hz := []string{} //herutu
	hz_f := []float64{}
	li := len(y)
	lf := float64(li)
	ycos := 0.0
	ysin := 0.0
	
	for i:=0; i<li; i++ {
		if i% 1000 == 0 {
			fmt.Println(i)
		}
		ycos = 0.0
		ysin = 0.0
		for k:=0; k<li; k++ {
			omega := 2.0 * math.Pi * float64(i) * float64(k) / lf // kaitensi
			ycos += y[k] * math.Cos(omega)
			ysin += y[k] * math.Sin(omega)
		}
		yr = append(yr, ycos)
		yi = append(yi, ysin)
		hz_f = append(hz_f, float64(i) * 44100.0 / lf)
		spt_f= append(spt_f, math.Sqrt(ycos*ycos + ysin*ysin))
		hz = append(hz, strconv.FormatFloat(hz_f[i], 'f', -4, 64))
		spt = append(spt, strconv.FormatFloat(spt_f[i], 'f', -4, 64))
		
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
	
	for l:=0; l< len(out1); l++ {
		writer.Write([]string{out1[l], out2[l]})
	}
}

func main() {
	
	if len(os.Args) != 2 {
		fmt.Println("input > dft xxx.csv")
		os.Exit(0)
	}

	var open_file string = os.Args[1]
	var output_file string = "dft_" + open_file
	
	
	x := []float64{}
	
	x = read_csv(open_file)
	
	//fmt.Println(x)
	
	out_hz := []string{}
	out_spt:= []string{}
	
	out_hz , out_spt = dft(x)

	//fmt.Println(out_hz[88], out_spt[88])

	out_csv(output_file, out_hz, out_spt)

}
