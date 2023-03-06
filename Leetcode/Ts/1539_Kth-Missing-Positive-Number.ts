function findKthPositive(arr: number[], k: number): number {
  const temp: number[] = [];
  const map = {};
  arr.forEach((item) => (map[item] = 1));

  let i = 1;
  while (!temp[k]) {
    if (!map[i]) {
      temp.push(i);
    }
    i++;
  }

  return temp[k - 1];
}

console.log(findKthPositive([2, 3, 4, 7, 11], 5));
console.log(findKthPositive([1, 2, 3, 4], 2));
console.log(findKthPositive([1, 2, 3, 4], 4));
console.log(findKthPositive([2], 1));
