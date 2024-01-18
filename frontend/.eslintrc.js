module.exports = {
  root: true,
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 2020,
    tsconfigRootDir: __dirname,
    sourceType: "module",
    project: "./tsconfig.json",
  },
  ignorePatterns: ["*.js"],
  extends: [
    "airbnb-typescript",
    "airbnb/hooks", // Lints React hooks
    "plugin:@typescript-eslint/recommended",
  ],
  plugins: ["react", "import", "@typescript-eslint"],
  env: {
    browser: true,
    es6: true,
  },
  rules: {
    "import/no-extraneous-dependencies": ["error", { devDependencies: true }],
    "no-use-before-define": "off",
    "@typescript-eslint/no-use-before-define": ["error"],
    "import/extensions": "off",
    "import/no-extraneous-dependencies": "off",
    "@typescript-eslint/no-use-before-define": "off",
    "@typescript-eslint/no-shadow": "off",
    "@typescript-eslint/quotes": "off",
    "@typescript-eslint/semi": "off",
    "@typescript-eslint/indent": "off",
    "@typescript-eslint/comma-dangle": "off",
  },
}
