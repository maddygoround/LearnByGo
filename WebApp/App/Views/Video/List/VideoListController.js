(function () {
    "use strict";
    IM.controller('VideoListController', ['$scope','$rootScope','$location', 'VideoListFactory',
        function ($scope, $rootScope,$location, VideoListFactory) {
            VideoListFactory.get(null, function (res) {
                $scope.videoList = res.VideoList;
            });

            $scope.playVideo = function (video) {
                $rootScope.video = video;
                $location.path('/PlayVideo/' + video.videoId)
            }
        }]);
}());