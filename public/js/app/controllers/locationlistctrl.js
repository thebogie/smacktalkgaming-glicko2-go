

app.controller( 'LocationListCtrl' , ['$scope', 'Lists', function( $scope , Lists) {

 	$scope.permlist = [];
 	$scope.activelist = {};
	$scope.index = 0;
	

    $scope.gPlace;
	
	
	$scope.$on('HideDIVEventFromEventCtrl',function(event, data){	
          $scope.hideDIV = data;
    });
	
	$scope.updateLocationList = function () {
		console.log("updateLocationList: ");
	};

	$scope.updateList =  function( item ) { 
		console.log("LIST: " + JSON.stringify(Lists.getList("locationList")));
		return $scope.permlist;
	};

  	$scope.removeItem = function( location ) { 
		Lists.removeList(location , "locationList");
		$scope.activelist = Lists.getList("locationList");
 	};

//lists.set =  $scope.activelist;

}]);