function singleNumber(nums: number[]): number {
  var single: number = 0;
  nums.forEach((number) => {
    const back = nums.lastIndexOf(number);
    const front = nums.indexOf(number);
    if (back === front) {
      single = nums[back];
    }
  });

  return single;
}
console.log(singleNumber([-1]));
