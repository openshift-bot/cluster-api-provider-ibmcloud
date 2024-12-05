package computing

import (
	"github.com/aquasecurity/trivy-checks/pkg/rules"
	"github.com/aquasecurity/trivy/pkg/iac/providers"
	"github.com/aquasecurity/trivy/pkg/iac/scan"
	"github.com/aquasecurity/trivy/pkg/iac/severity"
	"github.com/aquasecurity/trivy/pkg/iac/state"
)

var CheckAddSecurityGroupToInstance = rules.Register(
	scan.Rule{
		AVDID:       "AVD-NIF-0004",
		Aliases:     []string{"nifcloud-computing-add-security-group-to-instance"},
		Provider:    providers.NifcloudProvider,
		Service:     "computing",
		ShortCode:   "add-security-group-to-instance",
		Summary:     "Missing security group for instance.",
		Impact:      "A security group controls the traffic that is allowed to reach and leave the resources that it is associated with.",
		Resolution:  "Add security group for all instances",
		Explanation: "Need to add a security group to your instance.",
		Links: []string{
			"https://pfs.nifcloud.com/help/server/change_fw.htm",
		},
		Terraform: &scan.EngineMetadata{
			GoodExamples:        terraformAddSecurityGroupToInstanceGoodExamples,
			BadExamples:         terraformAddSecurityGroupToInstanceBadExamples,
			Links:               terraformAddSecurityGroupToInstanceLinks,
			RemediationMarkdown: terraformAddSecurityGroupToInstanceRemediationMarkdown,
		},
		Severity:   severity.Critical,
		Deprecated: true,
	},
	func(s *state.State) (results scan.Results) {
		for _, instance := range s.Nifcloud.Computing.Instances {
			if instance.Metadata.IsUnmanaged() {
				continue
			}
			if instance.SecurityGroup.IsEmpty() {
				results.Add(
					"Instance does not have a securiy group.",
					instance.SecurityGroup,
				)
			} else {
				results.AddPassed(&instance)
			}
		}
		return
	},
)