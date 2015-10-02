(function () {
    "use strict";
    IM.controller('PlayerController', ['$scope', '$rootScope', '$location', 'GlobalConstant',
        function ($scope, $rootScope, $location, GlobalConstant) {
            $scope.video = $rootScope.video;
            $('video').attr('src', GlobalConstant.videoHosting + $scope.video.videoName)

            //$scope.videoURL = GlobalConstant.videoHosting + $scope.video.videoName;
        }]);
}());