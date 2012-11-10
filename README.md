Yet Another GetOpt Library For Go
=================================

By Tim Henderson (tim.tadh@gmail.com)

Motivation
----------

The existing libraries[1,2,3,4] don't exactly work how I would like them to
work. Specifically, I want it to work exactly like how the Python `getopt`
module works. The reason being, that particular implementation has proven
through years of use to be incredibly scalable and flexible. While, there are
definitely easier to use approaches (see for instance `argparse`) none have the
flexibility of the simple `getopt`. 

The current solutions also try to do too much while not implementing the basic
functionality in a sensible way. This library will simply pull out a list of the
flags found, and their options (as strings). Parsing the arguments into Go
datatypes will be up the user of the library.


[1] flag
[2] https://github.com/droundy/goopt
[3] https://github.com/jteeuwen/go-pkg-optarg
[4] https://github.com/kesselborn/go-getopt

