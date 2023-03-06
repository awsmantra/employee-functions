import { StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as cdk from 'aws-cdk-lib';
import {CompositePrincipal, Effect, PolicyStatement, Role, ServicePrincipal} from "aws-cdk-lib/aws-iam";
import * as iam from 'aws-cdk-lib/aws-iam';

export class ApiGatewayAccessRole extends cdk.NestedStack  {
    private readonly _role: Role;
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, props);


        // Add scheduler assumeRole
        this._role  = new Role(this,  "APIGatewayAccessRole", {
            assumedBy: new CompositePrincipal(
                new iam.ServicePrincipal('apigateway.amazonaws.com'),
                new iam.ServicePrincipal('lambda.amazonaws.com')
            ),
            roleName: "api-gateway-access-role"
        })

        // Add policy
        this._role.addToPolicy(  new PolicyStatement( {
            sid: 'LambdaAccess',
            effect: Effect.ALLOW,
            actions: [
                "lambda:InvokeFunction"
            ],
            resources: ["*"], //Give the least privileges
        }))
    }

    get roleArn(): string {
        return this._role.roleArn;
    }
    get role(): Role {
        return this._role
    }
}
