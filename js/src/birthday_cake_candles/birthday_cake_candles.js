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

function birthdayCakeCandles(n, ar) {
    let max = ar[0], candlesToBlow = 0;
    for (let i = 0; i < n; i++) {
        if(ar[i] > max) {
            max = ar[i];
            candlesToBlow = 1;
        } else if(ar[i] == max) {
            candlesToBlow++;
        }
    }
    return candlesToBlow;
}

function main() {
    var n = parseInt(readLine());
    ar = readLine().split(' ');
    ar = ar.map(Number);
    var result = birthdayCakeCandles(n, ar);
    process.stdout.write("" + result + "\n");

}

module.exports.birthdayCakeCandles = birthdayCakeCandles;
