package easytemplate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Compile_EmptyTemplate_ShouldCompileWithoutErrors(t *testing.T) {
	_, err := Compile[int]("")

	assert.NoError(t, err)
}

func Test_Compile_TemplateBasedOnPrimitiveType_ShouldCompileWhenWeUseDotPlaceholder(t *testing.T) {
	tmpl, err := Compile[int]("value: {{.}}")

	assert.NoError(t, err)
	assert.Equal(t, "value: 123", tmpl.ExecuteSimple(123))
}

func Test_Compile_TemplateBasedOnPrimitiveType_ShouldNotCompile_WhenWeUseNamedPlaceholders(t *testing.T) {
	_, err := Compile[int]("value: {{.Name}}")

	assert.Error(t, err)
}

func Test_Compile_TemplateBasedOnStruct_ShouldCompileWhenWeDoNotUsePlaceholders(t *testing.T) {
	type Data struct {
		Name string
	}

	_, err := Compile[Data]("")

	assert.NoError(t, err)
}

func Test_Compile_TemplateBasedOnStruct_ShouldNotCompile_WhenWeHaveErrorInPlaceholderName(t *testing.T) {
	type Data struct {
		Name string
	}

	_, err := Compile[Data]("name: {{.Nme}}")

	assert.Error(t, err)
}

func Test_Compile_TemplateBasedOnStruct_ShouldCompile_WhenWeHaveCorrectPlaceholders(t *testing.T) {
	type Data struct {
		Name string
	}

	tmpl, err := Compile[Data]("name: {{.Name}}")

	assert.NoError(t, err)
	assert.Equal(t, "name: Alex", tmpl.ExecuteSimple(Data{"Alex"}))
}

func Test_Compile_TemplateWithPointerFields_ShouldCompile_WhenWeProvideProperlyInitializedObject(t *testing.T) {
	type User struct {
		Name string
	}
	type Data struct {
		User *User
	}

	tmpl, err := CompileWithExample("name: {{.User.Name}}", Data{
		User: &User{
			Name: "Alex",
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, "name: Alex", tmpl.ExecuteSimple(Data{
		User: &User{
			Name: "Alex",
		},
	}))
}
