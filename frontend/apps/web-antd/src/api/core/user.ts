import type { UserInfo } from '@vben/types';

import { requestClient } from '#/api/request';

/**
 * 获取用户信息
 * 对接后端 /api/profile
 */
export async function getUserInfoApi(): Promise<UserInfo> {
  const data = await requestClient.get('/profile');
  return {
    userId: data.user_id,
    username: data.username,
    realName: data.username,
    roles: ['admin'],
    avatar: '',
    desc: '',
    homePath: '/student/list',
  };
}
