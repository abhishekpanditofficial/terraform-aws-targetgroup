variable "target_group_name" {
  type = string
  description = "Name of the target group"
}
variable "target_group_port" {
  type = string
  description = "Port of the target group"
  default = 80
}
variable "target_group_protocol" {
  type = string
  description = "Protocol of the target group"
  default = "HTTP"
}
variable "vpc_id" {
  type = string
  description = "Vpc id for the target-group"
}