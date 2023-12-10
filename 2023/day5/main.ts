import fs from 'fs';

const file = fs.readFileSync('file.txt', 'utf-8');

const lines = file.split('\n');

const seeds = lines[0].split(':')[1].trim().split(' ');

const generateSourceAndDestinationMap = (
  desStart: number,
  sourceStart: number,
  range: number
): number[] => {
  return [sourceStart, sourceStart + range - 1, desStart, desStart + range - 1];
};

const getNumberInRange = (n: number, list: number[][]): number | undefined => {
  const row = list.find((row) => n >= row[0] && n <= row[1]);
  if (!row) return undefined;
  const distance = n - row[0];
  return row[2] + distance;
};

const getSourceFromDest = (n: number, list: number[][]): number | undefined => {
  const row = list.find((row) => n >= row[2] && n <= row[3]);
  if (!row) return undefined;
  const distance = n - row[2];
  return row[0] + distance;
};

const getSourceAndDestinationMap = (
  lines: string[],
  startingLine: string
): {
  list: number[][];
  get: (n: number) => number | undefined;
  getReverse: (n: number) => number | undefined;
} => {
  const l1 = lines.slice(lines.findIndex((l) => l.includes(startingLine)));

  const section = l1
    .slice(
      1,
      l1.findIndex((l) => l === '')
    )
    .reduce<number[][]>(
      (acc, line) =>
        line === ''
          ? acc
          : [
              ...acc,
              generateSourceAndDestinationMap(
                +line.split(' ')[0],
                +line.split(' ')[1],
                +line.split(' ')[2]
              ),
            ],
      []
    );

  return {
    list: section,
    get: (n: number) => getNumberInRange(n, section),
    getReverse: (n: number) => getSourceFromDest(n, section),
  };
};

const seedSoilMap = getSourceAndDestinationMap(lines, 'seed-to-soil map:');
const soilFertilizerMap = getSourceAndDestinationMap(
  lines,
  'soil-to-fertilizer map:'
);
const fertilizerWaterMap = getSourceAndDestinationMap(
  lines,
  'fertilizer-to-water map:'
);

const waterLightMap = getSourceAndDestinationMap(lines, 'water-to-light map:');

const lightTemperature = getSourceAndDestinationMap(
  lines,
  'light-to-temperature map:'
);

const temperatureHumidityMap = getSourceAndDestinationMap(
  lines,
  'temperature-to-humidity map:'
);

const humidityLocationMap = getSourceAndDestinationMap(
  lines,
  'humidity-to-location map:'
);

const mappingOperation = (seed: string): number => {
  const soil = seedSoilMap.get(+seed) ?? +seed;
  const fert = soilFertilizerMap.get(soil) ?? soil;
  const water = fertilizerWaterMap.get(fert) ?? fert;
  const light = waterLightMap.get(water) ?? water;
  const temp = lightTemperature.get(light) ?? light;
  const hum = temperatureHumidityMap.get(temp) ?? temp;
  const loc = humidityLocationMap.get(hum) ?? hum;
  return loc;
};

const mappingOperationReverse = (loc: number): number => {
  const hum = humidityLocationMap.getReverse(loc) ?? loc;
  const temp = temperatureHumidityMap.getReverse(hum) ?? hum;
  const light = lightTemperature.getReverse(temp) ?? temp;
  const water = waterLightMap.getReverse(light) ?? light;
  const fert = fertilizerWaterMap.getReverse(water) ?? water;
  const soil = soilFertilizerMap.getReverse(fert) ?? fert;
  const seed = seedSoilMap.getReverse(soil) ?? soil;

  return seed;
};

const res = seeds
  .map((seed) => mappingOperation(seed))
  .sort((a, b) => a - b)[0];

console.log('p1', res);

const findSeedInRange = (n: number) => {
  let found = false;
  for (let s = 0; s < seeds.length / 2; s++) {
    const startingRange = +seeds[s * 2];
    const endingRange = +seeds[s * 2] + +seeds[s * 2 + 1];

    if (n >= startingRange && n <= endingRange) found = true;
  }
  return found;
};
const locationSortAsc = humidityLocationMap.list.sort((a, b) => a[2] - b[2]);

let lowestLocation: number | undefined = undefined;
for (let guess = 0; true; guess++) {
  const seed = mappingOperationReverse(guess);
  if (findSeedInRange(seed)) {
    lowestLocation = guess;
    break;
  }
}

console.log('part2', lowestLocation);
