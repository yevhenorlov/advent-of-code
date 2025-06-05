/*
--- Part Two ---
During the second Go / No Go poll, the Elf in charge of the Rocket Equation Double-Checker stops the launch sequence. Apparently, you forgot to include additional fuel for the fuel you just added.

Fuel itself requires fuel just like a module - take its mass, divide by three, round down, and subtract 2. However, that fuel also requires fuel, and that fuel requires fuel, and so on. Any mass that would require negative fuel should instead be treated as if it requires zero fuel; the remaining mass, if any, is instead handled by wishing really hard, which has no mass and is outside the scope of this calculation.

So, for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you just calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative. For example:

A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2 divided by 3 and rounded down is 0, which would call for a negative fuel), so the total fuel required is still just 2.
At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.
The fuel required by a module of mass 100756 and its fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.
What is the sum of the fuel requirements for all of the modules on your spacecraft when also taking into account the mass of the added fuel? (Calculate the fuel requirements for each module separately, then add them all up at the end.)
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

// divide by 3, round down, subtract 2, repeat if the result is > 0
const getFuel = moduleMass => {
  let fuelMass = Math.floor(moduleMass / 3) - 2;
  if (fuelMass <= 0) {
    return 0;
  }
  if (fuelMass > 0) {
    fuelMass += getFuel(fuelMass);
  }
  return fuelMass;
};

const sum = input.map(getFuel).reduce((acc, el) => {
  acc += el;
  return acc;
}, 0);
