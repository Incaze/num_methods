package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Env struct {
	a float64

	x0 float64
	xN float64
	n  int

	t0 float64
	tM float64
	m  int

	precision int
}

func (env *Env) GetA() float64 {
	return env.a
}

func (env *Env) GetX0() float64 {
	return env.x0
}

func (env *Env) GetXN() float64 {
	return env.xN
}

func (env *Env) GetT0() float64 {
	return env.t0
}

func (env *Env) GetTm() float64 {
	return env.tM
}

func (env *Env) GetPrecision() int {
	return env.precision
}

func (env *Env) GetN() int {
	return env.n
}

func (env *Env) GetM() int {
	return env.m
}

func getExecutablePath() (string, error) {
	bin, err := os.Executable()
	if err != nil {
		return "", nil
	}
	return strings.TrimSuffix(filepath.Dir(bin), string(os.PathSeparator)), nil
}

func parseFloat64(key string) float64 {
	stringVal := os.Getenv(key)
	res, err := strconv.ParseFloat(stringVal, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func parseInt(key string) int {
	stringVal := os.Getenv(key)
	res, err := strconv.Atoi(stringVal)
	if err != nil {
		panic(err)
	}
	return res
}

func CreateEnv() *Env {
	path, err := getExecutablePath()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(fmt.Sprintf("%s/.env", path))
	if err != nil {
		panic(err)
	}

	env := new(Env)
	env.a = parseFloat64("a")
	env.x0 = parseFloat64("x0")
	env.xN = parseFloat64("xN")
	env.n = parseInt("n")
	env.t0 = parseFloat64("t0")
	env.tM = parseFloat64("tM")
	env.m = parseInt("m")
	env.precision = parseInt("precision")

	return env
}
