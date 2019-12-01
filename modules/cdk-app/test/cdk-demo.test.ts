import { expect as expectCDK, haveResource } from '@aws-cdk/assert';
import cdk = require('@aws-cdk/core');
import CdkDemo = require('../lib/cdk-demo-stack');

test('Bucket Stack', () => {
    const app = new cdk.App();
    const stack = new CdkDemo.DemoCdkStack(app, 'DemoTestStack');
    expectCDK(stack).to(haveResource('AWS::S3::Bucket', {
      VersioningConfiguration: {
        Status: "Enabled"
      },
    }));
});
