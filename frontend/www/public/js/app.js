console.log('ego tripping in the console');

var App = angular.module('App', []);

App.controller('Blog', function($scope, $http, $route, $routeParams) {
  blogs = ['babbage', 'freeexchange', 'schumpeter', 'gulliver', 'buttonwood', 'erasmus', 'prospero'];
  $http.get('http://localhost:9494/content/blogs/' + blogs[Math.floor(Math.random()*blogs.length)])
  .then(function(res){
    console.log('connected to content service for blog post');
    $scope.b = res.data;
    $scope.b.status = res.data.status ? "yes" : "no";
    $scope.b.created = moment.unix(res.data.created).format("dddd, MMMM Do YYYY");
  });
});

App.controller('Article', function($scope, $http, $route, $routeParams) {
  $http.get('http://localhost:9494/content/article/leaders/21648639-enduring-power-families-business-and-politics-should-trouble-believers')
  .then(function(res){
    console.log('connected to content service for article');
    $scope.b = res.data;
    $scope.b.status = res.data.status ? "yes" : "no";
    $scope.b.created = moment.unix(res.data.created).format("dddd, MMMM Do YYYY");
  });
});