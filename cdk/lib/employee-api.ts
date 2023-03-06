import {CfnOutput, StackProps} from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as cdk from 'aws-cdk-lib';
import * as apigateway from 'aws-cdk-lib/aws-apigateway';
import {EmployeePostFunction} from "./employee-post-stack";
import {GoFunction} from "@aws-cdk/aws-lambda-go-alpha";
import {Role} from "aws-cdk-lib/aws-iam";
import {Options} from "../types/options";

interface ApiProps extends cdk.NestedStackProps {
    employeePostFunc: GoFunction,
    employeeGetFunc: GoFunction,
    employeeDeleteFunc: GoFunction,
    employeePutFunc: GoFunction,
    apiGatewayAccessRoleArn: Role,
    options: Options,
}


export class EmployeeApi extends cdk.NestedStack  {
    constructor(scope: Construct, id: string, props: ApiProps) {
        super(scope, id, props);


        const api = new apigateway.RestApi(this, 'EmployeeApi', {
            description: 'example api gateway',
            deployOptions: {
                stageName: props.options.apiStageName,
            },
        });

        //Add Root Resource
        const employee = api.root.addResource('employee');

        //Add Post Method /employee
        employee.addMethod(
            'POST',
            new apigateway.LambdaIntegration(props?.employeePostFunc, {
                proxy: true,
                credentialsRole:props.apiGatewayAccessRoleArn
            }),
        );


        //Add Get Method. /employee/{id}
        const employeeId = employee.addResource('{id}');
        employeeId.addMethod(
            'GET',
            new apigateway.LambdaIntegration(props?.employeeGetFunc, {
                proxy: true,
                credentialsRole:props.apiGatewayAccessRoleArn
            }),
        );

        //Add Delete Method. /employee/{id}
        employeeId.addMethod(
            'DELETE',
            new apigateway.LambdaIntegration(props?.employeeDeleteFunc, {
                proxy: true,
                credentialsRole:props.apiGatewayAccessRoleArn
            }),
        );

        //Add Put Method. /employee/{id}
        employeeId.addMethod(
            'PUT',
            new apigateway.LambdaIntegration(props?.employeePutFunc, {
                proxy: true,
                credentialsRole:props.apiGatewayAccessRoleArn
            }),
        );

        new CfnOutput(this, 'EmployeeApiURL', {
            exportName: `employee-api-url`,
            value: api.url
        });
    }
}
