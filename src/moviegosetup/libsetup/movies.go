package moviegolib


import (
    "path"
    "strings"
	"gopkg.in/mgo.v2/bson"
	"time"
	"math/rand"
	"strconv"
)


func Get_mov_name(movname string) string {
    var mov_name string = ""
    _, fname := path.Split(movname)
    if strings.Contains(fname, "(") {
	   fi := strings.Index(fname, "(")
	   fdex := fi - 1
	   mov_name = fname[:fdex]
    } else {
        ddex := len(fname) - 11
        mov_name = fname[ddex:]
    }
	return mov_name
}

func get_movie_year(apath string) (movyr string) {
	_, filename := path.Split(apath)
	fi := strings.Index(filename, "(")
	fdex := fi + 1
	ldex := strings.LastIndex(filename, ")")
	movyr = filename[fdex:ldex]
	return
}

type MOVI struct {
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
}

func Get_movie_info(apath string, mov_picpath string, mov_picinfo string) (Mov_Info MOVI) {
    switch {
        case strings.Contains(apath, "SciFi"):
            dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "SciFi",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Cartoons"):
            dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Cartoons",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Godzilla"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Godzilla",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Kingsman"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Kingsman",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "StarTrek") && !strings.Contains(apath, " STTV "):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "StarTrek",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "StarWars"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "StarWars",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "SuperHeros"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "SuperHeros",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "IndianaJones"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "IndianaJones",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Action"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Action",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Comedy"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Comedy",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Drama"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Drama",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
        case strings.Contains(apath, "Jurassic Park"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Jurassic Park",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "John Wayne"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "John Wayne",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "John Wick"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "John Wick",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
        case strings.Contains(apath, "Men In Black"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Men In Black",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Harry Potter"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Harry Potter",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
        case strings.Contains(apath, "Tremors"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Tremors",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
		case strings.Contains(apath, "Misc"):
			dirp, _ := path.Split(apath)
			Mov_Info = MOVI{Id: bson.NewObjectId(),
        		DirPath: dirp,
        		Filepath: apath,
        		MediaId: movies_UUID(),
        		Movname: Get_mov_name(apath),
        		Genre: "Movies",
        		Catagory: "Misc",
        		MovPicPath: mov_picpath,
        		ThumbPath: mov_picinfo,
        		MovYear: get_movie_year(apath),
            }
	}
	return
}

func movies_UUID() (UUID string) {
	a_seed := time.Now()
	aseed := a_seed.UnixNano()
	rand.Seed(aseed)
	u := rand.Int63n(aseed)
	p := strconv.FormatInt(u, 10)
	return p
}

/*func Get_movie_info(pAth string) (Mov_info MOVI) {
    dirp, movname1 := path.Split(apath)
	var movname string = Get_mov_name(movname1)
    var movname string = Moname(Movname1)
	var Mov_genre string = get_gen(apath)
	var Mov_catagory string = get_movie_cat(apath)
	var Mov_picpath string = moviegolib.FindPicPaths(pAth, NO_ART_PIC_PATH)
	var Mov_picinfo string = moviegolib.CreateMoviesThumbnail(Mov_picpath)
	Mov_Info = MOVI{Id: bson.NewObjectId(),
		DirPath: dirp,
		Filepath: pAth,
		MediaId: movies_UUID(),
		Movname: Get_mov_name(apath),
		Genre: get_gen(apath),
		Catagory: Mov_catagory,
		MovPicPath: Mov_picpath,
		ThumbPath: Mov_picinfo,
		MovYear: get_movie_year(apath),
	}
    ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	err := MTc.Insert(&Mov_Info)
	if err != nil {
		fmt.Println(err)
	return 
}*/