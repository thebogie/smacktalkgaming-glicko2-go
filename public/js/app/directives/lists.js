app.directive('googleplace' , ['Lists',  function(Lists) {
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
					};
					
					Lists.addList(listObj, "locationList" );
					console.log("WORK?" + JSON.stringify(Lists.getList("locationList")  ) ) ;
                    model.$setViewValue(element.val());                
                });
            });
        }
    };
}]);

app.factory('Lists', function () {
    var userName = "John Doe";

    // Shared Models
	var locationList = [];
	var playerList = [];
	var gameList = [];
	
	
    return {
        getUserName: function () {
             return userName;                   
        },
		setPlayerResult: function (index, result) {
			playerList[index].playedin["Result"] = result;
            return true;                   
        },
		setPlayerPlace: function (index, place) {
			playerList[index].playedin["Place"] = place;
            return true;                   
        },
		
		//
		//TODO for playerlist only right now
		orderList: function(listType){
			
			for (var i = 0; i < playerList.length; i++) {
			
				if (i == 0) {
					//playerList[i].playedin["Result"] = "WON";
				}
				playerList[i].playedin["Place"] = (i+1).toString();
				   
			}
			
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
						Result:"", 
						Place:""};
			   		playerList.push(listObj);
				}
			   
				this.orderList("playerList");
			    break;
				
			  case "locationList":
		console.log("listobj" + JSON.stringify(listObj));
				var objalreadypresent  = false;
				for (var i = 0; i < locationList.length; i++) {
				    if (locationList[i].UUID == listObj.UUID ) {
						objalreadypresent = true;
					}
				    //Do something
				}
			   	if (objalreadypresent == false) {
			   		locationList.push(listObj);
				}
			
				
			   
				
			    break;
			  default:
			    console.log("Sorry, we are out of " + expr + ".");
			}   
			
		},
		removeList: function(uuid, listType) {
			switch (listType) {
 			case "locationList":

				locationList = [];

			   
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

