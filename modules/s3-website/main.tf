terraform {
  required_version = ">= 0.12"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE THE S3 BUCKET
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_s3_bucket" "example" {
  bucket = var.name
  acl    = "public-read"
  policy = data.aws_iam_policy_document.public_bucket_policy.json

  website {
    index_document = local.index_page
    error_document = local.error_page
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE AN IAM POLICY THAT ALLOWS THE BUCKET TO BE USED AS A PUBLIC WEBSITE
# ---------------------------------------------------------------------------------------------------------------------

data "aws_iam_policy_document" "public_bucket_policy" {
  statement {
    effect    = "Allow"
    actions   = ["s3:GetObject"]
    resources = ["arn:aws:s3:::${var.name}/*"]

    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# UPLOAD THE INDEX AND ERROR PAGES
# ---------------------------------------------------------------------------------------------------------------------

resource "aws_s3_bucket_object" "index_page" {
  bucket       = aws_s3_bucket.example.bucket
  key          = local.index_page
  source       = "${path.module}/example-website/${local.index_page}"
  content_type = "text/html"
}

resource "aws_s3_bucket_object" "error_page" {
  bucket       = aws_s3_bucket.example.bucket
  key          = local.error_page
  source       = "${path.module}/example-website/${local.error_page}"
  content_type = "text/html"
}

locals {
  index_page = "index.html"
  error_page = "error.html"
}