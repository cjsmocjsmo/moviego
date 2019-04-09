 package main

import (
	"os"
	"fmt"
//	"path"
	"time"
	"strings"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"moviegosetup/libsetup"
)

const (
	PATH_TO_MOVIES string = "/nfs/home/charlie/Videos/"
	USB_PATH string = "/nfs/usb/"
	NO_ART_PIC_PATH string = "/home/pi/MovieTime/static/images/animals.jpg"
	THUMB_BASE_PATH string = "./static/images/thumbnails/"
)

func DBcon() *mgo.Session {
	s, err := mgo.Dial("127.0.0.1")
//	s, err := mgo.Dial("mongodb://192.168.1.101:27017")
	if err != nil {
		fmt.Println("this is dial err")
		panic(err)
	}
	return s
}

func IsEmptyDir(name string) (bool, error) {
	entries, err :=  ioutil.ReadDir(name)
	if err != nil {
		return false, err	
	}
	return len(entries) == 0, nil
}

func Process_Movs(pAth string) {
	var Mov_picpath string = moviegolib.FindPicPaths(pAth, NO_ART_PIC_PATH)
	var Mov_picinfo string = moviegolib.CreateMoviesThumbnail(Mov_picpath)
	var MovI moviegolib.MOVI
	MovI = moviegolib.Get_movie_info(pAth, Mov_picpath, Mov_picinfo)

	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	err := MTc.Insert(&MovI)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Process_usb_movs(pAth string) {
    var Mov_picpath string = moviegolib.FindPicPaths(pAth, NO_ART_PIC_PATH)
	var Mov_picinfo string = moviegolib.CreateMoviesThumbnail(Mov_picpath)
	var MovI moviegolib.MOVI
	MovI = moviegolib.Get_movie_info(pAth, Mov_picpath, Mov_picinfo)

	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	err := MTc.Insert(&MovI)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Process_tvshow_info(pAth string) {
    var tvshow_picpath string = moviegolib.FindPicPaths(pAth, NO_ART_PIC_PATH)
	var tvshow_picinfo string = moviegolib.CreateTVShowsThumbnail(tvshow_picpath)
	var CJ_TV_info moviegolib.TvShowInfoS
	CJ_TV_info = moviegolib.Get_tvshow_info(pAth, tvshow_picpath, tvshow_picinfo)
	ses := DBcon()
	defer ses.Close()	
	TVSc := ses.DB("moviego").C("moviego")
	err := TVSc.Insert(&CJ_TV_info)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ProcessPath(pAth string) {
    switch {
        case strings.Contains(pAth, "TVShows"):
            Process_tvshow_info(pAth)
        case strings.Contains(pAth, "Movies"):
            Process_Movs(pAth)
    }
}

func my_dir_visit(pAth string, f os.FileInfo, err error) error {
    ext := filepath.Ext(pAth)
    if f.IsDir() {
        fmt.Println("its a dir")
    } else {
        switch {
            case ext == ".mp4":
                ProcessPath(pAth)
            case ext == ".mkv":
                ProcessPath(pAth)
            case ext == ".avi":
                 ProcessPath(pAth)
            case ext == ".m4v":
                ProcessPath(pAth)
        }
    }
    fmt.Println(pAth)
	return nil
}

func my_usb_dir_visit(pAth string, f os.FileInfo, err error) error {
    if !f.IsDir() {
        if strings.Contains(pAth, "usb") {
            Process_usb_movs(pAth)
        }
    }
//    fmt.Println(pAth)
    return nil
}

func main() {
	start_time := time.Now().Unix()	
	sess := DBcon()
	defer sess.Close()
	err := sess.DB("moviego").DropDatabase()
	if err != nil {
		panic(err)	
	}
	fmt.Println("Database has been droped")
	
    if _, err := os.Stat("./static/images/thumbnails" ); err != nil {
        if os.IsNotExist(err) {
            os.MkdirAll("./static/images/thumbnails", 0755)
        }
    }
	
	empty, err := IsEmptyDir("./static/images/thumbnails")
	if err != nil {
		fmt.Println(err)	
	}
	if empty {
		fmt.Println("Thumbnails dir is empty")	
	} else {
		jpgthumbs, err := filepath.Glob("*.jpg")
		if err !=  nil{
			fmt.Println(err)		
		}
		for _, j := range jpgthumbs {
			os.Remove(j)
		}
		pngthumbs, err := filepath.Glob("*.png")
		if err != nil {
			fmt.Println(err)		
		}
		for _, p := range pngthumbs {
			os.Remove(p)
		}
	}
	
	filepath.Walk(PATH_TO_MOVIES, my_dir_visit)
	
	filepath.Walk(USB_PATH, my_usb_dir_visit)

	fmt.Println("Mov wait group complete")
    for _, noart := range moviegolib.NoArtList {
		fmt.Printf("this is noartlist %s \n\n", noart)
	}
	fmt.Println(start_time)
	stop_time := time.Now().Unix()
	fmt.Println(stop_time)
	etime := stop_time - start_time
	fmt.Println(etime)
}