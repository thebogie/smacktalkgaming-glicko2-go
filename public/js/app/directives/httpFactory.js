//set this from the golang variables
URLplayerslistautocomplete = typeof(URLplayerslistautocomplete) == 'undefined' ? 0 : URLplayerslistautocomplete;
URLgameslistautocomplete = typeof(URLgameslistautocomplete) == 'undefined' ? 0 : URLgameslistautocomplete;

URLeventstatus = typeof(URLeventstatus) == 'undefined' ? 0 : URLeventstatus;
URLlasteventstatus = typeof(URLlasteventstatus) == 'undefined' ? 0 : URLlasteventstatus;

URLplayerstatus = typeof(URLplayerstatus) == 'undefined' ? 0 : URLplayerstatus;
URLplayeroverallstats = typeof(URLplayeroverallstats) == 'undefined' ? 0 : URLplayeroverallstats;
URLplayerlastlocation = typeof(URLplayerlastlocation) == 'undefined' ? 0 : URLplayerlastlocation;

URLgamesadd = typeof(URLgamesadd) == 'undefined' ? 0 : URLgamesadd;

URLstats = typeof(URLstats) == 'undefined' ? 0 : URLstats;


//google keys
timezoneMapKey = "&key=AIzaSyCXSL3n9tI-VTgRJOhXqJJJ42D1FO1EGBE";
geocodeMapKey = "&AIzaSyBvMnC_gxM_viymDC-Et4Jfr9UEMO9l-Hg";
//
//var ;
//var ;

app.factory('myHttpFactory', ['$http' ,'$filter',function($http, $filter) {
   return {
    	getPlayersAutoComplete: function(item) {
       //since $http.get returns a promise,
       //and promise.then() also returns a promise
       //that resolves to whatever value is returned in it's 
       //callback argument, we can return that.

       		return $http.get(URLplayerslistautocomplete.replace("<nil>", item)).then(function(result) {
		
           return result.data;
       	});
     } , 
		addGame: function(game) {
			console.log("game", game);
			console.log("URLgameadd", URLgamesadd.replace("<nil>", game));
			return $http.get(URLgamesadd.replace("<nil>", game)).then(function(result) {
				return result.data;
			});
     } ,
	 
	 	getGamesAutoComplete: function(item) {
       return $http.get(URLgameslistautocomplete.replace("<nil>", item)).then(function(gameresult) {
           return gameresult.data;
       });
     } ,
	  	getEvent: function(item) {
			return $http.get(URLeventstatus.replace("<nil>", item)).then(function(result) {
				return result.data;
			});
     } ,
	   	updatePlayer: function(updateplayercargo) {
			console.log("updateplayercargo:" + (updateplayercargo));

			return $http({method: 'POST', 
					url: '/players/update',  
					data:$filter('json')(updateplayercargo) 
					}).then(function(data, status, headers, config) {
					 return data;
			}); 
	 },
	   	getLastLocation: function(item) {
			console.log(item);
			return $http.get(URLplayerlastlocation.replace("<nil>", item)).then(function(result) {
				return result.data;
			});
	 },
	   	getPlayer: function(item) {
			//alert(item);
			return $http.get(URLplayerstatus.replace("<nil>", item)).then(function(result) {
				return result.data;
			});
	 },
	   	getPlayerOverallStats: function(item) {
			//alert(item);
			return $http.get(URLplayeroverallstats.replace("<nil>", item)).then(function(result) {
				return result.data;
			});
	 },
		httpGrabStats: function(cargo) {
			console.log("GrabStatsCargo:", cargo, URLstats);
			return $http.get(URLstats.replace("<nil>/<nil>", cargo)).then(function(result) {
				return result.data;
			});
	   },
		
		httpcommitEvent: function(eventcargo) {
			console.log("EVENTCARGO:", eventcargo);
			return $http({method: 'POST', 
					url: '/events/commit',  
					data:$filter('json')(eventcargo) 
					}).then(function(data, status, headers, config) {
					 return data;
			});
	   },
	   	httpTimeZoneOffset: function(mapstring) {
			//alert(item);
			return $http.get(mapstring).then(function(result) {
				console.log("fromfactoryresult:", result.data);
				return result.data.timeZoneId;
				//return parseInt(result.data.dstOffset + result.data.rawOffset) / 60;
			});
	   } ,
	   	httpGoogleAPITimeZoneOffset: function(mapstring) {

			return $http({method: 'POST', 
					url: '/stats/gapitimez/',  
					data:$filter('json')(mapstring) 
					}).then(function(data, status, headers, config) {
					 return data;
			}); 
	 }
	}
}]);