package setupgrid

import (
	"encoding/csv"
	"image"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
)

var totalsFile string = "../keylogger/key_totals.csv"

// GetImage appends all
func GetImage() {
	var points []heatmap.DataPoint
	file, err := os.OpenFile(totalsFile, os.O_RDWR, 0644)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal("Unable to open csv file: ", err)
		} else {
			return
		}
	}

	csvReader := csv.NewReader(file)

	for {
		val, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Problem reading csv file: ", err)
		}
		num, err := strconv.ParseUint(val[1], 10, 64)
		if err != nil {
			log.Fatal("Problem parsing number in csv file: ", err)
		}
		for j := uint64(0); j < num; j++ {
			center := centers[val[0]]
			if center != nil {
				// fmt.Printf("Center: %f, %f\n", center.X(), center.Y())
				points = append(points, center)
			}
		}
	}
	file.Close()

	scheme := schemes.Classic
	img := heatmap.Heatmap(image.Rect(0, 0, 1765, 395), points, 150, 64, scheme, false)
	// imgFile, err := os.Create("heatmap.png")
	// if err != nil {
	// 	log.Fatal("Error creating image: ", err)
	// }
	// if err := png.Encode(imgFile, img); err != nil {
	if err := imaging.Save(img, "heatmap.png"); err != nil {
		// imgFile.Close()
		log.Fatal("Error encoding image: ", err)
	}
	// imgFile.Close()
}
