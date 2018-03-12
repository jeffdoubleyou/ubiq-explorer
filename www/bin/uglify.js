var glob = require( 'glob' );  
var fs = require('fs');
var UglifyJS = require("uglify-js");
var uuid = require('uuid/v1');
var replace = require("replace");

var options = {
    "toplevel": true,
    "mangle": false,
    "compress": { "passes": 2 }
};

var outputFilename = 'ubiq-explorer.'+uuid()+'.js';

console.log("Minifying Javascript to " + outputFilename)
glob( '*.js', { cwd: "./scripts", matchBase:true }, function( err, files ) {
  if(err) {
      console.log("ERROR:", err)
  } else {
    var js = {};
    js["app.js"] = fs.readFileSync("app.js", "utf8");
    files.forEach(function(file) {
        console.log("Loading ", file);
        js["./scripts/"+file] = fs.readFileSync("./scripts/"+file, "utf8");
    });
    console.log("Done reading JS files");
    fs.writeFileSync(outputFilename, UglifyJS.minify(js, options).code, "utf8");
    var replaced = replace({
        regex: '.*<!--ubiq-explorer-js-->',
        replacement: '<script src="'+outputFilename+'"></script><!--ubiq-explorer-js-->',
        paths: ['index.html'],
        recursive: false,
        silent: false
    });
    console.log("Replaced:", replaced)
    console.log("Done")
  }
});


