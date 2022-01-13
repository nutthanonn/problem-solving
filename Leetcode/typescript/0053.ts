function maxSubArray(nums: number[]): number {
  let sum = 0;
  let maxNumber = nums[0];
  nums.forEach((item) => {
    sum += item;
    maxNumber = Math.max(maxNumber, sum);
    if (sum < 0) {
      sum = 0;
    }
  });
  return maxNumber;
}
