// Code generated by goa v3.1.2, DO NOT EDIT.
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
	ingestionc "github.com/fieldkit/cloud/server/api/gen/http/ingestion/client"
	modulesc "github.com/fieldkit/cloud/server/api/gen/http/modules/client"
	projectc "github.com/fieldkit/cloud/server/api/gen/http/project/client"
	stationc "github.com/fieldkit/cloud/server/api/gen/http/station/client"
	tasksc "github.com/fieldkit/cloud/server/api/gen/http/tasks/client"
	testc "github.com/fieldkit/cloud/server/api/gen/http/test/client"
	userc "github.com/fieldkit/cloud/server/api/gen/http/user/client"
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
ingestion (process- pending|process- station|process- ingestion|delete)
modules meta
project (add- update|delete- update|modify- update|invites|lookup- invite|accept- invite|reject- invite)
station (add|get|update|list- mine|list- project|photo)
tasks (five|refresh- device)
test (get|error|email)
user (roles|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` activity station --id 4850889053895604672 --page 4702121072547338975 --auth "Maiores ut qui quisquam molestiae."` + "\n" +
		os.Args[0] + ` following follow --id 7938282212268354146 --auth "Id quibusdam autem reprehenderit."` + "\n" +
		os.Args[0] + ` ingestion process- pending --auth "Voluptatem possimus sint sint aut."` + "\n" +
		os.Args[0] + ` modules meta` + "\n" +
		os.Args[0] + ` project add- update --body '{
      "body": "Et quis optio voluptatibus quas dolor."
   }' --project-id 520607661 --auth "Impedit quidem labore."` + "\n" +
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

		ingestionFlags = flag.NewFlagSet("ingestion", flag.ContinueOnError)

		ingestionProcessPendingFlags    = flag.NewFlagSet("process- pending", flag.ExitOnError)
		ingestionProcessPendingAuthFlag = ingestionProcessPendingFlags.String("auth", "REQUIRED", "")

		ingestionProcessStationFlags         = flag.NewFlagSet("process- station", flag.ExitOnError)
		ingestionProcessStationStationIDFlag = ingestionProcessStationFlags.String("station-id", "REQUIRED", "")
		ingestionProcessStationAuthFlag      = ingestionProcessStationFlags.String("auth", "REQUIRED", "")

		ingestionProcessIngestionFlags           = flag.NewFlagSet("process- ingestion", flag.ExitOnError)
		ingestionProcessIngestionIngestionIDFlag = ingestionProcessIngestionFlags.String("ingestion-id", "REQUIRED", "")
		ingestionProcessIngestionAuthFlag        = ingestionProcessIngestionFlags.String("auth", "REQUIRED", "")

		ingestionDeleteFlags           = flag.NewFlagSet("delete", flag.ExitOnError)
		ingestionDeleteIngestionIDFlag = ingestionDeleteFlags.String("ingestion-id", "REQUIRED", "")
		ingestionDeleteAuthFlag        = ingestionDeleteFlags.String("auth", "REQUIRED", "")

		modulesFlags = flag.NewFlagSet("modules", flag.ContinueOnError)

		modulesMetaFlags = flag.NewFlagSet("meta", flag.ExitOnError)

		projectFlags = flag.NewFlagSet("project", flag.ContinueOnError)

		projectAddUpdateFlags         = flag.NewFlagSet("add- update", flag.ExitOnError)
		projectAddUpdateBodyFlag      = projectAddUpdateFlags.String("body", "REQUIRED", "")
		projectAddUpdateProjectIDFlag = projectAddUpdateFlags.String("project-id", "REQUIRED", "")
		projectAddUpdateAuthFlag      = projectAddUpdateFlags.String("auth", "REQUIRED", "")

		projectDeleteUpdateFlags         = flag.NewFlagSet("delete- update", flag.ExitOnError)
		projectDeleteUpdateProjectIDFlag = projectDeleteUpdateFlags.String("project-id", "REQUIRED", "")
		projectDeleteUpdateUpdateIDFlag  = projectDeleteUpdateFlags.String("update-id", "REQUIRED", "")
		projectDeleteUpdateAuthFlag      = projectDeleteUpdateFlags.String("auth", "REQUIRED", "")

		projectModifyUpdateFlags         = flag.NewFlagSet("modify- update", flag.ExitOnError)
		projectModifyUpdateBodyFlag      = projectModifyUpdateFlags.String("body", "REQUIRED", "")
		projectModifyUpdateProjectIDFlag = projectModifyUpdateFlags.String("project-id", "REQUIRED", "")
		projectModifyUpdateUpdateIDFlag  = projectModifyUpdateFlags.String("update-id", "REQUIRED", "")
		projectModifyUpdateAuthFlag      = projectModifyUpdateFlags.String("auth", "REQUIRED", "")

		projectInvitesFlags    = flag.NewFlagSet("invites", flag.ExitOnError)
		projectInvitesAuthFlag = projectInvitesFlags.String("auth", "REQUIRED", "")

		projectLookupInviteFlags     = flag.NewFlagSet("lookup- invite", flag.ExitOnError)
		projectLookupInviteTokenFlag = projectLookupInviteFlags.String("token", "REQUIRED", "")
		projectLookupInviteAuthFlag  = projectLookupInviteFlags.String("auth", "REQUIRED", "")

		projectAcceptInviteFlags     = flag.NewFlagSet("accept- invite", flag.ExitOnError)
		projectAcceptInviteIDFlag    = projectAcceptInviteFlags.String("id", "REQUIRED", "")
		projectAcceptInviteTokenFlag = projectAcceptInviteFlags.String("token", "", "")
		projectAcceptInviteAuthFlag  = projectAcceptInviteFlags.String("auth", "REQUIRED", "")

		projectRejectInviteFlags     = flag.NewFlagSet("reject- invite", flag.ExitOnError)
		projectRejectInviteIDFlag    = projectRejectInviteFlags.String("id", "REQUIRED", "")
		projectRejectInviteTokenFlag = projectRejectInviteFlags.String("token", "", "")
		projectRejectInviteAuthFlag  = projectRejectInviteFlags.String("auth", "REQUIRED", "")

		stationFlags = flag.NewFlagSet("station", flag.ContinueOnError)

		stationAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		stationAddBodyFlag = stationAddFlags.String("body", "REQUIRED", "")
		stationAddAuthFlag = stationAddFlags.String("auth", "REQUIRED", "")

		stationGetFlags    = flag.NewFlagSet("get", flag.ExitOnError)
		stationGetIDFlag   = stationGetFlags.String("id", "REQUIRED", "")
		stationGetAuthFlag = stationGetFlags.String("auth", "REQUIRED", "")

		stationUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		stationUpdateBodyFlag = stationUpdateFlags.String("body", "REQUIRED", "")
		stationUpdateIDFlag   = stationUpdateFlags.String("id", "REQUIRED", "")
		stationUpdateAuthFlag = stationUpdateFlags.String("auth", "REQUIRED", "")

		stationListMineFlags    = flag.NewFlagSet("list- mine", flag.ExitOnError)
		stationListMineAuthFlag = stationListMineFlags.String("auth", "REQUIRED", "")

		stationListProjectFlags    = flag.NewFlagSet("list- project", flag.ExitOnError)
		stationListProjectIDFlag   = stationListProjectFlags.String("id", "REQUIRED", "")
		stationListProjectAuthFlag = stationListProjectFlags.String("auth", "REQUIRED", "")

		stationPhotoFlags    = flag.NewFlagSet("photo", flag.ExitOnError)
		stationPhotoIDFlag   = stationPhotoFlags.String("id", "REQUIRED", "")
		stationPhotoAuthFlag = stationPhotoFlags.String("auth", "REQUIRED", "")

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

		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userRolesFlags    = flag.NewFlagSet("roles", flag.ExitOnError)
		userRolesAuthFlag = userRolesFlags.String("auth", "REQUIRED", "")

		userDeleteFlags      = flag.NewFlagSet("delete", flag.ExitOnError)
		userDeleteUserIDFlag = userDeleteFlags.String("user-id", "REQUIRED", "")
		userDeleteAuthFlag   = userDeleteFlags.String("auth", "REQUIRED", "")
	)
	activityFlags.Usage = activityUsage
	activityStationFlags.Usage = activityStationUsage
	activityProjectFlags.Usage = activityProjectUsage

	followingFlags.Usage = followingUsage
	followingFollowFlags.Usage = followingFollowUsage
	followingUnfollowFlags.Usage = followingUnfollowUsage
	followingFollowersFlags.Usage = followingFollowersUsage

	ingestionFlags.Usage = ingestionUsage
	ingestionProcessPendingFlags.Usage = ingestionProcessPendingUsage
	ingestionProcessStationFlags.Usage = ingestionProcessStationUsage
	ingestionProcessIngestionFlags.Usage = ingestionProcessIngestionUsage
	ingestionDeleteFlags.Usage = ingestionDeleteUsage

	modulesFlags.Usage = modulesUsage
	modulesMetaFlags.Usage = modulesMetaUsage

	projectFlags.Usage = projectUsage
	projectAddUpdateFlags.Usage = projectAddUpdateUsage
	projectDeleteUpdateFlags.Usage = projectDeleteUpdateUsage
	projectModifyUpdateFlags.Usage = projectModifyUpdateUsage
	projectInvitesFlags.Usage = projectInvitesUsage
	projectLookupInviteFlags.Usage = projectLookupInviteUsage
	projectAcceptInviteFlags.Usage = projectAcceptInviteUsage
	projectRejectInviteFlags.Usage = projectRejectInviteUsage

	stationFlags.Usage = stationUsage
	stationAddFlags.Usage = stationAddUsage
	stationGetFlags.Usage = stationGetUsage
	stationUpdateFlags.Usage = stationUpdateUsage
	stationListMineFlags.Usage = stationListMineUsage
	stationListProjectFlags.Usage = stationListProjectUsage
	stationPhotoFlags.Usage = stationPhotoUsage

	tasksFlags.Usage = tasksUsage
	tasksFiveFlags.Usage = tasksFiveUsage
	tasksRefreshDeviceFlags.Usage = tasksRefreshDeviceUsage

	testFlags.Usage = testUsage
	testGetFlags.Usage = testGetUsage
	testErrorFlags.Usage = testErrorUsage
	testEmailFlags.Usage = testEmailUsage

	userFlags.Usage = userUsage
	userRolesFlags.Usage = userRolesUsage
	userDeleteFlags.Usage = userDeleteUsage

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
		case "ingestion":
			svcf = ingestionFlags
		case "modules":
			svcf = modulesFlags
		case "project":
			svcf = projectFlags
		case "station":
			svcf = stationFlags
		case "tasks":
			svcf = tasksFlags
		case "test":
			svcf = testFlags
		case "user":
			svcf = userFlags
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

		case "ingestion":
			switch epn {
			case "process- pending":
				epf = ingestionProcessPendingFlags

			case "process- station":
				epf = ingestionProcessStationFlags

			case "process- ingestion":
				epf = ingestionProcessIngestionFlags

			case "delete":
				epf = ingestionDeleteFlags

			}

		case "modules":
			switch epn {
			case "meta":
				epf = modulesMetaFlags

			}

		case "project":
			switch epn {
			case "add- update":
				epf = projectAddUpdateFlags

			case "delete- update":
				epf = projectDeleteUpdateFlags

			case "modify- update":
				epf = projectModifyUpdateFlags

			case "invites":
				epf = projectInvitesFlags

			case "lookup- invite":
				epf = projectLookupInviteFlags

			case "accept- invite":
				epf = projectAcceptInviteFlags

			case "reject- invite":
				epf = projectRejectInviteFlags

			}

		case "station":
			switch epn {
			case "add":
				epf = stationAddFlags

			case "get":
				epf = stationGetFlags

			case "update":
				epf = stationUpdateFlags

			case "list- mine":
				epf = stationListMineFlags

			case "list- project":
				epf = stationListProjectFlags

			case "photo":
				epf = stationPhotoFlags

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

		case "user":
			switch epn {
			case "roles":
				epf = userRolesFlags

			case "delete":
				epf = userDeleteFlags

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
		case "ingestion":
			c := ingestionc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "process- pending":
				endpoint = c.ProcessPending()
				data, err = ingestionc.BuildProcessPendingPayload(*ingestionProcessPendingAuthFlag)
			case "process- station":
				endpoint = c.ProcessStation()
				data, err = ingestionc.BuildProcessStationPayload(*ingestionProcessStationStationIDFlag, *ingestionProcessStationAuthFlag)
			case "process- ingestion":
				endpoint = c.ProcessIngestion()
				data, err = ingestionc.BuildProcessIngestionPayload(*ingestionProcessIngestionIngestionIDFlag, *ingestionProcessIngestionAuthFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = ingestionc.BuildDeletePayload(*ingestionDeleteIngestionIDFlag, *ingestionDeleteAuthFlag)
			}
		case "modules":
			c := modulesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "meta":
				endpoint = c.Meta()
				data = nil
			}
		case "project":
			c := projectc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add- update":
				endpoint = c.AddUpdate()
				data, err = projectc.BuildAddUpdatePayload(*projectAddUpdateBodyFlag, *projectAddUpdateProjectIDFlag, *projectAddUpdateAuthFlag)
			case "delete- update":
				endpoint = c.DeleteUpdate()
				data, err = projectc.BuildDeleteUpdatePayload(*projectDeleteUpdateProjectIDFlag, *projectDeleteUpdateUpdateIDFlag, *projectDeleteUpdateAuthFlag)
			case "modify- update":
				endpoint = c.ModifyUpdate()
				data, err = projectc.BuildModifyUpdatePayload(*projectModifyUpdateBodyFlag, *projectModifyUpdateProjectIDFlag, *projectModifyUpdateUpdateIDFlag, *projectModifyUpdateAuthFlag)
			case "invites":
				endpoint = c.Invites()
				data, err = projectc.BuildInvitesPayload(*projectInvitesAuthFlag)
			case "lookup- invite":
				endpoint = c.LookupInvite()
				data, err = projectc.BuildLookupInvitePayload(*projectLookupInviteTokenFlag, *projectLookupInviteAuthFlag)
			case "accept- invite":
				endpoint = c.AcceptInvite()
				data, err = projectc.BuildAcceptInvitePayload(*projectAcceptInviteIDFlag, *projectAcceptInviteTokenFlag, *projectAcceptInviteAuthFlag)
			case "reject- invite":
				endpoint = c.RejectInvite()
				data, err = projectc.BuildRejectInvitePayload(*projectRejectInviteIDFlag, *projectRejectInviteTokenFlag, *projectRejectInviteAuthFlag)
			}
		case "station":
			c := stationc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = stationc.BuildAddPayload(*stationAddBodyFlag, *stationAddAuthFlag)
			case "get":
				endpoint = c.Get()
				data, err = stationc.BuildGetPayload(*stationGetIDFlag, *stationGetAuthFlag)
			case "update":
				endpoint = c.Update()
				data, err = stationc.BuildUpdatePayload(*stationUpdateBodyFlag, *stationUpdateIDFlag, *stationUpdateAuthFlag)
			case "list- mine":
				endpoint = c.ListMine()
				data, err = stationc.BuildListMinePayload(*stationListMineAuthFlag)
			case "list- project":
				endpoint = c.ListProject()
				data, err = stationc.BuildListProjectPayload(*stationListProjectIDFlag, *stationListProjectAuthFlag)
			case "photo":
				endpoint = c.Photo()
				data, err = stationc.BuildPhotoPayload(*stationPhotoIDFlag, *stationPhotoAuthFlag)
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
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "roles":
				endpoint = c.Roles()
				data, err = userc.BuildRolesPayload(*userRolesAuthFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = userc.BuildDeletePayload(*userDeleteUserIDFlag, *userDeleteAuthFlag)
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
    `+os.Args[0]+` activity station --id 4850889053895604672 --page 4702121072547338975 --auth "Maiores ut qui quisquam molestiae."
`, os.Args[0])
}

func activityProjectUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] activity project -id INT64 -page INT64 -auth STRING

Project implements project.
    -id INT64: 
    -page INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` activity project --id 2588221656404062978 --page 8533501338479540862 --auth "Quam voluptatem illum ea dolorem adipisci."
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
    `+os.Args[0]+` following follow --id 7938282212268354146 --auth "Id quibusdam autem reprehenderit."
`, os.Args[0])
}

func followingUnfollowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] following unfollow -id INT64 -auth STRING

Unfollow implements unfollow.
    -id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` following unfollow --id 668004131381466168 --auth "Iusto porro qui."
`, os.Args[0])
}

func followingFollowersUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] following followers -id INT64 -page INT64

Followers implements followers.
    -id INT64: 
    -page INT64: 

Example:
    `+os.Args[0]+` following followers --id 6162833146368073442 --page 340396845702908308
`, os.Args[0])
}

// ingestionUsage displays the usage of the ingestion command and its
// subcommands.
func ingestionUsage() {
	fmt.Fprintf(os.Stderr, `Service is the ingestion service interface.
Usage:
    %s [globalflags] ingestion COMMAND [flags]

COMMAND:
    process- pending: ProcessPending implements process pending.
    process- station: ProcessStation implements process station.
    process- ingestion: ProcessIngestion implements process ingestion.
    delete: Delete implements delete.

Additional help:
    %s ingestion COMMAND --help
`, os.Args[0], os.Args[0])
}
func ingestionProcessPendingUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ingestion process- pending -auth STRING

ProcessPending implements process pending.
    -auth STRING: 

Example:
    `+os.Args[0]+` ingestion process- pending --auth "Voluptatem possimus sint sint aut."
`, os.Args[0])
}

func ingestionProcessStationUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ingestion process- station -station-id INT32 -auth STRING

ProcessStation implements process station.
    -station-id INT32: 
    -auth STRING: 

Example:
    `+os.Args[0]+` ingestion process- station --station-id 1858509964 --auth "Quis autem maxime quis nihil quia."
`, os.Args[0])
}

func ingestionProcessIngestionUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ingestion process- ingestion -ingestion-id INT64 -auth STRING

ProcessIngestion implements process ingestion.
    -ingestion-id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` ingestion process- ingestion --ingestion-id 8144715309246928875 --auth "Nobis eos."
`, os.Args[0])
}

func ingestionDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] ingestion delete -ingestion-id INT64 -auth STRING

Delete implements delete.
    -ingestion-id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` ingestion delete --ingestion-id 5552486977235119271 --auth "Blanditiis aut dolor est sequi vel."
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

// projectUsage displays the usage of the project command and its subcommands.
func projectUsage() {
	fmt.Fprintf(os.Stderr, `Service is the project service interface.
Usage:
    %s [globalflags] project COMMAND [flags]

COMMAND:
    add- update: AddUpdate implements add update.
    delete- update: DeleteUpdate implements delete update.
    modify- update: ModifyUpdate implements modify update.
    invites: Invites implements invites.
    lookup- invite: LookupInvite implements lookup invite.
    accept- invite: AcceptInvite implements accept invite.
    reject- invite: RejectInvite implements reject invite.

Additional help:
    %s project COMMAND --help
`, os.Args[0], os.Args[0])
}
func projectAddUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project add- update -body JSON -project-id INT32 -auth STRING

AddUpdate implements add update.
    -body JSON: 
    -project-id INT32: 
    -auth STRING: 

Example:
    `+os.Args[0]+` project add- update --body '{
      "body": "Et quis optio voluptatibus quas dolor."
   }' --project-id 520607661 --auth "Impedit quidem labore."
`, os.Args[0])
}

func projectDeleteUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project delete- update -project-id INT32 -update-id INT64 -auth STRING

DeleteUpdate implements delete update.
    -project-id INT32: 
    -update-id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` project delete- update --project-id 1303777347 --update-id 265494058706900424 --auth "Aut consequuntur et a."
`, os.Args[0])
}

func projectModifyUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project modify- update -body JSON -project-id INT32 -update-id INT64 -auth STRING

ModifyUpdate implements modify update.
    -body JSON: 
    -project-id INT32: 
    -update-id INT64: 
    -auth STRING: 

Example:
    `+os.Args[0]+` project modify- update --body '{
      "body": "Voluptatem perspiciatis libero minus."
   }' --project-id 1659700200 --update-id 5923090945204615108 --auth "Ea perspiciatis magnam fugiat repellendus doloremque aut."
`, os.Args[0])
}

func projectInvitesUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project invites -auth STRING

Invites implements invites.
    -auth STRING: 

Example:
    `+os.Args[0]+` project invites --auth "Dolorem placeat a et."
`, os.Args[0])
}

func projectLookupInviteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project lookup- invite -token STRING -auth STRING

LookupInvite implements lookup invite.
    -token STRING: 
    -auth STRING: 

Example:
    `+os.Args[0]+` project lookup- invite --token "Veniam qui." --auth "Minus totam illum veniam deserunt sint."
`, os.Args[0])
}

func projectAcceptInviteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project accept- invite -id INT64 -token STRING -auth STRING

AcceptInvite implements accept invite.
    -id INT64: 
    -token STRING: 
    -auth STRING: 

Example:
    `+os.Args[0]+` project accept- invite --id 5729175641368403504 --token "Aut voluptas." --auth "Nostrum possimus."
`, os.Args[0])
}

func projectRejectInviteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] project reject- invite -id INT64 -token STRING -auth STRING

RejectInvite implements reject invite.
    -id INT64: 
    -token STRING: 
    -auth STRING: 

Example:
    `+os.Args[0]+` project reject- invite --id 5438991930375446807 --token "Qui quo dolorem quo non earum." --auth "Neque qui."
`, os.Args[0])
}

// stationUsage displays the usage of the station command and its subcommands.
func stationUsage() {
	fmt.Fprintf(os.Stderr, `Service is the station service interface.
Usage:
    %s [globalflags] station COMMAND [flags]

COMMAND:
    add: Add implements add.
    get: Get implements get.
    update: Update implements update.
    list- mine: ListMine implements list mine.
    list- project: ListProject implements list project.
    photo: Photo implements photo.

Additional help:
    %s station COMMAND --help
`, os.Args[0], os.Args[0])
}
func stationAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] station add -body JSON -auth STRING

Add implements add.
    -body JSON: 
    -auth STRING: 

Example:
    `+os.Args[0]+` station add --body '{
      "device_id": "Occaecati sequi aut provident porro ut provident.",
      "location_name": "Voluptate deserunt beatae maxime magni qui.",
      "name": "Doloremque quasi sint.",
      "status_json": {
         "Et voluptatem laborum voluptatem.": "Distinctio reiciendis omnis quia.",
         "Laborum vero odio fugit.": "Maxime totam sapiente eligendi quo quam.",
         "Qui officiis natus est corrupti eligendi consequatur.": "Placeat provident iure non esse hic."
      },
      "status_pb": "Cumque odit omnis."
   }' --auth "Repellendus aut temporibus."
`, os.Args[0])
}

func stationGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] station get -id INT32 -auth STRING

Get implements get.
    -id INT32: 
    -auth STRING: 

Example:
    `+os.Args[0]+` station get --id 2098446527 --auth "Et omnis delectus sed."
`, os.Args[0])
}

func stationUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] station update -body JSON -id INT32 -auth STRING

Update implements update.
    -body JSON: 
    -id INT32: 
    -auth STRING: 

Example:
    `+os.Args[0]+` station update --body '{
      "location_name": "Porro quidem adipisci dolore unde maxime vero.",
      "name": "Quaerat dolorem voluptas aperiam sunt ipsam.",
      "status_json": {
         "Sequi distinctio unde.": "Repudiandae ab vel quia.",
         "Similique minus consequatur consequatur harum.": "Laborum laudantium."
      },
      "status_pb": "Fuga ut expedita eos nihil commodi expedita."
   }' --id 1585951539 --auth "Modi consectetur at voluptatem."
`, os.Args[0])
}

func stationListMineUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] station list- mine -auth STRING

ListMine implements list mine.
    -auth STRING: 

Example:
    `+os.Args[0]+` station list- mine --auth "Deserunt sequi est sunt qui occaecati."
`, os.Args[0])
}

func stationListProjectUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] station list- project -id INT32 -auth STRING

ListProject implements list project.
    -id INT32: 
    -auth STRING: 

Example:
    `+os.Args[0]+` station list- project --id 858855066 --auth "Odit minima animi est."
`, os.Args[0])
}

func stationPhotoUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] station photo -id INT32 -auth STRING

Photo implements photo.
    -id INT32: 
    -auth STRING: 

Example:
    `+os.Args[0]+` station photo --id 1174397453 --auth "Voluptatem nulla ipsa ea rerum laborum."
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
    `+os.Args[0]+` tasks refresh- device --device-id "Ex praesentium." --auth "Occaecati qui non."
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
    `+os.Args[0]+` test get --id 3456052859101749420
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
    `+os.Args[0]+` test email --address "Officiis eius quas at eligendi sed." --auth "Ab voluptatem quibusdam."
`, os.Args[0])
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `Service is the user service interface.
Usage:
    %s [globalflags] user COMMAND [flags]

COMMAND:
    roles: Roles implements roles.
    delete: Delete implements delete.

Additional help:
    %s user COMMAND --help
`, os.Args[0], os.Args[0])
}
func userRolesUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user roles -auth STRING

Roles implements roles.
    -auth STRING: 

Example:
    `+os.Args[0]+` user roles --auth "Recusandae enim numquam quae adipisci."
`, os.Args[0])
}

func userDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user delete -user-id INT32 -auth STRING

Delete implements delete.
    -user-id INT32: 
    -auth STRING: 

Example:
    `+os.Args[0]+` user delete --user-id 218502286 --auth "Aut aperiam soluta laboriosam sed qui."
`, os.Args[0])
}
