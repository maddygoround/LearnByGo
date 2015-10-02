(function () {
    "use strict";
    IM.factory("VideoListFactory", ['$resource', 'GlobalConstant',
        function ($resource, GlobalConstant) {
            return $resource(GlobalConstant.serviceURL + "videos", {},
                {
                    get: { method: "GET", isArray: false }
                });
        }
    ]);
}());