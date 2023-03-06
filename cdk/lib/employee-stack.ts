import {Stack, StackProps} from 'aws-cdk-lib';
import {Construct} from 'constructs';
import {ApiGatewayAccessRole } from "./scheduler-role";
import {EmployeeTable} from "./table-stack";
import {EmployeePostFunction} from "./employee-post-stack";
import {EmployeeApi} from "./employee-api";
import {EmployeeGetFunction} from "./employee-get-stack";
import {EmployeeDeleteFunction} from "./employee-delete-stack";
import {EmployeePutFunction} from "./employee-put-stack";
import {Options} from "../types/options";

interface EmployeeStackProps extends StackProps {
  options: Options,
}

export class EmployeeStack extends Stack {
  constructor(scope: Construct, id: string, props: EmployeeStackProps) {
    super(scope, id, props);


    // Create Role
    const apiGatewayAccessRole = new ApiGatewayAccessRole(this,"ApiGatewayAccessRoleStack")

    // Create Employee DynamoDB table
    const tableStack = new EmployeeTable(this, 'EmployeeTableStack', {})

    // Create EmployeePostFunction
    const employeePostFuncStack = new EmployeePostFunction(this, 'EmployeePostFunction', {
      table: tableStack.table,
      options: props?.options,
    })

    // Create EmployeeGetFunction
    const employeeGetFuncStack = new EmployeeGetFunction(this, 'EmployeeGetFunction', {
      table: tableStack.table,
      options: props?.options,
    })

    // Create EmployeeDeleteFunction
    const employeeDeleteFuncStack = new EmployeeDeleteFunction(this, 'EmployeeDeleteFunction', {
      table: tableStack.table,
      options: props?.options,
    })

    // Create EmployeePutFunction
    const employeePutFuncStack = new EmployeePutFunction(this, 'EmployeePutFunction', {
      table: tableStack.table,
      options: props?.options,
    })

    const employeeApi = new EmployeeApi(this, 'EmployeeApiStack', {
      employeePostFunc : employeePostFuncStack.function,
      employeeGetFunc : employeeGetFuncStack.function,
      employeeDeleteFunc: employeeDeleteFuncStack.function,
      employeePutFunc : employeePutFuncStack.function,
      apiGatewayAccessRoleArn: apiGatewayAccessRole.role,
      options: props?.options,
    })

  }
}
