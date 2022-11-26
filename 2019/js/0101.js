/*
--- Day 1: The Tyranny of the Rocket Equation ---
Santa has become stranded at the edge of the Solar System while delivering presents to other planets! To accurately calculate his position in space, safely align his warp drive, and return to Earth in time to save Christmas, he needs you to bring him measurements from fifty stars.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The Elves quickly load you into a spacecraft and prepare to launch.

At the first Go / No Go poll, every Elf is Go until the Fuel Counter-Upper. They haven't determined the amount of fuel required yet.

Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

For example:

For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.
The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.

What is the sum of the fuel requirements for all of the modules on your spacecraft?
*/
const input = [
  78207,
  89869,
  145449,
  73634,
  78681,
  81375,
  131482,
  126998,
  50801,
  115839,
  77949,
  53203,
  146099,
  56912,
  59925,
  132631,
  115087,
  89543,
  123234,
  108110,
  109873,
  81923,
  124264,
  87981,
  106554,
  147239,
  73615,
  72609,
  129684,
  84175,
  64915,
  98124,
  74391,
  55211,
  120961,
  119116,
  148275,
  89605,
  115986,
  120547,
  50299,
  137922,
  78906,
  145216,
  80424,
  122610,
  61408,
  97573,
  127533,
  116820,
  76068,
  77400,
  117943,
  85231,
  102442,
  62002,
  58761,
  56479,
  98200,
  85971,
  73985,
  88908,
  82719,
  120604,
  83378,
  88241,
  122574,
  76731,
  99810,
  137548,
  102617,
  105352,
  137585,
  83238,
  118817,
  149419,
  107629,
  63893,
  56049,
  70693,
  83844,
  76413,
  87021,
  90259,
  124289,
  102527,
  139625,
  106607,
  120241,
  101098,
  66142,
  96591,
  82277,
  142297,
  116671,
  131881,
  94861,
  79741,
  73561,
  115214
];

// divide by 3, round down, subtract 2
const getFuel = moduleMass => {
  return Math.floor(moduleMass / 3) - 2;
};

const sum = input.map(getFuel).reduce((acc, el) => {
  acc += el;
  return acc;
}, 0);
