(function () {
    "use strict";
    IM.config(["$stateProvider", "$urlRouterProvider", "$locationProvider",
        function ($stateProvider, $urlRouterProvider, $locationProvider) {

            $stateProvider
                .state('VideoList', {
                    url: "/VideoList",
                    templateUrl: "/App/Views/Video/List/VideoList.html",
                    controller: "VideoListController"
                    
                })
                .state('PlayVideo', {
                    url: "/PlayVideo/:id",
                    templateUrl: "/App/Views/Video/Player/Player.html",
                    controller: "PlayerController"
                })
            $urlRouterProvider.when("/", "/VideoList");

            $urlRouterProvider.otherwise("/");

            $locationProvider.html5Mode(false).hashPrefix("!");

        }]);
}());