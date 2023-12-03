const fs = require("fs");
const readline = require("readline");

function convertToIntNumber(number) {
  if (number === "one") {
    return "1";
  } else if (number === "two") {
    return "2";
  } else if (number === "three") {
    return "3";
  } else if (number === "four") {
    return "4";
  } else if (number === "five") {
    return "5";
  } else if (number === "six") {
    return "6";
  } else if (number === "seven") {
    return "7";
  } else if (number === "eight") {
    return "8";
  } else if (number === "nine") {
    return "9";
  } else {
    return number;
  }
}

async function processLineByLine() {
  const filestream = fs.createReadStream("puzzle.txt");
  const rl = readline.createInterface({
    input: filestream,
    crlfDelay: Infinity,
  });

  var regex = /(?=(one|two|three|four|five|six|seven|eight|nine|\d))/g;
  var sum = 0;
  lineNumber = 0;
  for await (const line of rl) {
    var numbers = Array.from(line.matchAll(regex), (m) => m[1]);

    console.log("Line: " + lineNumber + "\n" + line + "\n" + numbers);

    for (var i = 0; i < numbers.length; i++) {
      numbers[i] = convertToIntNumber(numbers[i]);
    }

    first = numbers[0];
    last = numbers[numbers.length - 1];
    pair = first + last;
    pair = parseInt(pair);
    console.log("pair: " + pair);
    sum += pair;
    console.log("sum: " + sum);
    lineNumber++;
  }
  console.log(sum);
}

processLineByLine();
