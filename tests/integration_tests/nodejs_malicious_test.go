package integrationtests_test

import (
	"strings"
	"testing"

	"github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/service"
)

func TestNodejsRunCommand(t *testing.T) {
	// Test case for run_command
	resp := service.RunNodeJsCode(`
const { spawn } = require( 'child_process' );
const ls = spawn( 'ls', [ '-lh', '/usr' ] );

ls.stdout.on( 'data', ( data ) => {
    console.log(data);
} );

ls.stderr.on( 'data', ( data ) => {
    console.log(data);
} );

ls.on( 'close', ( code ) => {
    console.log(cod