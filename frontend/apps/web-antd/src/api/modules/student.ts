import { requestClient } from '#/api/request';

export namespace StudentApi {
  /** 学生信息 */
  export interface Student {
    id?: number;
    name: string;
    gender?: string;
    birthday?: string;
    phone?: string;
    email?: string;
    address?: string;
    major?: string;
    class?: string;
    student_no: string;
    status?: number;
    created_at?: string;
    updated_at?: string;
  }

  /** 分页查询参数 */
  export interface ListParams {
    keyword?: string;
    page: number;
    pageSize: number;
  }

  /** 分页结果 */
  export interface PageResult {
    list: Student[];
    total: number;
    page: number;
    pageSize: number;
  }
}

/** 获取学生列表 */
export function getStudentListApi(params: StudentApi.ListParams) {
  return requestClient.get<StudentApi.PageResult>('/students', { params });
}

/** 获取学生详情 */
export function getStudentDetailApi(id: number) {
  return requestClient.get<StudentApi.Student>(`/students/${id}`);
}

/** 创建学生 */
export function createStudentApi(data: StudentApi.Student) {
  return requestClient.post('/students', data);
}

/** 更新学生 */
export function updateStudentApi(id: number, data: Partial<StudentApi.Student>) {
  return requestClient.put(`/students/${id}`, data);
}

/** 删除学生 */
export function deleteStudentApi(id: number) {
  return requestClient.delete(`/students/${id}`);
}
