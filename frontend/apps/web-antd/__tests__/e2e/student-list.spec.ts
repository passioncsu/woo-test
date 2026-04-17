import { expect, test } from '@playwright/test';

test.use({ storageState: 'node_modules/.e2e/auth.json' });

test.describe('学生管理', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/student/list');
    await page.waitForLoadState('networkidle');
    await page.waitForTimeout(1000);
  });

  test('列表页渲染正常', async ({ page }) => {
    // 验证页面标题
    await expect(page.getByText('新增学生')).toBeVisible();
    // 验证表格存在
    await expect(page.locator('table')).toBeVisible();
    // 验证有数据行（seed data）
    const rows = page.locator('table tbody tr');
    await expect(rows.first()).toBeVisible({ timeout: 5000 });
  });

  test('搜索功能', async ({ page }) => {
    // 获取初始行数
    const initialRows = await page.locator('table tbody tr').count();

    // 输入搜索关键词
    const searchInput = page.locator('input[placeholder="请输入姓名/学号/专业"]');
    await searchInput.fill('张三');
    await page.waitForTimeout(1500);

    // 验证搜索结果
    const rows = await page.locator('table tbody tr').count();
    expect(rows).toBeLessThanOrEqual(initialRows);
  });

  test('新增学生', async ({ page }) => {
    await page.getByRole('button', { name: '新增学生' }).click();
    await page.waitForTimeout(1000);

    // 填写表单
    await page.locator('input[placeholder="请输入学号"]').fill('2024E2E1');
    await page.locator('input[placeholder="请输入姓名"]').fill('E2E测试学生');

    // 提交
    await page.getByRole('button', { name: '提交' }).click();
    await page.waitForTimeout(2000);

    // 验证成功提示
    await expect(page.locator('.ant-message-success')).toBeVisible({ timeout: 5000 });

    // 验证列表中出现新学生
    await expect(page.locator('table')).toContainText('E2E测试学生', { timeout: 5000 });
  });

  test('编辑学生', async ({ page }) => {
    // 点击第一个编辑按钮
    await page.getByRole('button', { name: '编辑' }).first().click();
    await page.waitForTimeout(2000);

    // 验证弹窗和表单
    await expect(page.locator('input[placeholder="请输入姓名"]')).toBeVisible({ timeout: 5000 });

    // 修改姓名
    const nameInput = page.locator('input[placeholder="请输入姓名"]');
    await nameInput.clear();
    await nameInput.fill('编辑后名称');

    // 提交
    await page.getByRole('button', { name: '提交' }).click();
    await page.waitForTimeout(2000);

    await expect(page.locator('.ant-message-success')).toBeVisible({ timeout: 5000 });
  });

  test('删除学生', async ({ page }) => {
    // 先新增一个用于删除
    await page.getByRole('button', { name: '新增学生' }).click();
    await page.waitForTimeout(1000);
    await page.locator('input[placeholder="请输入学号"]').fill('2024DEL1');
    await page.locator('input[placeholder="请输入姓名"]').fill('待删除学生');
    await page.getByRole('button', { name: '提交' }).click();
    await page.waitForTimeout(2000);

    // 找到待删除学生的行，点删除
    const deleteRow = page.locator('table tbody tr').filter({ hasText: '待删除学生' });
    await deleteRow.getByRole('button', { name: '删除' }).click();

    // 确认弹窗
    await expect(page.locator('.ant-modal-confirm')).toBeVisible();
    await page.locator('.ant-modal-confirm').getByRole('button', { name: '确 定' }).click();
    await page.waitForTimeout(2000);

    await expect(page.locator('.ant-message-success')).toBeVisible({ timeout: 5000 });
  });
});
