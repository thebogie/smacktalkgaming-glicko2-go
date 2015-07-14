//set this from the golang variables
//var URLplayerslistautocomplete;
//var URLgameslistautocomplete;
//var URLeventstatus;

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
	
		httpstartEvent: function(eventcargo) {
			console.log("eventcargo:" + JSON.stringify(eventcargo));
			return $http({method: 'POST', 
					url: '/events/start',  
					data:$filter('json')(eventcargo) 
					}).then(function(data, status, headers, config) {
					 return data;
			});
	   } ,
		httpcommitEvent: function(eventcargo) {
				
			return $http({method: 'POST', 
					url: '/events/commit',  
					data:$filter('json')(eventcargo) 
					}).then(function(data, status, headers, config) {
					 return data;
			});
	   }
	}
}]);