package moviegolib

import (
	"bufio"
	"os"
	"fmt"
	"path"
	"time"
	"strconv"
	"strings"
	"math/rand"
	"encoding/base64"
	_"image/jpeg"
	_"image/png"
	"github.com/disintegration/imaging"
	"gopkg.in/mgo.v2/bson"
)

const (
    PIC_PATH string = "/nfs/home/charlie//Pictures/"
)

func FS(fi string) string {
	fuf, err := os.Open(fi)
	if err != nil {
		fmt.Printf("\n this is file Open error FS %v \n", fi)
		return "0"
	}
	defer fuf.Close()
	stat, err := fuf.Stat()
	if err != nil {
		fmt.Printf("\n this is os Stat error %v \n", stat)
		return "0"
	}
	size := stat.Size()
	return strconv.FormatInt(size, 10)
}

/*func FileExists(filePath string)(exists bool){
	exists = true
	if _, err := os.Stat(filePath); os.IsNotExist(err){
		exists = false	
	}
	return
}*/

var UUid string

func UUID() (UUid string) {
	a_seed := time.Now()
	aseed := a_seed.UnixNano()
	rand.Seed(aseed)
	u := rand.Int63n(aseed)
	UUid = strconv.FormatInt(u, 10)
	return
}


func FileExistence(mpath string) (result bool) {
    if _, err := os.Stat(mpath); err == nil {
        return true
    } else {
        return false
    }
}

var NoArtList []string
func FindPicPaths(mpath string, noartpicpath string) (result string) {
	newl := len(mpath) - 4
	var noextname string = mpath[0:newl]
	var jpgname string = noextname + ".jpg"
	var Jpgpicturespath string
	Jpgpicturespath = PIC_PATH + path.Base(noextname) + ".jpg"
	fmt.Println("\n")
    fmt.Println(noextname)
    fmt.Println(jpgname)
    fmt.Println(Jpgpicturespath)
    fmt.Println("\n")
	

	var pngname string = noextname + ".png"
	var PngPicturesPath string
	PngPicturesPath = PIC_PATH + path.Base(noextname) + ".png"
	fmt.Println(FileExistence(jpgname))
	fmt.Println(FileExistence(Jpgpicturespath))
	fmt.Println(FileExistence(pngname))
	fmt.Println(FileExistence(PngPicturesPath))
	
	switch {
		case FileExistence(jpgname):
			result = jpgname
		case FileExistence(Jpgpicturespath):
			result = Jpgpicturespath
		case FileExistence(pngname):
			result = pngname
		case FileExistence(PngPicturesPath):
			result = PngPicturesPath
		default:
            NoArtList = append(NoArtList, mpath)
			result = noartpicpath
	}
	return
}

func CreateTVShowsThumbnail(p string) string {
	_, fnn := path.Split(p)
	nfn1 := "./static/images/thumbnails/" + fnn
//	pic, err := imaging.Open(p)
/*	if err != nil {
		fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
	}*/
    
	if strings.Contains(p, "TVShows") {
        pic, err := imaging.Open(p)
        if err != nil {
            fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
        }
	    
	    cjImage := imaging.Thumbnail(pic, 100, 100, imaging.Lanczos)
        err = imaging.Save(cjImage, nfn1)
        if err != nil {
		  fmt.Println(err)
	   }
    }
    return nfn1[1:]
}

func CreateMoviesThumbnail(p string) string {
	_, fnn := path.Split(p)
	nfn1 := "./static/images/thumbnails/" + fnn
//	pic, err := imaging.Open(p)
/*	if err != nil {
		fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
	}*/
   fi, err := os.Lstat(p)
   if err != nil {
    fmt.Println(err)
    }
    fmt.Println("this is fi.Name")
    fmt.Println(fi.Name())
    pic, err := imaging.Open(p)
	if err != nil {
        fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
    }
//	   cjImage := imaging.Thumbnail(pic, 300, 300, imaging.Lanczos)
    cjImage := imaging.Resize(pic, 0, 225, imaging.Lanczos)

    err = imaging.Save(cjImage, nfn1)
    if err != nil {
	   fmt.Println(err)
    }
    return nfn1[1:]
}
 

type pngS struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Dirpath string `bson:"dirpath"`
	Filename string `bson:"filename"`
	FileId string `bson:"fileid"`
	Extension string `bson:"extension"`
	LFilesize string `bson:"lfilesize"`
	SFilesize string `bson:"sfilesize"`
	LB64Image string `bson:"lb64image"`
	SB64Image string `bson:"sb64image"`
}

func CreatePngThumbnail(p string) pngS {
	dpp, fnn := path.Split(p)
	pngEXT := path.Ext(p)
	pic, err := imaging.Open(p)
	if err != nil {
		fmt.Printf("\n this is file Open error pngthumb %v \n", p)
	}
	spImage := imaging.Resize(pic, 80, 0, imaging.Lanczos)
	lpImage := imaging.Resize(pic, 200, 0, imaging.Lanczos)
	//THIS NEEDS TO BE CHANGED LATER TO A TEMP DIR
	uuid := UUID()
	pnfn1 := "./tmp/S" + uuid + fnn
	pnfn2 := "./tmp/L" + uuid + fnn
	

	err = imaging.Save(spImage, pnfn1)
	if err != nil {
		fmt.Println("There was a problem saving simage png")
	}
	err = imaging.Save(lpImage, pnfn2)
	if err != nil {
		fmt.Println("There was a problem saving limage png")
	}
	stn, err := os.Open(pnfn1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)			
	}
	defer stn.Close()
	stnInfo, _ := stn.Stat()
	var Ssize int64 = stnInfo.Size()
	sbuf := make([]byte, Ssize)
	stnReader := bufio.NewReader(stn)
	stnReader.Read(sbuf)
	sb64img := base64.StdEncoding.EncodeToString(sbuf)
	ltn, err := os.Open(pnfn2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)			
	}
	defer ltn.Close()
	ltnInfo, _ := ltn.Stat()
	var Lsize int64 = ltnInfo.Size()
	lbuf := make([]byte, Lsize)
	ltnReader := bufio.NewReader(ltn)
	ltnReader.Read(lbuf)
	lb64img := base64.StdEncoding.EncodeToString(lbuf)
	pngI := pngS{Id: bson.NewObjectId(),
		Dirpath: dpp,
		Filename: fnn,
		FileId: UUID(),
		Extension: pngEXT,
		LFilesize: strconv.FormatInt(Lsize, 10),
		SFilesize: strconv.FormatInt(Ssize, 10),
		LB64Image: lb64img,
		SB64Image: sb64img,
	}
	err = os.Remove(pnfn1)
	if err != nil {
		fmt.Println("there was a problem deleting tempfile")	
	}
	err = os.Remove(pnfn2)
	if err != nil {
		fmt.Println("there was a problem deleting tempfile2")	
	}
	return pngI
}