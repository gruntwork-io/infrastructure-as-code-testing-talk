#!/usr/bin/env node
import cdk = require('@aws-cdk/core');
import { DemoCdkStack } from '../lib/cdk-demo-stack';

const app = new cdk.App();
new DemoCdkStack(app, 'DemoCdkStack');
