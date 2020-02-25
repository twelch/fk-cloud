// Code generated by goa v3.0.10, DO NOT EDIT.
//
// tasks HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	tasksc "github.com/fieldkit/cloud/server/api/gen/http/tasks/client"
	testc "github.com/fieldkit/cloud/server/api/gen/http/test/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `tasks five
test (get|error|json)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` tasks five` + "\n" +
		os.Args[0] + ` test get --id 6459967034483944467` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		tasksFlags = flag.NewFlagSet("tasks", flag.ContinueOnError)

		tasksFiveFlags = flag.NewFlagSet("five", flag.ExitOnError)

		testFlags = flag.NewFlagSet("test", flag.ContinueOnError)

		testGetFlags  = flag.NewFlagSet("get", flag.ExitOnError)
		testGetIDFlag = testGetFlags.String("id", "REQUIRED", "")

		testErrorFlags = flag.NewFlagSet("error", flag.ExitOnError)

		testJSONFlags  = flag.NewFlagSet("json", flag.ExitOnError)
		testJSONIDFlag = testJSONFlags.String("id", "REQUIRED", "")
	)
	tasksFlags.Usage = tasksUsage
	tasksFiveFlags.Usage = tasksFiveUsage

	testFlags.Usage = testUsage
	testGetFlags.Usage = testGetUsage
	testErrorFlags.Usage = testErrorUsage
	testJSONFlags.Usage = testJSONUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "tasks":
			svcf = tasksFlags
		case "test":
			svcf = testFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "tasks":
			switch epn {
			case "five":
				epf = tasksFiveFlags

			}

		case "test":
			switch epn {
			case "get":
				epf = testGetFlags

			case "error":
				epf = testErrorFlags

			case "json":
				epf = testJSONFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "tasks":
			c := tasksc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "five":
				endpoint = c.Five()
				data = nil
			}
		case "test":
			c := testc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get":
				endpoint = c.Get()
				data, err = testc.BuildGetPayload(*testGetIDFlag)
			case "error":
				endpoint = c.Error()
				data = nil
			case "json":
				endpoint = c.JSON()
				data, err = testc.BuildJSONPayload(*testJSONIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// tasksUsage displays the usage of the tasks command and its subcommands.
func tasksUsage() {
	fmt.Fprintf(os.Stderr, `Service is the tasks service interface.
Usage:
    %s [globalflags] tasks COMMAND [flags]

COMMAND:
    five: Five implements five.

Additional help:
    %s tasks COMMAND --help
`, os.Args[0], os.Args[0])
}
func tasksFiveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks five

Five implements five.

Example:
    `+os.Args[0]+` tasks five
`, os.Args[0])
}

// testUsage displays the usage of the test command and its subcommands.
func testUsage() {
	fmt.Fprintf(os.Stderr, `Service is the test service interface.
Usage:
    %s [globalflags] test COMMAND [flags]

COMMAND:
    get: Get implements get.
    error: Error implements error.
    json: JSON implements json.

Additional help:
    %s test COMMAND --help
`, os.Args[0], os.Args[0])
}
func testGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] test get -id INT64

Get implements get.
    -id INT64: 

Example:
    `+os.Args[0]+` test get --id 6459967034483944467
`, os.Args[0])
}

func testErrorUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] test error

Error implements error.

Example:
    `+os.Args[0]+` test error
`, os.Args[0])
}

func testJSONUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] test json -id INT64

JSON implements json.
    -id INT64: 

Example:
    `+os.Args[0]+` test json --id 459891176067806689
`, os.Args[0])
}