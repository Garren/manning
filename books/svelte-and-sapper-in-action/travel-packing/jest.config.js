module.exports = {
  bail: false,
  moduleFileExtensions: ['js', 'svelte'],
  testEnvironment: 'jsdom',
  transform: {
    '^.+\\.js$': 'babel-jest',
    '^.+\\.svelte$': 'svelte-jester'
  },
  verbose: true
};
