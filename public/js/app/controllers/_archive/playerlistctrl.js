app.controller( 'PlayerListCtrl' , ['$scope', '$http', 'Lists', 'myHttpFactory', function( $scope , $http, Lists, myHttpFactory) {

  	$scope.activelist = [];
	$scope.numplayers = 0;
	
	//event to hide div
	$scope.$on('HideDIVEventFromEventCtrl',function(event, data){
          $scope.hideDIV = data;
    });
	
	$scope.onSelect = function ($item, $model, $label) {
		Lists.addList($model , "playerList");
		$scope.activelist = Lists.getList("playerList");
		//console.log("LIST: " + JSON.stringify(Lists.getList("playerList")));
		$scope.numplayers = $scope.activelist .length;
	};
	
	$scope.updateList =  function( item ) { 
	
	   return myHttpFactory.getPlayersAutoComplete(item).then(function(data) {

 			for (x in data) {
    			console.log("HERE" +  JSON.stringify(data[x]) + x) ;
				//TODO add nickname
				var nicknameexist = data[x]["Nickname"];
				var nickname = " ";
				
				if (nicknameexist) {
					nickname = ' "' + nicknameexist + '" ';
					
				}  
				data[x]["selectName"] = data[x]["Firstname"] + nickname + data[x]["Surname"];
			}
			//console.log("DATA" + JSON.stringify(data));
		    return data;
		 });	
  	};
	
  	$scope.removeItem = function( uuid ) { 
		//console.log("removeItem" + uuid);
		Lists.removeList(uuid , "playerList");
		$scope.activelist = Lists.getList("playerList");
		$scope.numplayers = $scope.activelist.length;
	
		
 	}


}]);