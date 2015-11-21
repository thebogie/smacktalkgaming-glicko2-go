//app.controller( 'SmackTalkAppCtrl' , ['$scope', '$http', 'myHttpFactory', function( $scope , $http, myHttpFactory) {
app.controller('SmackTalkAppCtrl', ['$scope', '$mdBottomSheet','$mdSidenav', '$mdDialog', function($scope, $mdBottomSheet, $mdSidenav, $mdDialog){
	$scope.showRecordEvent = false;
	$scope.showDashboard = true;
	$scope.showAddGames = false;
	$scope.showOpponents = false;
	$scope.showEvents = false;
	$scope.showGames = false;
	
	
	$scope.menuEvent = function(action) {
		console.log("menuEvent", action);
		
		switch(action) {
			case "openRecordEvent":
			$scope.showRecordEvent = true;
			$scope.showDashboard = false;
			$scope.showAddGames = false;
			$scope.showOpponents = false;
			$scope.showEvents = false;
			$scope.showGames = false;
			break;
			
			case "openProfile":
			$scope.showRecordEvent = false;
			$scope.showDashboard = true;
			$scope.showAddGames = false;
			$scope.showOpponents = false;
			$scope.showEvents = false;
			$scope.showGames = false;
			break;
			
			case "openOpponents":
			$scope.showRecordEvent = false;
			$scope.showDashboard = false;
			$scope.showAddGames = false;
			$scope.showOpponents = true;
			$scope.showEvents = false;
			$scope.showGames = false;
			break;
			
			
			case "openEvents":
			$scope.showRecordEvent = false;
			$scope.showDashboard = false;
			$scope.showAddGames = false;
			$scope.showOpponents = false;
			$scope.showEvents = true;
			$scope.showGames = false;
			break;
			
			case "openAddGames":
			$scope.showRecordEvent = false;
			$scope.showDashboard = false;
			$scope.showAddGames = true;
			$scope.showOpponents = false;
			$scope.showEvents = false;
			$scope.showGames = false;
			break;

			case "openGames":

			$scope.showRecordEvent = false;
			$scope.showDashboard = false;
			$scope.showAddGames = false;
			$scope.showOpponents = false;
			$scope.showEvents = false;
			$scope.showGames = true;
			break;
				
		}
		
		//close the menu
		 $mdSidenav('left').toggle();
	};
	
	
	
  $scope.toggleSidenav = function(menuId) {
	  console.log("MENUID", menuId);
    $mdSidenav(menuId).toggle();
  };
    $scope.menu = [
    {
      link : '',
      title: 'Show Profile',
      icon: 'person',
	  action: 'openProfile'
    },
	{
      link : '',
      title: 'Show Opponents',
      icon: 'people',
	  action: 'openOpponents'
    },
	{
      link : '',
      title: 'Show Events',
      icon: 'event',
	  action: 'openEvents'
    },
	{
      link : '',
      title: 'Show Games',
      icon: 'extension',
	  action: 'openGames'
    },
    {
      link : '',
      title: 'Record Event',
      icon: 'star',
	  action: 'openRecordEvent'
    }
  ];
  $scope.admin = [
    {
      link : '',
      title: 'Add Games',
      icon: 'exposure_plus_1',
	  action: 'openAddGames'
    }
  ];
  $scope.activity = [
      {
        what: 'Brunch this weekend?',
        who: 'Ali Conners',
        when: '3:08PM',
        notes: " I'll be in your neighborhood doing errands"
      },
      {
        what: 'Summer BBQ',
        who: 'to Alex, Scott, Jennifer',
        when: '3:08PM',
        notes: "Wish I could come out but I'm out of town this weekend"
      },
      {
        what: 'Oui Oui',
        who: 'Sandra Adams',
        when: '3:08PM',
        notes: "Do you have Paris recommendations? Have you ever been?"
      },
      {
        what: 'Birthday Gift',
        who: 'Trevor Hansen',
        when: '3:08PM',
        notes: "Have any ideas of what we should get Heidi for her birthday?"
      },
      {
        what: 'Recipe to try',
        who: 'Brian Holt',
        when: '3:08PM',
        notes: "We should eat this: Grapefruit, Squash, Corn, and Tomatillo tacos"
      },
    ];
  $scope.alert = '';
  $scope.showListBottomSheet = function($event) {
    $scope.alert = '';
    $mdBottomSheet.show({
      template: '<md-bottom-sheet class="md-list md-has-header"> <md-subheader>Settings</md-subheader> <md-list> <md-item ng-repeat="item in items"><md-item-content md-ink-ripple flex class="inset"> <a flex aria-label="{{item.name}}" ng-click="listItemClick($index)"> <span class="md-inline-list-icon-label">{{ item.name }}</span> </a></md-item-content> </md-item> </md-list></md-bottom-sheet>',
      controller: 'ListBottomSheetCtrl',
      targetEvent: $event
    }).then(function(clickedItem) {
      $scope.alert = clickedItem.name + ' clicked!';
    });
  };
  
  $scope.showAdd = function(ev) {
    $mdDialog.show({
      controller: DialogController,
      template: '<md-dialog aria-label="Mango (Fruit)"> <md-content class="md-padding"> <form name="userForm"> <div layout layout-sm="column"> <md-input-container flex> <label>First Name</label> <input ng-model="user.firstName" placeholder="Placeholder text"> </md-input-container> <md-input-container flex> <label>Last Name</label> <input ng-model="theMax"> </md-input-container> </div> <md-input-container flex> <label>Address</label> <input ng-model="user.address"> </md-input-container> <div layout layout-sm="column"> <md-input-container flex> <label>City</label> <input ng-model="user.city"> </md-input-container> <md-input-container flex> <label>State</label> <input ng-model="user.state"> </md-input-container> <md-input-container flex> <label>Postal Code</label> <input ng-model="user.postalCode"> </md-input-container> </div> <md-input-container flex> <label>Biography</label> <textarea ng-model="user.biography" columns="1" md-maxlength="150"></textarea> </md-input-container> </form> </md-content> <div class="md-actions" layout="row"> <span flex></span> <md-button ng-click="answer(\'not useful\')"> Cancel </md-button> <md-button ng-click="answer(\'useful\')" class="md-primary"> Save </md-button> </div></md-dialog>',
      targetEvent: ev,
    })
    .then(function(answer) {
      $scope.alert = 'You said the information was "' + answer + '".';
    }, function() {
      $scope.alert = 'You cancelled the dialog.';
    });
  };
}]);

app.controller('ListBottomSheetCtrl', function($scope, $mdBottomSheet) {
  $scope.items = [
    { name: 'Share', icon: 'share' },
    { name: 'Upload', icon: 'upload' },
    { name: 'Copy', icon: 'copy' },
    { name: 'Print this page', icon: 'print' },
  ];
  
  $scope.listItemClick = function($index) {
    var clickedItem = $scope.items[$index];
    $mdBottomSheet.hide(clickedItem);
  };
});

function DialogController($scope, $mdDialog) {
  $scope.hide = function() {
    $mdDialog.hide();
  };
  $scope.cancel = function() {
    $mdDialog.cancel();
  };
  $scope.answer = function(answer) {
    $mdDialog.hide(answer);
  };
};

app.directive('userAvatar', function() {
  return {
    replace: true,
    template: '<svg class="user-avatar" viewBox="0 0 128 128" height="64" width="64" pointer-events="none" display="block" > <path fill="#FF8A80" d="M0 0h128v128H0z"/> <path fill="#FFE0B2" d="M36.3 94.8c6.4 7.3 16.2 12.1 27.3 12.4 10.7-.3 20.3-4.7 26.7-11.6l.2.1c-17-13.3-12.9-23.4-8.5-28.6 1.3-1.2 2.8-2.5 4.4-3.9l13.1-11c1.5-1.2 2.6-3 2.9-5.1.6-4.4-2.5-8.4-6.9-9.1-1.5-.2-3 0-4.3.6-.3-1.3-.4-2.7-1.6-3.5-1.4-.9-2.8-1.7-4.2-2.5-7.1-3.9-14.9-6.6-23-7.9-5.4-.9-11-1.2-16.1.7-3.3 1.2-6.1 3.2-8.7 5.6-1.3 1.2-2.5 2.4-3.7 3.7l-1.8 1.9c-.3.3-.5.6-.8.8-.1.1-.2 0-.4.2.1.2.1.5.1.6-1-.3-2.1-.4-3.2-.2-4.4.6-7.5 4.7-6.9 9.1.3 2.1 1.3 3.8 2.8 5.1l11 9.3c1.8 1.5 3.3 3.8 4.6 5.7 1.5 2.3 2.8 4.9 3.5 7.6 1.7 6.8-.8 13.4-5.4 18.4-.5.6-1.1 1-1.4 1.7-.2.6-.4 1.3-.6 2-.4 1.5-.5 3.1-.3 4.6.4 3.1 1.8 6.1 4.1 8.2 3.3 3 8 4 12.4 4.5 5.2.6 10.5.7 15.7.2 4.5-.4 9.1-1.2 13-3.4 5.6-3.1 9.6-8.9 10.5-15.2M76.4 46c.9 0 1.6.7 1.6 1.6 0 .9-.7 1.6-1.6 1.6-.9 0-1.6-.7-1.6-1.6-.1-.9.7-1.6 1.6-1.6zm-25.7 0c.9 0 1.6.7 1.6 1.6 0 .9-.7 1.6-1.6 1.6-.9 0-1.6-.7-1.6-1.6-.1-.9.7-1.6 1.6-1.6z"/> <path fill="#E0F7FA" d="M105.3 106.1c-.9-1.3-1.3-1.9-1.3-1.9l-.2-.3c-.6-.9-1.2-1.7-1.9-2.4-3.2-3.5-7.3-5.4-11.4-5.7 0 0 .1 0 .1.1l-.2-.1c-6.4 6.9-16 11.3-26.7 11.6-11.2-.3-21.1-5.1-27.5-12.6-.1.2-.2.4-.2.5-3.1.9-6 2.7-8.4 5.4l-.2.2s-.5.6-1.5 1.7c-.9 1.1-2.2 2.6-3.7 4.5-3.1 3.9-7.2 9.5-11.7 16.6-.9 1.4-1.7 2.8-2.6 4.3h109.6c-3.4-7.1-6.5-12.8-8.9-16.9-1.5-2.2-2.6-3.8-3.3-5z"/> <circle fill="#444" cx="76.3" cy="47.5" r="2"/> <circle fill="#444" cx="50.7" cy="47.6" r="2"/> <path fill="#444" d="M48.1 27.4c4.5 5.9 15.5 12.1 42.4 8.4-2.2-6.9-6.8-12.6-12.6-16.4C95.1 20.9 92 10 92 10c-1.4 5.5-11.1 4.4-11.1 4.4H62.1c-1.7-.1-3.4 0-5.2.3-12.8 1.8-22.6 11.1-25.7 22.9 10.6-1.9 15.3-7.6 16.9-10.2z"/> </svg>'
  };
});

app.config(function($mdThemingProvider) {
  var customBlueMap =       $mdThemingProvider.extendPalette('light-blue', {
    'contrastDefaultColor': 'light',
    'contrastDarkColors': ['50'],
    '50': 'ffffff'
  });
  $mdThemingProvider.definePalette('customBlue', customBlueMap);
  $mdThemingProvider.theme('default')
    .primaryPalette('customBlue', {
      'default': '500',
      'hue-1': '50'
    })
    .accentPalette('pink');
  $mdThemingProvider.theme('input', 'default')
        .primaryPalette('grey')
});


app.controller( 'OverAllStatsCtrl' , ['$scope', 'stats', function( $scope, stats ) {

	stats.setEventcargo(eventcargo);
	
	$scope.eventsinfo = {
		'total': stats.getNumberEvents(),
		'title': "EVENTS",
	};
	$scope.gamesinfo = {
		'total': gamesList.length,
		'title': "GAMES",
	};
	$scope.playersinfo = {
		'total': stats.getNumberCompetitors(),
		'title': "PLAYERS",
	};
	$scope.playersrating = {
		'rating': parseInt(playerrating.Rating),
		'ratingdeviation': parseInt(playerrating.RatingDeviation),
		'volatility': parseFloat(playerrating.Volatility).toFixed(3),
		
	}
	
	
	$scope.stats = {
		won: {
			'total': 0,
			'title': "WON or DEMOLISH",
			'DEMOLISH': 0,
			'WON': 0,
		},
		lost: {
			'total': 0,
			'title': "LOST or SKUNK",
			'LOST': 0,
			'SKUNK': 0,
		},
		other: {
			'total': 0,
			'title': "TIE or DROP or QUIT",
			'TIE': 0,
			'DROP': 0,
			'QUIT': 0,
			'RAGEQUIT': 0,
		}
		
	};
	
	for (var i = 0; i < playedinsList.length; i++) {
		switch(playedinsList[i].Result) {
			case "DEMOLISH":
				$scope.stats.won.total++;
				$scope.stats.won.DEMOLISH.total++;
				break;
			case "WON":
				$scope.stats.won.total++;
				$scope.stats.won.WON.total++;
				break;
			case "TIE":
				$scope.stats.other.total++;
				$scope.stats.other.TIE.total++;
				break;
			case "LOST":
				$scope.stats.lost.total++;
				$scope.stats.lost.LOST.total++;
				break;
			case "SKUNK":
				$scope.stats.lost.total++;
				$scope.stats.lost.SKUNK.total++;
				break;
			case "DROP":
				$scope.stats.other.total++;
				$scope.stats.other.DROP.total++;
				break;
			case "QUIT":
				$scope.stats.other.total++;
				$scope.stats.other.QUIT.total++;
				break;
			case "RAGE QUIT":
				$scope.stats.other.total++;
				$scope.stats.other.RAGEQUIT.total++;
				break;
				
		}
	}
	
	
	$scope.overallstats = {
		eventsinfo: $scope.eventsinfo,
		gamesinfo: $scope.gamesinfo,
		playersinfo: $scope.playersinfo,
		playersrating: $scope.playersrating,
		stats: $scope.stats,
	
	};

}]);

app.controller('PlayersProfileCtrl', ['$scope', 'myHttpFactory', 'stats', function($scope, myHttpFactory, stats) {

	players = [];
	
	$scope.showplayers = stats.getPrettyPlayerStats();
	$scope.selected = [];

	  $scope.query = {
		filter: '',
		order: '-rating',
		limit: 10,
		page: 1
	  };

	  
	  $scope.compare = 		function compare(a,b) {
				console.log("HERE",a.Eventname,b.Eventname);
				if (a.name < b.name)
					return -1;
				if (a.name > b.name)
					return 1;
				return 0;
		};

	  // in the future we may see a few built in alternate headers but in the mean time
	  // you can implement your own search header and do something like
	  $scope.search = function (predicate) {
		console.log("search");
		//$scope.filter = predicate;
		//$scope.deferred = $nutrition.desserts.get($scope.query, success).$promise;
	  };

	  $scope.onOrderChange = function (order) {
		console.log("onOrderChange", order );
	    //$scope.showevents = "fish";

		//return $nutrition.desserts.get($scope.query, success).$promise; 
	  };

	  $scope.onPaginationChange = function (page, limit) {
		console.log("onPaginationChange", page , limit);
		//return eventsList;
	  };
}]);

app.controller('GamesProfileCtrl', ['$scope', 'myHttpFactory', 'stats', function($scope, myHttpFactory, stats) {

	games = [];
	
	$scope.showgames = stats.getPrettyGameStats();
	console.log("showgames", $scope.showgames);
	$scope.selected = [];

	  $scope.query = {
		filter: '',
		order: '-WON',
		limit: 10,
		page: 1
	  };

	  
	  $scope.compare = 		function compare(a,b) {
				//console.log("HERE",a.Eventname,b.Eventname);
				if (a.name < b.name)
					return -1;
				if (a.name > b.name)
					return 1;
				return 0;
		};

	  // in the future we may see a few built in alternate headers but in the mean time
	  // you can implement your own search header and do something like
	  $scope.search = function (predicate) {
		console.log("search");
		//$scope.filter = predicate;
		//$scope.deferred = $nutrition.desserts.get($scope.query, success).$promise;
	  };

	  $scope.onOrderChange = function (order) {

	    //$scope.showevents = "fish";

		//return $nutrition.desserts.get($scope.query, success).$promise; 
	  };

	  $scope.onPaginationChange = function (page, limit) {
		console.log("onPaginationChange", page , limit);
		//return eventsList;
	  };
}]);



app.controller('EventsProfileCtrl', ['$scope', 'myHttpFactory', 'stats', function($scope, myHttpFactory, stats) {

	events = [];
	
	$scope.showevents = stats.getPrettyEventStats();
	console.log("showevents", $scope.showevents);
	$scope.selected = [];

	  $scope.query = {
		filter: '',
		order: '-Start',
		limit: 10,
		page: 1
	  };

	  
	  $scope.compare = 		function compare(a,b) {
			//console.log("HERE",a.Eventname,b.Eventname);
				if (a.Eventname < b.Eventname)
					return -1;
				if (a.Eventname > b.Eventname)
					return 1;
				return 0;
		};

	  // in the future we may see a few built in alternate headers but in the mean time
	  // you can implement your own search header and do something like
	  $scope.search = function (predicate) {
		console.log("search");
		//$scope.filter = predicate;
		//$scope.deferred = $nutrition.desserts.get($scope.query, success).$promise;
	  };

	  $scope.onOrderChange = function (order) {

	    //$scope.showevents = "fish";

		//return $nutrition.desserts.get($scope.query, success).$promise; 
	  };

	  $scope.onPaginationChange = function (page, limit) {
		console.log("onPaginationChange", page , limit);
		//return eventsList;
	  };
}]);



app.controller('ProfileCtrl', ['$scope', '$mdDialog', 'myHttpFactory', function($scope, myHttpFactory, $mdDialog) {


	if (typeof(playerinfo.UUID) == 'undefined') {
		alert("plaerUUID is undefined");
		return;
	} ;
	
	
	$scope.alignments = [
		"Lawful Good",	
		"Neutral Good",
		"Chaotic Good",
		"Lawful Neutral",	
		"(True) Neutral",
		"Chaotic Neutral",
		"Lawful Evil",	
		"Neutral Evil",	
		"Chaotic Evil"
	
	];
	

	$scope.project = {
		nickname: playerinfo.Nickname,
		birthdate: parseInt(playerinfo.Birthdate),
		alignment: playerinfo.Alignment
	};
	
	$scope.openProfileDialog = function (ev) {
		
		
	};
	
	
	$scope.update = function( playerinfodelta ) { 
		var playercargo = playerinfo;
		
				console.log("playerinfodelta", playerinfodelta);
		
		playercargo.Alignment = playerinfodelta.alignment.$viewValue;
		playercargo.Birthdate = playerinfodelta.birthdate.$viewValue;
		playercargo.Nickname = playerinfodelta.nickname.$viewValue;
		
		console.log("Playercargo", playercargo);
		
		myHttpFactory.updatePlayer(playercargo).then(function(results) {
					console.log("data:"  +JSON.stringify(results));

		});
	
		
 	};
   $scope.somethingchanged = function(projectForm) {
   
	   if (permission.readonly == "true") {
			return 0;
	   }
   
		var retval = 0;
		//console.log("alignment", projectForm.alignment);
		//|| projectForm.alignment.$viewValue != playerinfo.Alignment
		if (projectForm.birthdate.$viewValue != playerinfo.Birthdate || 
			projectForm.nickname.$viewValue != playerinfo.Nickname || 
			((typeof(projectForm.alignment) != 'undefined') && projectForm.alignment.$viewValue != playerinfo.Alignment) ) {
	
			retval = 1;
		}
         return retval;   
    };  

}]);

app.controller('CoolStats', ['$scope', 'myHttpFactory', '$mdDialog', 'stats',  function($scope, myHttpFactory , $mdDialog, stats) {



	var nemesis = "FISH";
	//NEMESIS
	
	$scope.shownemeses = false;
	$scope.showdominate = false;

	
	$scope.stats = {
		nemesis: [
			
		],
		
		dominate: [
			
		]
	
	
	};
	
	
	console.log("CoolStats");
	
	//TODO: fix top x. default to 3 for now
	var nemesislist = stats.getNemesis(3);
	$scope.stats.nemesis.push(nemesislist[0]);
	$scope.stats.nemesis.push(nemesislist[1]);
	$scope.stats.nemesis.push(nemesislist[2]);

	
	var domintatinglist = stats.getDominating(3);
	$scope.stats.dominate.push(domintatinglist[0]);
	$scope.stats.dominate.push(domintatinglist[1]);
	$scope.stats.dominate.push(domintatinglist[2]);
	

	$scope.showNemeses = function () {
		$scope.shownemeses = true;
		//$scope.showdominate = false;
	};
	$scope.showDominate = function () {
		//$scope.shownemeses = false;
		$scope.showdominate = true;
	};

	
	    $scope.showAlert = function(ev) {
    // Appending dialog to document.body to cover sidenav in docs app
    // Modal dialogs should fully cover application
    // to prevent interaction outside of dialog
    $mdDialog.show(
      $mdDialog.alert()
        .parent(angular.element(document.querySelector('#popupContainer')))
        .clickOutsideToClose(true)
        .title('This is an alert title')
        .content('You can specify some description text in here.')
        .ariaLabel('Alert Dialog Demo')
        .ok('Got it!')
        .targetEvent(ev)
    );
  };
  
   $scope.showStats = function(ev, choice) {
		
		var dialogURL = "";
		
		switch(choice) {
			case "nemeses":
				dialogURL = "/public/js/app/controllers/stats/nemesis.html";
				break;
			case "dominate":
				dialogURL = "/public/js/app/controllers/stats/dominate.html";
				break;
			
		}
		
        $mdDialog.show({	
			clickOutsideToClose: true,
			targetEvent: ev,
			templateUrl:dialogURL,
			controller: 'CoolStats'
        });
	};
  
  
	//locals: { employee: $scope.stats }
	//parent: angular.element(document.querySelector('#popupContainer')),
	//templateUrl:'/public/js/app/controllers/stats/nemesis.html',
    $scope.showAdvanced = function(ev) {
		
        $mdDialog.show({
			
			clickOutsideToClose: true,
			targetEvent: ev,
			templateUrl:'/public/js/app/controllers/stats/nemesis.html',
			controller: 'CoolStats'
        });
	};
	
	$scope.labels = ["January", "February", "March", "April", "May", "June", "July"];
		  $scope.series = ['Series A', 'Series B'];
		  $scope.data = [
			[65, 59, 80, 81, 56, 55, 40],
			[28, 48, 40, 19, 86, 27, 90]
		  ];
		  $scope.onClick = function (points, evt) {
			console.log(points, evt);
		  };
	
	
	$scope.showRatingHistory = function(ev) {
		console.log("ShowRatingHistory", ev);
		
          $mdDialog.show({	
			clickOutsideToClose: true,
			targetEvent: ev,
			templateUrl:'/public/js/app/controllers/stats/ratinghistory.html',
			controller: 'CoolStats'
        });
		
	};
	
	

}]);





app.controller( 'EventCtrl' ,  ['$scope', 'myHttpFactory' , 'event' , '$window' , function( $scope  , myHttpFactory, event, $window ) {

	console.log("event controller");
	$scope.recordable = {games:false,players:false,location:false};
	$scope.showInputMenu = true;
	
	$scope.showGameSetting = false;
	$scope.showPlayerSetting = false;
	$scope.showLocationSetting = false;
	
	$scope.showRecordButton = false;
	
	$scope.lockSetGames = "#FFF";
	$scope.lockSetPlayers = "#FFF";
	$scope.lockSetLocation = "#FFF";
	
	
	$scope.clickMenu = function(action) {
		console.log("action:", action);
		
		
		switch(action) {
			case "SetGames":
				$scope.showInputMenu = false;
				$scope.showGameSetting = true;
				$scope.showPlayerSetting = false;
				$scope.showLocationSetting = false;
				$scope.showgames = event.getList("gameList");
			break;
			case "SetPlayers":
				$scope.showInputMenu = false;
				$scope.showGameSetting = false;
				$scope.showPlayerSetting = true;
				$scope.showLocationSetting = false;
				$scope.showplayers = event.getList("playerList");
			break;
			case "SetLocation":
				$scope.startdisplay =  event.getEvent().Start; 
				$scope.stopdisplay = event.getEvent().Stop;
				
				console.log("STARTDISPALY", startdisplay, event.getEvent().Start);
				$scope.showlocations = event.getList("locationList");
				$scope.showInputMenu = false;
				$scope.showGameSetting = false;
				$scope.showPlayerSetting = false;
				$scope.showLocationSetting = true;

				
			break;
		
		};
		
	};
	
	
	$scope.cleandate = function(datetouse) {
		console.log("datetouse", datetouse);
		var currentDate = new Date(datetouse);

		var Day = currentDate.getDate();
        if (Day < 10) {
            Day = '0' + Day;
        } //end if
        var Month = currentDate.getMonth() + 1;
        if (Month < 10) {
            Month = '0' + Month;
        } //end if
        var Year = currentDate.getFullYear();
        var fullDate = Month + '/' + Day + '/' + Year;
		
		
		var Minutes = currentDate.getMinutes();
        if (Minutes < 10) {
            Minutes = '0' + Minutes;
        }
        var Hour = currentDate.getHours();
        if (Hour > 12) {
            Hour -= 12;
        } //end if
        var Time = Hour + ':' + Minutes;
        if (currentDate.getHours() <= 12) {
            Time += ' AM';
        } //end if
        if (currentDate.getHours() > 12) {
            Time += ' PM';
        } //end if
		
		
		return fullDate + " " + Time;
	};
	
	$scope.displayInputMenu = function(action) {
		
		
		$scope.showInputMenu = true;
		$scope.showGameSetting = false;
		$scope.showPlayerSetting = false;
		$scope.showLocationSetting = false;
		
		switch(action) {
			case "SetGames":
				$scope.lockSetGames = "#CCC";
				$scope.recordable["games"] = true;
			break;
			case "SetPlayers":
				$scope.lockSetPlayers = "#CCC";
				$scope.recordable["players"] = true;
			break;
			case "SetLocation":
				$scope.lockSetLocation = "#CCC";
				$scope.recordable["location"] = true;
				
				$scope.startdisplay = $scope.cleandate(event.getEvent().Start);
				$scope.stopdisplay = $scope.cleandate(event.getEvent().Stop);
			break;
		
		};

		console.log("recordable?", $scope.recordable);
		if ( Object.keys($scope.recordable).every(function(k){ return $scope.recordable[k] === true })) {
			$scope.showRecordButton = true;
		}
		
	};
	
	$scope.recordIt = function() {
	
		console.log("Record it!");	
			
		var eventload = {
				event: event.getEvent(),
				players: event.getList("playerList"),
				games: event.getList("gameList"),
				locations: event.getList("locationList"),
				playedin: event.getList("playedinList"),
				
					
		};
		//TODO: better place for this...
		eventload.locations[0].Locationtz = eventload.locations[0].Locationtz.toString();
		
		myHttpFactory.httpcommitEvent(eventload).then(function(results) {

		});
		
		//return to the main page
		$window.location.href = "/";
    }
	

	$scope.$on('RecordStage',function(){
		console.log("RecordStage");
		//inform the event to show record button
		$scope.showInputMenu = false;
		$scope.showAskQuestions = false;
		//TODO: check each part that the basics passed
		$scope.showPassedChecks = true;
		
		//console.log("getEvent:", event.getEvent().Start);
		//console.log("moment:", moment(event.getEvent().Start));
		
		$scope.showevent = {
			"start": event.getEvent().Start, 
			"end":event.getEvent().Stop, 
		};
		
		$scope.showgames = event.getList("gameList");
		$scope.showplayers = event.getList("playerList");
		$scope.showlocations = event.getList("locationList");
		//+ angular.toJson(event.getList("gameList")) + angular.toJson(event.getList("playerList")) + angular.toJson(event.getList("locationList")) + angular.toJson(event.getList("playedinList"));
		//+  + ;;
    });

}]);



app.controller( 'GameCtrl' , ['$scope', 'myHttpFactory' , 'event' , function( $scope  , myHttpFactory, event ) {

	
	$scope.selectedItemChange = function(item) {
		
		if (typeof item !== "undefined") { 
			console.log('Item changed to ' + JSON.stringify(item));
			event.addList(item , "gameList");
			console.log("Gamelist:", event.getList("gameList"));
			$scope.gamelist = event.getList("gameList");
			$scope.selectedItem = null;
			$scope.searchText = "";
		}
    }
	
	$scope.searchTextChange = function(item) {
		console.log('Text changed to ' + item);

    }
	
	$scope.addItem = function(item) {
      console.log('addItem ' + item);
    }
	
	$scope.removeItem = function(uuid) {
		console.log('removeItem ' + uuid);
		event.removeList(uuid , "gameList");
		$scope.gamelist = event.getList("gameList");
    }
	
//TODO: fix query to only search when query is one character, then save it for caching.
	$scope.querySearch   = function(query) {
		console.log("query: ", query);
	
		$scope.isLoading=true;
		return myHttpFactory.getGamesAutoComplete(query).then(function(data) {

					return data.filter( 
						function (value) {
							//console.log("Name + query",angular.lowercase(value.Name), angular.lowercase(query), //(angular.lowercase(value.Name)).indexOf(angular.lowercase(query)));
							retval = ((angular.lowercase(value.Name)).indexOf(angular.lowercase(query)) === 0);
							$scope.isLoading=false;
							return retval;
						}
					);
					

		});	
		
	}
	
	
}]);

app.controller( 'PlayerCtrl' , ['$scope', 'myHttpFactory', 'event',  function(   $scope  , myHttpFactory, event ) {
	
	$scope.possibleresults = event.getPossibleResults();
	$scope.possibleplaces = event.getPossiblePlaces();
	
	$scope.updatePlace = function(uuid, place) {
		event.setPlayerPlace(uuid, place.toString());
		$scope.playerlist = event.getList("playerList");
	}

	$scope.updateResult = function(uuid, result) {
		console.log("updateResult:", result);
		//console.log("RESULT:", result);
		event.setPlayerResult(uuid, result);
		//console.log("updateResult:AFTER:", event.getList("playerList"));
		$scope.playerlist = event.getList("playerList");
	}
	
	$scope.selectedItemChange = function(item) {
		console.log('Item changed to ' + JSON.stringify(item));
		
		if (typeof item !== "undefined") {
			event.addList(item , "playerList");
			console.log("playerList:", event.getList("playerList"));
			$scope.playerlist = event.getList("playerList");
		}
		$scope.selectedItem = null;
		$scope.searchText = "";
    }
	
	$scope.searchTextChange = function(text) {
      console.log('Text changed to ' + text);
    }
	
	$scope.addItem = function(item) {
      console.log('addItem ' + item);
    }

	$scope.removeItem = function(uuid) {
		console.log('removeItem ' + uuid);
		event.removeList(uuid , "playerList");
		$scope.playerlist = event.getList("playerList");
    }
	
//TODO: fix query to only search when query is one character, then save it for caching.
	$scope.querySearch   = function(query) {
		console.log("query: ", query);
		$scope.isLoading=true;

		return myHttpFactory.getPlayersAutoComplete(query).then(function(data) {

					return data.filter( 
						function (value) {
							console.log("Name + query",angular.lowercase(value.Firstname), angular.lowercase(query), (angular.lowercase(value.Firstname)).indexOf(angular.lowercase(query)));
							retval = ((angular.lowercase(value.Firstname)).indexOf(angular.lowercase(query)) === 0);
							$scope.isLoading=false;
							return retval;
						}
					);
					

		});	
		
	}
	
	
}]);



app.controller( 'LocationDateTimeCtrl' , ['$scope', 'event', 'myHttpFactory', function( $scope , event, myHttpFactory) {


	$scope.showNewLocation = true;
	$scope.showLastLocation = true;
	$scope.showNewLocationButton = true;
	$scope.showChosenLocation = false;
	$scope.showLockLocationTime = false;
	$scope.displayLocation  = "";


	$scope.startDateTime = 0;
	$scope.stopDateTime = 0;
	$scope.lastlocationlist = 0;
	
	$scope.lockedStartDate = 0;
	$scope.lockedStopDate = 0;

	startButtonText = "Lock Event Start";
	$scope.startButton = startButtonText;
	stopButtonText = "Lock Event Stop";
	$scope.stopButton = stopButtonText;
	
	$scope.start = {
		year: 0,
		month: 0,
		day: 0,
		hour: 0,
		minute: 0,
		offset: 0,
		PickedDate: new Date(),
	};
	
	$scope.stop = {
		year: 0,
		month: 0,
		day: 0,
		hour: 0,
		minute: 0,
		offset: 0,
		PickedDate: new Date(),
	};


	console.log("LocationDateTimeCtrl");
	$scope.years = [];
	for (var i=2015;i<3000;i++) {
		$scope.years.push(i);
	}
	
	$scope.months = ["01","02","03","04","05","06","07","08","09","10","11","12"];
	
	$scope.days = [];
	for (var i=1;i<32;i++) {
		$scope.days.push(i);
	}
	
	$scope.hours = ["00","01","02","03","04","05","06","07","08","09","10","11","12","13","14","15","16","17","18","19","20","21","22","23"];

	$scope.minutes = ["00","05","10","15","20","25","30","35","40","45","50","55"];
	
	$scope.showDateTime = true;

	

	
	
	$scope.testclick = function(){
		console.log("CAUGHT!?", $scope.gPlace );
		
	};
	
	$scope.getLastLocation = function() {
		console.log("Get Last Location:")
		return myHttpFactory.getLastLocation(playerinfo.UUID).then(function(data) {
			
			$scope.lastlocationlist = data;
			event.addList(data, "locationList");
	
		});
		
	
	};
	
	$scope.timeschanged = function(choice) {
		

			
			if (choice == "STOP") {
				
				
				var pickeddate = moment($scope.stop.PickedDate).format('YYYY-MM-DD');
				$scope.lockedStopDate = moment.tz(pickeddate + " " + $scope.stop.hour + ":" + $scope.stop.minute + " ", "YYYY-MM-DD HH:mm", $scope.start.offset);
				event.setStopDate($scope.lockedStopDate.format());
			}
			if (choice == "START") {
				
				var pickeddate = moment($scope.start.PickedDate).format('YYYY-MM-DD');
				$scope.lockedStartDate = moment.tz(pickeddate + " " + $scope.start.hour + ":" + $scope.start.minute + " ", "YYYY-MM-DD HH:mm", $scope.start.offset);
				event.setStartDate($scope.lockedStartDate.format());
			}
			
			console.log("TIME UPDATED:", event.getEvent().Start, event.getEvent().Stop );
		
			
	};
	
	$scope.initTime = function() {

		//assume the location list is correct
		currentloc = event.getList("locationList")[0];
		console.log("INIT TIME CURRENT LOCATION", currentloc);
			
		$scope.startDateTime = moment.tz(new Date(), currentloc.Locationtz);
		$scope.stopDateTime = moment.tz(new Date(), currentloc.Locationtz);
			
		minutes = Math.floor(($scope.startDateTime.minute())/5);
		
			
		$scope.start.hour = $scope.startDateTime.hour();
		$scope.start.minute = $scope.minutes[minutes];
		$scope.start.PickedDate = new Date($scope.startDateTime.format('MM-DD-YYYY'));
			
				
		if (($scope.stopDateTime.hour() + 1) == 24) {
			$scope.stop.hour = 0;
			$scope.stopDateTime = $scope.stopDateTime.add(1, 'days');
			//$scope.stopDateTime.day() = $scope.stopDateTime.day() + 1;
		} else {
			$scope.stop.hour = $scope.stopDateTime.hour() + 1;
		}
		$scope.stop.minute = $scope.minutes[minutes];
		$scope.stop.PickedDate = new Date($scope.stopDateTime.format('MM-DD-YYYY'));
			
		$scope.start.offset = currentloc.Locationtz;
			
		$scope.timeschanged("STOP");
		$scope.timeschanged("START");
			
		$scope.showDateTime = true;
		//$scope.showLockLocationTime = true;
		
	};
	
	$scope.onClick = function (choice) {
		//TODO: 
		/*
			take apart time and location. this click should be only for setting up location in eventlist
			
			if lastlocation was chosen and timezone isnt empty, then use lastlocation time zone 
				event.addList(lastlocation node, "locationList");
			else 
				get timezone and then event.addList(currentloc, "locationList");
		
			init time quesions
		*/
		
		pickedLocation = event.getList("locationList");	
		
		console.log("pickedLocation: ", pickedLocation.Locationname);
		
		
		if (choice == "lastlocation") {
			console.log("LAST LOCATION");

			$scope.showNewLocation = false;
			$scope.showLastLocation = false;
			$scope.showChosenLocation = true;
			$scope.showLockLocationTime = true;
		}		
			
		if (choice == "newlocation") {
			console.log("NEW LOCATION");
			$scope.showNewLocation = false;
			$scope.showLastLocation = false;
			$scope.showChosenLocation = true;
			$scope.showLockLocationTime = true;
		}
		
		if (choice == "reset") {
			console.log("Start over");
			$scope.showNewLocation = true;
			$scope.showLastLocation = true;
			$scope.showChosenLocation = false;
			$scope.showLockLocationTime = false;
					
		}
		
		if (pickedLocation.Locationtz != undefined ) {
			//the tz is already there!
			$scope.displayLocation  = pickedLocation.Locationname;
			return;
		}
		
		mapsString = {
			lat: pickedLocation[0].Locationlat,
			lng: pickedLocation[0].Locationlng,
			timestamp: "0",
			
		};
		
		return myHttpFactory.httpGoogleAPITimeZoneOffset(mapsString).then(function(result) {
		//return myHttpFactory.httpTimeZoneOffset( mapsString).then(function(result) {
			console.log("RESULT:",  result.data);
			pickedLocation = event.getList("locationList")[0];
			
			$scope.startDateTime = moment.tz(new Date(), result.data.timezone);
			$scope.stopDateTime = moment.tz(new Date(),  result.data.timezone);
			
			pickedLocation.Locationtz = result.data.timezone;
			event.addList(pickedLocation, "locationList");
			
			pickedLocation = event.getList("locationList")[0];
			console.log("******AFTER**********", pickedLocation);
			$scope.displayLocation = pickedLocation.Locationname;
			$scope.initTime();
			
		});

		
	};

}]);

app.controller( 'GamesCtrl' ,  ['$scope', 'myHttpFactory'   , function( $scope  , myHttpFactory ) {

	$scope.result="";
	$scope.showaddresults=false;
	$scope.showupdateresults=false;

	$scope.onClick = function () {
		console.log("games click", $scope.game);
		
		return myHttpFactory.addGame($scope.game).then(function(data) {
			$scope.result = data;
			
			if (data.Addedgames != null) {
				$scope.showaddresults=true;
			} else {
				$scope.showaddresults=false;
			}
			if (data.Updatedgames != null) {
				$scope.showupdateresults=true;
			} else {
				$scope.showupdateresults=false;
			
			}
			console.log("data",data);
			return data;

		});	

	};

}]);


//RETIRED
function FindNemesis(stats) {

	


	var rated = [];
	for (var i = 0; i < eventcargo.length; i++) {
		
		var meuuid = 666;
		var ident = 0;
		var sortA = [];
		for (j=0;j<eventcargo[i].Competitors.length;j++) {
			var playeruuid = eventcargo[i].Competitors[j].Player.UUID;
			
			var temprate = {
					name: eventcargo[i].Competitors[j].Player.Firstname + " " + eventcargo[i].Competitors[j].Player.Surname,
					uuid: playeruuid,
					rating: eventcargo[i].Competitors[j].Result.Place,
					beatthem: 0,
					beatenbythem: 0
			};

			sortA.push(temprate);
			
			sortA.sort(function(a, b) {
					if (a.rating < b.rating) {
						return -1;
					}
					if (a.rating > b.rating) {
						return 1;
					}
					return 0;
			});
		}

		//console.log("sorted",sortA);
		// after sorting the places, record wins/loses 
		// add the other player's stats to return array
		var foundme = false;
		
		
		
		
		//console.log("sortA", sortA);
		for (k=0;k<sortA.length;k++) {
		
			//console.log("rated:", rated, k);
			if (playerinfo.UUID == sortA[k].uuid ) {
				foundme = true;
			} else 
			{
				var ratedExist = false;
				var ratedLoc = 0;
				for (m=0;m<rated.length;m++) {
					if (rated[m].uuid == sortA[k].uuid) {
						ratedExist = true;
						ratedLoc = m;
					}
				}
				if (!ratedExist) {
					rated.push(sortA[k]);
					rated[rated.length-1].beatenbythem++;
				}  else {
			
					if (!foundme) {
						rated[ratedLoc].beatenbythem++; 
						
					} else {
						rated[ratedLoc].beatthem++; 
					}
				}
			}
			
		}
		
		
		
		
	}

	//ratedSorted = Object.keys(rated).sort(function(a,b){return rated[a]-rated[b]})
	//console.log("rated",rated);
	
	var top = 3;
	var bottom = 3;
	
	if (rated.length < 6) {
	//TODO: fix for less then 6. too tired
	
	} else {
	
	//sort the rated array of objects to top 3 and bottom 3
		rated.sort(function(a, b) {
					if (a.beatenbythem < b.beatenbythem) {
						return 1;
					}
					if (a.beatenbythem > b.beatenbythem) {
						return -1;
					}
					return 0;
			});

		stats.nemesis.push({ name:rated[0].name, record:rated[0].beatenbythem});
		stats.nemesis.push({ name:rated[1].name, record:rated[1].beatenbythem});
		stats.nemesis.push({ name:rated[2].name, record:rated[2].beatenbythem});

		//for charts
		//chart.nemesis.legend = true;
		//chart.nemesis.labels = [rated[0].name, rated[1].name, rated[2].name];
		//chart.nemesis.data = [rated[0].beatenbythem, rated[1].beatenbythem, rated[2].beatenbythem];
	
	
		rated.sort(function(a, b) {
					if (a.beatthem < b.beatthem) {
						return 1;
					}
					if (a.beatthem > b.beatthem) {
						return -1;
					}
					return 0;
			});
	
		stats.dominate.push({ name:rated[0].name, record:rated[0].beatthem});
		stats.dominate.push({ name:rated[1].name, record:rated[1].beatthem});
		stats.dominate.push({ name:rated[2].name, record:rated[2].beatthem});
	
	
		/*
		chart.dominate.legend = true;
		chart.dominate.labels = [rated[0].name, rated[1].name, rated[2].name];
		chart.dominate.data = [rated[0].beatthem, rated[1].beatthem, rated[2].beatthem];
		*/
	}
};