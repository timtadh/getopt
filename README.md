Yet Another GetOpt Library For Go
=================================

By Tim Henderson (tim.tadh@gmail.com)

BSD Licenced

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
datatypes will be up to the user of the library.

### Todo

1. Lots more tests (currently it only tests basic functionality, complex
   situations are probably missed.)

### Contributing

1. "Fork"
2. Make a feature branch.
3. Do your commits
4. Send "pull request". This can be
    1. A github pull request
    2. A issue with a pointer to your publicly readable git repo
    3. An email to me with a pointer to your publicly readable git repo

Docs
----

    PACKAGE

    package getopt
        import "github.com/timtadh/getopt"


    FUNCTIONS

    func GetOpt(
        args []string,
        shortopts string,
        longopts []string,
    ) (
        leftovers []string,
        optargs []OptArg,
        err error,
    )
        GetOpt works like `getopt` in python's `getopt` module in the stdlib
        (modulus implementation bugs).

        params

      args - the argv []string slice
      shortopts - a string of options (similar to what GNU's getopt excepts).
                  Options which desire an argument should have a colon, ":",
                  subsequent to them. There are no optional arguments at this
                  time.
                  ex. "hvx:r" would accept -h -v -x asdf -r
      longopts - a list of strings which describe the long options (eg those with
                 "--" in front). Placing an = on the end indicates a required
                 argument.
                 ex. []string{"help", "example="}
                   would accept --help --example=tom


    TYPES

    type OptArg interface {
        Opt() string
        Arg() string
    }
        The GetOpt function will return a list of OptArgs. If there is no arg
        then Arg() will return "". Opt will contain a leading "-" for short and
        "--" for long args.


#### Footnotes

- [1] flag
- [2] https://github.com/droundy/goopt
- [3] https://github.com/jteeuwen/go-pkg-optarg
- [4] https://github.com/kesselborn/go-getopt

