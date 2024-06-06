# flowforge frontend

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

> Do note that Flowforge project uses [Husky](https://typicode.github.io/husky/) to run all these checks before git commit.

## Environment Variables

We rely on the following environment variables for our app. Please create an `.env.local` in development and store these variables:

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

Here is the list of libraries that the app depende:

- [NextJS 14](https://nextjs.org/blog/next-14)
- [Radix UI](https://www.radix-ui.com/)
- [react-jsonschema-form](https://rjsf-team.github.io/react-jsonschema-form/docs/)
- [TanStack Query](https://tanstack.com/query/latest)
- [TanStack Table](https://tanstack.com/table/latest)
- [Axios](https://axios-http.com/docs/intro)
- [react-hook-form](https://www.react-hook-form.com/)
- [Zod](https://zod.dev/)
