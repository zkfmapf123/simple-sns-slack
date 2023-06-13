locals {
    config = jsondecode(file("../config.json"))

    vpc_id = local.config.vpc_id
    public_subnets = local.config.public_subnets
    lambda_arn = local.config.lambda.arn
}