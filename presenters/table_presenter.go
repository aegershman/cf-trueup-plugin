package presenters

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (p *Presenter) asTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{
		"Org",
		"Space",
		"Billable AIs",
		"AIs",
		"Stopped AIs",
		"SCS",
	})

	for _, orgDecorator := range p.Report.OrgDecorators {
		for _, SpaceReport := range orgDecorator.SpaceReport {
			table.Append([]string{
				orgDecorator.Name,
				SpaceReport.Name,
				strconv.Itoa(SpaceReport.BillableAppInstancesCount()),
				strconv.Itoa(SpaceReport.AppInstancesCount),
				strconv.Itoa(SpaceReport.StoppedAppInstancesCount),
				strconv.Itoa(SpaceReport.SpringCloudServicesCount()),
			})
		}
	}

	table.SetFooter([]string{
		"-",
		"Total",
		strconv.Itoa(p.Report.AggregateOrgDecorator.BillableAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorator.AppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorator.StoppedAppInstancesCount),
		strconv.Itoa(p.Report.AggregateOrgDecorator.SpringCloudServicesCount),
	})

	table.Render()

}
