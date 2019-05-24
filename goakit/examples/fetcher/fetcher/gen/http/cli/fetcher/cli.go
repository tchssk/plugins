// Code generated by goa v3.0.0, DO NOT EDIT.
//
// fetcher HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/fetcher/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/fetcher

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
	fetcherc "goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/http/fetcher/client"
	healthc "goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/http/health/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `health show
fetcher fetch
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` health show` + "\n" +
		os.Args[0] + ` fetcher fetch --url "http://keeblerherman.name/jazlyn.ruecker"` + "\n" +
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
) (endpoint.Endpoint, interface{}, error) {
	var (
		healthFlags = flag.NewFlagSet("health", flag.ContinueOnError)

		healthShowFlags = flag.NewFlagSet("show", flag.ExitOnError)

		fetcherFlags = flag.NewFlagSet("fetcher", flag.ContinueOnError)

		fetcherFetchFlags   = flag.NewFlagSet("fetch", flag.ExitOnError)
		fetcherFetchURLFlag = fetcherFetchFlags.String("url", "REQUIRED", "URL to be fetched")
	)
	healthFlags.Usage = healthUsage
	healthShowFlags.Usage = healthShowUsage

	fetcherFlags.Usage = fetcherUsage
	fetcherFetchFlags.Usage = fetcherFetchUsage

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
		case "health":
			svcf = healthFlags
		case "fetcher":
			svcf = fetcherFlags
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
		case "health":
			switch epn {
			case "show":
				epf = healthShowFlags

			}

		case "fetcher":
			switch epn {
			case "fetch":
				epf = fetcherFetchFlags

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
		endpoint endpoint.Endpoint
		err      error
	)
	{
		switch svcn {
		case "health":
			c := healthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "show":
				endpoint = c.Show()
				data = nil
			}
		case "fetcher":
			c := fetcherc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "fetch":
				endpoint = c.Fetch()
				data, err = fetcherc.BuildFetchPayload(*fetcherFetchURLFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// healthUsage displays the usage of the health command and its subcommands.
func healthUsage() {
	fmt.Fprintf(os.Stderr, `Service is the health service interface.
Usage:
    %s [globalflags] health COMMAND [flags]

COMMAND:
    show: Health check endpoint

Additional help:
    %s health COMMAND --help
`, os.Args[0], os.Args[0])
}
func healthShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] health show

Health check endpoint

Example:
    `+os.Args[0]+` health show
`, os.Args[0])
}

// fetcherUsage displays the usage of the fetcher command and its subcommands.
func fetcherUsage() {
	fmt.Fprintf(os.Stderr, `Service is the fetcher service interface.
Usage:
    %s [globalflags] fetcher COMMAND [flags]

COMMAND:
    fetch: Fetch makes a GET request to the given URL and stores the results in the archiver service which must be running or the request fails

Additional help:
    %s fetcher COMMAND --help
`, os.Args[0], os.Args[0])
}
func fetcherFetchUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fetcher fetch -url STRING

Fetch makes a GET request to the given URL and stores the results in the archiver service which must be running or the request fails
    -url STRING: URL to be fetched

Example:
    `+os.Args[0]+` fetcher fetch --url "http://keeblerherman.name/jazlyn.ruecker"
`, os.Args[0])
}
