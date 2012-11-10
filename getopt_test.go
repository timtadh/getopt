package getopt

import "testing"
import "reflect"

func Test_BuildShorts(t *testing.T) {
    expected := map[string]bool{ 
      "-h":false, "-v":false, "-e":false, "-x":true, "-y":true, "-z":true, }
    shorts, err := build_shorts("hvx:y:z:e") 
    if err != nil {
        t.Fatal(err)
    }
    if !reflect.DeepEqual(shorts, expected) {
        t.Log("got", shorts)
        t.Log("expected", expected)
        t.Fatal("Build shorts failed!")
    }
}

func Test_doubleBuildShorts(t *testing.T) {
    var err error
    _, err = build_shorts("hvh:x:y:z:e") 
    if err == nil {
        t.Fatal("expected an error...")
    }
    _, err = build_shorts("hvx:y:vz:e") 
    if err == nil {
        t.Fatal("expected an error...")
    }
}

func Test_Short_arg(t *testing.T) {
    shorts, err := build_shorts("hvx:y:z:e") 
    if err != nil {
        t.Fatal(err)
    }
    found, opt, hasarg := short("-y", shorts)
    if !found {
        t.Fatal("couldn't find -y")
    }
    if opt != "-y" {
        t.Fatal("opt != -y")
    }
    if !hasarg {
        t.Fatal("-y must have an arg")
    }
}

func Test_Short_noarg(t *testing.T) {
    shorts, err := build_shorts("hvx:y:z:e") 
    if err != nil {
        t.Fatal(err)
    }
    found, opt, hasarg := short("-h", shorts)
    if !found {
        t.Fatal("couldn't find -h")
    }
    if opt != "-h" {
        t.Fatal("opt != -h")
    }
    if hasarg {
        t.Fatal("-h must not have an arg")
    }
}

func Test_badShort(t *testing.T) {
    shorts, err := build_shorts("hvx:y:z:e") 
    if err != nil {
        t.Fatal(err)
    }
    found, opt, hasarg := short("-r", shorts)
    if found {
        t.Fatal("could find -r")
    }
    if opt != "" {
        t.Fatal("opt != ''")
    }
    if hasarg {
        t.Fatal("'' must not have an arg")
    }
}

func Test_BuildLongs(t *testing.T) {
    expected := map[string]bool{ 
      "--help":false, "--verbose":false, "--empty":false, "--example":true,
      "--yacc":true, "--zebra":true, }
    shorts, err := build_longs([]string{"help", "verbose", "empty", "example=",
      "yacc=", "zebra="})
    if err != nil {
        t.Fatal(err)
    }
    if !reflect.DeepEqual(shorts, expected) {
        t.Log("got", shorts)
        t.Log("expected", expected)
        t.Fatal("Build shorts failed!")
    }
}

func Test_doubleBuildLongs(t *testing.T) {
    var err error
    _, err = build_longs([]string{"help", "verbose", "empty", "example=",
      "yacc=", "zebra=", "yacc="})
    if err == nil {
        t.Fatal("expected an error...")
    }
    _, err = build_longs([]string{"help", "verbose", "help=", "empty", 
      "example=", "yacc=", "zebra="})
    if err == nil {
        t.Fatal("expected an error...")
    }
}

func Test_Long_noarg(t *testing.T) {
    longs, err := build_longs([]string{"help", "verbose", "empty", "example=",
      "yacc=", "zebra="})
    if err != nil {
        t.Fatal(err)
    }
    found, opt, arg, hasarg, err := long("--help", longs)
    if err != nil {
        t.Fatal(err)
    }
    if !found {
        t.Fatal("couldn't find --help")
    }
    if opt != "--help" {
        t.Fatal("opt != --help")
    }
    if arg != "" {
        t.Fatal("arg != ''")
    }
    if hasarg {
        t.Fatal("--help must not have arg")
    }
}

func Test_Long_arg(t *testing.T) {
    longs, err := build_longs([]string{"help", "verbose", "empty", "example=",
      "yacc=", "zebra="})
    if err != nil {
        t.Fatal(err)
    }
    found, opt, arg, hasarg, err := long("--example=help", longs)
    if err != nil {
        t.Fatal(err)
    }
    if !found {
        t.Fatal("couldn't find --example")
    }
    if opt != "--example" {
        t.Fatal("opt != --example")
    }
    if arg != "help" {
        t.Fatal("arg != 'help'")
    }
    if !hasarg {
        t.Fatal("--example must have arg")
    }
}

func Test_Long_argnext(t *testing.T) {
    longs, err := build_longs([]string{"help", "verbose", "empty", "example=",
      "yacc=", "zebra="})
    if err != nil {
        t.Fatal(err)
    }
    found, opt, arg, hasarg, err := long("--example", longs)
    if err != nil {
        t.Fatal(err)
    }
    if !found {
        t.Fatal("couldn't find --example")
    }
    if opt != "--example" {
        t.Fatal("opt != --example")
    }
    if arg != "" {
        t.Fatal("arg != ''")
    }
    if !hasarg {
        t.Fatal("--example must have arg")
    }
}

func Test_Long_arg_unexpected(t *testing.T) {
    longs, err := build_longs([]string{"help", "verbose", "empty", "example=",
      "yacc=", "zebra="})
    if err != nil {
        t.Fatal(err)
    }
    found, opt, arg, hasarg, err := long("--help=wat", longs)
    if err == nil {
        t.Fatal("expected an error")
    }
    if found {
        t.Fatal("shouldn't find --help=wat")
    }
    if opt != "" {
        t.Fatal("opt != ''")
    }
    if arg != "" {
        t.Fatal("arg != ''")
    }
    if hasarg {
        t.Fatal("--help must not have arg")
    }
}


func Test_Long_badopt(t *testing.T) {
    longs, err := build_longs([]string{"help", "verbose", "empty", "example=",
      "yacc=", "zebra="})
    if err != nil {
        t.Fatal(err)
    }
    found, opt, arg, hasarg, err := long("--wizard", longs)
    if err != nil {
        t.Fatal(err)
    }
    if found {
        t.Fatal("shouldn't find --wizard")
    }
    if opt != "" {
        t.Fatal("opt != ''")
    }
    if arg != "" {
        t.Fatal("arg != ''")
    }
    if hasarg {
        t.Fatal("--wizards shouldn't have arg")
    }
}

