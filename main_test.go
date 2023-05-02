package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/bep/helpers/envhelpers"
	"github.com/rogpeppe/go-internal/testscript"
	"golang.org/x/text/language"
)

func TestScripts(t *testing.T) {
	setup := testSetupFunc()
	testscript.Run(t, testscript.Params{
		Dir: "testscripts",
		// UpdateScripts: true, // Uncomment to rewrite the test scripts with
		// TestWork: true, // Uncomment to keep the test work dir.
		Setup: func(env *testscript.Env) error {
			return setup(env)
		},
	})
}

func TestMain(m *testing.M) {
	// This is about https://github.com/rogpeppe/go-internal/issues/200
	// There seem to be a timing issue between writing all the commands to disk and running them.
	// To test this, we need to have some commands that take some time to biuld and write to disk.
	commands := map[string]func() int{
		// This is the command we're calling.
		"myecho": func() int {
			// Include some large Go packages.
			en := language.English
			log.Println("Language:", en.String(), "Greek:", unicode.Is(unicode.Greek, 'A'))
			fmt.Println(strings.Join(os.Args[1:], " "))
			return 0
		},
	}

	for i := 0; i < 50; i++ {
		// Add some more dummy commands.
		commands[fmt.Sprintf("myecho%d_", i)] = commands["myecho"]
	}

	os.Exit(
		testscript.RunMain(m, commands),
	)
}

func testSetupFunc() func(env *testscript.Env) error {
	sourceDir, _ := os.Getwd()
	return func(env *testscript.Env) error {
		var keyVals []string
		// Add some environment variables to the test script.
		keyVals = append(keyVals, "SOURCE", sourceDir)
		envhelpers.SetEnvVars(&env.Vars, keyVals...)

		return nil
	}
}
