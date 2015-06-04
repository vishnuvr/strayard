'use strict';

// # Globbing
// for performance reasons we're only matching one level down:
// 'test/spec/{,*/}*.js'
// use this if you want to recursively match all subfolders:
// 'test/spec/**/*.js'

module.exports = function(grunt) {
  require('load-grunt-tasks')(grunt);
  require('time-grunt')(grunt);

  // Configurable paths for the application
  var appConfig = {
    app: require('./bower.json').appPath || 'app',
    dist: 'dist'
  };

  // Define the configuration for all the tasks
  grunt.initConfig({

    // Project settings
    yeoman: appConfig,

    // Empties folders to start fresh
    clean: {
      dist: {
        files: [{
          dot: true,
          src: [
            '.tmp',
            '<%= yeoman.dist %>/{,*/}*',
            '!<%= yeoman.dist %>/.git{,*/}*'
          ]
        }]
      },
      server: '.tmp'
    },

    // Add vendor prefixed styles
    autoprefixer: {
      options: {
        browsers: ['last 1 version']
      },
      dist: {
        files: [{
          expand: true,
          cwd: '.tmp/styles/',
          src: '{,*/}*.css',
          dest: '.tmp/styles/'
        }]
      }
    },

    // Automatically inject Bower components into the app
    wiredep: {
      app: {
        src: ['<%= yeoman.app %>/index.html'],
        ignorePath:  /\.\.\//,
        exclude: [
          'bower_components/messenger/build/css/messenger.css',
          'bower_components/messenger/build/css/messenger-theme-future.css',
          'bower_components/messenger/build/css/messenger-theme-block.css',
          'bower_components/messenger/build/css/messenger-theme-air.css',
          'bower_components/messenger/build/css/messenger-theme-ice.css',

          'bower_components/messenger/build/js/messenger-theme-future.js'
        ]
      }
    },

    less: {
      development: {
        files: {
          '.tmp/styles/main.css': '<%= yeoman.app %>/styles/main.less'
        },
        options: {
          paths: ['<%= yeoman.app %>/styles']
        }
      },
      production: {
        files: {
          'dist/css/main.css': '<%= yeoman.app %>/styles/main.less'
        },
        options: {
          cleancss: true,
          paths: ['<%= yeoman.app %>/styles']
        }
      }
    },

    // Renames files for browser caching purposes
    filerev: {
      dist: {
        src: [
          // We do not need digest filenames as our files are loaded from
          // a commit-based path
          //'<%= yeoman.dist %>/scripts/{,*/}*.js',
          //'<%= yeoman.dist %>/styles/{,*/}*.css',
          //'<%= yeoman.dist %>/images/{,*/}*.{png,jpg,jpeg,gif,webp,svg}',
          //'<%= yeoman.dist %>/styles/fonts/*'
        ]
      }
    },

    // Reads HTML for usemin blocks to enable smart builds that automatically
    // concat, minify and revision files. Creates configurations in memory so
    // additional tasks can operate on them
    useminPrepare: {
      html: '<%= yeoman.app %>/index.html',
      options: {
        dest: '<%= yeoman.dist %>',
        flow: {
          html: {
            steps: {
              js: ['concat', 'uglifyjs'],
              css: ['cssmin']
            },
            post: {
              css: [{
                name:'cssmin',
                createConfig: function(context, block) {
                  var generated = context.options.generated;
                  generated.options = {
                    keepBreaks: true,
                  };
                }
              }],

              js: [{
                name:'uglify',
                createConfig: function(context, block) {
                  var generated = context.options.generated;
                  generated.options = {
                    compress: {},
                    mangle: {},
                    beautify: {
                      beautify: true,
                      indent_level: 0,
                      space_colon: false,
                      width: 1000,
                    },
                  };
                }
              }]
            }
          }
        }
      }
    },

    // Performs rewrites based on filerev and the useminPrepare configuration
    usemin: {
      html: ['<%= yeoman.dist %>/{,*/}*.html'],
      css: ['<%= yeoman.dist %>/styles/{,*/}*.css'],
      options: {
        assetsDirs: ['<%= yeoman.dist %>','<%= yeoman.dist %>/images']
      }
    },

    imagemin: {
      dist: {
        files: [{
          expand: true,
          cwd: '<%= yeoman.app %>/images',
          src: '{,*/}*.{png,jpg,jpeg,gif}',
          dest: '<%= yeoman.dist %>/images'
        }]
      }
    },

    svgmin: {
      dist: {
        files: [{
          expand: true,
          cwd: '<%= yeoman.app %>/images',
          src: '{,*/}*.svg',
          dest: '<%= yeoman.dist %>/images'
        }]
      }
    },

    htmlmin: {
      dist: {
        options: {
          preserveLineBreaks: true,
          collapseWhitespace: true,
          conservativeCollapse: false,
          collapseBooleanAttributes: true,
          removeComments: true,
          removeCommentsFromCDATA: true,
          removeOptionalTags: false,
          keepClosingSlash: true
        },
        files: [{
          expand: true,
          cwd: '<%= yeoman.dist %>',
          src: ['*.html', 'views/{,*/}*.html'],
          dest: '<%= yeoman.dist %>'
        }]
      }
    },

    // ng-annotate tries to make the code safe for minification automatically
    // by using the Angular long form for dependency injection.
    ngAnnotate: {
      dist: {
        files: [{
          expand: true,
          cwd: '.tmp/concat/scripts',
          src: ['*.js', '!oldieshim.js'],
          dest: '.tmp/concat/scripts'
        }]
      }
    },

    // Replace Google CDN references
    cdnify: {
      dist: {
        html: ['<%= yeoman.dist %>/*.html']
      }
    },

    // Copies remaining files to places other tasks can use
    copy: {
      dist: {
        files: [{
          expand: true,
          dot: true,
          cwd: '<%= yeoman.app %>',
          dest: '<%= yeoman.dist %>',
          src: [
            '*.{ico,png,txt}',
            '.htaccess',
            '*.html',
            'views/{,*/}*.html',
            'images/{,*/}*.{webp}',
            'fonts/*',
            'styles/fonts/*'
          ]
        }, {
          expand: true,
          cwd: '.tmp/images',
          dest: '<%= yeoman.dist %>/images',
          src: ['generated/*']
        }, {
          expand: true,
          cwd: 'bower_components/patternfly/dist',
          src: 'fonts/*',
          dest: '<%= yeoman.dist %>/styles'
        }, {
          expand: true,
          cwd: 'bower_components/patternfly/components/font-awesome',
          src: 'fonts/*',
          dest: '<%= yeoman.dist %>/styles'
        }, {
          expand: true,
          cwd: 'bower_components/zeroclipboard/dist',
          src: 'ZeroClipboard.swf',
          dest: '<%= yeoman.dist %>/scripts'
        }]
      },
      styles: {
        files: [{
          expand: true,
          cwd: '<%= yeoman.app %>/styles',
          dest: '.tmp/styles/',
          src: '{,*/}*.css'
        }, {
          expand: true,
          cwd: 'bower_components/patternfly/dist',
          src: 'fonts/*',
          dest: '.tmp/styles'
        }, {
          expand: true,
          cwd: 'bower_components/patternfly/components/font-awesome',
          src: 'fonts/*',
          dest: '.tmp/styles'
        }]
      }
    },

    // Run some tasks in parallel to speed up the build process
    concurrent: {
      server: [
        'less:development',
        'copy:styles'
      ],
      test: [
        'less:development'
      ],
      dist: [
        'less:production',
        'imagemin',
        'svgmin'
      ]
    },
  });

  grunt.registerTask('build', [
    'clean:dist',
    'wiredep',
    'useminPrepare',
    'concurrent:dist',
    'autoprefixer',
    'concat',
    'ngAnnotate',
    'copy:dist',
//    'cdnify',
    'less',
    'cssmin',
    'uglify',
    'filerev',
    'usemin',
    'htmlmin'
  ]);
};
