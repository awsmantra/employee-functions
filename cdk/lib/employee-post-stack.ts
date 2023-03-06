import {Construct} from "constructs";
import {GoFunction} from "@aws-cdk/aws-lambda-go-alpha";

import * as path from "path";
import {Duration, Tags} from "aws-cdk-lib";
import {Table} from "aws-cdk-lib/aws-dynamodb";
import {Options} from "../types/options";


interface EmployeePostProps {
    table: Table,
    options: Options,
}

export class EmployeePostFunction extends Construct {
    private readonly _func: GoFunction;

    constructor(scope: Construct, id: string, props: EmployeePostProps) {
        super(scope, id);
        this._func = new GoFunction(this, `EmployeePostFunction`, {
            entry: path.join(__dirname, `../src/employee-post`),
            functionName: `employee-post`,
            timeout: Duration.seconds(10),
            environment: {
                "LOG_LEVEL": "debug",
                "TABLE_NAME": props.table.tableName,
                "REGION" : props.options.defaultRegion
            },
        });

        props.table.grantReadWriteData(this._func);
    }

    get function(): GoFunction {
        return this._func
    }
}
