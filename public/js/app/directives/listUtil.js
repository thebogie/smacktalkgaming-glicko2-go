app.directive('googleplace' , ['listUtil' , '$http',  function(listUtil, $http) {
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
						LocationUTCoffset: 0, 
					};

					
					listUtil.addList(listObj, "locationList" );
					//console.log("WORK?" + JSON.stringify(listUtil.getList("locationList")  ) ) ;
                    model.$setViewValue(element.val()); 
					
                });
            });
        }
    };
}]);

app.factory('listUtil', function () {
    var userName = "John Doe";

    // Shared Models
	var locationList = [1];
	var playerList = [];
	var gameList = [];
	
	
    return {
        getUserName: function () {
             return userName;                   
        },
		setPlayerResult: function (uuid, result) {
			for (var i = 0; i < playerList.length; i++) {
				if (playerList[i].UUID == uuid ) {
					playerList[i].playedin["Result"] = result;
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
				if (a.playedin["Place"] < b.playedin["Place"])
					return -1;
				if (a.playedin["Place"] > b.playedin["Place"])
					return 1;
				return 0;
			}	
			playerList.sort(compare);
			
			//for (var i = 0; i < playerList.length; i++) {
			//		playerList[i].playedin["Place"] = (i+1).toString();
						//Do something
			//}  
			
			
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

