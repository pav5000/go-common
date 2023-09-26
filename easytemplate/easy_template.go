// Package easytemplate is designed for easier usage of text/template in simple cases
// it's more type safe than vanilla text/template
// it checks on compilation if provided generic type param matches the template
package easytemplate

import (
	"bytes"
	"errors"
	"text/template"
)

type Template[T any] struct {
	tmpl *template.Template
}

// MustCompile compiles and tries to run a template with empty data object
// to check if the template is correct
// panics on error
func MustCompile[T any](templateString string) *Template[T] {
	tmpl, err := Compile[T](templateString)
	if err != nil {
		panic(err)
	}
	return tmpl
}

// Compile compiles and tries to run a template with empty data object
// to check if the template is correct
func Compile[T any](templateString string) (*Template[T], error) {
	var data T // we just use empty data object
	return CompileWithExample(templateString, data)
}

// CompileWithExample compiles and tries to run a template with the provided example data object
// to check if the template is correct
// you may need this func if your data object contains pointers, maps or any other objects you need to initialize
// but I recommend you to keep it simple and do not use them
func CompileWithExample[T any](templateString string, example T) (*Template[T], error) {
	rawTmpl, err := template.New("").Parse(templateString)
	if err != nil {
		return nil, err
	}
	tmpl := &Template[T]{
		tmpl: rawTmpl,
	}

	// we try to compile and execute a template with example data
	// to check if placeholders inside template are correct
	_, err = tmpl.Execute(example)
	if err != nil {
		return nil, errors.New("error while execution check with example data object: " + err.Error())
	}

	return tmpl, nil
}

// Execute generates a string based on the template and the provided data
// in most cases you won't get an error here because the template was checked on compilation
// but in some corner cases when you template is complicated you may get an error here
func (t *Template[T]) Execute(data T) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := t.tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ExecuteSimple does the same as Execute but ignores errors
func (t *Template[T]) ExecuteSimple(data T) string {
	str, _ := t.Execute(data)
	return str
}
