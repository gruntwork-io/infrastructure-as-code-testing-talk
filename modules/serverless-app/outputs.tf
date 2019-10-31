output "url" {
  value = aws_api_gateway_deployment.deployment.invoke_url
}

output "function_arn" {
  value = aws_lambda_function.web_app.arn
}

output "iam_role_arn" {
  value = aws_iam_role.lambda.arn
}