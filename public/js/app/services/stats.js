

app.factory('stats', function () {
	
	var eventcargo = "";
	var listofuniqueplayers = [];
	
	//for event table
	var showeventdata = [];
	
    return {
		
        setEventcargo: function (setcargo) {
			eventcargo = setcargo;

			//try to do one loop to get all the stats we need
			for (var i = 0; i < eventcargo.length; i++) {
				
				showeventdata[i] = eventcargo[i].Event;
				showeventdata[i].location = eventcargo[i].Location;
				
				
				var gamelist = "";
				var numofgames = (eventcargo[i].Games).length;
				for (var j = 0; j < numofgames ; j++) {
					var strtouse = gamelist + ", " + eventcargo[i].Games[j].Name  ;
					
					if ( j == 0 ) {
						strtouse = eventcargo[i].Games[j].Name  ;
					} else if  (j == ( numofgames -1 )) {
						strtouse = gamelist + ", " + eventcargo[i].Games[j].Name;
					} 
						
					gamelist = strtouse;
					
				}
				showeventdata[i].gamesplayed = gamelist;

				
				var prettyresults = "";
				//loop through competitors and grab stats
				for (j=0;j<eventcargo[i].Competitors.length;j++) {
					var uuid = eventcargo[i].Competitors[j].Player.UUID;
					if (-1 == listofuniqueplayers.indexOf(uuid)) {
						listofuniqueplayers.push(uuid);
					}
					
					var playername = 	eventcargo[i].Competitors[j].Player.Firstname + " "  +
								eventcargo[i].Competitors[j].Player.Surname;
								
					var status = eventcargo[i].Competitors[j].Result.Result;
					
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
				showeventdata[i].results = prettyresults;
				startdisplay = moment(showeventdata[i].Start).utc();
				stopdisplay = moment(showeventdata[i].Stop).utc();

				showeventdata[i].Start = startdisplay.format("YYYY-MM-DD HH:mm UTC");
				showeventdata[i].Stop = stopdisplay.format("YYYY-MM-DD HH:mm UTC");
				
			}
			
            return true;                   
        },
		
		
		getPrettyEventStats: function() {
            return showeventdata;
		},
		
		getNumberEvents: function() {
            return eventcargo.length;
		},
		
		getNumberCompetitors: function() {
			
			
			return listofuniqueplayers.length;
			
		}
	

    }
	    

});

