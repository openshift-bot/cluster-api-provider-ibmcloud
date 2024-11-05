# METADATA
# title: User data for EC2 instances must not contain sensitive AWS keys
# description: |
#   EC2 instance data is used to pass start up information into the EC2 instance. This userdata must not contain access key credentials. Instead use an IAM Instance Profile assigned to the instance to grant access to other AWS Services.
# scope: package
# schemas:
#   - input: schema["cloud"]
# related_resources:
#   - https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-add-user-data.html
# custom:
#   id: AVD-AWS-0129
#   avd_id: AVD-AWS-0129
#   provider: aws
#   service: ec2
#   severity: CRITICAL
#   short_code: no-secrets-in-launch-template-user-data
#   recommended_action: Remove sensitive data from the EC2 instance user-data generated by launch templates
#   input:
#     selector:
#       - type: cloud
#         subtypes:
#           - service: ec2
#             provider: aws
#   terraform:
#     links:
#       - https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#user_data
#     good_examples: checks/cloud/aws/ec2/as_no_secrets_in_user_data.tf.go
#     bad_examples: checks/cloud/aws/ec2/as_no_secrets_in_user_data.tf.go
#   cloudformation:
#     good_examples: checks/cloud/aws/ec2/as_no_secrets_in_user_data.cf.go
#     bad_examples: checks/cloud/aws/ec2/as_no_secrets_in_user_data.cf.go
package builtin.aws.ec2.aws0129

import rego.v1

deny contains res if {
	some tmpl in input.aws.ec2.launchtemplates
	scan_result := squealer.scan_string(tmpl.instance.userdata.value)
	scan_result.transgressionFound
	res := result.new(
		sprintf("Sensitive data found in user data: %s", [scan_result.description]),
		tmpl.instance.userdata,
	)
}
