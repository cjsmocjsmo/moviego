 package main

import (
	"os"
	"fmt"
	"path"
	"time"
	"strings"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

/*func get_movie_year(apath string) (movyr string) {
	_, filename := path.Split(apath)
	fi := strings.Index(filename, "(")
	fdex := fi + 1
	ldex := strings.LastIndex(filename, ")")
	movyr = filename[fdex:ldex]
	return
}*/

/*func Get_mov_name(fname string) string {
    var mov_name string = ""
    if strings.Contains(fname, "(") {
	   fi := strings.Index(fname, "(")
	   fdex := fi - 1
	   mov_name = fname[:fdex]
    } else {
        ddex := len(fname) - 11
        mov_name = fname[ddex:]
        fmt.Println("this is mov_name")
        fmt.Println(mov_name)
    }
	return mov_name
}*/

/*func get_gen(apath string) (CAT string) {
	switch {
		case strings.Contains(apath, "usb"):
			CAT = "usb"
		case strings.Contains(apath, "Movies"):
			CAT = "Movies"
		case strings.Contains(apath, "TVShows"):
			CAT = "TVShows"	
	}
	return
}*/

/*func get_movie_cat(apath string) (Mov_Cat string) {
	switch {
		case strings.Contains(apath, "SciFi"):	
			Mov_Cat = "SciFi"
		case strings.Contains(apath, "Cartoons"):
			Mov_Cat = "Cartoons"
		case strings.Contains(apath, "Godzilla"):
			Mov_Cat = "Godzilla"
		case strings.Contains(apath, "Kingsman"):
			Mov_Cat = "Kingsman"	
		case strings.Contains(apath, "StarTrek") && !strings.Contains(apath, " STTV "):
			Mov_Cat = "StarTrek"
		case strings.Contains(apath, "StarWars"):
			Mov_Cat = "StarWars"
		case strings.Contains(apath, "SuperHeros"):
			Mov_Cat = "SuperHeros"
		case strings.Contains(apath, "IndianaJones"):
			Mov_Cat = "IndianaJones"
		case strings.Contains(apath, "Action"):
			Mov_Cat = "Action"
		case strings.Contains(apath, "Comedy"):
			Mov_Cat = "Comedy"
		case strings.Contains(apath, "Drama"):
			Mov_Cat = "Drama"
        case strings.Contains(apath, "Jurassic Park"):
			Mov_Cat = "Jurassic Park"
		case strings.Contains(apath, "John Wayne"):		
			Mov_Cat = "John Wayne"
		case strings.Contains(apath, "John Wick"):
			Mov_Cat = "John Wick"
        case strings.Contains(apath, "Men In Black"):
			Mov_Cat = "Men In Black"
		case strings.Contains(apath, "Harry Potter"):
			Mov_Cat = "Harry Potter"
        case strings.Contains(apath, "Tremors"):
			Mov_Cat = "Tremors"
		case strings.Contains(apath, "Misc"):
			Mov_Cat = "Misc"
		default:
			Mov_Cat = "Misc"
	}
	return
}*/
		
/*
func get_usb_cat(apath string) (Usb_Cat string) {
	switch {
		case strings.Contains(apath, "SciFi"):
			Usb_Cat = "SciFi"
		case strings.Contains(apath, "Drama"):
			Usb_Cat = "Drama"
		case strings.Contains(apath, "Action"):
			Usb_Cat = "Action"
		case strings.Contains(apath, "Comedy"):
			Usb_Cat = "Comedy"
		case strings.Contains(apath, "Men In Black"):
			Usb_Cat = "Men In Black"
		case strings.Contains(apath, "Harry Potter"):
			Usb_Cat = "Harry Potter"
		case strings.Contains(apath, "Tremors"):
			Usb_Cat = "Tremors"
		case strings.Contains(apath, "Jurassic Park"):
			Usb_Cat = "Jurassic Park"
		case strings.Contains(apath, "John Wayne"):		
			Usb_Cat = "John Wayne"
		case strings.Contains(apath, "John Wick"):
			Usb_Cat = "John Wick"
		case strings.Contains(apath, "Misc"):
			Usb_Cat = "Misc"
		default:
			Usb_Cat = "Misc"
	}
	return
}
*/








/*type MOVI struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	DirPath string `bson: "dirpath"`
	Filepath string `bson: "filepath"`
	MediaId string `bson: "mediaid"`
	Movname string `bson: "movname"`
	Genre string `bson: "genre"`
	Catagory string `bson: "catagory"`
	MovPicPath string `bson: "movpicpath"`
	ThumbPath string `bson: "thumbpath"`
	MovYear string `bson: "movyear"`
}*/

/*func Get_mov_name(fname string) string {
    var mov_name string = ""
    if strings.Contains(fname, "(") {
	   fi := strings.Index(fname, "(")
	   fdex := fi - 1
	   mov_name = fname[:fdex]
    } else {
        ddex := len(fname) - 11
        mov_name = fname[ddex:]
    }
	return mov_name
}*/


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
	dirp, fname := path.Split(pAth)
	var movname string = moviegolib.Get_mov_name(fname)
	var Usb_picpath string = moviegolib.FindPicPaths(pAth, NO_ART_PIC_PATH)
	var Usb_picinfo string = moviegolib.CreateMoviesThumbnail(Usb_picpath)
	Usb_Mov_Info := moviegolib.MOVI{Id: bson.NewObjectId(),
		DirPath: dirp,
		Filepath: pAth,
		MediaId: moviegolib.UUID(),
		Movname: movname,
		Genre: "fart",
		Catagory: "Old",
		MovPicPath: Usb_picpath,
		ThumbPath: Usb_picinfo,
		MovYear: "1999",
	}
	ses := DBcon()
	defer ses.Close()	
	USBc := ses.DB("moviego").C("moviego")
	err := USBc.Insert(&Usb_Mov_Info)
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