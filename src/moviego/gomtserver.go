///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
	// LICENSE: GNU General Public License, version 2 (GPLv2)
	// Copyright 2016, Charlie J. Smotherman
	//
	// This program is free software; you can redistribute it and/or
	// modify it under the terms of the GNU General Public License v2
	// as published by the Free Software Foundation.
	//
	// This program is distributed in the hope that it will be useful,
	// but WITHOUT ANY WARRANTY; without even the implied warranty of
	// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	// GNU General Public License for more details.
	//
	// You should have received a copy of the GNU General Public License
	// along with this program; if not, write to the Free Software
	// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
package main

import (
	"log"
	"fmt"
	"net/http"
	"net/url"
	"html/template"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"moviego/lib"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func DBcon() *mgo.Session {
	s, err := mgo.Dial("127.0.0.1")
//	s, err := mgo.Dial("mongodb://192.168.1.101:27017")
	if err != nil {
		fmt.Println("this is dial err")
		panic(err)
	}
	return s
}

func ShowMovieGo(w http.ResponseWriter, r *http.Request) {
    tmppath := "/home/pi/GoBuild/moviego/static/template/movietime.template"
    tmpl := template.Must(template.ParseFiles(tmppath))
    tmpl.Execute(w, tmpl)
	
}

func IntActionHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var ActionMedia []map[string]string
	b1 := bson.M{"catagory":"Action"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ActionMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ActionMedia)
}

func IntCartoonsHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	b1 := bson.M{"catagory":"Cartoons"}
	b2 := bson.M{"_id": 0}
	var CartoonMedia []map[string]string
	err := MTc.Find(b1).Select(b2).All(&CartoonMedia)
	if err != nil {
		log.Println(err)	
	}
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CartoonMedia)
}

func IntComedyHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var ComedyMedia []map[string]string
	b1 := bson.M{"catagory":"Comedy"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ComedyMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ComedyMedia)
}

func IntDramaHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var DramaMedia []map[string]string
	b1 := bson.M{"catagory":"Drama"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DramaMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DramaMedia)
}

func IntGodzillaHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var GodzillaMedia []map[string]string
	b1 := bson.M{"catagory":"Godzilla"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&GodzillaMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GodzillaMedia)
}

func IntIndianaJonesHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var IndianaJonesMedia []map[string]string
	b1 := bson.M{"catagory":"IndianaJones"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&IndianaJonesMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IndianaJonesMedia)
}

func IntJohnWayneHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var JohnWayneMedia []map[string]string
	b1 := bson.M{"catagory":"John Wayne"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWayneMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JohnWayneMedia)
}

func IntJurassicParkHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var JurassicParkMedia []map[string]string
	b1 := bson.M{"catagory":"Jurassic Park"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JurassicParkMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JurassicParkMedia)
}

func IntKingsManHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var KingsmanMedia []map[string]string
	b1 := bson.M{"catagory":"Kingsman"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&KingsmanMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KingsmanMedia)
}

func IntHarryPotterHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var HarryPotterMedia []map[string]string
	b1 := bson.M{"catagory":"Harry Potter"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&HarryPotterMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HarryPotterMedia)
}

func IntMenInBlackHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var MenInBlackMedia []map[string]string
	b1 := bson.M{"catagory":"Men In Black"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MenInBlackMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MenInBlackMedia)
}

func IntMiscHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var MiscMedia []map[string]string
	b1 := bson.M{"catagory":"Misc"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MiscMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MiscMedia)
}

func IntSciFiHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var SciFiMedia []map[string]string
	b1 := bson.M{"catagory":"SciFi"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SciFiMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SciFiMedia)
}

func IntStarTrekHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var StarTrekMedia []map[string]string
	b1 := bson.M{"catagory":"StarTrek"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarTrekMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(StarTrekMedia)
}

func IntStarWarsHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var StarWarsMedia []map[string]string
	b1 := bson.M{"catagory":"StarWars"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarWarsMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(StarWarsMedia)
}

func IntSuperHerosHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var SuperHerosMedia []map[string]string
	b1 := bson.M{"catagory":"SuperHeros"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SuperHerosMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuperHerosMedia)
}

func IntTremorsHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TremorsMedia []map[string]string
	b1 := bson.M{"catagory":"Tremors"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TremorsMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TremorsMedia)
}

func STTVS1Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var STTVS1Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "series":"Star Trek", "season": "01"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&STTVS1Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(STTVS1Media)
}

func STTVS2Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var STTVS2Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "series":"Star Trek", "season": "02"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&STTVS2Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(STTVS2Media)
}

func STTVS3Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var STTVS3Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "series":"Star Trek", "season": "03"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&STTVS3Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(STTVS3Media)
}

func TNGS1Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TNGS1Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TNG", "season": "01"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TNGS1Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TNGS1Media)
}

func TNGS2Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TNGS2Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TNG", "season": "02"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TNGS2Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TNGS2Media)
}

func TNGS3Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TNGS3Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TNG", "season": "03"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TNGS3Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TNGS3Media)
}

func TNGS4Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TNGS4Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TNG", "season": "04"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TNGS4Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TNGS4Media)
}

func TNGS5Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TNGS5Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TNG", "season": "05"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TNGS5Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TNGS5Media)
}

func TNGS6Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TNGS6Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TNG", "season": "06"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TNGS6Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TNGS6Media)
}

func TNGS7Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TNGS7Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TNG", "season": "07"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TNGS7Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TNGS7Media)
}

func VOYS1Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var VOYS1Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Voyager", "season": "01"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&VOYS1Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VOYS1Media)
}

func VOYS2Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var VOYS2Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Voyager", "season": "02"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&VOYS2Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VOYS2Media)
}

func VOYS3Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var VOYS3Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Voyager", "season": "03"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&VOYS3Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VOYS3Media)
}

func VOYS4Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var VOYS4Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Voyager", "season": "04"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&VOYS4Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VOYS4Media)
}

func VOYS5Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var VOYS5Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Voyager", "season": "05"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&VOYS5Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VOYS5Media)
}

func VOYS6Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var VOYS6Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Voyager", "season": "06"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&VOYS6Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VOYS6Media)
}

func VOYS7Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var VOYS7Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Voyager", "season": "07"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&VOYS7Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VOYS7Media)
}

func ENTS1Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var ENTS1Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Enterprise", "season": "01"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&ENTS1Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ENTS1Media)
}

func ENTS2Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var ENTS2Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Enterprise", "season": "02"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&ENTS2Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ENTS2Media)
}

func ENTS3Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var ENTS3Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Enterprise", "season": "03"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&ENTS3Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ENTS3Media)
}

func ENTS4Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var ENTS4Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Enterprise", "season": "04"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&ENTS4Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ENTS4Media)
}







func DISS1Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var DISS1Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Discovery", "season": "01"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&DISS1Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DISS1Media)
}

func ORVS1Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var ORVS1Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"Orville", "season": "01"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&ORVS1Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ORVS1Media)
}

func TLSS1Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TLSS1Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TheLastShip", "season": "01"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TLSS1Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TLSS1Media)
}

func TLSS2Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TLSS2Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TheLastShip", "season": "02"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TLSS2Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TLSS2Media)
}

func TLSS3Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TLSS3Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TheLastShip", "season": "03"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TLSS3Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TLSS3Media)
}

func TLSS4Handler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var TLSS4Media []map[string]string
	b1 := bson.M{"genre":"TVShows", "catagory":"TheLastShip", "season": "04"}
	b2 := bson.M{"_id":0, "title":1, "mediaid":1, "episode":1, "thumbpath":1}
	err := MTc.Find(b1).Select(b2).Sort("episode").All(&TLSS4Media)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TLSS4Media)
}

func PlayMediaHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(err)	
	}
	q := u.Query()
	mid := q.Get("mid")
	ses := DBcon()
	defer ses.Close()	
	MTc := ses.DB("moviego").C("moviego")
	var MediaInfo map[string]string
	b1 := bson.M{"mediaid":mid}
	b2 := bson.M{"_id": 0, "moviename":0, "catagory":0, "artwork":0, "genre":0, "movieyear":0, "mediaid":0}
	err = MTc.Find(b1).Select(b2).One(&MediaInfo)
	if err != nil {
		log.Println(err)
	}

	omxplayer.Play(MediaInfo["filepath"])
	omxplayer.ParseOutput()
}
	
func PlayHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Resume()
}

func PauseHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Pause()
}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Stop()
}

//Seek forward
func NextHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Fwd()
}
//Seek backward
func PreviousHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Bwd()
}
//Next Chapter
func NextChapterHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Next()
}
//Previous Chapter
func PreviousChapterHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Prev()
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/static").Subrouter()
	r.HandleFunc("/moviego", ShowMovieGo)
	r.HandleFunc("/IntAction", IntActionHandler)
	r.HandleFunc("/IntCartoons", IntCartoonsHandler)
	r.HandleFunc("/IntComedy", IntComedyHandler)
	r.HandleFunc("/IntDrama", IntDramaHandler)
	r.HandleFunc("/IntGodzilla", IntGodzillaHandler)
	r.HandleFunc("/IntHarryPotter", IntHarryPotterHandler)
	r.HandleFunc("/IntIndianaJones", IntIndianaJonesHandler)
	r.HandleFunc("/IntMenInBlack", IntMenInBlackHandler)
	r.HandleFunc("/IntMisc", IntMiscHandler)
	r.HandleFunc("/IntJohnWayne", IntJohnWayneHandler)
	r.HandleFunc("/IntJurassicPark", IntJurassicParkHandler)
	r.HandleFunc("/IntKingsMan", IntKingsManHandler)
	r.HandleFunc("/IntSciFi", IntSciFiHandler)
	r.HandleFunc("/IntStarTrek", IntStarTrekHandler)
	r.HandleFunc("/IntStarWars", IntStarWarsHandler)
	r.HandleFunc("/IntSuperHeros", IntSuperHerosHandler)
	r.HandleFunc("/IntTremors", IntTremorsHandler)
	r.HandleFunc("/STTVs1", STTVS1Handler)
	r.HandleFunc("/STTVs2", STTVS2Handler)
	r.HandleFunc("/STTVs3", STTVS3Handler)
	r.HandleFunc("/TNGs1", TNGS1Handler)
	r.HandleFunc("/TNGs2", TNGS2Handler)
	r.HandleFunc("/TNGs3", TNGS3Handler)
	r.HandleFunc("/TNGs4", TNGS4Handler)	
	r.HandleFunc("/TNGs5", TNGS5Handler)	
	r.HandleFunc("/TNGs6", TNGS6Handler)
	r.HandleFunc("/TNGs7", TNGS7Handler)
	r.HandleFunc("/VOYs1", VOYS1Handler)
	r.HandleFunc("/VOYs2", VOYS2Handler)
	r.HandleFunc("/VOYs3", VOYS3Handler)
	r.HandleFunc("/VOYs4", VOYS4Handler)
	r.HandleFunc("/VOYs5", VOYS5Handler)
	r.HandleFunc("/VOYs6", VOYS6Handler)
	r.HandleFunc("/VOYs7", VOYS7Handler)
	r.HandleFunc("/ENTs1", ENTS1Handler)
	r.HandleFunc("/ENTs2", ENTS2Handler)
	r.HandleFunc("/ENTs3", ENTS3Handler)
	r.HandleFunc("/ENTs4", ENTS4Handler)
	r.HandleFunc("/DISs1", DISS1Handler)
	r.HandleFunc("/ORVs1", ORVS1Handler)
	r.HandleFunc("/TLSs1", TLSS1Handler)
	r.HandleFunc("/TLSs2", TLSS2Handler)
	r.HandleFunc("/TLSs3", TLSS3Handler)
	r.HandleFunc("/TLSs4", TLSS4Handler)
	r.HandleFunc("/PlayMedia", PlayMediaHandler)
	r.HandleFunc("/Play", PlayHandler)
	r.HandleFunc("/Pause", PauseHandler)
	r.HandleFunc("/Stop", StopHandler)
	r.HandleFunc("/Next", NextHandler)
	r.HandleFunc("/Previous", PreviousHandler)
	
	r.HandleFunc("/NextChapter", NextChapterHandler)
	r.HandleFunc("/PreviousChapter", PreviousChapterHandler)
	
	
	
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/media/"))))
	http.ListenAndServe(":8888", (r))
}
