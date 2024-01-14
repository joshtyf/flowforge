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
    "plugin:prettier/recommended",
    "prettier",
  ],
  plugins: ["react", "import", "prettier", "@typescript-eslint"],
  env: {
    browser: true,
    es6: true,
  },
  rules: {
    // Throws error when commiting with prettier issues
    "prettier/prettier": "error",
    "import/no-extraneous-dependencies": ["error", { devDependencies: true }],
    "no-use-before-define": "off",
    "@typescript-eslint/no-use-before-define": ["error"],
  },
};
