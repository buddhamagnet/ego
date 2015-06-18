console.log('ego tripping in the console');

var App = angular.module('App', []);

App.controller('Whitehouse', function($scope, $http, $route, $routeParams) {
  progress = ['economy', 'whatsnext', 'all'];
  $http.get('http://localhost:9494/content/whitehouse/' + progress[Math.floor(Math.random()*progress.length)])
  .then(function(res){
    console.log('connected to content service for whitehouse progress update');
    $scope.b = res.data;
  });
});
