output "vpc_id" {
    value= module.terraform_aws_vpc.vpc_id
}
output "vpc_cidr_block" {
    value= module.terraform_aws_vpc.vpc_cidr_block
}
output "target_group_arn" {
    value= module.terraform_aws_targetgroup.target_group_arn
}
output "target_group_name" {
    value= module.terraform_aws_targetgroup.target_group_name
}
output "target_group_vpc" {
    value= module.terraform_aws_targetgroup.target_group_vpc
}