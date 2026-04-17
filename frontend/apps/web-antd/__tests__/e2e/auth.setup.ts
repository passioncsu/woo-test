import { test as setup } from '@playwright/test';

const AUTH_FILE = 'node_modules/.e2e/auth.json';

setup('authenticate', async ({ page }) => {
  // 先通过 API 登录获取 token
  const res = await fetch('http://localhost:8888/api/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username: 'admin', password: 'admin123' }),
  });
  const data = await res.json();
  if (data.code !== 0) {
    throw new Error(`Login failed: ${data.message}`);
  }

  const token = data.data.accessToken;

  // 访问前端，注入 token 到 localStorage
  await page.goto('/auth/login');
  await page.waitForLoadState('networkidle');

  // 找到 Pinia persisted state 的 key 并注入 token
  await page.evaluate((t) => {
    const keys = Object.keys(localStorage);
    const accessKey = keys.find((k) => k.includes('access'));
    if (accessKey) {
      const stored = JSON.parse(localStorage.getItem(accessKey) || '{}');
      stored.accessToken = t;
      localStorage.setItem(accessKey, JSON.stringify(stored));
    }
  }, token);

  // 保存认证状态
  await page.context().storageState({ path: AUTH_FILE });
});
