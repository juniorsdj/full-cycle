const { sum, subtraction, multiply } = require("./sum");

test("adds 1 + 2 to equal 3", () => {
  expect(sum(1, 2)).toBe(3);
});
test("adds 2 - 1 to equal 1", () => {
  expect(subtraction(2, 1)).toBe(1);
});
test("adds 1 * 2 to equal 2", () => {
  expect(multiply(1, 2)).toBe(2);
});
