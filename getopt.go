package getopt

import "errors"
import "fmt"
import "strings"

type OptArg interface {
    Opt() string
    Arg() string
}

type optarg struct {
    opt string
    arg string
}

func new_optarg(opt, arg string) *optarg { return &optarg{opt, arg} }
func (self *optarg) Opt() string { return self.opt }
func (self *optarg) Arg() string { return self.arg }

func GetOpt(args []string, short string, long []string) ([]string, []OptArg) {
    var leftovers []string
    var optargs []OptArg


    return leftovers, optargs
}

func build_longs(long []string) (map[string]bool, error) {
    longs := make(map[string]bool)
    for _, opt := range long {
        hasarg := false
        if opt[len(opt)-1] == '=' {
            opt = opt[:len(opt)-1]
            hasarg = true
        }
        opt = "--" + opt
        if _, has := longs[opt]; has {
            msg := fmt.Sprintf(
              "Option %v entered more than one in longs", opt)
            return nil, errors.New(msg)
        } else {
            longs[opt] = hasarg
        }
    }
    return longs, nil
}

func build_shorts(short string) (map[string]bool, error) {
    shorts := make(map[string]bool)
    for i, rc := range short {
        c := string(rc)
        if c == ":" { continue }
        if _, has := shorts["-"+c]; has {
            msg := fmt.Sprintf(
              "Option %v entered more than one in shorts", c)
            return nil, errors.New(msg)
        } else {
            shorts["-"+c] = false
            if i+1 < len(short) {
                nc := string(short[i+1])
                if nc == ":" {
                    shorts["-"+c] = true
                }
            }
        }
    }
    return shorts, nil
}

func short(arg string, shorts map[string]bool) (found bool, opt string, hasarg bool) {
    if hasarg, has := shorts[arg]; has {
        return true, arg, hasarg
    }
    return false, "", false
}

func long(arg string, longs map[string]bool) (
  found bool,
  opt, rarg string,
  hasarg bool,
  err error,
) {
    if i := strings.Index(arg, "="); i != -1 {
        opt = arg[:i]
        rarg = arg[i+1:]
    } else {
        opt = arg
        rarg = ""
    }
    if hasarg, has := longs[opt]; has {
        if !hasarg && rarg != "" {
            msg := fmt.Sprintf(
              "Option %v received an arg, %v, and did not expect one", opt, rarg)
            return false, "", "", false, errors.New(msg)
        }
        return true, opt, rarg, hasarg, nil
    }
    return false, "", "", false, nil
}

