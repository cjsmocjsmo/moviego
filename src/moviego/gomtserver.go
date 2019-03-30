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
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"moviego/lib"
)


var moviegoForm = `
<!DOCTYPE html>
<html>
<head>
	<title>MovieTime</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link type="text/css" rel="stylesheet" href="static/css/jquery.mobile-1.4.5.min.css" />
	<link type="text/css" rel="stylesheet" href="static/css/jquery.mobile.theme-1.4.5.min.css" />
	<link type="text/css" rel="stylesheet" href="static/css/movietime.css" />
	<script src="static/js/jquery-2.1.4.min.js"></script>
	<script src="static/js/jquery.mobile-1.4.5.min.js"></script>
	<script src="static/js/movietime.js"></script>
	<script src="static/js/movietime2.js"></script>
</head>
<body>
<div class='amp' data-role="page" id="intro" data-theme="b">
	<div id="sttvPan" data-role='panel'>
		<a id="sttv1" href="#sttv" class="ui-btn ui-corner-all show-page-loading-msg">Season 1</a>	
		<a id="sttv2" href="#sttv" class="ui-btn ui-corner-all show-page-loading-msg">Season 2</a>
		<a id="sttv3" href="#sttv" class="ui-btn ui-corner-all show-page-loading-msg">Season 3</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
	<div id="tngPan" data-role='panel'>
		<a id="tng1" href="#tng" class="ui-btn ui-corner-all">Season 1</a>	
		<a id="tng2" href="#tng" class="ui-btn ui-corner-all">Season 2</a>
		<a id="tng3" href="#tng" class="ui-btn ui-corner-all">Season 3</a>
		<a id="tng4" href="#tng" class="ui-btn ui-corner-all">Season 4</a>	
		<a id="tng5" href="#tng" class="ui-btn ui-corner-all">Season 5</a>
		<a id="tng6" href="#tng" class="ui-btn ui-corner-all">Season 6</a>
		<a id="tng7" href="#tng" class="ui-btn ui-corner-all">Season 7</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
	<div id="voyPan" data-role='panel'>
		<a id="voy1" href="#voy" class="ui-btn ui-corner-all">Season 1</a>	
		<a id="voy2" href="#voy" class="ui-btn ui-corner-all">Season 2</a>
		<a id="voy3" href="#voy" class="ui-btn ui-corner-all">Season 3</a>
		<a id="voy4" href="#voy" class="ui-btn ui-corner-all">Season 4</a>	
		<a id="voy5" href="#voy" class="ui-btn ui-corner-all">Season 5</a>
		<a id="voy6" href="#voy" class="ui-btn ui-corner-all">Season 6</a>
		<a id="voy7" href="#voy" class="ui-btn ui-corner-all">Season 7</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
	<div id="entPan" data-role='panel'>
		<a id="ent1" href="#ent" class="ui-btn ui-corner-all">Season 1</a>	
		<a id="ent2" href="#ent" class="ui-btn ui-corner-all">Season 2</a>
		<a id="ent3" href="#ent" class="ui-btn ui-corner-all">Season 3</a>
		<a id="ent4" href="#ent" class="ui-btn ui-corner-all">Season 4</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
	<div id="disPan" data-role='panel'>
		<a id="dis1" href="#dis" class="ui-btn ui-corner-all">Season 1</a>	
		<a id="dis2" href="#dis" class="ui-btn ui-corner-all">Season 2</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
	<div id="orvPan" data-role='panel'>
		<a id="orv1" href="#orv" class="ui-btn ui-corner-all">Season 1</a>	
		<a id="orv2" href="#orv" class="ui-btn ui-corner-all">Season 2</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
	<div id="tlsPan" data-role="panel">
		<a id="tls1" href="#tls" class="ui-btn ui-corner-all">Season 1</a>
		<a id="tls2" href="#tls" class="ui-btn ui-corner-all">Season 2</a>
		<a id="tls3" href="#tls" class="ui-btn ui-corner-all">Season 3</a>
		<a id="tls4" href="#tls" class="ui-btn ui-corner-all">Season 4</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
	
<!--	<div id="playerPan" data-role="panel">
		<a href="" id="PannextBtn" class="ui-btn ui-corner-all">Seek 60 >></a>
		<a href="" id="PanpChapterBtn" class="ui-btn ui-corner-all">Prev Chapter</a>
		<a href="" id="PannChapterBtn" class="ui-btn ui-corner-all">Next Chapter</a>
		<a href="" id="PanprevBtn" class="ui-btn ui-corner-all"><< Seek 30</a>
		<a class="ui-btn ui-corner-all" data-rel="close">Close</a>
	</div>
-->
	
	
	
	
	
	
	
	
	<div data-role="header" data-theme="b" data-position="fixed">
	<!--	<a href="#" id="shutdownBtn" data-ajax='false' class="ui-btn-left ui-btn ui-btn-inline ui-mini ui-corner-all">Power Off</a>-->
		<a href="#updatePg" id="updateBtn" data-ajax='false' class="ui-btn-left ui-btn ui-btn-inline ui-mini ui-corner-all">Update</a>
		<h1 id="introH1">MovieTime</h1>
	</div>
	
	<div data-role="tabs" id="tags">
		<div data-role="navbar">
			<ul>
				<li><a href="#Moovs" data-ajax="false"><img id="movimg" src="static/images/movie-7.png"></img></a></li>
<!--				<li><a href="http://192.168.1.102:8080/ampnado" data-ajax='false'>Music</a></li>-->
				<li><a href="#TVVs" data-ajax="false"><img id="tvimg" src="static/images/tv.png"></img></a></li>
			</ul>
		</div>
		<a href="" id="playerPanBtn" class="ui-btn ui-corner-all">Player Controls</a>
		<div id="playerControls" data-role="navbar">
			<ul>
				<li><button id="playBtn" class="ui-btn ui-corner-all">Play</button></li>
				<li><button id="pauseBtn" class="ui-btn ui-corner-all">Pause</button></li>
				<li><button id="stopBtn" class="ui-btn ui-corner-all">Stop</button></li>
			</ul>

			<ul>
				<li><button id="prevBtn" class="ui-btn ui-icon-carat-l ui-btn-icon-top ui-corner-all">Seek</button></li>
				<li><button id="pChapterBtn" class="ui-btn ui-icon-carat-l ui-btn-icon-top ui-corner-all">Chapter</button></li>
				<li><button id="nChapterBtn" class="ui-btn ui-icon-carat-r ui-btn-icon-top ui-corner-all">Chapter</button></li>
				<li><button id="nextBtn" class="ui-btn ui-icon-carat-r ui-btn-icon-top ui-corner-all">Seek</button></li>
			</ul>
		</div>

		
		
<!--		<a href="#" data-role="button" data-inline='true' id="playBtn">Play</button>
		<a href="#" data-role="button" data-inline='true' id="pauseBtn">Pause</button>
		<a href="#" data-role="button" data-inline='ture' id="stopBtn">Stop</a>
		<a href-"#" data-role="button" data-inline='true' id="nextBtn">Next</button>
		<a href="#" data-role="button" data-inline='true' id="prevBtn">Previous</button>-->
		<div id="Moovs" class="ui-body-d ui-content">
			<a id="actionBtn" href="#action" class="ui-btn ui-corner-all show-page-loading-msg">Action</a>
			<a id="cartoonsBtn" href="#cartoons" class="ui-btn ui-corner-all show-page-loading-msg">Cartoons</a>
			<a id="comedyBtn" href="#comedy" class="ui-btn ui-corner-all show-page-loading-msg">Comedy</a>
			<a id="dramaBtn" href="#drama" class="ui-btn ui-corner-all show-page-loading-msg">Drama</a>
			<a id="godzillaBtn" href="#godzilla" class="ui-btn ui-corner-all show-page-loading-msg">Godzilla</a>
			<a id="harrypotterBtn" href="#harrypotter" class="ui-btn ui-corner-all show-page-loading-msg">Harry Potter</a>
			<a id="indianajonesBtn" href="#indianajones" class="ui-btn ui-corner-all show-page-loading-msg">Indiana Jones</a>
			<a id="johnwayneBtn" href="#johnwayne" class="ui-btn ui-corner-all show-page-loading-msg">John Wayne</a>
			<a id="jurassicparkBtn" href="#jurassicpark" class="ui-btn ui-corner-all show-page-loading-msg">Jurassic Park</a>
			<a id="kingsmanBtn" href="#kingsman" class="ui-btn ui-corner-all show-page-loading-msg">KingsMan</a>
			<a id="meninblackBtn" href="#meninblack" class="ui-btn ui-corner-all">Men In Black</a>
			<a id="scifiBtn" href="#scifi" class="ui-btn ui-corner-all show-page-loading-msg">SciFi</a>
			<a id="startrekMBtn" href="#startrekM" class="ui-btn ui-corner-all show-page-loading-msg">Star Trek</a>
			<a id="starwarsBtn" href="#starwars" class="ui-btn ui-corner-all show-page-loading-msg">Star Wars</a>
			<a id="superherosBtn" href="#superheros" class="ui-btn ui-corner-all show-page-loading-msg">Super Heros</a>
			<a id="tremorsBtn" href="#tremors" class="ui-btn ui-corner-all show-page-loading-msg">Tremors</a>
			<a id="miscBtn" href="#misc" class="ui-btn ui-corner-all show-page-loading-msg">Misc</a>
		</div>
		<div id="TVVs" class="ui-body-d ui-content">
			<a href="#sttvPan" class="ui-btn ui-corner-all">Star Trek</a>
			<a href="#tngPan" class="ui-btn ui-corner-all">The Next Generation</a>
			<a href="#voyPan" class="ui-btn ui-corner-all">Voyager</a>
			<a href="#entPan" class="ui-btn ui-corner-all">Enterprise</a>
			<a href="#disPan" class="ui-btn ui-corner-all">Discovery</a>
			<a href="#orvPan" class="ui-btn ui-corner-all">Orville</a>
			<a href="#tlsPan" class="ui-btn ui-corner-all">The Last Ship</a>
		</div>
	</div>
</div>


<!--	
	<div data-role="header" data-theme="b" data-position="fixed">
		<a href="#" id="updateBtn" data-ajax='false' class="ui-btn-left ui-btn ui-btn-inline ui-mini ui-corner-all">Update</a>
		<h1 id="introH1">MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><button id="pauseBtn" class="ui-btn ui-corner-all">Pause</button></li>
				<li><button id="stopBtn" class="ui-btn ui-corner-all">Stop</button></li>
				<li><button id="nextBtn" class="ui-btn ui-corner-all">Next</button></li>
				<li><button id="prevBtn" class="ui-btn ui-corner-all">Previous</button></li>
			</ul>		
		</div>
	</div>
	<div id="Moovs" class="ui-grid-a">
		<div class="ui-block-a">
		
			<a href="#action" class="ui-btn ui-corner-all">Action</a>
			<a href="#cartoons" class="ui-btn ui-corner-all">Cartoons</a>
			<a href="#comedy" class="ui-btn ui-corner-all">Comedy</a>
			
			<a href="#drama" class="ui-btn ui-corner-all">Drama</a>
			
			<a href="#godzilla" class="ui-btn ui-corner-all">Godzilla</a>
			<a href="#harrypotter" class="ui-btn ui-corner-all">Harry Potter</a>
			<a href="#indianajones" class="ui-btn ui-corner-all">Indiana Jones</a>
			<a href="#johnwayne" class="ui-btn ui-corner-all">John Wayne</a>
			<a href="#jurassicpark" class="ui-btn ui-corner-all">Jurassic Park</a>
			<a href="#kingsman" class="ui-btn ui-corner-all">KingsMan</a>
			<a href="#meninblack" class="ui-btn ui-corner-all">Men In Black</a>
			<a href="#scifi" class="ui-btn ui-corner-all">SciFi</a>
			<a href="#startrekM" class="ui-btn ui-corner-all">Star Trek</a>
			<a href="#starwars" class="ui-btn ui-corner-all">Star Wars</a>
			<a href="#superheros" class="ui-btn ui-corner-all">Super Heros</a>
			<a href="#tremors" class="ui-btn ui-corner-all">Tremors</a>
			<a href="#misc" class="ui-btn ui-corner-all">Misc</a>
		</div>
		<div class="ui-block-b">
			<a href="#sttvPan" class="ui-btn ui-corner-all">Star Trek</a>
			<a href="#tngPan" class="ui-btn ui-corner-all">The Next Generation</a>
			<a href="#voyPan" class="ui-btn ui-corner-all">Voyager</a>
			<a href="#entPan" class="ui-btn ui-corner-all">Enterprise</a>
			<a href="#disPan" class="ui-btn ui-corner-all">Discovery</a>
			<a href="#orvPan" class="ui-btn ui-corner-all">Orville</a>
			<a href="#tlsPan" class="ui-btn ui-corner-all">The Last Ship</a>
		</div>
	</div>
</div>
-->


<div data-role="page" id="updatePg" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
<!--		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>-->
	</div>
	<div id='updateMain' role='main' class='ui-content' data-theme='b'>
		<span id="update1">Updating and adding new content</span>
		<span id="update2">Update is complete</span>
		<a href="#intro" id="doneBtn" class="ui-btn ui-corner-all">Done</a>
	</div>		
</div>





<div data-role="page" id="scifi" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id='scifiMain' role='main' class='ui-content' data-theme='b'>
	</div>		
</div>





<div data-role="page" id="action" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id='actionMain' role='main' class='ui-content' data-theme='b'>
	</div>		
</div>



<div data-role="page" id="comedy" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id='comedyMain' role='main' class='ui-content' data-theme='b'>
	</div>		
</div>




<div data-role="page" id="drama" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id='dramaMain' role='main' class='ui-content' data-theme='b'>
	</div>		
</div>






















<div data-role="page" id="cartoons" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id="cartoonsMain" role="main" class="ui-content">
	</div>
</div>

<div data-role="page" id="kingsman" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id="kingsmanMain" role="main" class="ui-content">
	</div>
</div>

<div data-role="page" id="godzilla" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id="godzillaMain" role="main" class="ui-content">
	</div>
</div>

<div data-role="page" id="startrekM" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id="startrekMain" role="main" class="ui-content">
	</div>
</div>

<div data-role="page" id="starwars" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role="navbar">
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>			
			</ul>
		</div>
	</div>
	<div id="starwarsMain" role="main" class="ui-content">
	</div>
</div>

<div class='amp' data-role="page" id="superheros" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="superherosMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div data-role="page" id="sttv" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>
	</div>
	<div id="sttvMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div class='amp' data-role="page" id="harrypotter" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="harrypotterMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div class='amp' data-role="page" id="indianajones" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="indianajonesMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div class='amp' data-role="page" id="jurassicpark" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="jurassicparkMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div class='amp' data-role="page" id="meninblack" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="meninblackMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>


<div class='amp' data-role="page" id="tremors" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="tremorsMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div class='amp' data-role="page" id="johnwayne" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="johnwayneMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div class='amp' data-role="page" id="misc" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>	
	</div> 
	<div id="miscMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>

<div data-role="page" id="tng" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>
	</div>
	<div id="tngMain" role="main" class="ui-content" data-theme="b">
	</div>
</div>
<div data-role="page" id="voy" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>
	</div>
	<div id="voyMain" role="main" class="ui-content" data-theme="b">

	</div>
</div>
<div data-role="page" id="ent" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>
	</div>
	<div id="entMain" role="main" class="ui-content" data-theme="b">

	</div>
</div>
<div data-role="page" id="dis" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>
	</div>
	<div id="disMain" role="main" class="ui-content" data-theme="b">

	</div>
</div>
<div data-role="page" id="orv" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>
	</div>
	<div id="orvMain" role="main" class="ui-content" data-theme="b">

	</div>
</div>
<div data-role="page" id="tls" data-theme="b">
	<div data-role="header" data-theme="b" data-position="fixed">
		<h1>MovieTime</h1>
		<div data-role='navbar'>
			<ul>
				<li><a href="#intro" data-theme='a'>Home</a></li>
			</ul>
		</div>
	</div>
	<div id="tlsMain" role="main" class="ui-content" data-theme="b">

	</div>
</div>





</div>
`

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
	fmt.Fprintf(w, "%s", moviegoForm)
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