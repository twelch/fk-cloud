// Code generated by goa v3.1.1, DO NOT EDIT.
//
// activity HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	activityc "github.com/fieldkit/cloud/server/api/gen/http/activity/client"
	followingc "github.com/fieldkit/cloud/server/api/gen/http/following/client"
	modulesc "github.com/fieldkit/cloud/server/api/gen/http/modules/client"
	projectc "github.com/fieldkit/cloud/server/api/gen/http/project/client"
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
	return `activity (station|project)
following (follow|unfollow|followers)
project update
tasks (five|refresh- device)
test (get|error|email)
modules meta
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` activity station --id 3245699681067865054 --page 6214483488355277271 --auth "Enim occaecati et perferendis asperiores vitae."` + "\n" +
		os.Args[0] + ` following follow --id 343528989186680944 --auth "Perspiciatis doloremque."` + "\n" +
		os.Args[0] + ` project update --body '{
      "body": "Nobis tempore."
   }' --id 6171153226389699817 --auth "Aliquam in quod laborum."` + "\n" +
		os.Args[0] + ` tasks five` + "\n" +
		os.Args[0] + ` test get --id 7577586275964425592` + "\n" +
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
		activityFlags = flag.NewFlagSet("activity", flag.ContinueOnError)

		activityStationFlags    = flag.NewFlagSet("station", flag.ExitOnError)
		activityStationIDFlag   = activityStationFlags.String("id", "REQUIRED", "")
		activityStationPageFlag = activityStationFlags.String("page", "", "")
		activityStationAuthFlag = activityStationFlags.String("auth", "", "")

		activityProjectFlags    = flag.NewFlagSet("project", flag.ExitOnError)
		activityProjectIDFlag   = activityProjectFlags.String("id", "REQUIRED", "")
		activityProjectPageFlag = activityProjectFlags.String("page", "", "")
		activityProjectAuthFlag = activityProjectFlags.String("auth", "", "")

		followingFlags = flag.NewFlagSet("following", flag.ContinueOnError)

		followingFollowFlags    = flag.NewFlagSet("follow", flag.ExitOnError)
		followingFollowIDFlag   = followingFollowFlags.String("id", "REQUIRED", "")
		followingFollowAuthFlag = followingFollowFlags.String("auth", "", "")

		followingUnfollowFlags    = flag.NewFlagSet("unfollow", flag.ExitOnError)
		followingUnfollowIDFlag   = followingUnfollowFlags.String("id", "REQUIRED", "")
		followingUnfollowAuthFlag = followingUnfollowFlags.String("auth", "", "")

		followingFollowersFlags    = flag.NewFlagSet("followers", flag.ExitOnError)
		followingFollowersIDFlag   = followingFollowersFlags.String("id", "REQUIRED", "")
		followingFollowersPageFlag = followingFollowersFlags.String("page", "", "")

		projectFlags = flag.NewFlagSet("project", flag.ContinueOnError)

		projectUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		projectUpdateBodyFlag = projectUpdateFlags.String("body", "REQUIRED", "")
		projectUpdateIDFlag   = projectUpdateFlags.String("id", "REQUIRED", "")
		projectUpdateAuthFlag = projectUpdateFlags.String("auth", "", "")

		tasksFlags = flag.NewFlagSet("tasks", flag.ContinueOnError)

		tasksFiveFlags = flag.NewFlagSet("five", flag.ExitOnError)

		tasksRefreshDeviceFlags        = flag.NewFlagSet("refresh- device", flag.ExitOnError)
		tasksRefreshDeviceDeviceIDFlag = tasksRefreshDeviceFlags.String("device-id", "REQUIRED", "")
		tasksRefreshDeviceAuthFlag     = tasksRefreshDeviceFlags.String("auth", "REQUIRED", "")

		testFlags = flag.NewFlagSet("test", flag.ContinueOnError)

		testGetFlags  = flag.NewFlagSet("get", flag.ExitOnError)
		testGetIDFlag = testGetFlags.String("id", "REQUIRED", "")

		testErrorFlags = flag.NewFlagSet("error", flag.ExitOnError)

		testEmailFlags       = flag.NewFlagSet("email", flag.ExitOnError)
		testEmailAddressFlag = testEmailFlags.String("address", "REQUIRED", "")
		testEmailAuthFlag    = testEmailFlags.String("auth", "REQUIRED", "")

		modulesFlags = flag.NewFlagSet("modules", flag.ContinueOnError)

		modulesMetaFlags = flag.NewFlagSet("meta", flag.ExitOnError)
	)
	activityFlags.Usage = activityUsage
	activityStationFlags.Usage = activityStationUsage
	activityProjectFlags.Usage = activityProjectUsage

	followingFlags.Usage = followingUsage
	followingFollowFlags.Usage = followingFollowUsage
	followingUnfollowFlags.Usage = followingUnfollowUsage
	followingFollowersFlags.Usage = followingFollowersUsage

	projectFlags.Usage = projectUsage
	projectUpdateFlags.Usage = projectUpdateUsage

	tasksFlags.Usage = tasksUsage
	tasksFiveFlags.Usage = tasksFiveUsage
	tasksRefreshDeviceFlags.Usage = tasksRefreshDeviceUsage

	testFlags.Usage = testUsage
	testGetFlags.Usage = testGetUsage
	testErrorFlags.Usage = testErrorUsage
	testEmailFlags.Usage = testEmailUsage

	modulesFlags.Usage = modulesUsage
	modulesMetaFlags.Usage = modulesMetaUsage

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
		case "activity":
			svcf = activityFlags
		case "following":
			svcf = followingFlags
		case "project":
			svcf = projectFlags
		case "tasks":
			svcf = tasksFlags
		case "test":
			svcf = testFlags
		case "modules":
			svcf = modulesFlags
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
		case "activity":
			switch epn {
			case "station":
				epf = activityStationFlags

			case "project":
				epf = activityProjectFlags

			}

		case "following":
			switch epn {
			case "follow":
				epf = followingFollowFlags

			case "unfollow":
				epf = followingUnfollowFlags

			case "followers":
				epf = followingFollowersFlags

			}

		case "project":
			switch epn {
			case "update":
				epf = projectUpdateFlags

			}

		case "tasks":
			switch epn {
			case "five":
				epf = tasksFiveFlags

			case "refresh- device":
				epf = tasksRefreshDeviceFlags

			}

		case "test":
			switch epn {
			case "get":
				epf = testGetFlags

			case "error":
				epf = testErrorFlags

			case "email":
				epf = testEmailFlags

			}

		case "modules":
			switch epn {
			case "meta":
				epf = modulesMetaFlags

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
		case "activity":
			c := activityc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "station":
				endpoint = c.Station()
				data, err = activityc.BuildStationPayload(*activityStationIDFlag, *activityStationPageFlag, *activityStationAuthFlag)
			case "project":
				endpoint = c.Project()
				data, err = activityc.BuildProjectPayload(*activityProjectIDFlag, *activityProjectPageFlag, *activityProjectAuthFlag)
			}
		case "following":
			c := followingc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "follow":
				endpoint = c.Follow()
				data, err = followingc.BuildFollowPayload(*followingFollowIDFlag, *followingFollowAuthFlag)
			case "unfollow":
				endpoint = c.Unfollow()
				data, err = followingc.BuildUnfollowPayload(*followingUnfollowIDFlag, *followingUnfollowAuthFlag)
			case "followers":
				endpoint = c.Followers()
				data, err = followingc.BuildFollowersPayload(*followingFollowersIDFlag, *followingFollowersPageFlag)
			}
		case "project":
			c := projectc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "update":
				endpoint = c.Update()
				data, err = projectc.BuildUpdatePayload(*projectUpdateBodyFlag, *projectUpdateIDFlag, *projectUpdateAuthFlag)
			}
		case "tasks":
			c := tasksc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "five":
				endpoint = c.Five()
				data = nil
			case "refresh- device":
				endpoint = c.RefreshDevice()
				data, err = tasksc.BuildRefreshDevicePayload(*tasksRefreshDeviceDeviceIDFlag, *tasksRefreshDeviceAuthFlag)
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
			case "email":
				endpoint = c.Email()
				data, err = testc.BuildEmailPayload(*testEmailAddressFlag, *testEmailAuthFlag)
			}
		case "modules":
			c := modulesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "meta":
				endpoint = c.Meta()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// activityUsage displays the usage of the activity command and its subcommands.
func activityUsage() {
	fmt.Fprintf(os.Stderr, `Service is the activity service interface.
Usage:
    %s [globalflags] activity COMMAND [flags]

COMMAND:
    station: Station implements station.
    project: Project implements project.

Additional help:
    %s activity COMMAND --help
`, os.Args[0], os.Args[0])
}
func activityStationUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] activity station -id INT64 -page INT64 -auth STRING

Station implements station.
    -id INT64: 
    -page INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` activity station --id 3245699681067865054 --page 6214483488355277271 --auth "Enim occaecati et perferendis asperiores vitae."
`, os.Args[0])
}

func activityProjectUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] activity project -id INT64 -page INT64 -auth STRING

Project implements project.
    -id INT64: 
    -page INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` activity project --id 3564629015996628995 --page 6264850308329228339 --auth "Voluptatem a in provident ullam facilis."
`, os.Args[0])
}

// followingUsage displays the usage of the following command and its
// subcommands.
func followingUsage() {
	fmt.Fprintf(os.Stderr, `Service is the following service interface.
Usage:
    %s [globalflags] following COMMAND [flags]

COMMAND:
    follow: Follow implements follow.
    unfollow: Unfollow implements unfollow.
    followers: Followers implements followers.

Additional help:
    %s following COMMAND --help
`, os.Args[0], os.Args[0])
}
func followingFollowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] following follow -id INT64 -auth STRING

Follow implements follow.
    -id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` following follow --id 343528989186680944 --auth "Perspiciatis doloremque."
`, os.Args[0])
}

func followingUnfollowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] following unfollow -id INT64 -auth STRING

Unfollow implements unfollow.
    -id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` following unfollow --id 3639159256254407470 --auth "Non sed a nostrum."
`, os.Args[0])
}

func followingFollowersUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] following followers -id INT64 -page INT64

Followers implements followers.
    -id INT64: 
    -page INT64: 

Example:
    `+os.Args[0]+` following followers --id 5065645685327175683 --page 8922159228118147378
`, os.Args[0])
}

// projectUsage displays the usage of the project command and its subcommands.
func projectUsage() {
	fmt.Fprintf(os.Stderr, `Service is the project service interface.
Usage:
    %s [globalflags] project COMMAND [flags]

COMMAND:
    update: Update implements update.

Additional help:
    %s project COMMAND --help
`, os.Args[0], os.Args[0])
}
func projectUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project update -body JSON -id INT64 -auth STRING

Update implements update.
    -body JSON: 
    -id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` project update --body '{
      "body": "Nobis tempore."
   }' --id 6171153226389699817 --auth "Aliquam in quod laborum."
`, os.Args[0])
}

// tasksUsage displays the usage of the tasks command and its subcommands.
func tasksUsage() {
	fmt.Fprintf(os.Stderr, `Service is the tasks service interface.
Usage:
    %s [globalflags] tasks COMMAND [flags]

COMMAND:
    five: Five implements five.
    refresh- device: RefreshDevice implements refresh device.

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

func tasksRefreshDeviceUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks refresh- device -device-id STRING -auth STRING

RefreshDevice implements refresh device.
    -device-id STRING: 
    -auth STRING: 

Example:
    `+os.Args[0]+` tasks refresh- device --device-id "Fugit dolor consequatur aliquam dignissimos quia nobis." --auth "Quia omnis fugiat."
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
    email: Email implements email.

Additional help:
    %s test COMMAND --help
`, os.Args[0], os.Args[0])
}
func testGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] test get -id INT64

Get implements get.
    -id INT64: 

Example:
    `+os.Args[0]+` test get --id 7577586275964425592
`, os.Args[0])
}

func testErrorUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] test error

Error implements error.

Example:
    `+os.Args[0]+` test error
`, os.Args[0])
}

func testEmailUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] test email -address STRING -auth STRING

Email implements email.
    -address STRING: 
    -auth STRING: 

Example:
    `+os.Args[0]+` test email --address "Dolore earum perspiciatis." --auth "Esse et consequatur at amet perferendis quibusdam."
`, os.Args[0])
}

// modulesUsage displays the usage of the modules command and its subcommands.
func modulesUsage() {
	fmt.Fprintf(os.Stderr, `Service is the modules service interface.
Usage:
    %s [globalflags] modules COMMAND [flags]

COMMAND:
    meta: Meta implements meta.

Additional help:
    %s modules COMMAND --help
`, os.Args[0], os.Args[0])
}
func modulesMetaUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] modules meta

Meta implements meta.

Example:
    `+os.Args[0]+` modules meta
`, os.Args[0])
}
