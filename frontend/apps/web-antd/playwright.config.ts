import type { PlaywrightTestConfig } from '@playwright/test';

const config: PlaywrightTestConfig = {
  expect: { timeout: 5000 },
  forbidOnly: !!process.env.CI,
  outputDir: 'node_modules/.e2e/test-results/',
  projects: [
    {
      name: 'chromium',
      use: { viewport: { width: 1440, height: 900 } },
    },
  ],
  reporter: [['list']],
  retries: process.env.CI ? 2 : 0,
  testDir: './__tests__/e2e',
  timeout: 30_000,
  use: {
    actionTimeout: 0,
    baseURL: 'http://localhost:5666',
    headless: true,
    trace: 'retain-on-failure',
  },
  workers: process.env.CI ? 1 : undefined,
};

export default config;
