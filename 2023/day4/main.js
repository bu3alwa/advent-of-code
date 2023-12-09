const fs = require('fs')

const f = fs.readFileSync('./file.txt', 'utf-8')
const lines = f.split('\n')
console.log(lines)

const cardScore = (n, point) => {
  if (point === 0) return n
  if (n === 0) return cardScore(1, point - 1)
  else return cardScore(n * 2, point - 1)
}

const total = lines.reduce((acc, line) => {
  if (line === "") { return acc }
  const card = line.split(":")[1]
  console.log(line.split(':')[0], card)
  const winning = new Set(card.split("|")[0].trim().split(/\s+/))
  console.log('winning numbers', winning)
  const draw = card.split('|')[1].split(/\s+/).map((num) => winning.has(num))
  const won = draw.filter((d) => d === true).length
  const score = won === 0 ? 0 : 2 ** (won - 1)

  console.log('draw', draw)
  console.log('cur score', score)
  return acc + score
}, 0)


console.log('part1', total)

