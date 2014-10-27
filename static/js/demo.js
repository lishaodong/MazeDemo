/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function TaskCtrl($scope, $http) {
  $scope.tasks = [];
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/task/').
      success(function(data) { $scope.tasks = data.Tasks; }).
      error(logError);
  };

  $scope.createRoom = function() {
    $http.post('/create/', {uname: $scope.uname}).
      error(logError).
      success(function() {
        refresh().then(function() {
        })
      });
  };

  refresh().then(function() { $scope.working = false; });
}