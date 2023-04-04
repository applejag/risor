package vm

import (
	"github.com/cloudcmds/tamarin/evaluator"
	"github.com/cloudcmds/tamarin/internal/compiler"
	modJson "github.com/cloudcmds/tamarin/modules/json"
	modMath "github.com/cloudcmds/tamarin/modules/math"
	modPgx "github.com/cloudcmds/tamarin/modules/pgx"
	modRand "github.com/cloudcmds/tamarin/modules/rand"
	modStrconv "github.com/cloudcmds/tamarin/modules/strconv"
	modStrings "github.com/cloudcmds/tamarin/modules/strings"
	modTime "github.com/cloudcmds/tamarin/modules/time"
	modUuid "github.com/cloudcmds/tamarin/modules/uuid"
	"github.com/cloudcmds/tamarin/object"
	"github.com/cloudcmds/tamarin/parser"
)

type Interpreter struct {
	c    *compiler.Compiler
	main *compiler.Scope
	// vm   *vm.VM
}

func NewInterpreter(builtins []*object.Builtin) *Interpreter {

	bmap := map[string]object.Object{}
	for _, b := range evaluator.GlobalBuiltins() {
		bmap[b.Key()] = b
	}

	bmap["math"] = modMath.Module()
	bmap["json"] = modJson.Module()
	bmap["strings"] = modStrings.Module()
	bmap["time"] = modTime.Module()
	bmap["uuid"] = modUuid.Module()
	bmap["rand"] = modRand.Module()
	bmap["strconv"] = modStrconv.Module()
	bmap["pgx"] = modPgx.Module()

	s := compiler.NewScope("main")

	c := compiler.New(compiler.Options{
		Builtins: bmap,
		Name:     "main",
		Scope:    s,
	})

	return &Interpreter{c: c, main: s}
}

func (i *Interpreter) Eval(code string) (object.Object, error) {
	ast, err := parser.Parse(code)
	if err != nil {
		return nil, err
	}
	pos := len(i.c.Instructions())
	_, err = i.c.Compile(ast)
	if err != nil {
		return nil, err
	}
	v := NewWithOffset(i.main, pos)
	if err := v.Run(); err != nil {
		return nil, err
	}
	result, exists := v.TOS()
	if exists {
		return result, nil
	}
	return object.Nil, nil
}
