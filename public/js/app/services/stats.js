

app.factory('stats', function () {
	
	var eventcargo = "";
	var listofuniqueplayers = [];
	var scores = [];
	var prettyscores = [];
	var gamescores = [];
	var orgplayeruuid = playerinfo.UUID;
	
	//for event table
	var showeventdata = [];

	
    return {
		
		getNemesis: function (topx){
			var retVal = [];
			scores.sort(function(a, b) {
					if (a.beatenbythem < b.beatenbythem) {
						return 1;
					}
					if (a.beatenbythem > b.beatenbythem) {
						return -1;
					}
					return 0;
			});
		
			for (m=0;m<scores.length;m++) {	
				retVal.push( { name:scores[m].name, record:scores[m].beatenbythem} );
			}
			console.log("RETVAL:", retVal);
			
			//pad the return
			if (retVal.length < topx) {
				for (j=0;j<(topx-retVal.length);j++) {
					retVal.push( { name:"Not Enough Players", record:0} );
				}
				
			}

			return (retVal);
			
		},
		getDominating: function (topx){
			var retVal = [];
			scores.sort(function(a, b) {
					if (a.beatthem < b.beatthem) {
						return 1;
					}
					if (a.beatthem > b.beatthem) {
						return -1;
					}
					return 0;
			});

			for (m=0;m<scores.length;m++) {	
				retVal.push( { name:scores[m].name, record:scores[m].beatthem} );
			}
			
			
			//pad the return
			if (retVal.length < topx) {
				for (j=0;j<(topx-retVal.length);j++) {
					retVal.push( { name:"Not Enough Players", record:0} );
				}
				
			}
			return (retVal);
			
		},
		
        setEventcargo: function (setcargo) {
			eventcargo = setcargo;

			//try to do one loop to get all the stats we need
			for (var i = 0; i < eventcargo.length; i++) {
				
				showeventdata[i] = eventcargo[i].Event;
				showeventdata[i].location = eventcargo[i].Location;
				
				
				//loop through game list
				var gamelocforthisevent = [];
				var gamelist = "";
				var numofgames = (eventcargo[i].Games).length;
				for (var j = 0; j < numofgames ; j++) {
					
					var gameuuid = eventcargo[i].Games[j].UUID;
					var gameuuidExists = false;
					for (k=0;k<gamescores.length;k++) {
						if (gamescores[k].uuid == gameuuid) {
							gameuuidExists = true;
							gamelocforthisevent.push(k);
						}
					}
					
					if (!gameuuidExists) {
						var gamescorehistory = {
							name: eventcargo[i].Games[j].Name + " " + eventcargo[i].Games[j].Published,
							uuid: gameuuid,
							DEMOLISH: 0,
							WON: 0,
							TIE: 0,
							LOST: 0,
							SKUNK: 0,
							DROP: 0,
							QUIT: 0,
							RAGEQUIT: 0
						};
						gamescores.push(gamescorehistory);
						gamelocforthisevent.push(gamescores.length-1);
						
						
					}
					
					var strtouse = gamelist + ", " + eventcargo[i].Games[j].Name  ;
					
					if ( j == 0 ) {
						strtouse = eventcargo[i].Games[j].Name  ;
					} else if  (j == ( numofgames -1 )) {
						strtouse = gamelist + ", " + eventcargo[i].Games[j].Name;
					} 
						
					gamelist = strtouse;
					
				}
				
				
				showeventdata[i].gamesplayed = gamelist;

				var meuuid = 666;
				var ident = 0;
				var sortA = [];
				var prettyresults = "";
				//loop through competitors and grab stats
				for (j=0;j<eventcargo[i].Competitors.length;j++) {
					var playeruuid = eventcargo[i].Competitors[j].Player.UUID;
					var temprate = {
						name: eventcargo[i].Competitors[j].Player.Firstname + " " + eventcargo[i].Competitors[j].Player.Surname,
						uuid: playeruuid,
						ranking: eventcargo[i].Competitors[j].Result.Place,
						beatthem: 0,
						beatenbythem: 0,
						tiedthem: 0,
						rating: 0,
						ratingdeviation: 0,
						volatility: 0
					};
					
					var uuidExists = false;
					for (k=0;k<scores.length;k++) {
					
						if (scores[k].uuid == playeruuid) {
							uuidExists = true;
						}
					}
					if (!uuidExists) {
						var scorehistory = {
							name: eventcargo[i].Competitors[j].Player.Firstname + " " + eventcargo[i].Competitors[j].Player.Surname,
							uuid: playeruuid,
							beatthem: 0,
							beatenbythem: 0,
							tiedthem: 0,
							rating: eventcargo[i].Competitors[j].Rating.Rating,
							ratingdeviation: eventcargo[i].Competitors[j].Rating.RatingDeviation,
							volatility: eventcargo[i].Competitors[j].Rating.Volatility,
							prettyrating: parseInt(eventcargo[i].Competitors[j].Rating.Rating),
							prettyratingdeviation: parseInt(eventcargo[i].Competitors[j].Rating.RatingDeviation),
							prettyvolatility: parseFloat(eventcargo[i].Competitors[j].Rating.Volatility).toFixed(3),
							
						};
						scores.push(scorehistory);
					}
					
					sortA.push(temprate);

					var playername = 	eventcargo[i].Competitors[j].Player.Firstname + " "  +
								eventcargo[i].Competitors[j].Player.Surname;
								
					var status = eventcargo[i].Competitors[j].Result.Result;
					
					//update the game stats with the results
					if (playerinfo.UUID == playeruuid) {
						for (var q = 0; q < gamelocforthisevent.length ; q++) {
								gamescores[gamelocforthisevent[q]][status]++;
						}
						
					}
					
					
					var place = parseInt(eventcargo[i].Competitors[j].Result.Place);
					switch (  place ) {
						case 1:
							place = place + "st";
							break;
						case 2:
							place = place + "nd";
							break;
						case 3:
							place = place + "rd";
							break;
						default:
							place = place + "th";
					}
					var comma = "";
					
					if ( j == eventcargo[i].Competitors.length-1 ) {
						comma = "";
					
					} else {
						comma = ", ";
					}
					prettyresults = prettyresults + playername + " "  + status + " " + place + comma;
					
				}
				
				
					
				sortA.sort(function(a, b) {
							if (a.ranking < b.ranking) {
								return -1;
							}
							if (a.ranking > b.ranking) {
								return 1;
							}
							return 0;
				});
				
				var foundme = false;
				for (k=0;k<sortA.length;k++) {
		
					//console.log("rated:", rated, k);
					if (playerinfo.UUID == sortA[k].uuid ) {
						foundme = true;
					} else 
					{
						var ratedExist = false;
						var ratedLoc = 0;
						for (m=0;m<scores.length;m++) {
							if (scores[m].uuid == sortA[k].uuid) {
								ratedExist = true;
								ratedLoc = m;
							}
						}
						if (!ratedExist) {
							
							//scores[rated.length-1].beatenbythem++;
						}  else {
					
							if (!foundme) {
								scores[ratedLoc].beatenbythem++; 
								
							} else {
								scores[ratedLoc].beatthem++; 
							}
						}
					}
					
				}
				
				
				showeventdata[i].results = prettyresults;
				startdisplay = moment(showeventdata[i].Start).utc();
				stopdisplay = moment(showeventdata[i].Stop).utc();

				showeventdata[i].Start = startdisplay.format("YYYY-MM-DD HH:mm UTC");
				showeventdata[i].Stop = stopdisplay.format("YYYY-MM-DD HH:mm UTC");
				
			}
			
			//pretty scores
			for (k=0;k<scores.length;k++) {
				if (scores[k].uuid == orgplayeruuid) {
							scores[k].prettybeatthem = "N/A";
							scores[k].prettybeatenbythem = "N/A";
							scores[k].prettytiedthem = "N/A";
				} else {
							scores[k].prettybeatthem = scores[k].beatthem;
							scores[k].prettybeatenbythem = scores[k].beatenbythem ;
							scores[k].prettytiedthem = scores[k].tiedthem;
					
					
				}

			}
										
			
			console.log("PRETTY SCORES", prettyscores);
			
			console.log("scores", scores);
            return true;                   
        },
		
		
		getPrettyGameStats: function() {
            return gamescores;
		},
		
		getPrettyEventStats: function() {
            return showeventdata;
		},
		
		getPrettyPlayerStats: function() {

            return scores;
		},
		
		getNumberEvents: function() {
            return eventcargo.length;
		},
		
		getNumberCompetitors: function() {
			
			
			//return listofuniqueplayers.length;
			return scores.length;
			
		}
	

    }
	    

});

