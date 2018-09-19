package lib

import (
	"fmt"
	"os"
	"regexp"
)

type KVArg struct {
	Key      string
	value    string
	Regexp   string
	Required bool
	Extended bool
	bvalue   bool
	set      bool
}

func (kva KVArg) GetKey() string {
	return kva.Key
}

func (kva KVArg) IsRequired() bool {
	return kva.Required
}

func (kva KVArg) GetValue() interface{} {
	if kva.Extended {
		return kva.value
	}
	return kva.bvalue
}

func (kva *KVArg) SetValue(args []string, i int, l int) (err error) {
	if kva.Extended {
		if i+1 == l {
			return fmt.Errorf("GetArgs: missing value for '%s'", kva.GetKey())
		}
		var re *regexp.Regexp
		if re, err = regexp.Compile(kva.Regexp); err != nil {
			return
		}
		if re.Match([]byte(args[i+1])) {
			kva.value = args[i+1]
		} else {
			return fmt.Errorf("Incorrect value given for key '%s'", kva.GetKey())
		}
	} else {
		kva.bvalue = true
	}
	kva.set = true
	return
}

func (kva KVArg) IsValid() bool {
	return kva.set
}

func getArgReq(str string, argsFmt []*KVArg) (f *KVArg, err error) {
	for _, f = range argsFmt {
		if str == f.GetKey() {
			return
		}
	}
	err = fmt.Errorf("'%s' is not in args list", str)
	return
}

func isValid(args []*KVArg) (err error) {
	for _, a := range args {
		if a.IsRequired() && !a.IsValid() {
			err = fmt.Errorf("Missing arg '%s'", a.GetKey())
			return
		}
	}
	return nil
}

func GetArgs(argsFmt []*KVArg) (err error) {
	// mand, opt := splitArgs(argsFmt)
	args := os.Args[1:]
	var arg *KVArg
	skipNext := false
	l := len(args)
	for i, str := range args {
		if skipNext {
			skipNext = false
			continue
		}
		arg, err = getArgReq(str, argsFmt)
		if err != nil {
			return
		}
		if err = arg.SetValue(args, i, l); err != nil {
			return
		}
		if arg.Extended {
			skipNext = true
		}
	}
	if err = isValid(argsFmt); err != nil {
		return
	}
	return nil
}
