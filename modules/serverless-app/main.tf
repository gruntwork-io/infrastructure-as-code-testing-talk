terraform {
  # This module is now only being tested with Terraform 0.15.x. However, to make upgrading easier, we are setting
  # 0.12.26 as the minimum version, as that version added support for required_providers with source URLs, making it
  # forwards compatible with 0.15.x code.
  required_version = ">= 0.12.26"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE THE LAMBDA FUNCTION
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_lambda_function" "web_app" {
  function_name = var.name
  role          = aws_iam_role.lambda.arn

  filename         = data.archive_file.source_code.output_path
  source_code_hash = data.archive_file.source_code.output_base64sha256

  runtime = var.runtime
  handler = var.handler

  memory_size = var.memory_size
  timeout     = var.timeout

  environment {
    variables = var.environment_variables
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE A DEPLOYMENT PACKAGE FOR THE LAMBDA FUNCTION BY ZIPPING UP THE SOURCE CODE
# ---------------------------------------------------------------------------------------------------------------------

data "archive_file" "source_code" {
  type        = "zip"
  source_dir  = var.source_dir
  output_path = "${path.module}/deployment-package.zip"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE AN IAM ROLE FOR THE LAMBDA FUNCTION
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role" "lambda" {
  name               = var.name
  assume_role_policy = data.aws_iam_policy_document.lambda_role.json
}

data "aws_iam_policy_document" "lambda_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE AN API GATEWAY TO PROXY REQUESTS TO THE LAMBDA FUNCTION
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_api_gateway_rest_api" "proxy" {
  name = var.name
}

# ---------------------------------------------------------------------------------------------------------------------
# PROXY THE ROOT URL (/)
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_api_gateway_method" "root" {
  rest_api_id   = aws_api_gateway_rest_api.proxy.id
  resource_id   = aws_api_gateway_rest_api.proxy.root_resource_id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "root" {
  rest_api_id = aws_api_gateway_rest_api.proxy.id
  resource_id = aws_api_gateway_rest_api.proxy.root_resource_id
  http_method = aws_api_gateway_method.root.http_method
  type        = "AWS_PROXY"

  # You can only invoke Lambdas with a POST
  integration_http_method = "POST"
  uri                     = aws_lambda_function.web_app.invoke_arn
}

# ---------------------------------------------------------------------------------------------------------------------
# PROXY ALL OTHER URLS (/foo/bar)
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.proxy.id
  parent_id   = aws_api_gateway_rest_api.proxy.root_resource_id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = aws_api_gateway_rest_api.proxy.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.proxy.id
  resource_id = aws_api_gateway_resource.proxy.id
  http_method = aws_api_gateway_method.proxy.http_method
  type        = "AWS_PROXY"

  # You can only invoke Lambdas with a POST
  integration_http_method = "POST"
  uri                     = aws_lambda_function.web_app.invoke_arn
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY THE API GATEWAY EACH TIME WE RUN APPLY
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_api_gateway_deployment" "deployment" {
  rest_api_id = aws_api_gateway_rest_api.proxy.id
  stage_name  = "live"

  depends_on = [
    aws_api_gateway_integration.root,
    aws_api_gateway_integration.proxy
  ]

  # forces to 'create' a new deployment each run. Otherwise the deployment doesn't get created after the initial run
  stage_description = timestamp()
}

# ---------------------------------------------------------------------------------------------------------------------
# GIVE THE API GATEWAY PERMISSIONS TO INVOKE THE LAMBDA FUNCTION
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_lambda_permission" "allow_api_gateway" {
  function_name = aws_lambda_function.web_app.arn
  statement_id  = "AllowExecutionFromApiGateway"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.proxy.execution_arn}/*/*/*"

  depends_on = [
    aws_api_gateway_rest_api.proxy,
    aws_api_gateway_resource.proxy
  ]
}
