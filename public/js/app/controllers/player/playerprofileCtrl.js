

app.controller( 'PlayerProfileCtrl' , ['$scope', '$http', 'myHttpFactory', function( $scope , $http, myHttpFactory) {
	if (typeof(playerUUID) == 'undefined') {
		alert("plaerUUID is undefined");
		return;
	} 
	
	myHttpFactory.getPlayer(playerUUID).then(function(data) {
				
		//alert("HERE" + JSON.stringify(data));
		$scope.playerprofile = data;
	});
	
	$scope.update = function( uuid ) { 
		alert("HEER");
	
		
 	}
  	
}]);