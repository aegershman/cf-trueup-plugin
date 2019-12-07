package main

import (
	"flag"
	"fmt"

	"github.com/aegershman/cf-report-usage-plugin/presenters"
	"github.com/aegershman/cf-report-usage-plugin/report"
	"github.com/cloudfoundry/cli/plugin"
	log "github.com/sirupsen/logrus"
)

// ReportUsageCmd -
type ReportUsageCmd struct {
	cli plugin.CliConnection
}

type flags struct {
	OrgNames []string
}

func (f *flags) String() string {
	return fmt.Sprint(f.OrgNames)
}

func (f *flags) Set(value string) error {
	f.OrgNames = append(f.OrgNames, value)
	return nil
}

// UsageReportCommand -
func (cmd *ReportUsageCmd) UsageReportCommand(args []string) {
	var (
		userFlags    flags
		formatFlag   string
		logLevelFlag string
	)

	flagss := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagss.Var(&userFlags, "o", "")
	flagss.StringVar(&formatFlag, "format", "table", "")
	flagss.StringVar(&logLevelFlag, "log-level", "info", "")

	err := flagss.Parse(args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	logLevel, err := log.ParseLevel(logLevelFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(logLevel)

	reporter := report.NewReporter(cmd.cli, userFlags.OrgNames)
	summaryReport, err := reporter.GetSummaryReport()
	if err != nil {
		log.Fatalln(err)
	}

	presenter := presenters.NewPresenter(*summaryReport, formatFlag) // todo hacky pointer
	presenter.Render()
}

// Run -
func (cmd *ReportUsageCmd) Run(cli plugin.CliConnection, args []string) {
	cmd.cli = cli
	switch args[0] {
	case "report-usage":
		cmd.UsageReportCommand(args)
	}
}

// GetMetadata -
func (cmd *ReportUsageCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-report-usage-plugin",
		Version: plugin.VersionType{
			Major: 2,
			Minor: 12,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "report-usage",
				HelpText: "View AIs, SIs and memory usage for orgs and spaces",
				UsageDetails: plugin.Usage{
					Usage: "cf report-usage [-o orgName...] --format formatChoice",
					Options: map[string]string{
						"o":         "organization(s) included in report. Flag can be provided multiple times.",
						"format":    "format to print as (options: string,table,json) (default: table)",
						"log-level": "(options: info,debug,trace) (default: info)",
					},
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(ReportUsageCmd))
}
