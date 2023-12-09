import fs from 'fs';

const file = fs.readFileSync('file.txt', 'utf-8');
const lines = file.split('\n');

const instructions = lines[0];

const map = new Map<string, string[]>();

lines.slice(2).forEach((line) => {
  const node = line.split('=')[0].trim();
  const tuple = line
    .split('=')[1]
    .trim()
    .replace('(', '')
    .replace(')', '')
    .split(',');
  const left = tuple[0].trim();
  const right = tuple[1].trim();
  map.set(node, [left, right]);
});

const getStepsEndWithZZZ = (node: string) => {
  let current = node;
  let counter = 0;

  while (!current.startsWith('ZZZ')) {
    for (const instruction of instructions) {
      const elem = map.get(current);
      if (elem === undefined) {
        console.log(current);
        throw 'Error';
      }

      current = elem[instruction === 'L' ? 0 : 1].slice();
      counter += 1;

      if (current.endsWith('ZZZ')) break;
    }
  }
  return counter;
};

const p1Res = getStepsEndWithZZZ('AAA');
console.log('part1', p1Res);

const getStepsEndWithZ = (node: string): number => {
  let current = node;
  let counter = 0;

  while (!current.endsWith('Z')) {
    for (const instruction of instructions) {
      const elem = map.get(current);
      if (elem === undefined) {
        console.log(current);
        throw 'Error';
      }

      current = elem[instruction === 'L' ? 0 : 1].slice();
      counter += 1;

      if (current.endsWith('Z')) break;
    }
  }
  return counter;
};

const gcdOp = (a: number, b: number): number => (b === 0 ? a : gcdOp(b, a % b));
const lcmOp = (a: number, b: number) => (a * b) / gcdOp(a, b);

const lcm = (n: number[]) => n.reduce((acc, n) => lcmOp(acc, n));

const startingNodes = [...map.keys()].filter((k) => k.endsWith('A'));
const stepsToLowestZ = startingNodes.map((n) => getStepsEndWithZ(n));

console.log('part2', lcm(stepsToLowestZ));
