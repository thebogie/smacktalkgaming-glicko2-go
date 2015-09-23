app.directive('googleplace' , ['event' , '$http',  function(event, $http) {
    return {
        require: 'ngModel',
        link: function(scope, element, attrs, model) {
            var options = {
                types: []
            };
            scope.gPlace = new google.maps.places.Autocomplete(element[0], options);

            google.maps.event.addListener(scope.gPlace, 'place_changed', function() {
                scope.$apply(function() {
					var place = scope.gPlace.getPlace();
					
					var listObj = {
						Locationname: element.val(),
						Locationlng: place.geometry.location.lng().toString(),
						Locationlat: place.geometry.location.lat().toString(),
						Locationtz: 0, 
					};

					
					event.addList(listObj, "locationList" );
					
                    model.$setViewValue(element.val()); 
					
                });
            });
        }
    };
}]);

app.factory('event', function () {
    var userName = "John Doe";

	var possibleresults = [ "DEMOLISH", "WON" , "TIE", "LOST", "SKUNK", "DROP", "QUIT", "RAGE QUIT" ];
	var possibleplaces = [ "1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20" ];
	
	
    // Shared Models
	var locationList = [1];
	var playerList = [];
	var gameList = [];
	var playedinList = [];
	var startDateTime = "";
	var stopDateTime = "";
	
	
    return {
		
		
		
        getPossibleResults: function () {
             return possibleresults;                   
        },
		
        getPossiblePlaces: function () {
             return possibleplaces;                   
        },
		
        getUserName: function () {
             return userName;                   
        },
		setPlayerResult: function (uuid, result) {
			
			for (var i = 0; i < playerList.length; i++) {
				
				if (playerList[i].UUID == uuid ) {
					
					//playerList[i].playedin["Result"] = "";
					playerList[i].playedin["Result"] = result;
					
					
					break;
				}
						//Do something
			}  
			//this.orderList("playerList");
            return true;                   
        },
		setPlayerPlace: function (uuid, place) {
			
			for (var i = 0; i < playerList.length; i++) {
				if (playerList[i].UUID == uuid ) {
					playerList[i].playedin["Place"] = place;
				}
						//Do something
			}    
			this.orderList("playerList");
			return true;			
        },
		
		//TODO for playerlist only right now
		orderList: function(listType){
			function compare(a,b) {
				if (parseInt(a.playedin["Place"]) < parseInt(b.playedin["Place"]))
					return -1;
				if (parseInt(a.playedin["Place"]) >= parseInt(b.playedin["Place"]))
					return 1;
				return 0;
			}	
			playerList.sort(compare);
			
			//for (var i = 0; i < playerList.length; i++) {
			//		playerList[i].playedin["Place"] = (i+1).toString();
						//Do something
			//}  
			
			
		},
		getEvent: function () {
			return {"Numplayers":(playerList.length).toString(), "Start": startDateTime, "Stop": stopDateTime};
			
		},
		getList: function (listType) {
			switch (listType) {
			  case "gameList":

			    return gameList;
			    break;
			  case "playerList":

			    return playerList;
			    break;
			  case "locationList":

			    return locationList;
			    break;
			  case "playedinList":
				playedinList = [];
				for (var i = 0; i < playerList.length; i++) {
					console.log("name" + playerList[i].Firstname);
					console.log("Result" + playerList[i].playedin["Result"]);
					console.log("Place" + playerList[i].playedin["Place"]);
				
					var playedinObj = { Result:"", Place:""};
					playedinObj.Result = playerList[i].playedin["Result"];
					playedinObj.Place = playerList[i].playedin["Place"];
					playedinList.push(playedinObj);
				}
			    return playedinList;
			    break;
			  default:
			   
			}                  
        },
		addList: function(listObj, listType) {
			switch (listType) {
			  case "gameList":
			  
			    console.log("GAMELIST LEGNGHT: " , gameList.length);
				var objalreadypresent  = false;
				
				for (var i = 0; i < gameList.length; i++) {
						if (gameList[i].UUID == listObj.UUID ) {
							objalreadypresent = true;
						}
						//Do something
				}
				
			   	if (objalreadypresent == false) {
			   		gameList.push(listObj);
				}
			   
				
			    break;
			  case "playerList":
				var objalreadypresent  = false;
				for (var i = 0; i < playerList.length; i++) {
				    if (playerList[i].UUID == listObj.UUID ) {
						objalreadypresent = true;
					
					}
				   
				}
			   	if (objalreadypresent == false) {
					
					listObj["playedin"] = {
						Result: (playerList.length == 0 ? 'WON' : 'LOST'), 
						Place: (playerList.length+1).toString()};
			   		playerList.push(listObj);
				}
			   
				this.orderList("playerList");
			    break;
				
			  case "locationList":
		
				locationList[0] = listObj;

			    break;
			  default:
			    console.log("Sorry, we are out of " + expr + ".");
			}   
			
		},
		
		setStartDate: function (startdate) {

			//TODO: validate form
			startDateTime = startdate;
			console.log("SETSTARTDATE: " , startdate);
					
		},
		setStopDate: function (stopdate) {
			//TODO: validate form
			stopDateTime = stopdate;

			console.log("SETSTOPDATE: " , stopdate);
					
		},
		
		convertDate: function (state, datecargo) {
				//convert the date from local area to neo4j format:
				//2014-09-27T21:00:00-05:00
			console.log("CONVERTDATE: " , datecargo);
			
			var finalDateStr = datecargo.year + "-";
			var monthedit = datecargo.month;
			if (datecargo.month < 10) {
				monthedit = "0" + datecargo.month;
			}
			finalDateStr = finalDateStr + monthedit + "-";
			
			var dayedit = datecargo.day;
			if (datecargo.day < 10) {
				dayedit = "0" + datecargo.day;
			}
			finalDateStr = finalDateStr + dayedit + "T";
			
			var houredit = datecargo.hour;
			if (datecargo.hour < 10) {
				houredit = "0" + datecargo.hour;
			}
			finalDateStr = finalDateStr + houredit + ":";
			
			var minuteedit = datecargo.minute;
			if (datecargo.minute < 10) {
				minuteedit = "0" + datecargo.minute;
			}
			finalDateStr = finalDateStr + minuteedit + ":00";
			
			var offsetmintues = "00";
			var offsetedit = datecargo.offset;
			//negative
			var offsetsymbol = "+";
			if (datecargo.offset < 0) {
				offsetsymbol = "-";
			}
			finalDateStr = finalDateStr + offsetsymbol;
			
			//pesky 30 minutes...
			if ( (datecargo.offset % 2) == .5) {
				offsetedit = Math.abs(datecargo.offset) - .5
				offsetmintues = "30";
			}
			
			//pad digits
			if (Math.abs(datecargo.offset) < 10) {
				offsetedit = "0" + Math.abs(datecargo.offset);
			}
			finalDateStr = finalDateStr + offsetedit + ":" + offsetmintues ;
			
			if (state == "start") {
				startDateTime = finalDateStr;
			} 
			if (state == "stop") {
				stopDateTime = finalDateStr;
			} 
			
			return finalDateStr;
		},
		removeList: function(uuid, listType) {
			switch (listType) {
 			case "locationList":

				locationList[0] = null;

			   
			    break;
				
 			case "gameList":

				for (var i = 0; i < gameList.length; i++) {
				    if (gameList[i].UUID == uuid ) {
						gameList.splice(i,1);
						break;
					}
				    //Do something
				}

			   
			    break;
			  case "playerList":

				for (var i = 0; i < playerList.length; i++) {
				    if (playerList[i].UUID == uuid ) {
						playerList.splice(i,1);
						break;
					}
				    //Do something
				}

			   
			    break;
			  default:
			    console.log("Sorry, we are out of " + expr + ".");
			}   
			
		}

    }
	    

});

