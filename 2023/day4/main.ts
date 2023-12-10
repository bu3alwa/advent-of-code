import fs from 'fs';

const f = fs.readFileSync('./file.txt', 'utf-8');
const lines = f.split('\n');

const getWinningNumbers = (card: string) => {
  const winning = new Set(card.split('|')[0].trim().split(/\s+/));
  const draw = card
    .split('|')[1]
    .split(/\s+/)
    .map((num) => winning.has(num));
  const won = draw.filter((d) => d === true).length;
  return won;
};

const total = lines.reduce((acc, line) => {
  if (line === '') {
    return acc;
  }
  const card = line.split(':')[1];
  const won = getWinningNumbers(card);
  const score = won === 0 ? 0 : 2 ** (won - 1);

  return acc + score;
}, 0);

console.log('part1', total);

const cardPoints = lines.map((line) => {
  const card = line.split(':')[1];
  return getWinningNumbers(card);
});

let cardsNumber = 0 + lines.length;
lines.forEach((_, i) => {
  const stack: number[] = [];
  stack.push(i);
  while (stack.length !== 0) {
    stack.sort((a, b) => a - b);
    const current = stack.pop();

    if (current === undefined) {
      throw 'Error stack empty';
    }

    // return if bigger than array
    if (current > lines.length) break;

    const points = cardPoints[current];
    cardsNumber += points;

    for (let n = 1; n <= points; n++) {
      stack.push(current + n);
    }
  }
});

console.log('part2', cardsNumber);
