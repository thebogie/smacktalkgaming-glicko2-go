
app.controller( 'GameListCtrl' , ['$scope', '$http', 'Lists', 'myHttpFactory', function( $scope , $http, Lists, myHttpFactory) {

 	$scope.activelist = [];

	$scope.updateList = function(item) {

		 return myHttpFactory.getGamesAutoComplete(item).then(function(data) {
			for (x in data) {
    			//console.log("HERE" +  JSON.stringify(data[x]) + x) ;
				data[x]["selectName"] = data[x]["Name"] + " (" + data[x]["Published"] + ")";
			}
			//console.log("DATA" + JSON.stringify(data));
		    //return [{"Name":"Carcassonne","Published":"2000","UUID":"e3ed6302-f868-437f-b754-6646061c1363"},{"Name":"Caesar & Cleopatra","Published":"1997","UUID":"77b3f5f2-d710-469b-b4c8-cefa5b2974ff"}];
			return data;
		 });	
  	};

	$scope.onSelect = function ($item, $model, $label) {
		Lists.addList($model , "gameList");
		$scope.activelist = Lists.getList("gameList");
		//console.log("LIST: " + JSON.stringify(Lists.getList("gameList")));
	};

  	$scope.removeItem = function( uuid ) { 
		//console.log("removeItem" + uuid);
		Lists.removeList(uuid , "gameList");
		$scope.activelist = Lists.getList("gameList");

 	}
	
	$scope.$on('HideDIVEventFromEventCtrl',function(event, data){
          $scope.hideDIV = data;
    });

}]);