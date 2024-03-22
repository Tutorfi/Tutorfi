import { test, expect } from '@playwright/test';

test('homepage-visiblity', async ({ page }) => {
  await page.goto('http://localhost:8000/');
  await expect(page.getByRole('img', { name: 'Logo' })).toBeVisible();
  await expect(page.getByRole('link', { name: 'About' })).toBeVisible();
  await expect(page.getByRole('link', { name: 'Log in' })).toBeVisible();
  await page.getByText('Tutor ManagementMade').click();
  await expect(page.locator('body')).toContainText('Creating easy tutor management for organizations while also reducing the complexity of getting a tutor for students');
  await expect(page.locator('body')).toContainText('Made Easy');
  await expect(page.locator('body')).toContainText('Tutor Management');
});

test('login-visiblity', async({page}) =>{
    await page.goto('http://localhost:8000/');
    await page.getByRole('link', { name: 'Log in' }).click();
    await expect(page.getByRole('link', { name: 'Tutorfi' })).toBeVisible();
    await expect(page.getByRole('heading', { name: 'Sign in' })).toBeVisible();
    await expect(page.getByPlaceholder('Email')).toBeVisible();
    await expect(page.getByPlaceholder('Password')).toBeVisible();
    await expect(page.getByRole('button', { name: 'Submit' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'Create an account' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'Forgot password?' })).toBeVisible();
});

test('create-visiblity', async({page}) =>{
    await page.goto('http://localhost:8000/');
    await page.getByRole('link', { name: 'Log in' }).click();
    await page.getByRole('link', { name: 'Create an account' }).click();
    await expect(page.getByRole('link', { name: 'Tutorfi' })).toBeVisible();
    await expect(page.getByRole('heading', { name: 'Create Account' })).toBeVisible();
    await expect(page.getByPlaceholder('First Name')).toBeVisible();
    await expect(page.getByPlaceholder('Last Name')).toBeVisible();
    await expect(page.getByPlaceholder('Email')).toBeVisible();
    await expect(page.getByPlaceholder('password')).toBeVisible();
    await expect(page.getByRole('link', { name: 'Sign in' })).toBeVisible();
    await expect(page.getByRole('button', { name: 'Submit' })).toBeVisible();
});