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

func GetOpt(
  args []string,
  shortopts string,
  longopts []string,
) (
  leftovers []string,
  optargs []OptArg,
  err error,
) {
    shorts, err := build_shorts(shortopts)
    if err != nil {
        return nil, nil, err
    }
    longs, err := build_longs(longopts)
    if err != nil {
        return nil, nil, err
    }
    leftovers = args
    skip := false
    emitopt := ""
    for i, arg := range args {
        leftovers = leftovers[1:]
        if arg == "--" {
            if skip {
                msg := "expected an argument got --"
                return nil, nil, errors.New(msg)
            }
            break
        } else if skip {
            if arg[0] == '-' {
                msg := fmt.Sprintf("expected an argument got %v", arg)
                return nil, nil, errors.New(msg)
            }
            optargs = append(optargs, new_optarg(emitopt, arg))
            skip = false
            continue
        }

        if found, opt, hasarg := short(arg, shorts); found {
            if hasarg {
                skip = true
                emitopt = opt
            } else {
                optargs = append(optargs, new_optarg(opt, ""))
            }
        } else if found, opt, oarg, hasarg, err := long(arg, longs); found {
            if err != nil {
                return nil, nil, err
            } else if oarg != "" {
                optargs = append(optargs, new_optarg(opt, oarg))
            } else if hasarg {
                skip = true
                emitopt = opt
            } else {
                optargs = append(optargs, new_optarg(opt, ""))
            }
        } else {
            leftovers = args[i:]
            break
        }
    }

    return leftovers, optargs, nil
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

