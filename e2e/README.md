# E2E Tests
We use playwright for our E2E testing.
https://playwright.dev/docs/intro 

## Run tests
npx playwright test 
npx playwright show-report
# Creating new Tests

## Simple tests 
Use codegen to generate simple tests, such as the navigation and visibility tests.
https://playwright.dev/docs/codegen-intro
npx playwright codegen localhost:8000

## Complex tests

For tests like the login test or create account test,