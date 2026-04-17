import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    meta: {
      icon: 'lucide:graduation-cap',
      order: 10,
      title: '学生管理',
    },
    name: 'Student',
    path: '/student',
    children: [
      {
        name: 'StudentList',
        path: '/student/list',
        component: () => import('#/views/student/list/index.vue'),
        meta: {
          title: '学生列表',
          icon: 'lucide:list',
        },
      },
    ],
  },
];

export default routes;
