Unixify
-------

**Unixify** is a simple command line interface tool for changing file and folder names to be more UNIX like. This is done by converting space white signs to floors and all uppercase letters to lowercase. Unixify also has the ability to rename folders and hidden files.

Installation
------------

**Unixify** is completely written in go with no dependencies, so all you have to do is install git, go and execute below commands.

    $ go get github.com/thinkofher/unixify

Make sure you have `$GOPATH/bin` added to your `$PATH` variable.

Usage
----

    $ unixify -help
    USAGE:
        unixify [OPTIONS] [ARGS...]

    INFO:
        Simple portable command line interface tool
        for changing file and folder names to be more
        unixy.

        Examples:

        * "Some File"         -> "some_file"
        * "sOme File"         -> "some_file"
        * "SOMEFILE"          -> "somefile"
        * ".some hidden file" -> ".some_hidden_file"

    AUTHOR:
        Beniamin Dudek <beniamin.dudek@yahoo.com>

    OPTIONS:
        -dirs
            include directories
        -hidden
            include hidden files
        -interactive
            asks before every rename
        -verbose
            print current actions
        -version
            show version and exit
        -yes
            skip asking to rename at startup

Sample session with unixify:

    $ ls
    'Some File'  'Some Folder'  'Some Other File'
    $ unixify
    unixify: Do you want to rename 2 files? (y/N): y
    $ ls
    'Some Folder'   some_file   some_other_file

Using unixify in a folder with hidden files and directories:

    $ ls -a
     .   '.Some file'   '.Some hidden folder'   some_other_file
     ..  'Some Folder'   some_file
    $ unixify -dirs -hidden -yes
    $ ls -a
     .  ..  .some_file  some_file  some_folder  .some_hidden_folder  some_other_file

When providing arguments, unixify will change only files (and folders!) with given filenames. Example:

    $ ls -a
     .   ..  '.Some Hidden Folder'  'Some Other File'   SomeFile
    $ unixify -yes ".Some Hidden Folder" "Some Other File"
    $ ls -a
    .  ..  SomeFile  .some_hidden_folder  some_other_file

Using unixify, when there is nothing more to rename:

    $ unixify
    Files in '/home/user/somedir' are already unixified.

**Remember**: renaming files may **damage** your data! Be careful when using this tool, it comes without any warranty.
