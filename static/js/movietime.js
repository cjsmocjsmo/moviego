function makeGrid1(aimglist) {
	var bb1 = "<div class='oddone ui-body ui-body-b'>"
	var blkAImg1 = "<a href='#intro'><img id='" + aimglist.mediaid + "' class='blkAImg' src='" + aimglist.thumbpath + "'></img></a>";
	var closeA1 = "</div>";
	var mg1 = bb1 + blkAImg1 + closeA1;
	return mg1;
};
/*
function makeGrid2(aimglist) {
	var grid = "<div class='ui-grid-a'>";
	var blkA = "<div class='ui-block-a'>";
	var barA = "<div class='ui-bar ui-bar-b'>";
	var blkAImg = "<a href='#intro'><img id='" + aimglist[0].MediaId + "' class='blkAImg' src='" + aimglist[0].Artwork + "'></img>";
	var closeA = "</div></div>";
	var blkB = "<div class='ui-block-b'>";
	var barB = "<div class='ui-bar ui-bar-b'>";
	var blkBImg = "<a href='#intro'><img id='" + aimglist[1].MediaId + "' class='blkBImg' src='" + aimglist[1].Artwork + "'></img></a>";
	var closeB = "</div></div>";
	var closeGrid = "</div>";
	var nblk2 = grid + blkA + barA + blkAImg + closeA + blkB + barB + blkBImg + closeB + closeGrid;
	return nblk2;
};*/

/*function makeGrid3(aimglist) {
	var grid = "<div class='ui-grid-b'>";
	var blkA = "<div class='ui-block-a'>";
	var barA = "<div class='ui-bar ui-bar-b'>";
	var blkAImg = "<a href='#intro'><img id='" + aimglist[0].MediaId + "' class='blkAImg' src='" + aimglist[0].HttpArtwork + "'></img></a>";
	var closeA = "</div></div>";
	var blkB = "<div class='ui-block-b'>";
	var barB = "<div class='ui-bar ui-bar-b'>";
	var blkBImg = "<a href='#intro'><img id='" + aimglist[1].MediaId + "' class='blkBImg' src='" + aimglist[1].HttpArtwork + "'></img></a>";
	var closeB = "</div></div>"
	var blkC = "<div class='ui-block-c'>";
	var barC = "<div class='ui-bar ui-bar-b'>";
	var blkCImg = "<a href='#intro'><img id='" + aimglist[2].MediaId + "' class='blkCImg' src='" + aimglist[2].HttpArtwork + "'></img></a>";
	var closeC = "</div></div>";
	var closeGrid = "</div>";
	var n = grid + blkA + barA + blkAImg + closeA;
	var n2 = n + blkB + barB + blkBImg + closeB;
	var n3 = n2 + blkC + barC + blkCImg + closeC + closeGrid;
	return n3;
};*/

function initSciFi() {
	if ($('#scifiMain').children().length === 0){
		
		$.get('IntSciFi', function (data) {
			$.each(data, function ( SciFikey, SciFival ) {

/*			var SciFileng = SciFival.length;
			if (SciFileng == 2) {
				var SciFiarr2 = makeGrid2(SciFival);
				$('#scifiMain').append(SciFiarr2);
			}
			if (SciFileng == 1) {*/
				var SciFiarr1 = makeGrid1(SciFival);
				$('#scifiMain').append(SciFiarr1);		
			})
		})
		$.mobile.loading("hide");
	}
};			

function initAction() {
	if ($('#actionMain').children().length === 0){
		$.get('IntAction', function (data) {
			$.each(data, function ( Actionkey, Actionval ) {
				var Actionarr1 = makeGrid1(Actionval);
				$('#actionMain').append(Actionarr1);
			})
		})
		/*  $.mobile.loading("hide");*/
	}
};

function initComedy() {
	if ($('#comedyMain').children().length === 0){
		$.get('IntComedy', function (data) {
			$.each(data, function ( Comedykey, Comedyval ) {
				var Comedyarr1 = makeGrid1(Comedyval);
				$('#comedyMain').append(Comedyarr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initDrama() {
	if ($('#dramaMain').children().length === 0){
		$.get('IntDrama', function (data) {
			$.each(data, function ( Dramakey, Dramaval ) {
				var Dramaarr3 = makeGrid1(Dramaval);
				$('#dramaMain').append(Dramaarr3);
			})
		})
		$.mobile.loading("hide");
	}
};

function initCartoons() {
	if ($('#cartoonsMain').children().length === 0){
		$.get('IntCartoons', function (data) {
			$.each(data, function ( Cartoonskey, Cartoonsval ) {
				var Cartoonsarr3 = makeGrid1(Cartoonsval);
				$('#cartoonsMain').append(Cartoonsarr3);
			})
		})
	}
};

function initKingsman() {
	if ($('#kingsmanMain').children().length === 0){
		$.get('IntKingsMan', function (data) {
			$.each(data, function ( Kingkey, Kingval ) {
				var Kingsmanarr2 = makeGrid1(Kingval);
				$('#kingsmanMain').append(Kingsmanarr2);
			})
		})
		$.mobile.loading("hide");
	}
};

function initGodzilla() {
	if ($('#godzillaMain').children().length === 0){
		$.get('IntGodzilla', function (data) {
			$.each(data, function ( GodzillaKey, GodzillaVal) {
				var Godzillaarr1 = makeGrid1(GodzillaVal);
				$('#godzillaMain').append(Godzillaarr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initStarTrek() {
	if ($('#startrekMain').children().length === 0){
		$.get('IntStarTrek', function (data) {
			$.each(data, function ( StarTrekkey, StarTrekval ) {
				var StarTrekarr1 = makeGrid1(StarTrekval);
				$('#startrekMain').append(StarTrekarr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initStarWars() {
	if ($('#starwarsMain').children().length === 0){
		$.get('IntStarWars', function (data) {
			$.each(data, function ( key, val ) {
				var arr1 = makeGrid1(val);
				$('#starwarsMain').append(arr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initSuperHeros() {
	if ($('#superherosMain').children().length === 0){
		$.get('IntSuperHeros', function (data) {
			$.each(data, function ( shkey, shval ) {
				var sharr1 = makeGrid1(shval);
				$('#superherosMain').append(sharr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initIndianaJones() {
	if ($('#indianajonesMain').children().length === 0){
		$.get('IntIndianaJones', function (data) {
			$.each(data, function ( ijkey, ijval ) {
				var IndianaJonesarr1 = makeGrid1(ijval);
				$('#indianajonesMain').append(IndianaJonesarr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initHarryPotter() {
	if ($('#harrypotterMain').children().length === 0){
		$.get('IntHarryPotter', function (data) {
			$.each(data, function ( HPkey, HPval ) {
				var HParr1 = makeGrid1(HPval);
				$('#harrypotterMain').append(HParr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initTremors() {
	if ($('#tremorsMain').children().length === 0){
		$.get('IntTremors', function (data) {
			$.each(data, function (Tremorskey, Tremorsval) {
				var HParr1 = makeGrid1(Tremorsval);
				$('#tremorsMain').append(HParr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initJohnWayne() {
	if ($('#johnwayneMain').children().length === 0){
		$.get('IntJohnWayne', function (data) {
			$.each(data, function ( JohnWaynekey, JohnWayneval ) {
				var JohnWaynearr1 = makeGrid1(JohnWayneval);
				$('#johnwayneMain').append(JohnWaynearr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initJurassicPark() {
	if ($('#jurassicparkMain').children().length === 0){
		$.get('IntJurassicPark', function (data) {
			$.each(data, function (jurkey, jurval) {
				var JParr1 = makeGrid1(jurval);
				$('#jurassicparkMain').append(JParr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initMisc() {
	if ($('#miscMain').children().length === 0){
		$.get('IntMisc', function (data) {
			$.each(data, function ( Misckey, Miscval ) {
				var Miscarr1 = makeGrid1(Miscval);
				$('#miscMain').append(Miscarr1);
			})
		})
		$.mobile.loading("hide");
	}
};

function initMenInBlack() {
	if ($('#meninblackMain').children().length === 0){
		$.get('IntMenInBlack', function (data) {
			$.each(data, function (mibkey, mibval) {
				var MIBarr1 = makeGrid1(mibval);
				$('#meninblackMain').append(MIBarr1);	
			})
		})
		$.mobile.loading("hide");
	}
};


var initMtime = function () {
	$("#foo3, #foo4, #doneBtn, #update2, #playerControls").hide();
//	initSciFi();
/*	initAction();
	initComedy();
	initDrama();
	initCartoons();*/
/*	initKingsman();
	initGodzilla();*/
/*	initStarTrek();
	initStarWars();
	initSuperHeros();*/
/*	initIndianaJones();
	initHarryPotter();*/
/*	initJohnWayne();
	initJurassicPark();*/
/*	initTremors();
	initMenInBlack();
	initMisc();*/
};
/*$(window).load(initMtime)*/
$(window).on("load", function () {
	initMtime();
});
