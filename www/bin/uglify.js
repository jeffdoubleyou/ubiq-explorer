var glob = require( 'glob' );  
var fs = require('fs');
var UglifyJS = require("uglify-js");
var uuid = require('uuid/v1');
var replace = require("replace");
var ini = require("ini");
var validator = require("html-angular-validate");

var config = ini.parse(fs.readFileSync('../conf/app.conf', 'utf-8'));
var customConfig = ini.parse(fs.readFileSync('../conf/custom.conf', 'utf-8'));

if(customConfig.base_href)
    config.base_href = customConfig.base_href

var options = {
    "toplevel": true,
    "mangle": false,
    "compress": { "passes": 2 }
};

var buildId = uuid();
var outputFilename = 'ubiq-explorer.'+buildId+'.js';
console.log("Build ID", buildId);

glob('ubiq-explorer.*.js', function(err, files) {
    console.log("Removing old minified javascript files");
    if(err) {
        console.log("ERROR:", err)
    } else {
        files.forEach(function(file) {
            console.log("Removing ", file)
            fs.unlinkSync(file)
        });
    }
});

glob( '*.js', { cwd: "./scripts", matchBase:true }, function( err, files ) {
  console.log("Minifying Javascript to " + outputFilename);
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
        silent: true
    });
    replace({
       regex: '.*base href.* />',
       replacement: '    <base href="'+config.base_href+'" />',
       paths: ['index.html'],
       recursive: false,
       silent: true
    });
    console.log("Done")
  }
});

replace({
    regex: "src=\"'(.*).html.*'\"></div>",
    replacement: "src=\"'$1.html#buildId="+buildId+"'\"></div>",
    paths: ['index.html', 'views', 'template'],
    recursive: true,
    silent: true
});

replace({
    regex: "ng-include=\"'(.*).html.*'\"></div>",
    replacement: "ng-include=\"'$1.html#buildId="+buildId+"'\"></div>",
    paths: ['index.html', 'views', 'template'],
    recursive: true,
    silent: false
});

validator.validate(
    ['template/**', 'views/**'],
    {
        tmplext: 'html',
        customattrs: [
            'md-cell',
            'md-row',
            'md-head',
            'md-table',
            'md-body',
            'md-column',
            'blockie',
            'address', // For rendering blockies
            'size', // For choosing blockie size
            'chart-data',
            'chart-labels',
            'chart-colors',
            'chart-options',
            'chart-legend',
            'chart-series',
            'flex', // Not sure why all of these throw errors - pretty sure they are valid for divs
            'layout',
            'layout-fill',
            'layout-xs',
            'flex-xs',
            'flex-gt-xs',
            'hide-gt-xs',
            'hide-xs',
            'layout-margin',
            'layout-padding',
            'layout-align',
            'uib-tab-content-transclude',
            'uib-tab-heading-transclude'
        ],
        relaxerror: [
            'Consider adding a “lang” attribute to the “html” start tag to declare the language of this document.',
            'Possible misuse of “aria-label”. (If you disagree with this warning, file an issue report or send e-mail to www-validator@w3.org.)'
        ],
        wrapping: {
            'li': '<ul>{0}</ul>'
        }
    }
).then(function(result) {
    if(result.allpassed == true) {
        console.log("All HTML validation tests passed");
    } else {
        console.log("Failed HTML validation tests...", result);
        process.exit(1);
    }
}, function(err) {
    console.log(err);
    process.exit(1);
});


validator.validate(
    ['index.html'],
    {
        customattrs: [
            'md-cell',
            'md-row',
            'md-head',
            'md-table',
            'md-body',
            'md-column',
            'blockie',
            'address', // For rendering blockies
            'size', // For choosing blockie size
            'chart-data',
            'chart-labels',
            'chart-colors',
            'chart-options',
            'chart-legend',
            'chart-series',
            'flex', // Not sure why all of these throw errors - pretty sure they are valid for divs
            'layout',
            'layout-fill',
            'layout-xs',
            'flex-xs',
            'flex-gt-xs',
            'hide-gt-xs',
            'hide-xs',
            'layout-margin',
            'layout-padding',
            'layout-align',
            'uib-tab-content-transclude',
            'uib-tab-heading-transclude',
        ],
        relaxerror: [
            'Possible misuse of “aria-label”. (If you disagree with this warning, file an issue report or send e-mail to www-validator@w3.org.)'
        ],
    }
).then(function(result) {
    if(result.allpassed == true) {
        console.log("HTML validation of index.html passed");
    } else {
        console.log("Failed HTML validation of index.html...", result);
        console.log(result.failed[0]);
        process.exit(1);
    }
}, function(err) {
    console.log(err);
    process.exit(1);
});
