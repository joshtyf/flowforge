# Frontend Dev Guide

The app is built with [NextJS 14](https://nextjs.org/blog/next-14), Typescript and TailwindCSS as our web app framework.

We also use [shadcn/ui](https://ui.shadcn.com/) as our underlying UI library.

## Commands

Run the following to install all the dependencies:

```bash
npm ci
```

### NextJS Commands

To run development server:

```bash
npm run dev
```

To create the production app:

```bash
npm run build
```

To start the standalone production app after building:

```bash
npm run start
```

> Note that the `public` or `.next/static` folders are not copied to the standalone build folder during building based on [NextJS's default output after build](https://nextjs.org/docs/pages/api-reference/next-config-js/output#automatically-copying-traced-files). Do copy them manually into `standalone/public` and `standalone/.next/static` if necessary.

### Formatting/Checking Commands

To run Type check:

```bash
npm run check-types
```

To run Prettier format check:

```bash
npm run check-format
```

To run Prettier format:

```bash
npm run format
```

To run ESLint:

```
npm run check-lint
```

To run all checks:

```
npm run test-all
```

> Do note that Flowforge repository implemented pre-commit hooks to run all these checks with [Husky](https://typicode.github.io/husky/).

## Environment Variables

We utilize the following environment variables for the Flowforge app. Please create an `.env.local` in development to store these variables:

```bash
NEXT_PUBLIC_APP_ENV=dev
# Add your own Auth0 Domain
NEXT_PUBLIC_AUTH0_DOMAIN=
# Add your own Auth0 Client ID
NEXT_PUBLIC_AUTH0_CLIENT_ID=
# Add your own Auth0 Client Secret
NEXT_PUBLIC_AUTH0_CLIENT_SECRET=
NEXT_PUBLIC_APP_BASE_URL=http://localhost:3000
NEXT_PUBLIC_AUTH0_AUDIENCE=https://flowforge.com
```

## Credits

Here are the libraries that the frontend uses:

- [NextJS 14](https://nextjs.org/blog/next-14)
- [Radix UI](https://www.radix-ui.com/)
- [react-jsonschema-form](https://rjsf-team.github.io/react-jsonschema-form/docs/)
- [TanStack Query](https://tanstack.com/query/latest)
- [TanStack Table](https://tanstack.com/table/latest)
- [Axios](https://axios-http.com/docs/intro)
- [react-hook-form](https://www.react-hook-form.com/)
- [Zod](https://zod.dev/)
