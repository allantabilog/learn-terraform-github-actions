const inc = require('./index.js');
test('increment 1', () => {
  expect(inc(0)).toBe(1);
});