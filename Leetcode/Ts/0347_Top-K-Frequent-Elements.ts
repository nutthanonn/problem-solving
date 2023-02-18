function topKFrequent(nums: number[], k: number): number[] {
  const map = new Map<number, number>();
  for (const num of nums) {
    map.set(num, (map.get(num) || 0) + 1);
  }
  const arr = Array.from(map.entries());
  arr.sort((a, b) => b[1] - a[1]);
  return arr.slice(0, k).map((item) => item[0]);
}
