#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { EmployeeStack} from '../lib/employee-stack';
import {getConfig} from "./config";

const app = new cdk.App();
const options = getConfig();

new EmployeeStack(app, 'EmployeeStack', {
    options: options,
});
