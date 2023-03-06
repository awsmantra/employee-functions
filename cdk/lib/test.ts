const jsonString = '{\n' +
    '  "Service": "calendar-service",\n' +
    '  "Cluster": "ECSFGCluster",\n' +
    '  "DesiredCount": 1\n' +
    '}';


const JSobj = JSON.parse(jsonString);

console.log(JSobj);
console.log(typeof JSobj);

const JSON_string = JSON.stringify(JSobj);

console.log(JSON_string);
console.log(typeof JSON_string);
