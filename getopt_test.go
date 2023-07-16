package getopt

import "testing"
import "reflect"

func Test_BuildShorts(t *testing.T) {
	expected := map[string]bool{
		"-h": false, "-v": false, "-e": false, "-x": true, "-y": true, "-z": true}
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
		"--help": false, "--verbose": false, "--empty": false, "--example": true,
		"--yacc": true, "--zebra": true}
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

func Test_Getopt_one_short_no_arg_no_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-h",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if len(args) != 0 {
		t.Fatal("should be no leftovers")
	}
	if optargs[0].Opt() != "-h" && optargs[0].Arg() != "" {
		t.Fatal("expected to find -h")
	}
}

func Test_Getopt_one_short_arg_no_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-y", "its a yacc!",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if len(args) != 0 {
		t.Fatal("should be no leftovers")
	}
	if optargs[0].Opt() != "-y" && optargs[0].Arg() != "its a yacc!" {
		t.Fatal("expected to find -y 'its a yacc!'")
	}
}

func Test_Getopt_one_short_arg_one_leftover(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-y", "its a yacc!", "fizzy",
	}
	expected_leftovers := []string{
		"fizzy",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if optargs[0].Opt() != "-y" && optargs[0].Arg() != "its a yacc!" {
		t.Fatal("expected to find -y 'its a yacc!'")
	}
}

func Test_Getopt_one_short_arg_several_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-y", "its a yacc!", "fizzy", "bears", "are", "so", "--tasty",
	}
	expected_leftovers := []string{
		"fizzy", "bears", "are", "so", "--tasty",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if optargs[0].Opt() != "-y" && optargs[0].Arg() != "its a yacc!" {
		t.Fatal("expected to find -y 'its a yacc!'")
	}
}

func Test_Getopt_one_long_no_arg_no_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"--help",
	}
	expected_leftovers := []string{}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if optargs[0].Opt() != "--help" && optargs[0].Arg() != "" {
		t.Fatal("expected to find --help ''")
	}
}

func Test_Getopt_one_long_one_narg_no_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"--example", "charles",
	}
	expected_leftovers := []string{}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if optargs[0].Opt() != "--help" && optargs[0].Arg() != "charles" {
		t.Fatal("expected to find --help 'charles'")
	}
}

func Test_Getopt_one_long_one_earg_no_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"--example=charles",
	}
	expected_leftovers := []string{}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if optargs[0].Opt() != "--help" && optargs[0].Arg() != "charles" {
		t.Fatal("expected to find --help 'charles'")
	}
}

func Test_Getopt_one_long_one_narg_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"--example", "charles",
		"fizzy", "bears", "are", "so", "--tasty",
	}
	expected_leftovers := []string{
		"fizzy", "bears", "are", "so", "--tasty",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if optargs[0].Opt() != "--help" && optargs[0].Arg() != "charles" {
		t.Fatal("expected to find --help 'charles'")
	}
}

func Test_Getopt_one_long_one_earg_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"--example=charles",
		"fizzy", "bears", "are", "so", "--tasty",
	}
	expected_leftovers := []string{
		"fizzy", "bears", "are", "so", "--tasty",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if optargs[0].Opt() != "--help" && optargs[0].Arg() != "charles" {
		t.Fatal("expected to find --help 'charles'")
	}
}

func Test_Getopt_two_short_arg_several_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-h", "-y", "its a yacc!",
		"fizzy", "bears", "are", "so", "--tasty",
	}
	expected_leftovers := []string{
		"fizzy", "bears", "are", "so", "--tasty",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if len(optargs) != 2 {
		t.Log(optargs)
		t.Fatal("expected two optargs")
	}
	if optargs[0].Opt() != "-h" && optargs[0].Arg() != "" {
		t.Fatal("expected to find -h ''")
	}
	if optargs[1].Opt() != "-y" && optargs[1].Arg() != "its a yacc!" {
		t.Fatal("expected to find -y 'its a yacc!'")
	}
}

func Test_Getopt_two_rshort_arg_several_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-hy", "its a yacc!",
		"fizzy", "bears", "are", "so", "--tasty",
	}
	expected_leftovers := []string{
		"fizzy", "bears", "are", "so", "--tasty",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if len(optargs) != 2 {
		t.Log(optargs)
		t.Fatal("expected two optargs")
	}
	if optargs[0].Opt() != "-h" && optargs[0].Arg() != "" {
		t.Fatal("expected to find -h ''")
	}
	if optargs[1].Opt() != "-y" && optargs[1].Arg() != "its a yacc!" {
		t.Fatal("expected to find -y 'its a yacc!'")
	}
}

func Test_Getopt_two_rshort_arg_bad_several_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-zy", "its a yacc!",
		"fizzy", "bears", "are", "so", "--tasty",
	}
	_ = []string{
		"fizzy", "bears", "are", "so", "--tasty",
	}
	_, _, err := GetOpt(input, short, long)
	t.Log(err)
	if err == nil {
		t.Fatal(err)
	}
}

func Test_Getopt_three_short_arg_several_leftovers(t *testing.T) {
	short := "hvx:y:z:e"
	long := []string{
		"help", "verbose", "example=", "yacc=", "zebra=", "empty",
	}
	input := []string{
		"-z", "zebra", "-hy", "its a yacc!",
		"fizzy", "bears", "are", "so", "--tasty",
	}
	expected_leftovers := []string{
		"fizzy", "bears", "are", "so", "--tasty",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if len(optargs) != 3 {
		t.Log(optargs)
		t.Fatal("expected two optargs")
	}
	if optargs[0].Opt() != "-z" && optargs[0].Arg() != "zebra" {
		t.Fatal("expected to find -z 'zebra'")
	}
	if optargs[1].Opt() != "-h" && optargs[1].Arg() != "" {
		t.Fatal("expected to find -h ''")
	}
	if optargs[2].Opt() != "-y" && optargs[2].Arg() != "its a yacc!" {
		t.Fatal("expected to find -y 'its a yacc!'")
	}
}

func Test_Getopt_no_last_arg(t *testing.T) {
	short := "hx:"
	long := []string{
		"help", "example=",
	}
	input := []string{
		"-x",
	}
	_, _, err := GetOpt(input, short, long)
	if err == nil {
		t.Fatal("expected parse error got nil")
	}
	t.Logf("expected err %v", err)
}

func Test_Getopt_empty_arg(t *testing.T) {
	short := ""
	long := []string{}
	input := []string{""}
	expected_leftovers := []string{
		"",
	}
	args, optargs, err := GetOpt(input, short, long)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(args, expected_leftovers) {
		t.Log("got", args)
		t.Log("expected", expected_leftovers)
		t.Log("optargs", optargs[0], len(optargs))
		t.Fatal("recieved wrong leftovers")
	}
	if len(optargs) != 0 {
		t.Log(optargs)
		t.Fatal("expected zero optargs")
	}
}
