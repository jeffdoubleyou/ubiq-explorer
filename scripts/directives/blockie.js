angular.module('Explorer').directive('blockie', function() {
    var blockie = {};
    blockie.restrict = 'A';
    blockie.template = "&nbsp;";
    blockie.compile = function(element, attributes) {
      var linkFunction = function($scope, element, attributes) {
         element.css("background-image", 'url(' + blockies.create({ seed: attributes.address, size: 8, scale: 3}).toDataURL()+')');
         element.css("width", "24px");
         element.css("height", "24px");
         element.css("margin", "10px");
         element.css("overflow", "hidden");
         element.css("border-radius", "10%");
         element.css("box-shadow", "inset rgba(255, 255, 255, 0.2) 0 2px 2px, inset rgba(0, 0, 0, 0.3) 0 -1px 8px");
      }
      return linkFunction;
    }
    return blockie;
});

