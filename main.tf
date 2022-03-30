module "terraform_aws_vpc" {
  source= "./modules/vpc"
}
module "terraform_aws_targetgroup" {
  source= "./modules/target-group"
  vpc_id= module.terraform_aws_vpc.vpc_id
  target_group_name = var.target_group_name
}