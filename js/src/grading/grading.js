#!/usr/bin/env node

process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////
const constantGrades = [0, 1, 2, 5, 6, 7];

function solve(grades){
    return grades.map((grade) => {
        const onesDigit = grade % 10;
        if (grade < 38 || constantGrades.includes(onesDigit))
            return grade;
        return grade + Math.min(10 - onesDigit, Math.abs(5 - onesDigit));
    });
}

function main() {
    var n = parseInt(readLine());
    var grades = [];
    for(var grades_i = 0; grades_i < n; grades_i++){
       grades[grades_i] = parseInt(readLine());
    }
    var result = solve(grades);
    console.log(result.join("\n"));
}

module.exports.grade = solve;
