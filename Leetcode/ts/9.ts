function isPalindrome(x: number): boolean {
  if (x < 0) {
    return false;
  }
  const check = parseInt(x.toString().split("").reverse().join(""));
  if (x === check) {
    return true;
  } else {
    return false;
  }
}
