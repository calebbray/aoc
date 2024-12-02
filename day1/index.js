const { test } = require("./test");
const { input } = require("./input");

function main() {
  const [left, right] = input.split("\n").reduce(
    (acc, line) => {
      const [l, r] = line.split("   ");
      acc[0].push(l);
      acc[1].push(r);
      return acc;
    },
    [[], []],
  );
  left.sort();
  right.sort();

  const frequencies = left.reduce((acc, num) => {
    if (acc[num]) {
      acc[num].times++;
      return acc;
    }
    acc[num] = { count: countNumberOccurrences(right, Number(num)), times: 1 };
    return acc;
  }, {});

  console.log("frequencies", frequencies);

  const total = Object.entries(frequencies).reduce((acc, [num, entry]) => {
    acc += Number(num) * entry.count * entry.times;
    return acc;
  }, 0);

  console.log(total);
}

function one() {
  const [left, right] = input.split("\n").reduce(
    (acc, line) => {
      const [l, r] = line.split("   ");
      acc[0].push(Number(l));
      acc[1].push(Number(r));
      return acc;
    },
    [[], []],
  );
  left.sort((a, b) => b - a);
  right.sort((a, b) => b - a);

  sum = 0;
  while (left.length > 0) {
    sum += Math.abs(left.pop() - right.pop());
  }

  console.log(sum);
}

function countNumberOccurrences(haystack, needle) {
  let count = 0;
  haystack.forEach((n) => needle == n && count++);
  console.log("count", count);
  return count;
}

main();
