import {Construct} from "constructs";
import {GoFunction} from "@aws-cdk/aws-lambda-go-alpha";

import * as path from "path";
import {Duration, Tags} from "aws-cdk-lib";
import {Table} from "aws-cdk-lib/aws-dynamodb";
import {Options} from "../types/options";


interface EmployeeDeleteProps {
    table: Table,
    options: Options,
}

export class EmployeeDeleteFunction extends Construct {
    private readonly _func: GoFunction;

    constructor(scope: Construct, id: string, props: EmployeeDeleteProps) {
        super(scope, id);
        this._func = new GoFunction(this, `EmployeeDeleteFunction`, {
            entry: path.join(__dirname, `../src/employee-delete-by-id`),
            functionName: `employee-delete-by-id`,
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
